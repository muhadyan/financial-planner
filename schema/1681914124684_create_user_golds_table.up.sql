CREATE TABLE user_golds (
  id serial PRIMARY KEY,
  user_id int4 NOT NULL,
  buy_price float NOT NULL,
  sell_price float,
  buy_date date NOT NULL,
  sell_date date,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_user_golds_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);