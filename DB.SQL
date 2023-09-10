create table users(
    id integer, -- primary key
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255),
    password varchar(255),
    created_at date,
    updated_at date,
    deleted_at date,
    access_level integer,
);

create table reservations(
    id integer, -- primary key
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255),
    phone varchar(255),
    room_id integer, -- foreign key from rooms table
    start_date date,
    send_date date,
    created_at date,
    updated_at date,
    deleted_at date,
);
create table rooms(
    id integer, -- primary key
    room_name varchar(255),
    created_at date,
    updated_at date,
    deleted_at date,
);
create table room_restrictions(
    id integer, -- primary key
    start_date date,
    send_date date,
    room_id integer, -- foreign key from rooms table
    reservation_id integer, -- foreign key from reservation table
    created_at date,
    updated_at date,
    deleted_at date,
    restrictions_id integer,

);

create table restrictions {
    id integer, -- primary key
    restriction_name varchar(255),
    created_at date,
    updated_at date,
    deleted_at date,
}



-- create table users(

-- );