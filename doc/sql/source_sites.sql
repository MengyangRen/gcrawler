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

 Date: 24/11/2020 14:22:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for source_sites
-- ----------------------------
DROP TABLE IF EXISTS `source_sites`;
CREATE TABLE `source_sites`  (
  `id` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站名',
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站地址',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '1=正常,2=不抓取',
  `code` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站简码',
  `api_conf` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '爬取的目标网站' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of source_sites
-- ----------------------------
INSERT INTO `source_sites` VALUES (1, '火力直播', 'http://www.huolisport.cn', 1, 'huoli', '{\"liveUrl\":\"http://www.huolisport.cn\",\"host\":\"www.huolisport.cn\",\"iframeHost\":\"live.iixxix.cn\"}');
INSERT INTO `source_sites` VALUES (2, '火速直播', 'http://huosu.tv', 2, 'huosu', '');
INSERT INTO `source_sites` VALUES (3, '好看体育', 'http://www.haokan.tv', 2, 'haokan', '');
INSERT INTO `source_sites` VALUES (4, 'jrs直播', 'http://www.jrszbz.com', 2, 'jrszb', '');
INSERT INTO `source_sites` VALUES (5, 'jrs看球网', 'http://www.jrskq.com/', 1, 'jrskq', '');

SET FOREIGN_KEY_CHECKS = 1;
