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

 Date: 27/01/2024 22:01:09
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
) ENGINE = MyISAM AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Records of checkin_record
-- ----------------------------
INSERT INTO `checkin_record` VALUES (15, 3, 1, 1, '2024-01-27 12:45:03');
INSERT INTO `checkin_record` VALUES (14, 2, 2, 1, '2024-01-27 12:41:46');

-- ----------------------------
-- Table structure for integral
-- ----------------------------
DROP TABLE IF EXISTS `integral`;
CREATE TABLE `integral`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `integral` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of integral
-- ----------------------------
INSERT INTO `integral` VALUES (11, 2, 10);
INSERT INTO `integral` VALUES (12, 3, 5);

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
-- Records of integral_record
-- ----------------------------
INSERT INTO `integral_record` VALUES (1, 2, 5, '签到', '2024-01-27 12:41:45');
INSERT INTO `integral_record` VALUES (2, 3, 5, '签到', '2024-01-27 12:45:03');

-- ----------------------------
-- Table structure for task_record
-- ----------------------------
DROP TABLE IF EXISTS `task_record`;
CREATE TABLE `task_record`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `task_id` int NOT NULL,
  `isFinished` tinyint NOT NULL DEFAULT 1 COMMENT '0 means not completed, 1 means completed',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = FIXED;

-- ----------------------------
-- Records of task_record
-- ----------------------------
INSERT INTO `task_record` VALUES (2, 2, 1, 1);
INSERT INTO `task_record` VALUES (3, 2, 5, 2);

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_cs_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tasks
-- ----------------------------
INSERT INTO `tasks` VALUES (1, 1, '参与任意抽奖', 4, 0);
INSERT INTO `tasks` VALUES (2, 1, '订阅签到提醒', 3, 0);
INSERT INTO `tasks` VALUES (3, 1, '发起任意抽奖', 3, 0);
INSERT INTO `tasks` VALUES (4, 2, '参与3个首页抽奖', 5, 0);
INSERT INTO `tasks` VALUES (5, 2, '观看完整视频1次', 8, 0);
INSERT INTO `tasks` VALUES (6, 2, '邀请5位好友参与抽奖成功', 10, 0);
INSERT INTO `tasks` VALUES (7, 3, '参与首页抽奖30次以上', 30, 0);
INSERT INTO `tasks` VALUES (8, 3, '发起抽奖并超过10人参与', 20, 0);
INSERT INTO `tasks` VALUES (9, 3, '给晒单的锦鲤点1个赞', 3, 0);

SET FOREIGN_KEY_CHECKS = 1;
