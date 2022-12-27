create table if not exists `round` (
    round_id varchar(250) PRIMARY KEY,
    schedule_id varchar(250) not null,
    description varchar(250) not null,
    created_on datetime(3) not null default CURRENT_TIMESTAMP(3),
    updated_on datetime(3) not null default CURRENT_TIMESTAMP(3)
);

create table if not exists schedule (
    schedule_id varchar(250) PRIMARY KEY,
    description varchar(250) not null,
    created_on datetime(3) not null default CURRENT_TIMESTAMP(3),
    updated_on datetime(3) not null default CURRENT_TIMESTAMP(3),
    constraint idx_description_schedule unique (description)
);





