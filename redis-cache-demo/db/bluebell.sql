CREATE DATABASE  IF NOT EXISTS `bluebell` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `bluebell`;
-- MySQL dump 10.13  Distrib 8.0.31, for macos12 (x86_64)
--
-- Host: localhost    Database: bluebell
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comment_id` bigint unsigned NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `post_id` bigint NOT NULL,
  `author_id` bigint NOT NULL,
  `parent_id` bigint NOT NULL DEFAULT '0',
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_comment_id` (`comment_id`),
  KEY `idx_author_Id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `community`
--

DROP TABLE IF EXISTS `community`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `community` (
  `id` int NOT NULL AUTO_INCREMENT,
  `community_id` int unsigned NOT NULL,
  `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `community`
--

LOCK TABLES `community` WRITE;
/*!40000 ALTER TABLE `community` DISABLE KEYS */;
INSERT INTO `community` VALUES (1,1,'Go Topic','Golang','2016-11-01 00:10:10','2023-05-05 00:28:54'),(2,2,'Leetcode Topic','？？？？','2020-01-01 00:00:00','2023-05-05 00:28:54'),(3,3,'C++ Topic','C ++ ','2023-05-05 00:31:41','2023-05-05 00:31:41');
/*!40000 ALTER TABLE `community` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS `post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL COMMENT '帖子id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint NOT NULL COMMENT '作者的用户id',
  `community_id` bigint NOT NULL COMMENT '所属社区',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES `post` WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
INSERT INTO `post` VALUES (1,63542771746603009,'holy shit ','holy shit !',63542625180844033,2,1,'2023-04-23 08:40:43','2023-04-23 08:40:43'),(2,63543686792740865,'123','321',63542625180844033,1,1,'2023-04-23 08:49:48','2023-04-23 08:49:48'),(3,63547823483781121,'1','2',63542625180844033,5,1,'2023-04-23 09:30:54','2023-05-04 07:44:57'),(4,63940350946836481,'gogogo','shit shit shit',63542625180844033,2,1,'2023-04-26 02:30:19','2023-04-26 02:30:19'),(5,65132075690229761,'5.4','5.4',63542625180844033,1,1,'2023-05-04 07:49:02','2023-05-04 07:49:02'),(6,65132733440983041,'66','66',63542625180844033,1,1,'2023-05-04 07:55:34','2023-05-04 07:55:34'),(7,65132756559986689,'66','66',63542625180844033,1,1,'2023-05-04 07:55:48','2023-05-04 07:55:48'),(8,65132782313013249,'66','666',63542696282685441,1,1,'2023-05-04 07:56:03','2023-05-04 07:56:03'),(9,65132859806973953,'44','44',63542625180844033,1,1,'2023-05-04 07:56:49','2023-05-04 07:56:49'),(10,65132899669639169,'11','11',63542625180844033,1,1,'2023-05-04 07:57:13','2023-05-04 07:57:13'),(11,65153146984333313,'test-test','test-test',65094992003072001,1,1,'2023-05-04 11:18:21','2023-05-04 11:18:21'),(12,65231969750876161,'test2023-test2023','test2023-test2023',65094992003072001,2,1,'2023-05-05 00:21:23','2023-05-05 00:21:23'),(13,65232007466057729,'test2023-test2023','test2023-test2023',65094992003072001,2,1,'2023-05-05 00:21:46','2023-05-05 00:21:46'),(14,65232924257026049,'2023-5-5','2023-5-5',63542625180844033,2,1,'2023-05-05 00:30:52','2023-05-05 00:30:52');
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `gender` tinyint NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,63542625180844033,'feng','31323334f5d77a10ae47e3738837865e6a831793','feng@126.com',0,'2023-04-23 08:39:16','2023-06-05 01:01:25'),(2,63542696282685441,'dev','31323334f5d77a10ae47e3738837865e6a831793','dev@126.com',0,'2023-04-23 08:39:58','2023-06-05 01:01:25'),(3,63948201509519361,'liang','31323334f5d77a10ae47e3738837865e6a831793','liang@126.com',0,'2023-04-26 03:48:18','2023-06-05 01:01:25'),(4,64078576667852801,'holy','31323334f5d77a10ae47e3738837865e6a831793','holy',1,'2023-04-27 01:23:28','2023-04-27 01:23:28'),(5,65094992003072001,'feng121','3132333435368cab425cf81f49fef958042bfb16029b','123@123.com',1,'2023-05-04 01:40:38','2023-05-04 01:40:38'),(6,65095113201680385,'feng1211','3132333435368cab425cf81f49fef958042bfb16029b','123@123.com',1,'2023-05-04 01:41:50','2023-05-04 01:41:50'),(7,65097239747362817,'feng12112','3132333435368cab425cf81f49fef958042bfb16029b','555@11.com',1,'2023-05-04 02:02:58','2023-06-06 11:48:26');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-30 17:55:45
