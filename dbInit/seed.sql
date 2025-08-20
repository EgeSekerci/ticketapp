INSERT INTO users (email, password, role, name) VALUES 
('usertest@example.com', '$argon2id$v=19$m=65536,t=1,p=20$t4GHNZDZAW1/9JAa/otkag$HxL4lMzHQHofw08oHwo0oTv3F1Kht4MhswaEFgEJJsg', 'user', 'Test User'),
('admintest@example.com', '$argon2id$v=19$m=65536,t=1,p=20$sF6BA5AJh1xwwt8eKLBZTQ$klE7uHGjN1xollKOb6DfZF/dmgyT0iUt2G3EPpnHHI8', 'admin', 'Test Admin');
