CREATE TABLE borrower (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    amount NUMERIC(10, 2) CHECK (amount >= 0),
    tax_rate NUMERIC(5, 4) DEFAULT 0.0000 CHECK (tax_rate >= 0 AND tax_rate <= 1),
    tax_amount NUMERIC(10, 2) DEFAULT 0.00 CHECK (tax_amount >= 0),
    final_amount NUMERIC(12, 2) GENERATED ALWAYS AS (amount + tax_amount) STORED,
    is_settled BOOLEAN DEFAULT FALSE,
    transaction_date timestamp DEFAULT null,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    borrow_id INT NOT NULL,
    total_payment NUMERIC(10, 2) CHECK (total_payment >= 0),
    is_paid BOOLEAN DEFAULT FALSE,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_borrower FOREIGN KEY (borrow_id) REFERENCES borrower(id)
);