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

 Date: 21/02/2024 14:31:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for clock_task
-- ----------------------------
DROP TABLE IF EXISTS `clock_task`;
CREATE TABLE `clock_task` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lottery_id` int NOT NULL COMMENT '抽奖ID',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '任务类型 1: 体验小程序 2： 浏览指定公众号文章 3: 浏览图片（微信图片二维码等） 4： 浏览视频号视频',
  `seconds` int NOT NULL DEFAULT '0' COMMENT '任务秒数',
  `applet_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'type=1时该字段才有意义，1小程序链接 2小程序路径',
  `page_link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=1 并且 applet_type=1时 该字段才有意义，配置要跳转小程序的页面链接 （如 #小程序://抽奖/oM....）',
  `app_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=1 并且 applet_type=2时 该字段才有意义，配置要跳转的小程序AppID',
  `page_path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=1 并且 applet_type=2时 该字段才有意义，配置要跳转的小程序路径（如：/pages/index）',
  `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=3时 该字段才有意义，添加要查看的图片',
  `video_account_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=4时 该字段才有意义，视频号ID',
  `video_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=4时 该字段才有意义，视频ID',
  `article_link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=2时 该字段才有意义，公众号文章链接',
  `copywriting` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '引导参与者完成打卡任务的文案',
  `chance_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '概率类型 1: 随机 2: 指定 ',
  `increase_multiple` int NOT NULL DEFAULT '1' COMMENT 'chance_type=2时 该字段才有意义，概率增加倍数',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='打卡任务表';

SET FOREIGN_KEY_CHECKS = 1;
