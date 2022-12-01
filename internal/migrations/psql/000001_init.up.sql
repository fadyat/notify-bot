create type frequency_type as enum ('once', 'daily', 'weekly', 'monthly');
create sequence reminders_id_seq;

create table reminders
(
    id         int            not null default nextval('reminders_id_seq'),
    user_id    int            not null,
    name       varchar(255)   not null,
    frequency  frequency_type not null,
    completed  boolean        not null default false,
    created_at timestamp      not null default now(),
    updated_at timestamp      not null default now(),
    deadline   timestamp      not null,

    primary key (id)
);

alter sequence reminders_id_seq owned by reminders.id;

