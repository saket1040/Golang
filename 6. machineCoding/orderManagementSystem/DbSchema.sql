CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    house VARCHAR(255),
    pincode VARCHAR(20)
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    order_no SERIAL UNIQUE,
    user_id UUID REFERENCES users(id),
    amount DECIMAL(10,2) NOT NULL,
    status INT NOT NULL, -- 1: Created, 2: Paid, ...
    payment_id UUID REFERENCES payments(id),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE items (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    order_id UUID REFERENCES orders(id)
);

CREATE TABLE payments (
    id UUID PRIMARY KEY,
    amount DECIMAL(10,2) NOT NULL,
    type INT NOT NULL, -- 1: Cash, 2: UPI
    created_at TIMESTAMP NOT NULL DEFAULT now()
);