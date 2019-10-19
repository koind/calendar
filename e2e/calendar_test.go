package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/streadway/amqp"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	queueName    = "events"
	exchangeName = "notify_about_event"
)

var amqpDSN = os.Getenv("TESTS_AMQP_DSN")

type Event struct {
	ID             int           `json:"id" db:"id"`                                       //ID - уникальный идентификатор события
	Title          string        `json:"title" db:"title"`                                 //Заголовок - короткий текст
	Datetime       time.Time     `json:"datetime" db:"datetime"`                           //Дата и время события
	Duration       time.Duration `json:"duration" db:"duration"`                           //Длительность события
	Description    string        `json:"description,omitempty" db:"description"`           //Описание события, опционально
	UserId         int           `json:"user_id" db:"user_id"`                             //Пользователь, владелец события
	TimeSendNotify time.Time     `json:"time_send_notify,omitempty" db:"time_send_notify"` //За сколько времени высылать уведомление, опционально
}

type notifyTest struct {
	conn          *amqp.Connection
	ch            *amqp.Channel
	messages      [][]byte
	messagesMutex sync.RWMutex
	stopSignal    chan struct{}

	responseStatusCode int
	responseBody       []byte
	eventID            int
	event              Event
	events             []Event
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (test *notifyTest) startConsuming(interface{}) {
	test.messages = make([][]byte, 0)
	test.messagesMutex = sync.RWMutex{}
	test.stopSignal = make(chan struct{})

	var err error

	test.conn, err = amqp.Dial(amqpDSN)
	panicOnErr(err)

	test.ch, err = test.conn.Channel()
	panicOnErr(err)

	// Consume
	_, err = test.ch.QueueDeclare(queueName, true, false, true, false, nil)
	panicOnErr(err)

	err = test.ch.QueueBind(queueName, "", exchangeName, false, nil)
	panicOnErr(err)

	events, err := test.ch.Consume(queueName, "", true, false, false, false, nil)
	panicOnErr(err)

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				return
			case event := <-events:
				test.messagesMutex.Lock()
				test.messages = append(test.messages, event.Body)
				test.messagesMutex.Unlock()
			}
		}
	}(test.stopSignal)
}

func (test *notifyTest) stopConsuming(interface{}, error) {
	test.stopSignal <- struct{}{}

	panicOnErr(test.ch.Close())
	panicOnErr(test.conn.Close())
	test.messages = nil
}

func (test *notifyTest) iSendRequestToWithData(httpMethod, url, contentType string, data *gherkin.DocString) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodPost:
		replacer := strings.NewReplacer("\n", "", "\t", "")
		cleanJson := replacer.Replace(data.Content)
		r, err = http.Post(url, contentType, bytes.NewReader([]byte(cleanJson)))
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)

	err = json.Unmarshal(test.responseBody, &test.event)
	test.eventID = test.event.ID

	return
}

func (test *notifyTest) iSendRequestToWithDataForUpdate(httpMethod, url, contentType string, data *gherkin.DocString) (err error) {
	var r *http.Response

	if strings.HasSuffix(url, ":id") {
		id := strconv.Itoa(test.event.ID)
		url = strings.ReplaceAll(url, ":id", id)
	}

	switch httpMethod {
	case http.MethodPut:
		replacer := strings.NewReplacer("\n", "", "\t", "")
		cleanJson := replacer.Replace(data.Content)

		client := &http.Client{}

		req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(cleanJson)))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Accept", contentType)
		r, err = client.Do(req)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)

	err = json.Unmarshal(test.responseBody, &test.event)
	test.eventID = test.event.ID

	return
}

func (test *notifyTest) theResponseCodeShouldBe(code int) error {
	if test.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", test.responseStatusCode, code)
	}
	return nil
}

func (test *notifyTest) theResponseShouldMatchTitle(title string) error {
	if test.event.Title != title {
		return fmt.Errorf("title do not match: %s != %s", test.event.Title, title)
	}
	return nil
}

func (test *notifyTest) iSendRequestTo(httpMethod, url string) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodGet:
		r, err = http.Get(url)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	test.responseStatusCode = r.StatusCode
	test.responseBody, err = ioutil.ReadAll(r.Body)

	err = json.Unmarshal(test.responseBody, &test.events)

	return
}

func (test *notifyTest) theResponseShouldMatchJson(data *gherkin.DocString) error {
	var events []Event

	err := json.Unmarshal([]byte(data.Content), &events)
	if err != nil {
		return err
	}

	if len(test.events) != len(events) {
		return fmt.Errorf("number of events must match: %d != %d", len(test.events), len(events))
	}

	return nil
}

func (test *notifyTest) iChangeTimeForEvent() error {
	var event struct {
		TimeSendNotify time.Time `json:"time_send_notify"`
	}

	event.TimeSendNotify = time.Now().Add(time.Duration(2) * time.Minute)

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	var data = gherkin.DocString{
		Content: string(b),
	}

	err = test.iSendRequestToWithDataForUpdate("PUT", "http://api:7766/event/:id", "application/json", &data)
	if err != nil {
		return err
	}

	return nil
}

func (test *notifyTest) iReceiveEventWithTitle(title string) error {
	time.Sleep(120 * time.Second)

	test.messagesMutex.RLock()
	defer test.messagesMutex.RUnlock()

	for _, msg := range test.messages {
		var event Event

		err := json.Unmarshal(msg, &event)
		if err != nil {
			return err
		}

		if event.Title == title {
			return nil
		}
	}
	return fmt.Errorf("event with titile '%s' was not found in %s", title, test.messages)
}

func FeatureContext(s *godog.Suite) {
	test := new(notifyTest)

	s.BeforeScenario(test.startConsuming)

	s.Step(`^1. I send "([^"]*)" request to "([^"]*)" with "([^"]*)" data:$`, test.iSendRequestToWithData)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)

	s.Step(`^2. I send "([^"]*)" request to "([^"]*)" with "([^"]*)" data:$`, test.iSendRequestToWithDataForUpdate)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match title "([^"]*)"$`, test.theResponseShouldMatchTitle)

	s.Step(`^3. I send "([^"]*)" request to "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match json:$`, test.theResponseShouldMatchJson)

	s.Step(`^4. I change time_send_notify for event$`, test.iChangeTimeForEvent)
	s.Step(`^I receive event with title "([^"]*)"$`, test.iReceiveEventWithTitle)

	s.AfterScenario(test.stopConsuming)
}
