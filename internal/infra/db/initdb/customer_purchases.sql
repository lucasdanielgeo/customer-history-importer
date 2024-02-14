CREATE TABLE IF NOT EXISTS customers_purchase_history (
    id SERIAL NOT NULL,
    cpf VARCHAR(20) NOT NULL,
    is_cpf_valid BOOLEAN DEFAULT false,
    "private" BOOLEAN NULL,
    incomplete BOOLEAN NULL,
    last_purchase_date DATE NULL,
    average_ticket NUMERIC(12, 2) NULL,
    last_purchase_ticket NUMERIC(12, 2) NULL,
    most_frequent_store VARCHAR(20) NULL,
    last_purchase_store VARCHAR(20) NULL,
    is_most_frequent_store_valid BOOLEAN DEFAULT false,
    is_last_purchase_store_valid BOOLEAN DEFAULT false
);