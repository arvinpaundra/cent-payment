BEGIN;

CREATE TYPE payment_source AS ENUM ('midtrans', 'others');

CREATE TYPE payment_status AS ENUM ('pending', 'paid', 'failed', 'expired');

CREATE TYPE payment_purpose AS ENUM ('donation', 'others');

CREATE TYPE payment_method AS ENUM ('gopay', 'shopeepay', 'qris', 'none', 'others');

CREATE TABLE IF NOT EXISTS payments (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	code CHAR(10) UNIQUE NOT NULL,
	source payment_source DEFAULT 'midtrans'::payment_source NOT NULL,
	status payment_status DEFAULT 'pending'::payment_status NOT NULL,
	method payment_method DEFAULT 'none'::payment_method NOT NULL,
	purpose payment_purpose DEFAULT 'donation'::payment_purpose NOT NULL,
	amount NUMERIC(10, 2) NOT NULL,
	reference VARCHAR(255),
	currency VARCHAR(5),
	qr_code TEXT,
	payment_link TEXT,
	expired_at TIMESTAMP,
	paid_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
