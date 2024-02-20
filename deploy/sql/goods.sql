/*
 Navicat MySQL Data Transfer

 Source Server         : go-zero-looklook
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : localhost:33069
 Source Schema         : shop

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 20/02/2024 15:41:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `goods_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `category_id` int NOT NULL DEFAULT 0,
  `precoupon _price` decimal(10, 2) NOT NULL DEFAULT 0.00,
  `aftercoupon _price` decimal(10, 2) NOT NULL DEFAULT 0.00,
  `goods_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `wish_points` int NOT NULL DEFAULT 0,
  `coupon_start_time` int NOT NULL DEFAULT 0,
  `coupon_end_time` int NOT NULL DEFAULT 0,
  `coupon_discount` int NOT NULL DEFAULT 0,
  `coupon_remain_quantity` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
