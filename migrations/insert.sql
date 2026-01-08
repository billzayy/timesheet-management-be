	BEGIN;
	
	INSERT INTO branches (name,display_name,code,color,created_by) VALUES
	('HN1','HN1','HN1','#0000FF','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('HN2','HN2','HN2','#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('HN3','HN3','HN3','blue', '155dd490-6528-41ba-aa16-d4ec616120fb'),
	('SG1','SG1','SG1','#2196f3','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('SG2','SG2','SG2','#17a2b8','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('ĐN','ĐN','ĐN','purple','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Vinh','Vinh','Vinh','#6c757d','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('QN','QN','QN','Brown','155dd490-6528-41ba-aa16-d4ec616120fb')
	ON CONFLICT (code) DO NOTHING;
	
	INSERT INTO positions (name, short_name, code, color, created_by) VALUES
  ('Super Admin','Admin','SAD','#000000','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('dev','dev','dev','#c81919','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('IT','IT','IT','#389951','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('PM','PM','PM','#2bb65b','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('HR','HR','HR','#1468d7','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Sale','Sale','Sale','#9d30c5','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Tester','Tester','Tester','#1eccc9','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Test Support','Test Support','Test Sup','#009dff','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Sale Support','Sale Support','Sale Sup','#1cf220','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Sing','Sing','Sing','#a42d2d','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('PO','PO','PO','#d93636','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('BA','BA','BA','#ba3636','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Art','Art','Art','#971c1c','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('DevOps','DevOps','DevOps','#c91818','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Auto test','Auto test','Auto test','#dc3e09','155dd490-6528-41ba-aa16-d4ec616120fb')
	ON CONFLICT (code) DO NOTHING;
	
	INSERT INTO user_type (name, code, created_by) VALUES
	('Super Admin','sad','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Staff','sta','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Intern','int','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Colaborator','col','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Probation','pro','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Vendor','ven','155dd490-6528-41ba-aa16-d4ec616120fb')
	ON CONFLICT (code) DO NOTHING;
	
	INSERT INTO levels (name, display_name, code, color, created_by) VALUES
  ('Administrator','Administrator','ADMIN','#FF0000','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Intern 0', 'Intern_0', 'int_0', '#ff9800', '155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Intern 1', 'Intern_1', 'int_1', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Intern 2', 'Intern_2', 'int_2', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Intern 3', 'Intern_3', 'int_3', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Fresher Low', 'Fresher-', 'fre_low', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Fresher', 'Fresher', 'fre_mid', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Fresher High', 'Fresher+', 'fre_high', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Junior Low', 'Junior-', 'jun_low', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Junior', 'Junior', 'jun_mid', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Junior High', 'Junior+', 'jun_high', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Middle Low', 'Middle-', 'mid_low', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Middle', 'Middle', 'mid_mid', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Middle High', 'Middle+', 'mid_high', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Senior Low', 'Senior-', 'sen_low', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Senior', 'Senior', 'sen_mid', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb'),
	('Senior High', 'Senior+', 'sen_high', '#ff9800','155dd490-6528-41ba-aa16-d4ec616120fb')
	ON CONFLICT (code) DO NOTHING;
	
INSERT INTO working_times (
    entity_type,
    entity_id,
    shift_name,
    start_time,
    end_time
) VALUES
('branch', '7f9d7362-c70a-40f6-bc47-b8f3ab53c6bd', 'morning', '08:00', '12:00'),
('branch', '0cf442b5-78d8-4c0b-aebd-0967fa291c0f', 'morning', '08:00', '12:00'),
('branch', '52cbefc4-1411-4ad3-81b2-bcaed4550eca', 'morning', '08:00', '12:00'),
('branch', '7d4e7529-b514-4a33-a411-1369c4e1bc5a', 'morning', '08:00', '12:00'),
('branch', 'de156456-58eb-4756-bf8b-c26d063e56e6', 'morning', '08:00', '12:00'),
('branch', '02c4a702-a1a2-4880-9ad4-8529c46909cd', 'morning', '08:00', '12:00'),
('branch', '8c6cb11f-e9d6-4158-abb8-0e4a56373424', 'morning', '08:00', '12:00'),
('branch', '55112fe0-8590-496b-902a-a11678448118', 'morning', '08:00', '12:00');

INSERT INTO working_times (
    entity_type,
    entity_id,
    shift_name,
    start_time,
    end_time
) VALUES
('branch', '7f9d7362-c70a-40f6-bc47-b8f3ab53c6bd', 'afternoon', '13:00', '17:00'),
('branch', '0cf442b5-78d8-4c0b-aebd-0967fa291c0f', 'afternoon', '13:00', '17:00'),
('branch', '52cbefc4-1411-4ad3-81b2-bcaed4550eca', 'afternoon', '13:00', '17:00'),
('branch', '7d4e7529-b514-4a33-a411-1369c4e1bc5a', 'afternoon', '13:00', '17:00'),
('branch', 'de156456-58eb-4756-bf8b-c26d063e56e6', 'afternoon', '13:00', '17:00'),
('branch', '02c4a702-a1a2-4880-9ad4-8529c46909cd', 'afternoon', '13:00', '17:00'),
('branch', '8c6cb11f-e9d6-4158-abb8-0e4a56373424', 'afternoon', '13:00', '17:00'),
('branch', '55112fe0-8590-496b-902a-a11678448118', 'afternoon', '13:00', '17:00');COMMIT;

-- Permissions
START TRANSACTION;

INSERT INTO permissions (name, display_name, created_by) VALUES
('Admin','Admin','155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyProfile','My profile','155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyTimesheet','My timesheets','155dd490-6528-41ba-aa16-d4ec616120fb'),
('ManageTimesheet','Manage Timesheets','155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project','Projets', '155dd490-6528-41ba-aa16-d4ec616120fb');

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('Admin.Users', 'Users', (SELECT id FROM permissions WHERE name='Admin'), '155dd490-6528-41ba-aa16-d4ec616120fb');

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('Admin.Users.View', 'View users', (SELECT id FROM permissions WHERE name='Admin.Users'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Admin.Users.AddNew', 'Add new user', (SELECT id FROM permissions WHERE name='Admin.Users'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Admin.Users.Edit', 'Edit user', (SELECT id FROM permissions WHERE name='Admin.Users'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Admin.Users.EditRole', 'Edit user role', (SELECT id FROM permissions WHERE name='Admin.Users'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Admin.Users.Delete', 'Delete user', (SELECT id FROM permissions WHERE name='Admin.Users'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Admin.Users.ResetPassword', 'Reset password', (SELECT id FROM permissions WHERE name='Admin.Users'), '155dd490-6528-41ba-aa16-d4ec616120fb');

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('Project.View', 'View my projects', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.ViewAll', 'View all projects', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.AddNew', 'Add new project', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.Edit', 'Edit project', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.Delete', 'Delete project', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.ChangeStatus', 'Change project status', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.ViewDetail', 'View project detail', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('Project.Export', 'Export project', (SELECT id FROM permissions WHERE name='Project'), '155dd490-6528-41ba-aa16-d4ec616120fb');

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('MyProfile.View', 'View profile', (SELECT id FROM permissions WHERE name='MyProfile'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyProfile.RequestUpdateInfo', 'Request update info', (SELECT id FROM permissions WHERE name='MyProfile'), '155dd490-6528-41ba-aa16-d4ec616120fb');

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('MyTimesheet.View', 'View my timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyTimesheet.AddNew', 'Add new timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyTimesheet.Edit', 'Edit timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyTimesheet.Delete', 'Delete timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('MyTimesheet.Submit', 'Submit timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb');


INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('ManagementTimesheet.View', 'View all timesheets', (SELECT id FROM permissions WHERE name='ManageTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('ManagementTimesheet.Approve', 'Approve timesheet', (SELECT id FROM permissions WHERE name='ManageTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('ManagementTimesheet.Reject', 'Reject timesheet', (SELECT id FROM permissions WHERE name='ManageTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb'),
('ManagementTimesheet.Export', 'Export timesheet', (SELECT id FROM permissions WHERE name='ManageTimesheet'), '155dd490-6528-41ba-aa16-d4ec616120fb');

COMMIT;
ROLLBACK;

CREATE UNIQUE INDEX IF NOT EXISTS ux_permissions_name ON permissions(name);

-- Roles
INSERT INTO roles(name,display_name, description, created_by) VALUES
('Admin','Admin','Administrator','155dd490-6528-41ba-aa16-d4ec616120fb'),
('All','All','Full Permission','155dd490-6528-41ba-aa16-d4ec616120fb'),
('Approve Review Intern','ApproveReviewInten','Approve Review Intern','155dd490-6528-41ba-aa16-d4ec616120fb'),
('Basic User','Basic User','Basic User','155dd490-6528-41ba-aa16-d4ec616120fb'),
('Branch Director','Branch Director','Branch Director','155dd490-6528-41ba-aa16-d4ec616120fb'),
('Create Review Intern','Create Review Intern','Create Review Intern','155dd490-6528-41ba-aa16-d4ec616120fb'),
('HDQT','HDQT','HDQT','155dd490-6528-41ba-aa16-d4ec616120fb'),
('HPM','Head PM','Head PM','155dd490-6528-41ba-aa16-d4ec616120fb'),
('PorjectManager','PM','Project Manager','155dd490-6528-41ba-aa16-d4ec616120fb'),
('PQA','PQA','PQA','155dd490-6528-41ba-aa16-d4ec616120fb'),
('View All Project','View All Project','View All Project','155dd490-6528-41ba-aa16-d4ec616120fb'),
('View User','View User','View User','155dd490-6528-41ba-aa16-d4ec616120fb'),
('Test Review Intern','Test Review Intern','Test Review Intern','155dd490-6528-41ba-aa16-d4ec616120fb');

-- Role_Permission
-- Admin gets all permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'Admin';

-- All gets all permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'All';

-- Basic User Permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
JOIN permissions p ON p.name IN (
    'MyProfile',
    'MyProfile.View',
    'MyProfile.RequestUpdateInfo',
    'MyTimesheet',
    'MyTimesheet.View',
    'MyTimesheet.AddNew',
    'MyTimesheet.Edit',
    'MyTimesheet.Delete',
    'MyTimesheet.Submit',
    'Project.View',
    'Project.ViewDetail'
)
WHERE r.name = 'Basic User';

-- HeadPM/PM permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
JOIN permissions p ON p.name LIKE 'Project.%'
   OR p.name LIKE 'ManagementTimesheet.%'
WHERE r.name IN ('PorjectManager', 'HPM');

-- Branch Director/HDQT permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
JOIN permissions p ON p.name IN (
    'Project.ViewAll',
    'Project.ViewDetail',
    'Project.Export',
    'ManagementTimesheet.View',
    'ManagementTimesheet.Approve',
    'ManagementTimesheet.Export'
)
WHERE r.name IN ('Branch Director', 'HDQT');

-- Create Review Intern permission
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
JOIN permissions p ON p.name IN (
    'Project.View',
    'Project.ViewDetail',
    'MyTimesheet.View',
    'MyTimesheet.AddNew',
    'MyTimesheet.Submit'
)
WHERE r.name = 'Create Review Intern';

-- Approve/Test Review Intern permission
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
JOIN permissions p ON p.name IN (
    'Project.ViewAll',
    'Project.ViewDetail',
    'ManagementTimesheet.View',
    'ManagementTimesheet.Approve'
)
WHERE r.name IN ('Approve Review Intern', 'Test Review Intern');

-- View permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    '155dd490-6528-41ba-aa16-d4ec616120fb'
FROM roles r
JOIN permissions p ON p.name LIKE '%.View%'
WHERE r.name IN ('View All Project', 'View User');
