create table users (
   id           serial primary key,
   first_name   varchar(100) not null,
   last_name    varchar(100) not null,
   email        varchar(255) unique not null,
   phone_number varchar(20)
);