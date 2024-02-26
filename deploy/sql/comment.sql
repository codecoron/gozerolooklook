/*
 Navicat Premium Data Transfer

 Source Server         : lottery-backend
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : 127.0.0.1:33069
 Source Schema         : comment

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 24/02/2024 21:50:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `lottery_id` int NOT NULL COMMENT '抽奖id',
  `prize_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '奖品名称',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '晒单评论内容',
  `pics` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '晒单评论图片',
  `praise_count` int NOT NULL COMMENT '点赞数量',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `del_state` tinyint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 17, 111, '派部选油', 'voluptate occaecat Excepteur minim', 'http://dummyimage.com/400x400', 123, '2024-02-24 21:42:17', '2024-02-24 21:42:17', NULL, 0);
INSERT INTO `comment` VALUES (2, 18, 111, '派部选油', 'voluptate occaecat Excepteur minim', 'http://dummyimage.com/400x400', 123, '2024-02-24 21:45:14', '2024-02-24 21:45:14', NULL, 0);

-- ----------------------------
-- Table structure for praise
-- ----------------------------
DROP TABLE IF EXISTS `praise`;
CREATE TABLE `praise`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '评论者id',
  `comment_id` int NOT NULL COMMENT '评论id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NULL DEFAULT NULL,
  `del_state` tinyint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `comment_id`(`comment_id` ASC) USING BTREE COMMENT '评论id索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of praise
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
