create table if not exists users (
    id serial primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(50) not null,
    password text not null,
    role int not null, /* 1 client, 2 seller, 3 admin */
    created_at timestamp,
    updated_at timestamp
);

create table if not exists item (
    id serial primary key ,
    name varchar(50) not null,
    price numeric(12,2) not null,
    rating numeric(3,2) not null,
    number_of_ratings int not null,
    category varchar(50) not null,
    seller_id int references users(id) not null,
    created_at timestamp,
    updated_at timestamp
);

insert into item (name, price, rating, number_of_ratings, category, seller_id)
values ('Makhabbat', 120, 3.4, 1, 'book', 1);
insert into item (name, price, rating, number_of_ratings, category, seller_id)
values ('Laptop', 1200, 4.1, 1, 'Gadget', 1);