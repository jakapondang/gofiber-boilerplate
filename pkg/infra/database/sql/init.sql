CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
                       id UUID PRIMARY KEY,  -- Removed DEFAULT uuid_generate_v4()
                       username VARCHAR(50) UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       first_name VARCHAR(50),
                       last_name VARCHAR(50),
                       phone_number VARCHAR(20),
                       is_verify_email BOOLEAN DEFAULT FALSE,
                       is_verify_phone_number BOOLEAN DEFAULT FALSE,
                       is_active BOOLEAN DEFAULT TRUE,
                       is_admin BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       last_login TIMESTAMPTZ
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_updated_at ON users(updated_at);

CREATE TABLE roles (
                       id UUID PRIMARY KEY ,
                       name VARCHAR(50) UNIQUE NOT NULL,
                       description TEXT,
                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_created_at ON roles(created_at);
CREATE INDEX idx_roles_updated_at ON roles(updated_at);

CREATE TABLE user_roles (
                            user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                            role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
                            PRIMARY KEY (user_id, role_id)
);

CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);

CREATE TABLE password_resets (
                                 user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                 reset_token UUID PRIMARY KEY,  -- No default value for UUID
                                 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                 expires_at TIMESTAMPTZ NOT NULL,
                                 used BOOLEAN DEFAULT FALSE,
                                 CONSTRAINT reset_token_expiry CHECK (expires_at > CURRENT_TIMESTAMP)
);