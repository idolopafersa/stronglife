/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-11.5.2-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: FitCalendar
-- ------------------------------------------------------
-- Server version	11.5.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

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
(60,8),
(33,9),
(38,9),
(45,9),
(50,10),
(60,10),
(33,11),
(57,11),
(38,12),
(51,12),
(48,13),
(54,16);
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
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Days`
--

LOCK TABLES `Days` WRITE;
/*!40000 ALTER TABLE `Days` DISABLE KEYS */;
INSERT INTO `Days` VALUES
(33,13,'2024-09-29',10),
(34,13,'2024-09-27',NULL),
(35,13,'2024-09-28',NULL),
(36,13,'2024-09-18',NULL),
(37,13,'2024-09-19',NULL),
(38,18,'2024-09-29',11),
(39,18,'2024-10-25',NULL),
(40,18,'2024-10-24',NULL),
(41,18,'2024-10-17',NULL),
(42,18,'2024-09-13',NULL),
(43,18,'2024-09-21',NULL),
(44,18,'2024-09-01',NULL),
(45,13,'2024-09-30',10),
(46,13,'2024-09-25',NULL),
(47,13,'2024-09-24',NULL),
(48,13,'2024-09-17',10),
(49,13,'2024-09-03',NULL),
(50,13,'2024-09-09',13),
(51,13,'2024-09-10',4),
(52,13,'2024-09-11',NULL),
(53,13,'2024-09-08',NULL),
(54,13,'2024-10-03',10),
(55,13,'2024-10-02',NULL),
(56,13,'2024-10-10',NULL),
(57,13,'2024-09-01',9),
(58,13,'2024-09-22',NULL),
(59,13,'2024-09-23',NULL),
(60,13,'2024-10-15',10);
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
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Exercises`
--

LOCK TABLES `Exercises` WRITE;
/*!40000 ALTER TABLE `Exercises` DISABLE KEYS */;
INSERT INTO `Exercises` VALUES
(6,'Push Ups',3,12,'A basic upper body exercise.','https://example.com/images/pushups.jpg'),
(8,'Squat',4,12,'A fundamental lower body exercise.','http://example.com/squat.jpg'),
(9,'SquatPutilla',4,12,'A fundamental lower body exercise.','http://example.com/squat.jpg'),
(10,'Bench Press',3,10,'A strength training exercise that targets the chest, shoulders, and triceps.','https://example.com/photos/bench_press.jpg'),
(11,'Deadlift',4,8,'A compound exercise that works the entire posterior chain.','https://example.com/photos/deadlift.jpg'),
(12,'Squat',3,12,'A foundational exercise that targets the quadriceps, hamstrings, and glutes.','https://example.com/photos/squat.jpg'),
(13,'Pull-Up',3,6,'An upper body exercise that targets the back and biceps.','https://example.com/photos/pull_up.jpg'),
(14,'Lunges',3,10,'A lower body exercise that targets the legs and glutes.','https://example.com/photos/lunges.jpg'),
(15,'Shoulder Press',4,10,'An exercise that targets the shoulders and triceps using a barbell or dumbbells.','https://example.com/photos/shoulder_press.jpg'),
(16,'Plank',3,NULL,'A core strengthening exercise that involves holding a position similar to a push-up.','https://example.com/photos/plank.jpg'),
(17,'Bicep Curl',3,12,'An isolation exercise that targets the biceps using dumbbells or a barbell.','https://example.com/photos/bicep_curl.jpg'),
(18,'Tricep Dips',3,10,'An exercise that targets the triceps, performed using parallel bars or a bench.','https://example.com/photos/tricep_dips.jpg'),
(19,'Leg Press',4,10,'A compound exercise that targets the quadriceps, hamstrings, and glutes using a leg press machine.','https://example.com/photos/leg_press.jpg'),
(20,'Ejercicio de prueba 3',32,34,'Ejercicio de prueba 3','Ejercicio de prueba 3');
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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
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
(19,'Comidita Rica',2372,23,232,23232,'Es una rutina muy chula','Comida de prueba');
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
(10,11);
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
(4,'Cardio Blast cocomelon','High-intensity workout focusing on cardiovascular endurance.',''),
(9,'Rutina de prueba 6','Hola caracola',''),
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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Users`
--

LOCK TABLES `Users` WRITE;
/*!40000 ALTER TABLE `Users` DISABLE KEYS */;
INSERT INTO `Users` VALUES
(8,'jose','$2a$10$Fs7LFhWGiHlNgBJRot4DGO3D4ArFIsrEcU7dRVLEfXP4VFWZPCpWa\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mauricio@gg'),
(9,'jose2','$2a$10$5Yi51eGQtthg1XtjLRD/kuADTKXowiQ8C3PiYFvJaWt8a5osZl1jK\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mauricio@gg'),
(10,'jose22','$2a$10$lClctsr8Axn6b69zphgiRO97CbLQUo6Th6WfTGfpvpfpCWHrwvJKS\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mauricio@gg'),
(11,'jose222','$2a$10$l72gb5h52nYOf.OPVDQCkOM5LqzTNa1IS.gV7wf2uq7FKnOjir8u.\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mauricio@gg'),
(12,'front','$2a$10$Qk9bDFPIE4rdZ5.utuW.eOPbG3Wb.7qwsDDtBNM5ZLcPsDpkVlSwq\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mauricio@gg'),
(13,'fr','$2a$10$UQifRZ/rafVJtX.c5AryNeLC7divaYZOIUAlIYgJ.mzbPadJJa4ja\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','mauricio@gg'),
(14,'Kalhan','$2a$10$upRHORx/p1xzmFFA3qrre.kCPa.k6E1RDe/64FWw/vJyAlbhYBQTK\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','kalhan@gmail.com'),
(15,'pepe','$2a$10$gOxKVmJZXCz.QXfSYOMHEu85RRP9nClxZUU6UpznnJFVQZ5PsYdpm\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','cruasan'),
(16,'miguel','$2a$10$An3Atq6wqdn5cqT7RCCkte4Soh8X5IEh.2LUR9X5WlPKEplivO2XW\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','fernandez@gmail.com'),
(17,'migensio','$2a$10$sb77RqEHjjKsjq.rxU6.CuHQmdOmOWGYKZVouG9XLu6QWOY2qLbWS\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','ddd'),
(18,'TheScouting','$2a$10$cb.frunIU6.UeaDH6lwj0exDDU3BKeWOQiWBrv.ON2uiiScyPyB0O\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','askdjalks');
/*!40000 ALTER TABLE `Users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2024-10-04  9:54:05
