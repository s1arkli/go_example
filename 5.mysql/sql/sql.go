package sql

var (
	Insert = "INSERT INTO `users` (`name`,`age`,`created_at`,`updated_at`) VALUES (?, ?, NOW(), NOW());"
	Select = "SELECT * FROM `users` WHERE `id` = ?;"
	Delete = "DELETE FROM `users` WHERE `id` = ?;"
	Update = "UPDATE `users` SET `name` = ?,`age` = ?,`updated_at`=NOW() WHERE `id` = ?;"
	Create = "CREATE TABLE `users` ( `id` bigint NOT NULL AUTO_INCREMENT,`name` varchar(32) NOT NULL,`age` int NOT NULL,`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,PRIMARY KEY (`id`))ENGINE=InnoDB"
)
