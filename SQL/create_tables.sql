CREATE TABLE `user` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` varchar(40) NOT NULL,
	`Surname` varchar(40) NOT NULL,
	`available_funds` INT NOT NULL DEFAULT '0',
	`lang` varchar(2) NOT NULL DEFAULT 'en',
	`theme` varchar(5) NOT NULL DEFAULT 'light',
	PRIMARY KEY (`id`)
);

CREATE TABLE `transactions` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`user` INT NOT NULL,
	`date` DATE NOT NULL,
	`desc` varchar(255),
	PRIMARY KEY (`id`)
);

CREATE TABLE `m_boxes` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`user` INT NOT NULL,
	`target_sum` INT NOT NULL DEFAULT '0',
	`current_sum` INT NOT NULL DEFAULT '0',
	`auto_adj` BOOLEAN NOT NULL DEFAULT '0',
	`adj_amount` INT NOT NULL DEFAULT '0',
	`desc` varchar(255),
	`name` varchar(40) NOT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `transactions` ADD CONSTRAINT `transactions_fk0` FOREIGN KEY (`user`) REFERENCES `user`(`id`);

ALTER TABLE `m_boxes` ADD CONSTRAINT `m_boxes_fk0` FOREIGN KEY (`user`) REFERENCES `user`(`id`);



