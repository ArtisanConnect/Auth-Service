CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       username VARCHAR(255) UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       full_name VARCHAR(255) NOT NULL,
                       user_type VARCHAR(50) NOT NULL,
                       bio TEXT,
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE user_tokens (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             user_id UUID REFERENCES users(id),
                             access_token VARCHAR(255) NOT NULL,
                             refresh_token VARCHAR(255) NOT NULL,
                             expires_in INTEGER NOT NULL,
                             created_at TIMESTAMP DEFAULT now()
);