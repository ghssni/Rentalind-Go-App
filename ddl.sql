CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  deposit_amount DECIMAL(10,2) DEFAULT 0
);

CREATE TABLE rentals (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  availability INT NOT NULL,
  rental_cost DECIMAL(10,2) NOT NULL,
  category VARCHAR(255) NOT NULL
);

CREATE TABLE rental_history (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  rental_id INT NOT NULL,
  rental_start_date DATETIME NOT NULL,
  rental_end_date DATETIME NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (rental_id) REFERENCES rentals(id)
);

-- dbdiagram.io code

Table users {
  id INT [pk, increment]
  name string
  address string
  email VARCHAR(255) [not null, unique]
  password VARCHAR(255) [not null]
  deposit DECIMAL(10,2) [default: 0]
}

Table rentals {
  id INT [pk, increment]
  name VARCHAR(255) [not null]
  availability INT [not null]
  rental_cost DECIMAL(10,2) [not null]
  category VARCHAR(255) [not null]
}

Table rental_history {
  id INT [pk, increment]
  user_id INT [not null]
  rental_id INT [not null]
  payment_id INT [not null]
  rental_start_date DATETIME [not null]
  rental_end_date DATETIME [not null]
}

Table payment {
  id INT [pk, increment]
  amount DECIMAL(10,2)
  status string
  url string
}

ref: payment.id - rental_history.payment_id
ref: rental_history.user_id > users.id
ref: rental_history.rental_id > rentals.id