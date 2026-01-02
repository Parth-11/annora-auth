CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	email TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL,
	role TEXT NOT NULL DEFAULT 'user',
	email_verified BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT now()
);