
create table if not exists bookstore
(
    id         serial,
    author     varchar(50) NOT NULL ,
    name       varchar(50) NOT NULL ,
    price      int NOT NULL ,
    isSold     boolean default false,
    sellerID   int references users(id) on delete cascade NOT NULL ,
    created_at timestamp,
    updated_at timestamp
);

create table if not exists users
(
    id          serial primary key,
    first_name  varchar(50) NOT NULL ,
    last_name   varchar(50) NOT NULL ,
    email       varchar(50) NOT NULL ,
    password    varchar(50) NOT NULL ,
    accesslevel int default 1,
    balance     int NOT NULL ,
    created_at  timestamp,
    updated_at  timestamp
);

insert into users (first_name, last_name, email, password, balance) values ('tttt', 'yyyy','ali@gmail.com', 444444, 19999)