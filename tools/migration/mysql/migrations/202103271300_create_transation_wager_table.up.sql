CREATE table transaction_wagers (
    id bigint not null AUTO_INCREMENT,
    buying_price decimal(10, 2) not null,
    wager_id BIGINT not null,
    bought_at TIMESTAMP not null,
    PRIMARY KEY (id),
    FOREIGN KEY (wager_id) REFERENCES wagers(id)
)
