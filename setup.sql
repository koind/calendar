create table events (
    id serial primary key,
    title text not null,
    datetime timestamp not null,
    duration bigint not null,
    description text,
    user_id bigint not null,
    time_send_notify timestamp
);
create index user_idx on events (user_id);
create index start_idx on events using btree (datetime, duration);