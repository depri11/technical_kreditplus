CREATE TABLE customer_limits (
    customer_id BIGINT UNIQUE,
    tenor_1_month NUMERIC(15, 2) NOT NULL,
    tenor_2_months NUMERIC(15, 2) NOT NULL,
    tenor_3_months NUMERIC(15, 2) NOT NULL,
    tenor_4_months NUMERIC(15, 2) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);