-- MySQL dump 10.13  Distrib 5.7.30, for Linux (x86_64)
--
-- Host: localhost    Database: bookreader
-- ------------------------------------------------------
-- Server version	5.7.30-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin_login_tokens`
--

DROP TABLE IF EXISTS `admin_login_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_login_tokens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`),
  UNIQUE KEY `admin_login_tokens_token_unique` (`token`),
  KEY `admin_login_tokens_user_id_foreign` (`user_id`),
  CONSTRAINT `admin_login_tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `admins` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_login_tokens`
--

LOCK TABLES `admin_login_tokens` WRITE;
/*!40000 ALTER TABLE `admin_login_tokens` DISABLE KEYS */;
INSERT INTO `admin_login_tokens` VALUES (1,4,'a95ea0064f2c34addf9093831786d6bb5c90930a','2020-10-16 19:48:43','2020-10-16 19:48:43');
/*!40000 ALTER TABLE `admin_login_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admins`
--

DROP TABLE IF EXISTS `admins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admins` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admins`
--

LOCK TABLES `admins` WRITE;
/*!40000 ALTER TABLE `admins` DISABLE KEYS */;
INSERT INTO `admins` VALUES (4,'lsg.seritili@gmail.com','$2a$04$79OB34nfOjhazj2A4tGG/ejm28IKeiI.JNhG0Tv6jV3HuUVjZqFCW','2020-09-16 09:37:02','2020-09-16 09:37:02');
/*!40000 ALTER TABLE `admins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `publish_date` year(4) DEFAULT NULL,
  `isbn` varchar(255) NOT NULL,
  `cover_page` varchar(255) NOT NULL,
  `description` text,
  `book` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `isbn` (`isbn`),
  UNIQUE KEY `books_isbn_unique` (`isbn`),
  UNIQUE KEY `book` (`book`),
  UNIQUE KEY `books_book_unique` (`book`),
  UNIQUE KEY `cover_page` (`cover_page`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (2,'Learning MySQL','Seyed M.M. “Saied” Tahaghoghi and Hugh E. Williams',2006,'9780596008642','/data/images/book_covers/cover_page_KRI0lCbYYxvBlfCS.png','This book is primarily for people who don’t know much about deploying and using an actual database-management system, or about developing applications that use a da- tabase. We provide a readable introduction to relational databases, the MySQL data- base management system, the Structured Query Language (SQL), and the PHP and Perl programming languages. We also cover some quite advanced material that will be of interest even to experienced database users. Readers with some exposure to these topics should be able to use this book to expand their repertoire and deepen their under- standing of MySQL in particular, and database techniques in general.','/data/books/book_36quylV9bpsh5qhF.pdf','2020-09-21 13:23:25','2020-09-21 13:39:37'),(3,'Natural Language Processing in Action','Hobson Lane , Cole Howard , Hannes Max Hapke',2019,'9781617294631','/data/images/book_covers/cover_page_U26WEakcjbbii7Np.png','Natural Language Processing in Action is a practical guide to processing and generating natural language text in the real world. In this book we provide you with all the tools and techniques you need to build the backend NLP systems to support a virtual assistant (chatbot), spam filter, forum moderator, sentiment analyzer, knowledge base builder, natural language text miner, or nearly any other NLP application you can imagine. Natural Language Processing in Action is aimed at intermediate to advanced Python developers. Readers already capable of designing and building complex systems will also find most of this book useful, since it provides numerous best-practice examples and insight into the capabilities of state-of-the art NLP algorithms. While knowledge of object-oriented Python development may help you build better systems, it’s not required to use what you learn in this book. For special topics, we provide sufficient background material and cite resources (both text and online) for those who want to gain an in-depth understanding.','/data/books/book_AlePU1uEJv9cZgR4.pdf','2020-09-21 13:34:40','2020-09-21 13:34:40'),(4,'THE NEW WORLD ORDER','A. Ralph Epperson',1990,'9780961413514','/data/images/book_covers/cover_page_Dxb4U0mt6tSDMrb2.png','','/data/books/book_J6vj79v5Z8G7ZYj5.pdf','2020-10-08 21:04:58','2020-10-08 21:04:58'),(5,'The Four: The Hidden DNA of Amazon, Apple, Facebook, and Google','Scott Galloway',2017,'9780525501220','/data/images/book_covers/cover_page_tFU8QTfKQjOqOS64.png','Amazon, Apple, Facebook, and Google are the four most influential companies on the planet. Just about everyone thinks they know how they got there. Just about everyone is wrong.','/data/books/book_iXKVlgFwdAzB8mxc.pdf','2020-10-08 21:14:47','2020-10-08 21:14:47'),(6,'SVG Essentials: Producing Scalable Vector Graphics with XML 2nd Edition','J. David Eisenberg',2014,'9781449374358','/data/images/book_covers/cover_page_xzvx9dlTGdfWFBXe.png','','/data/books/book_Yr1EuQ0kdJ1mBru4.pdf','2020-10-08 21:21:28','2020-10-08 21:21:28');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
  CONSTRAINT `courses_faculty_id_foreign` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE,
  CONSTRAINT `courses_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,2,3,'Bcom Economics','2020-09-20 06:39:56','2020-09-20 06:39:56'),(2,2,3,'Bcom Law','2020-09-20 06:40:44','2020-09-20 06:40:44'),(4,2,3,'Bcom Auditing','2020-09-20 06:41:18','2020-09-20 06:41:18'),(5,2,3,'Bcom Accounting','2020-09-20 06:43:33','2020-09-20 06:43:33'),(6,2,3,'Bcom Information Systems','2020-10-07 06:41:48','2020-10-07 06:41:48'),(7,2,6,'Llb Criminal Law','2020-10-07 06:43:39','2020-10-07 06:43:39'),(8,2,7,'Mechanical Engineering','2020-10-07 06:43:58','2020-10-07 06:43:58'),(9,2,7,'Industrial Engineering','2020-10-07 06:44:09','2020-10-07 06:44:09'),(10,5,14,'Bsc Mathematics','2020-10-07 06:45:17','2020-10-07 06:45:17');
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `enrolled`
--

DROP TABLE IF EXISTS `enrolled`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `enrolled` (
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `enrolled`
--

LOCK TABLES `enrolled` WRITE;
/*!40000 ALTER TABLE `enrolled` DISABLE KEYS */;
INSERT INTO `enrolled` VALUES (2,1,1,'2020-10-15 21:41:38','2020-10-15 21:41:38'),(3,2,1,'2020-10-15 21:55:31','2020-10-15 21:55:31'),(4,3,2,'2020-10-15 21:58:37','2020-10-15 21:58:37'),(5,1,2,'2020-10-15 21:58:44','2020-10-15 21:58:44'),(6,2,2,'2020-10-15 21:58:49','2020-10-15 21:58:49'),(7,5,2,'2020-10-15 21:58:56','2020-10-15 21:58:56');
/*!40000 ALTER TABLE `enrolled` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `faculties`
--

DROP TABLE IF EXISTS `faculties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `faculties` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint(20) unsigned NOT NULL,
  `faculty` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `faculties_school_id_foreign` (`school_id`),
  CONSTRAINT `faculties_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `faculties`
--

LOCK TABLES `faculties` WRITE;
/*!40000 ALTER TABLE `faculties` DISABLE KEYS */;
INSERT INTO `faculties` VALUES (3,2,'Commerce','2020-09-20 04:11:52','2020-09-20 04:11:52'),(6,2,'Law','2020-09-20 04:15:22','2020-09-20 04:15:22'),(7,2,'Engineering','2020-09-20 04:16:45','2020-09-20 04:16:45'),(8,2,'Arts','2020-09-20 04:38:59','2020-09-20 04:38:59'),(9,2,'Science','2020-10-06 17:21:53','2020-10-06 17:21:53'),(10,2,'Health Sciences','2020-10-06 17:27:38','2020-10-06 17:30:29'),(11,3,'Education','2020-10-06 17:37:25','2020-10-06 17:37:25'),(12,3,'Engineering','2020-10-06 17:38:10','2020-10-06 17:38:10'),(13,3,'Commerce','2020-10-06 17:38:19','2020-10-06 17:38:19'),(14,5,'Science','2020-10-07 06:44:40','2020-10-07 06:44:40');
/*!40000 ALTER TABLE `faculties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `login_tokens`
--

DROP TABLE IF EXISTS `login_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `login_tokens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`),
  UNIQUE KEY `login_tokens_token_unique` (`token`),
  KEY `login_tokens_user_id_foreign` (`user_id`),
  CONSTRAINT `login_tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `login_tokens`
--

LOCK TABLES `login_tokens` WRITE;
/*!40000 ALTER TABLE `login_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `login_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `modules`
--

DROP TABLE IF EXISTS `modules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
  CONSTRAINT `modules_course_id_foreign` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `modules_faculty_id_foreign` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE,
  CONSTRAINT `modules_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `modules`
--

LOCK TABLES `modules` WRITE;
/*!40000 ALTER TABLE `modules` DISABLE KEYS */;
INSERT INTO `modules` VALUES (1,2,3,1,'Economics 1000','2020-09-20 07:51:40','2020-09-20 07:51:40'),(2,2,3,1,'Economics 1006','2020-09-20 07:52:21','2020-09-20 07:52:21'),(3,2,3,1,'Economics 2000','2020-09-20 07:52:49','2020-09-20 07:52:49'),(5,2,6,7,'Criminal Law 1000','2020-10-08 10:26:20','2020-10-08 10:26:20'),(6,2,6,7,'Criminal Law 2000','2020-10-08 10:27:22','2020-10-08 10:27:22');
/*!40000 ALTER TABLE `modules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recommended`
--

DROP TABLE IF EXISTS `recommended`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recommended` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `book_id` bigint(20) unsigned NOT NULL,
  `module_id` bigint(20) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `recommended_book_id_foreign` (`book_id`),
  KEY `recommended_module_id_foreign` (`module_id`),
  CONSTRAINT `recommended_book_id_foreign` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE,
  CONSTRAINT `recommended_module_id_foreign` FOREIGN KEY (`module_id`) REFERENCES `modules` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recommended`
--

LOCK TABLES `recommended` WRITE;
/*!40000 ALTER TABLE `recommended` DISABLE KEYS */;
INSERT INTO `recommended` VALUES (1,5,1,'2020-10-16 22:46:42','2020-10-16 22:46:42'),(2,5,2,'2020-10-16 22:52:22','2020-10-16 22:52:22'),(3,4,5,'2020-10-16 22:53:15','2020-10-16 22:53:15');
/*!40000 ALTER TABLE `recommended` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `required`
--

DROP TABLE IF EXISTS `required`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `required` (
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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `required`
--

LOCK TABLES `required` WRITE;
/*!40000 ALTER TABLE `required` DISABLE KEYS */;
/*!40000 ALTER TABLE `required` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schools`
--

DROP TABLE IF EXISTS `schools`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schools` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `school` varchar(255) NOT NULL,
  `school_icon` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `school` (`school`),
  UNIQUE KEY `schools_unique` (`school`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schools`
--

LOCK TABLES `schools` WRITE;
/*!40000 ALTER TABLE `schools` DISABLE KEYS */;
INSERT INTO `schools` VALUES (2,'UNIVERSITY OF CAPE TOWN','/data/images/institutions/school_0z6k2KHvLmY5damr.png','2020-09-18 14:08:19','2020-10-05 12:36:55'),(3,'UNIVERSITY OF PRETORIA','/data/images/institutions/school_FWTIW1vcWCNAmZ4q.png','2020-10-05 12:29:06','2020-10-05 12:29:06'),(4,'UNIVERSITY OF THE FREE STATE','/data/images/institutions/school_nfU4R2Oxbd86sgdV.png','2020-10-05 12:30:09','2020-10-05 12:30:09'),(5,'STANFORD UNIVERSITY','/data/images/institutions/school_uosBXqBORzPn6GvU.png','2020-10-05 12:30:47','2020-10-05 12:30:47'),(6,'CAMBRIDGE UNIVERSITY','/data/images/institutions/school_S1DgdKAmbKO9KNJR.png','2020-10-05 12:31:24','2020-10-05 12:31:24'),(7,'UNIVERSITY OF THE WITWATERSRAND','/data/images/institutions/school_VBpgVlRCjIYdDwNB.jpg','2020-10-06 10:22:14','2020-10-06 10:22:14');
/*!40000 ALTER TABLE `schools` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `school_id` bigint(20) unsigned NOT NULL,
  `faculty_id` bigint(20) unsigned NOT NULL,
  `course_id` bigint(20) unsigned NOT NULL,
  `student_nr` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `surname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `picture` varchar(255) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `student_nr` (`student_nr`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `users_student_nr_unique` (`student_nr`),
  UNIQUE KEY `users_email_unique` (`email`),
  UNIQUE KEY `student_nr_2` (`student_nr`),
  KEY `users_school_id_foreign` (`school_id`),
  KEY `users_faculty_id_foreign` (`faculty_id`),
  KEY `users_course_id_foreign` (`course_id`),
  CONSTRAINT `users_course_id_foreign` FOREIGN KEY (`course_id`) REFERENCES `courses` (`id`) ON DELETE CASCADE,
  CONSTRAINT `users_faculty_id_foreign` FOREIGN KEY (`faculty_id`) REFERENCES `faculties` (`id`) ON DELETE CASCADE,
  CONSTRAINT `users_school_id_foreign` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,2,3,1,'7172748','Peter','Gregory','7172748@uct.ac.za','/data/users/user_default.jpg','$2a$04$.FFJC6t/m61hak5fGIcYkO40jbW0T9xc6HypjshBKtBL8iCJS1ZPS','2020-09-23 11:11:10','2020-09-23 11:11:10'),(2,2,3,6,'8484848','Lesego','Seritili','8484848@uct.ac.za','/data/users/user_default.jpg','$2a$04$OzJXFeOda40TUaKyEfyoj.j3hGVteLjFJX.H6MP904.Zzkf9ul2Fu','2020-10-15 09:28:23','2020-10-15 09:28:23'),(3,2,7,8,'123456','John','Doe','123456@uct.ac.za','/data/users/user_default.jpg','$2a$04$lvoN1kXFov1xOv0H4HQL0eFYaHDgOjAujlI69cj6d/AZYARq2fLIK','2020-10-15 10:17:19','2020-10-15 10:17:19'),(4,2,3,4,'7384849','Lerato','Maseko','7384849@uct.ac.za','/data/users/user_default.jpg','$2a$04$EEwx84GsvqyFS/XcZuChpOB2YMoypaoM5riQBjwBgDS1N6ZBC5rSq','2020-10-15 10:19:00','2020-10-15 10:19:00');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-05 12:31:22
