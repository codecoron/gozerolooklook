/*
 Navicat Premium Data Transfer

 Source Server         : gozero-looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : lottery

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 17/01/2024 23:07:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lottery
-- ----------------------------
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `user_id` int(0) NOT NULL DEFAULT 0 COMMENT '发起抽奖用户ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等奖名称',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认取一等经配图',
  `publish_time` datetime(0) NULL DEFAULT NULL COMMENT '发布抽奖时间',
  `join_number` int(0) NOT NULL DEFAULT 0 COMMENT '自动开奖人数',
  `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '抽奖说明',
  `award_deadline` datetime(0) NOT NULL COMMENT '领奖截止时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_selected` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否精选 1是 0否',
  `announce_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '开奖设置：1按时间开奖 2按人数开奖 3即抽即中',
  `announce_time` datetime(0) NULL DEFAULT NULL COMMENT '开奖时间',
  `is_order_by_people_draw_prize` tinyint(1) NOT NULL DEFAULT 0 COMMENT '按照人数依次解锁奖品',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '抽奖表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for prize
-- ----------------------------
DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `lottery_id` int(0) NOT NULL DEFAULT 0 COMMENT '抽奖ID',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '奖品类型：1奖品 2优惠券 3兑换码 4商城 5微信红包封面 6红包',
  `name` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品名称',
  `level` int(0) NOT NULL DEFAULT 1 COMMENT '几等奖 默认1',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品图',
  `count` int(0) NOT NULL DEFAULT 0 COMMENT '奖品份数',
  `grant_type` tinyint(1) NOT NULL COMMENT '奖品发放方式：1快递邮寄 2让中奖者联系我 3中奖者填写信息 4跳转到其他小程序',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '奖品表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
