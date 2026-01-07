CREATE INDEX idx_working_times_user ON working_times (entity_type, entity_id);

CREATE INDEX idx_users_branch_id ON users(branch_id);
CREATE INDEX idx_users_level_id ON users(level_id);
CREATE INDEX idx_users_position_id ON users(position_id);
CREATE INDEX idx_users_user_type_id ON users(user_type_id);
