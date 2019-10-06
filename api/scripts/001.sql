create table events (
    id serial primary key,
    title text not null,
    datetime timestamp not null,
    duration bigint not null,
    description text,
    user_id bigint not null,
    time_send_notify timestamp
);