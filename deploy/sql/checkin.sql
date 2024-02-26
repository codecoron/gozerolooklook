/*
 Navicat Premium Data Transfer

 Source Server         : lottery
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : checkin

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 24/02/2024 18:22:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for checkin_record
-- ----------------------------
DROP TABLE IF EXISTS `checkin_record`;
CREATE TABLE `checkin_record`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `continuous_checkin_days` tinyint NOT NULL DEFAULT 0 COMMENT 'Number of consecutive check-in days',
  `state` tinyint NOT NULL DEFAULT 0 COMMENT 'Whether to sign in on the day, 1 means signed, 0 means not signed.',
  `last_checkin_date` datetime NULL DEFAULT NULL COMMENT 'Date of last check-in',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Table structure for integral
-- ----------------------------
DROP TABLE IF EXISTS `integral`;
CREATE TABLE `integral`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `integral` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for integral_record
-- ----------------------------
DROP TABLE IF EXISTS `integral_record`;
CREATE TABLE `integral_record`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `integral` int NOT NULL COMMENT 'points added or subtracted',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_cs_0900_ai_ci NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Table structure for task_progress
-- ----------------------------
DROP TABLE IF EXISTS `task_progress`;
CREATE TABLE `task_progress`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `isParticipatedLottery` int NOT NULL DEFAULT 0,
  `isInitiatedLottery` int NOT NULL DEFAULT 0,
  `is_sub_checkin` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for task_record
-- ----------------------------
DROP TABLE IF EXISTS `task_record`;
CREATE TABLE `task_record`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` tinyint NOT NULL,
  `user_id` int NOT NULL,
  `task_id` int NOT NULL,
  `isFinished` tinyint NOT NULL DEFAULT 1 COMMENT '0 means not completed, 1 means completed',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` tinyint NOT NULL COMMENT '1 for novice, 2 for daily, 3 for weekly',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_cs_0900_ai_ci NOT NULL,
  `integral` int NOT NULL COMMENT 'Increased wish value after completion',
  `del_state` tinyint NOT NULL DEFAULT 0,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_cs_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
