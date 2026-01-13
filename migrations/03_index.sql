CREATE INDEX idx_working_times_user ON working_times (entity_type, entity_id);

CREATE INDEX idx_users_branch_id ON users(branch_id);
CREATE INDEX idx_users_level_id ON users(level_id);
CREATE INDEX idx_users_position_id ON users(position_id);
CREATE INDEX idx_users_user_type_id ON users(user_type_id);
CREATE INDEX idx_users_created_by ON users(created_by);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_mezon_id ON users(mezon_id);

CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);
CREATE INDEX idx_role_permissions_created_by ON role_permissions(created_by);
