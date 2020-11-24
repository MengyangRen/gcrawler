/*
 Navicat Premium Data Transfer

 Source Server         : 本地-Mysql
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3306
 Source Schema         : live_program

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 24/11/2020 14:21:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for live_class
-- ----------------------------
DROP TABLE IF EXISTS `live_class`;
CREATE TABLE `live_class`  (
  `id` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名',
  `code` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类简码',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '1=显示,2=隐藏',
  `sort` tinyint(3) UNSIGNED NOT NULL DEFAULT 100 COMMENT '排序;0-255,越小越靠前',
  `icon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标url',
  `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '1=体育,2=电竞',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '直播分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of live_class
-- ----------------------------
INSERT INTO `live_class` VALUES (1, '足球', 'soccer', 1, 1, 'http://172.21.34.101:3010/soccer.png', 1);
INSERT INTO `live_class` VALUES (2, '篮球', 'basketball', 1, 2, 'http://172.21.34.101:3010/basketball.png', 1);
INSERT INTO `live_class` VALUES (3, '综合', 'complex_sports', 1, 3, 'http://172.21.34.101:3010/esports.png', 1);
INSERT INTO `live_class` VALUES (4, 'DOTA2', 'dota2', 2, 100, 'http://172.21.34.101:3010/esports.png', 2);
INSERT INTO `live_class` VALUES (5, '英雄联盟', 'lol', 2, 100, 'http://172.21.34.101:3010/esports.png', 2);
INSERT INTO `live_class` VALUES (6, '电竞', 'esports', 1, 4, 'http://172.21.34.101:3010/esports.png', 2);

SET FOREIGN_KEY_CHECKS = 1;
