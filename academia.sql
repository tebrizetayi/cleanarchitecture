CREATE TABLE `Author` (
  `Id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `Name` varchar(256) NOT NULL,
  PRIMARY KEY (`Id`)
);

CREATE TABLE `Article` (
  `id` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Id` (`id`),
  KEY `index_name` (`Name`)
)