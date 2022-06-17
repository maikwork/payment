CREATE TABLE payment(
    id serial PRIMARY KEY NOT NULL,
    id_client INT NOT NULL,
    email varchar(50) NOT NULL, 
    amount money not null,
    currency varchar(1) not null,
    time_created timestamp not null,
    time_changed timestamp not null,
    status_pay int
);

CREATE TABLE client(
    id serial PRIMARY KEY NOT NULL,
    email VARCHAR(50) NOT NULL
);
