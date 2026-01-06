CREATE TYPE genderType AS ENUM ('male', 'female', 'do not tell');
CREATE TYPE entityType AS ENUM('user', 'branch');
CREATE TYPE shiftType AS ENUM('morning', 'afternoon')

CREATE TABLE Users(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  sur_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100)  NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  dob DATE NOT NULL,
  gender genderType NOT NULL DEFAULT 'do not tell',
  phone VARCHAR(11) NOT NULL,
  current_address VARCHAR(255) NULL,
  address VARCHAR(255) NOT NULL,
  avatar_path VARCHAR(255) NULL,
  bank_account VARCHAR(14) NOT NULL,
  identify_number VARCHAR(12) NOT NULL UNIQUE,
  identify_issue_date DATE NOT NULL,
  identify_place VARCHAR(255) NOT NULL,
  emergency_contact VARCHAR(255) NULL,
  emergency_contact_phone VARCHAR(11) NULL,
  tax_code VARCHAR(10) NULL UNIQUE,
  is_active BOOLEAN NOT NULL DEFAULT(false),
  mezon_id VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by uuid REFERENCES Users(id),
  FOREIGN KEY (created_by) REFERENCES Users(id)
);

CREATE TABLE Levels(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  display_name VARCHAR(100) NOT NULL,
  code VARCHAR(10) NOT NULL UNIQUE,
  color VARCHAR(20) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by uuid NOT NULL,
  FOREIGN KEY(created_by) REFERENCES Users(id)
);

CREATE TABLE Branches(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  display_name VARCHAR(100) NOT NULL,
  code VARCHAR(10) NOT NULL UNIQUE,
  color VARCHAR(20) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by uuid NOT NULL,
  FOREIGN KEY(created_by) REFERENCES Users(id)
);

CREATE TABLE Positions(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  short_name VARCHAR(100) NOT NULL,
  code VARCHAR(10) NOT NULL UNIQUE,
  color VARCHAR(20) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by uuid NOT NULL,
  FOREIGN KEY(created_by) REFERENCES Users(id)
);

CREATE TABLE User_Type(
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  code VARCHAR(10) NOT NULL UNIQUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_by uuid NOT NULL,
  FOREIGN KEY(created_by) REFERENCES Users(id)
);

CREATE TABLE working_times (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entity_type entityType NOT NULL, -- 'user' or 'branch'
    entity_id UUID NOT NULL,          -- user.id or branch.id
    shift_name shiftType NOT NULL, -- 'morning' or 'afternoon' (or other future shifts)
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    working_hours NUMERIC(3,1) NOT NULL GENERATED ALWAYS AS (ROUND(EXTRACT(EPOCH FROM (end_time - start_time)) / 3600, 1)) STORED,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by uuid NOT NULL,
    CHECK (end_time > start_time),
    UNIQUE (entity_type, entity_id, shift_name)
);

INSERT INTO user_type (name,code,created_by)
VALUES (
  'Super Admin',
  'SAD',
  NULL
)
ON CONFLICT (code) DO NOTHING;

INSERT INTO levels (name,display_name,code,color,created_by)
VALUES (
  'Administrator',
  'Administrator',
  'ADMIN',
  '#FF0000',
  NULL
)
ON CONFLICT (code) DO NOTHING;

INSERT INTO branches (
  name,
  display_name,
  code,
  color,
  created_by
)
VALUES (
  'Head Office',
  'Head Office',
  'HO',
  '#0000FF',
  NULL
)
ON CONFLICT (code) DO NOTHING;

INSERT INTO positions (
  name,
  short_name,
  code,
  color,
  created_by
)
VALUES (
  'Super Admin',
  'Admin',
  'SAD',
  '#000000',
  NULL
)
ON CONFLICT (code) DO NOTHING;

INSERT INTO users (
  sur_name,
  last_name,
  email,
  dob,
  gender,
  phone,
  address,
  bank_account,
  identify_number,
  identify_issue_date,
  identify_place,
  mezon_id,
  level_id,
  branch_id,
  position_id,
  user_type_id,
  is_active,
  created_by
)
VALUES (
  'Admin',
  'Root',
  'admin@company.com',
  '1990-01-01',
  'do not tell',
  '0123456789',
  'Head Office',
  '00000000000000',
  '123456789012',
  '2010-01-01',
  'Government',
  'MEZON_SUPER_ADMIN',
  (SELECT id FROM levels WHERE code = 'ADMIN'),
  (SELECT id FROM branches WHERE code = 'HO'),
  (SELECT id FROM positions WHERE code = 'SAD'),
  (SELECT id FROM user_type WHERE code = 'SAD'),
  TRUE,
  NULL
);

ALTER TABLE users
ADD COLUMN level_id BIGINT NOT NULL,
ADD COLUMN branch_id BIGINT NOT NULL,
ADD COLUMN position_id BIGINT NOT NULL,
ADD COLUMN user_type_id BIGINT NOT NULL;

ALTER TABLE users
ADD CONSTRAINT fk_level FOREIGN KEY (level_id) REFERENCES levels(id),
ADD CONSTRAINT fk_branch FOREIGN KEY (branch_id) REFERENCES branches(id),
ADD CONSTRAINT fk_position FOREIGN KEY (position_id) REFERENCES positions(id),
ADD CONSTRAINT fk_user_type FOREIGN KEY (user_type_id) REFERENCES user_Type(id);
