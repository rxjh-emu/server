-- --------------------------------------------------------
-- Host:                         localhost
-- Server version:               11.2.2-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table db_world.characters
CREATE TABLE IF NOT EXISTS `characters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` int(11) NOT NULL,
  `name` varchar(16) NOT NULL,
  `slot` int(1) NOT NULL DEFAULT 0,
  `class` int(2) NOT NULL,
  `level` int(3) NOT NULL DEFAULT 1,
  `ki` int(11) NOT NULL DEFAULT 0,
  `spr` int(3) NOT NULL DEFAULT 0,
  `str` int(3) NOT NULL DEFAULT 0,
  `stm` int(3) NOT NULL DEFAULT 0,
  `dex` int(3) NOT NULL DEFAULT 0,
  `fame` int(3) NOT NULL DEFAULT 0,
  `morals` int(3) NOT NULL DEFAULT 0,
  `hp` int(3) NOT NULL DEFAULT 0,
  `mp` int(3) NOT NULL DEFAULT 0,
  `rp` int(3) NOT NULL DEFAULT 0,
  `exp` int(3) NOT NULL DEFAULT 0,
  `gender` int(1) NOT NULL,
  `hair` int(10) NOT NULL,
  `hair_color` int(10) NOT NULL,
  `face` int(10) NOT NULL,
  `voice` int(1) DEFAULT NULL,
  `map` int(11) NOT NULL,
  `x` float NOT NULL,
  `y` float NOT NULL,
  `z` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- Data exporting was unselected.

-- Dumping structure for function db_world.getCharacterFreeSlot
DELIMITER //
CREATE FUNCTION `getCharacterFreeSlot`(`account_id` INT
) RETURNS int(11)
BEGIN
	DECLARE available_slot INT;

	SELECT MIN(slot + 1) INTO available_slot
	FROM characters
	WHERE account = account_id
	AND (slot + 1) NOT IN (SELECT slot FROM characters WHERE account = account_id)
	AND (slot + 1) > 0;
	
	RETURN available_slot;
END//
DELIMITER ;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
