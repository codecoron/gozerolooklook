/*
 Navicat Premium Data Transfer

 Source Server         : go-zero-looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:33069
 Source Schema         : lottery

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 04/03/2024 15:41:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for clock_task_record
-- ----------------------------
DROP TABLE IF EXISTS `clock_task_record`;
CREATE TABLE `clock_task_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lottery_id` int NOT NULL COMMENT '抽奖ID',
  `clock_task_id` int NOT NULL COMMENT '打卡任务ID',
  `user_id` int NOT NULL COMMENT '用户id',
  `increase_multiple` int NOT NULL DEFAULT '1' COMMENT '概率增加倍数',
  `del_state` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='打卡任务记录表';

SET FOREIGN_KEY_CHECKS = 1;
