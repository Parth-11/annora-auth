CREATE TABLE refresh_tokens (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID REFERENCES users(id) ON DELETE CASCADE,
	token TEXT UNIQUE NOT NULL,
	expires_at TIMESTAMP NOT NULL,
	revoked BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_refresh_token ON refresh_tokens(token);