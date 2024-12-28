CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    customer_id VARCHAR(36) NOT NULL,
    contract_number VARCHAR(50) NOT NULL,
    otr DECIMAL(15, 2) NOT NULL,
    admin_fee DECIMAL(15, 2) NOT NULL,
    installment_amount DECIMAL(15, 2) NOT NULL,
    interest_amount DECIMAL(15, 2) NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);