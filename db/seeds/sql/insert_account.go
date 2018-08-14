package sql

// InsertAccounts insert account sql
var InsertAccounts = `insert into accounts
(id,name,email,password,created_by,updated_by,created_at,updated_at,deleted_at)
values
('139f2b41-c1b7-42a8-962f-3ad306ea81a5','テストユーザー','test@example.com','$2a$10$4OPY6NlJf9uUoruhZBGVVeMyIvXpXrRByBdcktKw7CnKOT40Clv/S',null,null,now(),now(),null);
`
