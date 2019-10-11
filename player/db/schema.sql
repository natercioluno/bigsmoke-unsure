create table smoke_cursors
(
    id         varchar(255) not null,
    `cursor`   bigint       not null,
    updated_at datetime(3)  not null,

    primary key (id)
);

create table smoke_events
(
    id         bigint      not null auto_increment,
    foreign_id bigint      not null,
    timestamp  datetime(3) not null,
    type       int         not null,

    primary key (id)
);

create table player_parts
(
    id         bigint      not null auto_increment,
    round_id   bigint      not null,
    player_id  varchar(16) not null,
    part       int         not null,
    created_at datetime(3) not null,

    primary key (id)
);

create table rounds
(
    id         bigint      not null auto_increment,
    round_id   int         not null,
    included   bool        not null,

    status     int         not null,
    created_at datetime(3) not null,
    updated_at datetime(3) not null,

    primary key (id)
);
