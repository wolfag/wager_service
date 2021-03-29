CREATE table wagers (
    id bigint not null AUTO_INCREMENT,
    total_wager_value int unsigned not null,
    odds int unsigned not null,
    selling_percentage int unsigned not null,
    selling_price decimal(10, 2) not null,
    current_selling_price decimal(10, 2) not null default(selling_price),
    percentage_sold decimal(5 ,2),
    amount_sold decimal(10, 2),
    placed_at TIMESTAMP not null,
    PRIMARY KEY(id),
    CONSTRAINT `total_wager_value_greater_0` CHECK(total_wager_value > 0),
    CONSTRAINT `odds_value_greater_0` CHECK(odds > 0),
	CONSTRAINT `selling_percentage_between_1_100` CHECK(100 >= selling_percentage AND selling_percentage >= 1),
	CONSTRAINT `selling_price_value` CHECK(selling_price > total_wager_value * (selling_percentage / 100))
)