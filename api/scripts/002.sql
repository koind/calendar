create index user_idx on events (user_id);
create index start_idx on events using btree (datetime, duration);