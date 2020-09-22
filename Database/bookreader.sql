DROP TABLE IF EXISTS `admin`;

CREATE TABLE `admins` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`email` varchar(255) NOT NULL UNIQUE,
	`password` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `email_unique` (`email`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `schools`;

CREATE TABLE `schools` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school` varchar(255) NOT NULL UNIQUE,
	`school_icon` varchar(255) NULL DEFAULT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `schools_unique` (`school`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `faculties`;

CREATE TABLE `faculties` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school_id` bigint(20) unsigned NOT NULL,
	`faculty` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `faculties_school_id_foreign` (`school_id`),
	CONSTRAINT `faculties_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `courses`;

CREATE TABLE `courses` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school_id` bigint(20) unsigned NOT NULL,
	`faculty_id` bigint(20) unsigned NOT NULL,
	`course` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `courses_school_id_foreign` (`school_id`),
	KEY `courses_faculty_id_foreign` (`faculty_id`),
	CONSTRAINT `courses_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE,
	CONSTRAINT `courses_faculty_id_foreign` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school_id` bigint(20) unsigned NOT NULL,
	`faculty_id` bigint(20) unsigned NOT NULL,
	`course_id` bigint(20) unsigned NOT NULL,
	`student_nr` varchar(255) NOT NULL UNIQUE,
	`name` varchar(255) NOT NULL,
	`surname` varchar(255) NOT NULL,
	`email` varchar(255) NOT NULL UNIQUE,
	`picture` varchar(255) NULL DEFAULT NULL,
	`password` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `users_student_nr_unique` (`student_nr`),
	UNIQUE KEY `users_email_unique` (`email`),
	KEY `users_school_id_foreign` (`school_id`),
	KEY `users_faculty_id_foreign` (`faculty_id`),
	KEY `users_course_id_foreign` (`course_id`),
	CONSTRAINT `users_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE,
	CONSTRAINT `users_faculty_id_foreign` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE,
	CONSTRAINT `users_course_id_foreign` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `modules`;

CREATE TABLE `modules` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`school_id` bigint(20) unsigned NOT NULL,
	`faculty_id` bigint(20) unsigned NOT NULL,
	`course_id` bigint(20) unsigned NOT NULL,
	`module` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `modules_school_id_foreign` (`school_id`),
	KEY `modules_faculty_id_foreign` (`faculty_id`),
	KEY `modules_course_id_foreign` (`course_id`),
	CONSTRAINT `modules_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE,
	CONSTRAINT `modules_faculty_id_foreign` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE,
	CONSTRAINT `modules_course_id_foreign` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE					
) ENGINE=InnoDB;

/*
	Modules that the user is enrolled in
*/
DROP TABLE IF EXISTS `enrolled`;
 
CREATE TABLE `enrolled`(
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`module_id` bigint(20) unsigned NOT NULL,
	`user_id` bigint(20) unsigned NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `enrolled_module_id_foreign` (`module_id`),
	KEY `enrolled_user_id_foreign` (`user_id`),
	CONSTRAINT `enrolled_module_id_foreign` FOREIGN KEY (`module_id`) REFERENCES `modules` (`id`) ON DELETE CASCADE,
	CONSTRAINT `enrolled_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE		
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`title` varchar(255) NOT NULL,
	`author` varchar(255) NOT NULL,
	`publish_date` year NULL DEFAULT NULL,
	`isbn` varchar(20) NOT NULL UNIQUE,
	`cover_page` varchar(255) NOT NULL UNIQUE,
	`description` text,
	`book` varchar(255) NOT NULL UNIQUE,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `books_isbn_unique` (`isbn`),
	UNIQUE KEY `books_cover_page_unique` (`cover_page`),
	UNIQUE KEY `books_book_unique` (`book`)	 	
) ENGINE=InnoDB;

/*
	Books that are required for the course
	that the student has enrolled in
*/
DROP TABLE IF EXISTS `required`;

CREATE TABLE `required`(
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`book_id` bigint(20) unsigned NOT NULL,
	`module_id` bigint(20) unsigned NOT NULL,
	`user_id` bigint(20) unsigned NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `required_book_id_foreign` (`book_id`),
	KEY `required_module_id_foreign` (`module_id`),
	KEY `required_user_id_foreign` (`user_id`),
	CONSTRAINT `required_book_id_foreign` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE,
	CONSTRAINT `required_module_id_foreign` FOREIGN KEY (`module_id`) REFERENCES `modules` (`id`) ON DELETE CASCADE,
	CONSTRAINT `required_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB;


DROP TABLE IF EXISTS `login_tokens`;

CREATE TABLE `login_tokens` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint(20) unsigned NOT NULL,
	`token` varchar(255) NOT NULL UNIQUE,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `login_tokens_token_unique` (`token`),
	KEY `login_tokens_user_id_foreign` (`user_id`),
	CONSTRAINT `login_tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE			
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `admin_login_tokens`;

CREATE TABLE `admin_login_tokens` (
	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint(20) unsigned NOT NULL,
	`token` varchar(255) NOT NULL UNIQUE,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `admin_login_tokens_token_unique` (`token`),
	KEY `admin_login_tokens_user_id_foreign` (`user_id`),
	CONSTRAINT `admin_login_tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `admins` (`id`) ON DELETE CASCADE			
) ENGINE=InnoDB;