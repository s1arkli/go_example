package sql

var (
	Insert = "INSERT INTO `users` (`name`,`age`,`created_at`,`updated_at`) VALUES (?, ?, NOW(), NOW());"
	Select = "SELECT * FROM `users` WHERE `id` = ?;"
	Delete = "DELETE FROM `users` WHERE `id` = ?;"
	Update = "UPDATE `users` SET `name` = ?,`age` = ?,`updated_at`=NOW() WHERE `id` = ?;"
)
