/*
 Navicat Premium Data Transfer

 Source Server         : 【抽奖项目】docker-composer-mysql-looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:33069
 Source Schema         : vote

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 24/01/2024 21:42:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for vote_config
-- ----------------------------
DROP TABLE IF EXISTS `vote_config`;
CREATE TABLE `vote_config`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `lottery_id` int NOT NULL COMMENT '抽奖ID',
  `enable_vote` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否启用投票功能 1是 0否',
  `vote_config` json NULL COMMENT '投票配置字段说明: {\"title\": \"投票标题\", \"description\": \"投票描述【非必填】\", \"winner_selection\": \"中奖者设置：1从所有投票者中抽取 2从票数最多的一方中抽取\", \"type\": \"投票类型：1单选 2多选\", \"min_votes\": \"最小投票范围\", \"max_votes\": \"最大投票范围\", \"options\": [{\"text\": \"张三\", \"image\": \"path/to/zhangsan.jpg\"}, {\"text\": \"李四\", \"image\": \"path/to/lisi.jpg\"}, {\"text\": \"王五\", \"image\": \"path/to/wangwu.jpg\"}]}',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '投票表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vote_record
-- ----------------------------
DROP TABLE IF EXISTS `vote_record`;
CREATE TABLE `vote_record`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `lottery_id` int NOT NULL COMMENT '抽奖ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `selected_option` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户选择的投票选项',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '投票时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '投票记录表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
