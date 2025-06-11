BEGIN;

CREATE TYPE payment_source AS ENUM ('midtrans', 'others');

CREATE TYPE payment_status AS ENUM ('pending', 'succeed', 'failed', 'expired');

CREATE TYPE payment_type AS ENUM ('donation', 'others');

CREATE TYPE payment_method AS ENUM ('gopay', 'shopeepay', 'qris', 'bank_transfer', 'none', 'others');

CREATE TABLE IF NOT EXISTS payments (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	code CHAR(10) UNIQUE NOT NULL,
	source payment_source DEFAULT 'midtrans'::payment_source NOT NULL,
	status payment_status DEFAULT 'pending'::payment_status NOT NULL,
	type payment_type DEFAULT 'donation'::payment_type NOT NULL,
	method payment_method DEFAULT 'none'::payment_method NOT NULL,
	amount NUMERIC(10, 2) NOT NULL,
	currency VARCHAR(5),
	bank_name VARCHAR(8),
	va_number VARCHAR(50),
	qr_code TEXT,
	payment_link TEXT,
	expired_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
