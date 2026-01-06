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
	
	COMMIT;
