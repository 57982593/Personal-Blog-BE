/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 80021
Source Host           : localhost:3306
Source Database       : personal-blog-db

Target Server Type    : MYSQL
Target Server Version : 80021
File Encoding         : 65001

Date: 2021-09-18 16:40:19
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `user_id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) DEFAULT NULL,
  `update_at` int DEFAULT NULL,
  `create_at` int DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('a', 'b', '1', '123', null, null);
INSERT INTO `user` VALUES ('test', '123456', '4', '1234567891', '1631871218', '1631871218');
INSERT INTO `user` VALUES ('', '2134567891', '8', '1234567891', '1631949205', '1631949205');
