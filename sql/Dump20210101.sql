CREATE DATABASE  IF NOT EXISTS `account_book` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `account_book`;
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: account_book
-- ------------------------------------------------------
-- Server version	8.0.22

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
-- Table structure for table `account_book`
--

DROP TABLE IF EXISTS `account_book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_book` (
  `book_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `cover` varchar(100) NOT NULL,
  `uid` int NOT NULL,
  `target_id` int NOT NULL DEFAULT '0',
  `cycle_type` tinyint NOT NULL DEFAULT '1' COMMENT '1月, 2周，3年',
  `cycle_start` varchar(45) NOT NULL DEFAULT '1' COMMENT '周期开始时间',
  `cycle_end` varchar(45) NOT NULL DEFAULT '31' COMMENT '周期结束时间',
  `create_time` int NOT NULL,
  `update_time` int NOT NULL DEFAULT '0',
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_book`
--

LOCK TABLES `account_book` WRITE;
/*!40000 ALTER TABLE `account_book` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_flow_2021`
--

DROP TABLE IF EXISTS `account_flow_2021`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_flow_2021` (
  `flow_id` bigint NOT NULL,
  `cost` int NOT NULL,
  `uid` int NOT NULL,
  `book_id` int NOT NULL,
  `member_id` int NOT NULL DEFAULT '0' COMMENT '用于哪个成员',
  `tip` varchar(100) NOT NULL,
  `review_star` tinyint NOT NULL DEFAULT '0' COMMENT '星级，用于复盘时对这项支出评价，便于下次决定',
  `review_time` int NOT NULL DEFAULT '0' COMMENT '复盘时间',
  `create_time` int NOT NULL,
  PRIMARY KEY (`flow_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_flow_2021`
--

LOCK TABLES `account_flow_2021` WRITE;
/*!40000 ALTER TABLE `account_flow_2021` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_flow_2021` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_item`
--

DROP TABLE IF EXISTS `account_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_item` (
  `item_id` int NOT NULL,
  `uid` int NOT NULL DEFAULT '0',
  `item_name` varchar(45) NOT NULL,
  `create_time` int DEFAULT NULL,
  `flow` tinyint NOT NULL DEFAULT '2' COMMENT '1收入2支出',
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='收入与支出项';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_item`
--

LOCK TABLES `account_item` WRITE;
/*!40000 ALTER TABLE `account_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_member`
--

DROP TABLE IF EXISTS `account_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_member` (
  `member_id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `nickname` varchar(45) NOT NULL,
  `create_time` int NOT NULL,
  `relation` varchar(45) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_member`
--

LOCK TABLES `account_member` WRITE;
/*!40000 ALTER TABLE `account_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_target`
--

DROP TABLE IF EXISTS `account_target`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_target` (
  `target_id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `content` varchar(100) NOT NULL,
  `status` varchar(45) NOT NULL DEFAULT '0',
  `pic` varchar(100) NOT NULL,
  `expert_start_time` int NOT NULL,
  `expert_finish_time` int NOT NULL,
  `review_time` int NOT NULL DEFAULT '0' COMMENT '复盘时间',
  `review_result` varchar(100) NOT NULL COMMENT '结果总结',
  `review_star` tinyint NOT NULL DEFAULT '0' COMMENT '目标完成星级',
  `create_time` int NOT NULL DEFAULT '0',
  `update_time` int NOT NULL,
  PRIMARY KEY (`target_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='目标';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_target`
--

LOCK TABLES `account_target` WRITE;
/*!40000 ALTER TABLE `account_target` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_target` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_user`
--

DROP TABLE IF EXISTS `account_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_user` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `nickname` varchar(45) NOT NULL,
  `password` varchar(100) NOT NULL,
  `gender` tinyint NOT NULL DEFAULT '0' COMMENT '0未填写，1女，2男',
  `birthday` varchar(15) NOT NULL DEFAULT '0',
  `work` varchar(45) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  `create_time` int NOT NULL,
  `update_time` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_user`
--

LOCK TABLES `account_user` WRITE;
/*!40000 ALTER TABLE `account_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `account_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-01 19:02:37
