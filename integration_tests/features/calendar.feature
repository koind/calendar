# file: features/notification.feature

# http://localhost:7766/
# http://api:7766/

Feature: microservices for calendar implementation
  It should be possible to create, update events.
  Also receive event lists.
  It should also be possible to send notifications and clear old events.

  Scenario: Api create new event
    When 1. I send "POST" request to "http://api:7766/event" with "application/json" data:
    """
    {
        "title": "test",
        "datetime": "2019-08-18T10:00:00Z",
        "duration": 15,
        "description": "test description",
        "user_id": 3,
        "time_send_notify": "2019-08-18T10:00:00Z"
    }
    """
    Then The response code should be 200

  Scenario: Api update event
    When 2. I send "PUT" request to "http://api:7766/event/:id" with "application/json" data:
    """
    {
        "title": "Changed test title"
    }
    """
    Then The response code should be 200
    And The response should match title "Changed test title"

  Scenario: Api get all event list
    When  3. I send "GET" request to "http://api:7766/events"
    Then The response code should be 200
    And The response should match json:
    """
    [
      {
        "id": 1,
        "title": "Changed test title",
        "datetime": "2019-08-18T10:00:00Z",
        "duration": 15,
        "description": "test description",
        "user_id": 3,
        "time_send_notify": "2019-08-18T10:00:00Z"
      }
    ]
    """

  Scenario: Scheduler sent event
    When 4. I change time_send_notify for event
    Then I receive event with title "Changed test title"