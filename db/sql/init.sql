CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL
);

CREATE TABLE borrow (
    id SERIAL4 PRIMARY KEY,
    user_id int NOT NULL,
    amount NUMERIC(10, 2) CHECK (amount >= 0),
    tax_rate NUMERIC(5, 4) DEFAULT 0.0000 CHECK (tax_rate >= 0 AND tax_rate <= 1),
    tax_amount NUMERIC(10, 2) DEFAULT 0.00 CHECK (tax_amount >= 0),
    final_amount NUMERIC(12, 2) GENERATED ALWAYS AS (amount + tax_amount) STORED,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

