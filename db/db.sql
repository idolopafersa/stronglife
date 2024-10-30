-- MariaDB dump 10.19  Distrib 10.11.6-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: FitCalendar
-- ------------------------------------------------------
-- Server version	10.11.6-MariaDB-0+deb12u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `DayMeals`
--

DROP TABLE IF EXISTS `DayMeals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `DayMeals` (
  `day_id` int(11) NOT NULL,
  `meal_id` int(11) NOT NULL,
  PRIMARY KEY (`day_id`,`meal_id`),
  KEY `meal_id` (`meal_id`),
  CONSTRAINT `DayMeals_ibfk_1` FOREIGN KEY (`day_id`) REFERENCES `Days` (`id`) ON DELETE CASCADE,
  CONSTRAINT `DayMeals_ibfk_2` FOREIGN KEY (`meal_id`) REFERENCES `Meals` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DayMeals`
--

LOCK TABLES `DayMeals` WRITE;
/*!40000 ALTER TABLE `DayMeals` DISABLE KEYS */;
INSERT INTO `DayMeals` VALUES
(67,7),
(67,11),
(78,7),
(78,10),
(79,8),
(79,10),
(97,9),
(97,11);
/*!40000 ALTER TABLE `DayMeals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Days`
--

DROP TABLE IF EXISTS `Days`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Days` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `date` date NOT NULL,
  `routine_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `routine_id` (`routine_id`),
  CONSTRAINT `Days_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`),
  CONSTRAINT `Days_ibfk_3` FOREIGN KEY (`routine_id`) REFERENCES `Routines` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Days`
--

LOCK TABLES `Days` WRITE;
/*!40000 ALTER TABLE `Days` DISABLE KEYS */;
INSERT INTO `Days` VALUES
(61,20,'2024-10-06',NULL),
(62,20,'2024-10-03',NULL),
(63,20,'2024-10-02',NULL),
(64,25,'2024-10-06',10),
(65,25,'2024-10-01',NULL),
(66,21,'2024-10-06',NULL),
(67,25,'2025-04-30',NULL),
(68,25,'2024-10-08',NULL),
(69,25,'2024-10-15',NULL),
(70,25,'2024-10-05',NULL),
(71,25,'2024-10-11',NULL),
(72,25,'2024-10-13',NULL),
(73,25,'2024-10-20',NULL),
(74,25,'2025-02-02',NULL),
(75,25,'2024-09-08',NULL),
(76,25,'2024-09-01',NULL),
(77,27,'2024-10-06',10),
(78,28,'2024-10-06',13),
(79,20,'2024-10-07',10),
(80,25,'2024-10-07',NULL),
(81,25,'2024-10-14',NULL),
(82,25,'2024-10-27',NULL),
(83,25,'2024-10-29',NULL),
(84,25,'2024-10-30',NULL),
(85,25,'2024-10-23',NULL),
(86,25,'2024-10-16',NULL),
(87,25,'2024-10-09',NULL),
(88,25,'2024-10-02',NULL),
(89,25,'2024-10-26',NULL),
(90,25,'2024-10-28',NULL),
(91,25,'2024-09-30',NULL),
(92,21,'2024-10-07',NULL),
(93,21,'2024-10-20',NULL),
(94,21,'2024-10-13',NULL),
(95,21,'2024-10-21',NULL),
(96,21,'2024-10-22',NULL),
(97,30,'2024-10-07',NULL),
(98,21,'2024-10-08',NULL),
(99,30,'2024-10-08',NULL),
(100,30,'2024-10-15',NULL),
(101,30,'2024-10-06',NULL);
/*!40000 ALTER TABLE `Days` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Exercises`
--

DROP TABLE IF EXISTS `Exercises`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Exercises` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `sets` int(11) DEFAULT NULL,
  `repetitions` int(11) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Exercises`
--

LOCK TABLES `Exercises` WRITE;
/*!40000 ALTER TABLE `Exercises` DISABLE KEYS */;
INSERT INTO `Exercises` VALUES
(6,'Push Ups',3,12,'A basic upper body exercise.','https://example.com/images/pushups.jpg'),
(8,'Squat',4,12,'A fundamental lower body exercise.','http://example.com/squat.jpg'),
(10,'Bench Press',3,10,'A strength training exercise that targets the chest, shoulders, and triceps.','https://example.com/photos/bench_press.jpg'),
(11,'Deadlift',4,8,'A compound exercise that works the entire posterior chain.','https://example.com/photos/deadlift.jpg'),
(12,'Squat',3,12,'A foundational exercise that targets the quadriceps, hamstrings, and glutes.','https://example.com/photos/squat.jpg'),
(13,'Pull-Up',3,6,'An upper body exercise that targets the back and biceps.','https://example.com/photos/pull_up.jpg'),
(15,'Shoulder Press',4,10,'An exercise that targets the shoulders and triceps using a barbell or dumbbells.','https://example.com/photos/shoulder_press.jpg'),
(16,'Plank',3,NULL,'A core strengthening exercise that involves holding a position similar to a push-up.','https://example.com/photos/plank.jpg'),
(17,'Bicep Curl',3,12,'An isolation exercise that targets the biceps using dumbbells or a barbell.','https://example.com/photos/bicep_curl.jpg'),
(18,'Tricep Dips',4,10,'An exercise that targets the triceps, performed using parallel bars or a bench.','https://example.com/photos/tricep_dips.jpg'),
(19,'Leg Press',4,10,'A compound exercise that targets the quadriceps, hamstrings, and glutes using a leg press machine.','https://example.com/photos/leg_press.jpg'),
(21,'Lunges',3,10,'A lower body exercise that targets the legs and glutes.','https://example.com/photos/lunges.jpg');
/*!40000 ALTER TABLE `Exercises` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Meals`
--

DROP TABLE IF EXISTS `Meals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Meals` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `description` text DEFAULT NULL,
  `calories` int(11) DEFAULT NULL,
  `proteins` int(11) DEFAULT NULL,
  `fats` int(11) DEFAULT NULL,
  `carbs` int(11) DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Meals_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Meals`
--

LOCK TABLES `Meals` WRITE;
/*!40000 ALTER TABLE `Meals` DISABLE KEYS */;
INSERT INTO `Meals` VALUES
(7,'A healthy salad with grilled chicken, mixed greens, and vinaigrette.',400,30,15,20,'https://example.com/photos/grilled_chicken_salad.jpg','Grilled Chicken Salad'),
(8,'Tender beef strips stir-fried with vegetables and soy sauce.',471,354,20,30,'https://example.com/photos/beef_stir_fry.jpg','Beef Stir Fry'),
(9,'A nutritious bowl with quinoa, black beans, corn, and avocado.',400,15,12,60,'https://example.com/photos/vegetarian_quinoa_bowl.jpg','Vegetarian Quinoa Bowl'),
(10,'Pasta with fresh vegetables, olive oil, and parmesan cheese.',500,18,10,80,'https://example.com/photos/pasta_primavera.jpg','Pasta Primavera'),
(11,'Baked salmon fillet served with roasted asparagus.',600,40,25,10,'https://example.com/photos/salmon_asparagus.jpg','Salmon with Asparagus'),
(12,'Toasted bread topped with avocado and poached eggs.',300,12,20,25,'https://example.com/photos/egg_avocado_toast.jpg','Egg and Avocado Toast'),
(13,'Spicy chili made with ground turkey, beans, and tomatoes.',450,30,15,40,'https://example.com/photos/turkey_chili.jpg','Turkey Chili'),
(14,'A refreshing smoothie bowl topped with fresh fruits and granola.',350,5,8,65,'https://example.com/photos/mango_smoothie_bowl.jpg','Mango Smoothie Bowl'),
(16,'A rich and creamy protein shake with chocolate flavor.',250,30,5,20,'https://example.com/photos/chocolate_protein_shake.jpg','Chocolate Protein Shake'),
(20,'',0,0,0,0,'','');
/*!40000 ALTER TABLE `Meals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RoutineExercises`
--

DROP TABLE IF EXISTS `RoutineExercises`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RoutineExercises` (
  `routine_id` int(11) NOT NULL,
  `exercise_id` int(11) NOT NULL,
  PRIMARY KEY (`routine_id`,`exercise_id`),
  KEY `exercise_id` (`exercise_id`),
  CONSTRAINT `RoutineExercises_ibfk_1` FOREIGN KEY (`routine_id`) REFERENCES `Routines` (`id`),
  CONSTRAINT `RoutineExercises_ibfk_2` FOREIGN KEY (`exercise_id`) REFERENCES `Exercises` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RoutineExercises`
--

LOCK TABLES `RoutineExercises` WRITE;
/*!40000 ALTER TABLE `RoutineExercises` DISABLE KEYS */;
INSERT INTO `RoutineExercises` VALUES
(10,10),
(10,11),
(11,6);
/*!40000 ALTER TABLE `RoutineExercises` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Routines`
--

DROP TABLE IF EXISTS `Routines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Routines` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Routines`
--

LOCK TABLES `Routines` WRITE;
/*!40000 ALTER TABLE `Routines` DISABLE KEYS */;
INSERT INTO `Routines` VALUES
(10,'Full Body Workout','A balanced routine targeting all major muscle groups.','https://example.com/photos/full_body_workout.jpg'),
(11,'Upper Body Strength','Focuses on building strength in the chest, shoulders, and arms.','https://example.com/photos/upper_body_strength.jpg'),
(12,'Lower Body Blast','A high-intensity routine designed to strengthen legs and glutes.','https://example.com/photos/lower_body_blast.jpg'),
(13,'Core Conditioning','Targets the core muscles for stability and strength.','https://example.com/photos/core_conditioning.jpg');
/*!40000 ALTER TABLE `Routines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password_hash` binary(200) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Users`
--

LOCK TABLES `Users` WRITE;
/*!40000 ALTER TABLE `Users` DISABLE KEYS */;
INSERT INTO `Users` VALUES
(19,'pablo','$2a$10$J86alRXHBnxbWDyJoAjOT.xrxNuc8eVB9Gn/kx8o8JdxCG1nRMw2y\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','pafersa2003@gmail.com'),
(20,'pablo2','$2a$10$84wLB.yPaoVnF/Wonw1A.Oq62FHdTLl1mCAK71ux/PpHHfVCXpH2i\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','ddd'),
(21,'Salvador','$2a$10$dYWw5lbuw1m5IdpDPuQV4e21zvz0xw.mNrJ7YlkV6etwXxm9sny8W\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','salva_ov@hotmail.com'),
(22,'Jua nca','$2a$10$CVmDt4TbizAO9pizhp9NKuf2nYd8aJiXPT8kO83i2wgzCblT3Grx6\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','juanca@gmail.com'),
(23,'Pabloeldemon','$2a$10$azk3glsm8N28N.b0iedJ5.n2uhFoNjzMSCspMScSOmToI8qJzsuzW\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','lastetasdetriviño@gmail.com'),
(24,'Lastetasdetriviño','$2a$10$Etl3A2fyyx0P.d1/UNvKQ.PUpE39Ns4GpplXi4fXscl6mB4pZFajW\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','pablopafersa'),
(25,'123','$2a$10$uxIzvGXp6Tq008At6Vv3uOOywtfI6tquqOAy09crYs1K.j3jN5.QS\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','123'),
(26,'Pacocraisy','$2a$10$UAyEO7PF/1nSM8L6IVNBvuieAlFMXLXuv1HSwGP4QTK9aRTtn0Mty\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','pacopacman'),
(27,'Mifersa','$2a$10$ngRXxS66ELQkH2jqbjHKZO4db/VxlluX6EUrZ05vmBW8VN.wP6gN.\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mifersa2007@gmail.com'),
(28,'Pacoputo','$2a$10$8xbnuXzd3/L0oOKaVsRIS.f5aMSrRCkicRpSHoYUrAKJwIoIDWTk2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','pacopelo2003@gmail.com'),
(29,'Jdjdj','$2a$10$xZ4SLjvKQIpGLuZTpEwGs.uGR5LtCXcoKLlvwz0cU1PVpN2MPhZ7y\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','jjzjd@gmail.som'),
(30,'jose','$2a$10$95OrfEs4T7FfyRE6H4u66ueGE/YROGeS0AHyBLwxekt03Anz/IY6S\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','jose'),
(31,'jose2','$2a$10$NZbAgO.qQAgrxnMvcf3HFuGpVWOeLOR2Z1QXHhP7xFqCobsybZJP.\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','jose2@jose.com'),
(32,'Chuuuchooo','$2a$10$EQdlkHJNU3lHM89Da3csBeoFi/quA7gWC9ik8cpUFK0XBJySsrbcy\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','kali.n.moylan@gmail.com'),
(33,'Chucho','$2a$10$cOClrrPAy4ATpd8YuwnYS.6FG/CFY1NE5Dv314KAmkLh5wVjfTE1y\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','kali.n.moylan@gmail.com');
/*!40000 ALTER TABLE `Users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-19 12:14:09
