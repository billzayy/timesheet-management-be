	BEGIN;
	
	INSERT INTO branches (name,display_name,code,color,created_by) VALUES
	('HN1','HN1','HN1','#0000FF',(SELECT id FROM users WHERE sur_name = 'Admin')),
	('HN2','HN2','HN2','#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('HN3','HN3','HN3','blue', (SELECT id FROM users WHERE sur_name='Admin')),
	('SG1','SG1','SG1','#2196f3',(SELECT id FROM users WHERE sur_name='Admin')),
	('SG2','SG2','SG2','#17a2b8',(SELECT id FROM users WHERE sur_name='Admin')),
	('ĐN','ĐN','ĐN','purple',(SELECT id FROM users WHERE sur_name='Admin')),
	('Vinh','Vinh','Vinh','#6c757d',(SELECT id FROM users WHERE sur_name='Admin')),
	('QN','QN','QN','Brown',(SELECT id FROM users WHERE sur_name='Admin'))
	ON CONFLICT (code) DO NOTHING;
	
	INSERT INTO positions (name, short_name, code, color, created_by) VALUES
	('dev','dev','dev','#c81919',(SELECT id FROM users WHERE sur_name='Admin')),
	('IT','IT','IT','#389951',(SELECT id FROM users WHERE sur_name='Admin')),
	('PM','PM','PM','#2bb65b',(SELECT id FROM users WHERE sur_name='Admin')),
	('HR','HR','HR','#1468d7',(SELECT id FROM users WHERE sur_name='Admin')),
	('Sale','Sale','Sale','#9d30c5',(SELECT id FROM users WHERE sur_name='Admin')),
	('Tester','Tester','Tester','#1eccc9',(SELECT id FROM users WHERE sur_name='Admin')),
	('Test Support','Test Support','Test Sup','#009dff',(SELECT id FROM users WHERE sur_name='Admin')),
	('Sale Support','Sale Support','Sale Sup','#1cf220',(SELECT id FROM users WHERE sur_name='Admin')),
	('Sing','Sing','Sing','#a42d2d',(SELECT id FROM users WHERE sur_name='Admin')),
	('PO','PO','PO','#d93636',(SELECT id FROM users WHERE sur_name='Admin')),
	('BA','BA','BA','#ba3636',(SELECT id FROM users WHERE sur_name='Admin')),
	('Art','Art','Art','#971c1c',(SELECT id FROM users WHERE sur_name='Admin')),
	('DevOps','DevOps','DevOps','#c91818',(SELECT id FROM users WHERE sur_name='Admin')),
	('Auto test','Auto test','Auto test','#dc3e09',(SELECT id FROM users WHERE sur_name='Admin'))
	ON CONFLICT (code) DO NOTHING;
	
	INSERT INTO user_types (name, code, created_by) VALUES
	('Staff','sta',(SELECT id FROM users WHERE sur_name='Admin')),
	('Intern','int',(SELECT id FROM users WHERE sur_name='Admin')),
	('Colaborator','col',(SELECT id FROM users WHERE sur_name='Admin')),
	('Probation','pro',(SELECT id FROM users WHERE sur_name='Admin')),
	('Vendor','ven',(SELECT id FROM users WHERE sur_name='Admin'))
	ON CONFLICT (code) DO NOTHING;
	
	INSERT INTO levels (name, display_name, code, color, created_by) VALUES
	('Intern 0', 'Intern_0', 'int_0', '#ff9800', (SELECT id FROM users WHERE sur_name='Admin')),
	('Intern 1', 'Intern_1', 'int_1', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Intern 2', 'Intern_2', 'int_2', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Intern 3', 'Intern_3', 'int_3', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Fresher Low', 'Fresher-', 'fre_low', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Fresher', 'Fresher', 'fre_mid', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Fresher High', 'Fresher+', 'fre_high', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Junior Low', 'Junior-', 'jun_low', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Junior', 'Junior', 'jun_mid', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Junior High', 'Junior+', 'jun_high', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Middle Low', 'Middle-', 'mid_low', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Middle', 'Middle', 'mid_mid', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Middle High', 'Middle+', 'mid_high', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Senior Low', 'Senior-', 'sen_low', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Senior', 'Senior', 'sen_mid', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin')),
	('Senior High', 'Senior+', 'sen_high', '#ff9800',(SELECT id FROM users WHERE sur_name='Admin'))
	ON CONFLICT (code) DO NOTHING;
	
INSERT INTO working_times (
    entity_type,
    entity_id,
    shift_name,
    start_time,
    end_time,
    created_by
) VALUES
('branch', (SELECT uuid FROM branches WHERE name='HN1'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='HN2'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='HN3'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='SG1'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='SG2'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='Vinh'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='ĐN'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin')),
('branch', (SELECT uuid FROM branches WHERE name='QN'), 'morning', '08:00', '12:00', (SELECT id FROM users WHERE sur_name = 'Admin'));

INSERT INTO working_times (
    entity_type,
    entity_id,
    shift_name,
    start_time,
    end_time,
    created_by
) VALUES
('branch', (SELECT uuid FROM branches WHERE name='HN1'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='HN2'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='HN3'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='SG1'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='SG2'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='Vinh'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='ĐN'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin')),
('branch', (SELECT uuid FROM branches WHERE name='QN'), 'afternoon', '13:00', '17:00', (SELECT id FROM users WHERE sur_name='Admin'));

COMMIT;

-- Permissions
START TRANSACTION;

INSERT INTO permissions (name, display_name, created_by) VALUES
('Admin','Admin',(SELECT id FROM users WHERE sur_name='Admin')),
('MyProfile','My profile',(SELECT id FROM users WHERE sur_name='Admin')),
('MyTimesheet','My timesheets',(SELECT id FROM users WHERE sur_name='Admin')),
('ManageTimesheet','Manage Timesheets',(SELECT id FROM users WHERE sur_name='Admin')),
('Project','Projets', (SELECT id FROM users WHERE sur_name='Admin'));

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('Admin.Users', 'Users', (SELECT id FROM permissions WHERE name='Admin'), (SELECT id FROM users WHERE sur_name='Admin'));

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('Admin.Users.View', 'View users', (SELECT id FROM permissions WHERE name='Admin.Users'), (SELECT id FROM users WHERE sur_name='Admin')),
('Admin.Users.AddNew', 'Add new user', (SELECT id FROM permissions WHERE name='Admin.Users'), (SELECT id FROM users WHERE sur_name='Admin')),
('Admin.Users.Edit', 'Edit user', (SELECT id FROM permissions WHERE name='Admin.Users'), (SELECT id FROM users WHERE sur_name='Admin')),
('Admin.Users.EditRole', 'Edit user role', (SELECT id FROM permissions WHERE name='Admin.Users'), (SELECT id FROM users WHERE sur_name='Admin')),
('Admin.Users.Delete', 'Delete user', (SELECT id FROM permissions WHERE name='Admin.Users'), (SELECT id FROM users WHERE sur_name='Admin')),
('Admin.Users.ResetPassword', 'Reset password', (SELECT id FROM permissions WHERE name='Admin.Users'), (SELECT id FROM users WHERE sur_name='Admin'));

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('Project.View', 'View my projects', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.ViewAll', 'View all projects', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.AddNew', 'Add new project', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.Edit', 'Edit project', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.Delete', 'Delete project', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.ChangeStatus', 'Change project status', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.ViewDetail', 'View project detail', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin')),
('Project.Export', 'Export project', (SELECT id FROM permissions WHERE name='Project'), (SELECT id FROM users WHERE sur_name='Admin'));

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('MyProfile.View', 'View profile', (SELECT id FROM permissions WHERE name='MyProfile'), (SELECT id FROM users WHERE sur_name='Admin')),
('MyProfile.RequestUpdateInfo', 'Request update info', (SELECT id FROM permissions WHERE name='MyProfile'), (SELECT id FROM users WHERE sur_name='Admin'));

INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('MyTimesheet.View', 'View my timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('MyTimesheet.AddNew', 'Add new timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('MyTimesheet.Edit', 'Edit timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('MyTimesheet.Delete', 'Delete timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('MyTimesheet.Submit', 'Submit timesheet', (SELECT id FROM permissions WHERE name='MyTimesheet'), (SELECT id FROM users WHERE sur_name='Admin'));


INSERT INTO permissions (name, display_name, parent_id, created_by) VALUES
('ManagementTimesheet.View', 'View all timesheets', (SELECT id FROM permissions WHERE name='ManageTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('ManagementTimesheet.Approve', 'Approve timesheet', (SELECT id FROM permissions WHERE name='ManageTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('ManagementTimesheet.Reject', 'Reject timesheet', (SELECT id FROM permissions WHERE name='ManageTimesheet'), (SELECT id FROM users WHERE sur_name='Admin')),
('ManagementTimesheet.Export', 'Export timesheet', (SELECT id FROM permissions WHERE name='ManageTimesheet'), (SELECT id FROM users WHERE sur_name='Admin'));

COMMIT;

CREATE UNIQUE INDEX IF NOT EXISTS ux_permissions_name ON permissions(name);

-- Roles
INSERT INTO roles(name,display_name, description, created_by) VALUES
('Admin','Admin','Administrator',(SELECT id FROM users WHERE sur_name='Admin')),
('All','All','Full Permission',(SELECT id FROM users WHERE sur_name='Admin')),
('Approve Review Intern','ApproveReviewInten','Approve Review Intern',(SELECT id FROM users WHERE sur_name='Admin')),
('Basic User','Basic User','Basic User',(SELECT id FROM users WHERE sur_name='Admin')),
('Branch Director','Branch Director','Branch Director',(SELECT id FROM users WHERE sur_name='Admin')),
('Create Review Intern','Create Review Intern','Create Review Intern',(SELECT id FROM users WHERE sur_name='Admin')),
('HDQT','HDQT','HDQT',(SELECT id FROM users WHERE sur_name='Admin')),
('HPM','Head PM','Head PM',(SELECT id FROM users WHERE sur_name='Admin')),
('PorjectManager','PM','Project Manager',(SELECT id FROM users WHERE sur_name='Admin')),
('PQA','PQA','PQA',(SELECT id FROM users WHERE sur_name='Admin')),
('View All Project','View All Project','View All Project',(SELECT id FROM users WHERE sur_name='Admin')),
('View User','View User','View User',(SELECT id FROM users WHERE sur_name='Admin')),
('Test Review Intern','Test Review Intern','Test Review Intern',(SELECT id FROM users WHERE sur_name='Admin'));

-- Role_Permission
-- Admin gets all permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    (SELECT id FROM users WHERE sur_name='Admin')
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'Admin';

-- All gets all permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    (SELECT id FROM users WHERE sur_name='Admin')
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'All';

-- Basic User Permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    (SELECT id FROM users WHERE sur_name='Admin')
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
    (SELECT id FROM users WHERE sur_name='Admin')
FROM roles r
JOIN permissions p ON p.name LIKE 'Project.%'
   OR p.name LIKE 'ManagementTimesheet.%'
WHERE r.name IN ('PorjectManager', 'HPM');

-- Branch Director/HDQT permissions
INSERT INTO role_permissions (role_id, permission_id, created_by)
SELECT
    r.id,
    p.id,
    (SELECT id FROM users WHERE sur_name='Admin')
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
    (SELECT id FROM users WHERE sur_name='Admin')
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
    (SELECT id FROM users WHERE sur_name='Admin')
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
    (SELECT id FROM users WHERE sur_name='Admin')
FROM roles r
JOIN permissions p ON p.name LIKE '%.View%'
WHERE r.name IN ('View All Project', 'View User');
