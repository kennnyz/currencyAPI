CREATE TABLE operations
(
    currency_from varchar(255) not null,
    currency_to varchar(255) not null,
    exchange_rate numeric not null,
    updated_at timestamp
);