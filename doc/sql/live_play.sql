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

 Date: 24/11/2020 14:21:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for live_play
-- ----------------------------
DROP TABLE IF EXISTS `live_play`;
CREATE TABLE `live_play`  (
  `id` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `live_id` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'lives表id',
  `sites_id` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '线路1,线路2;source_sites表id',
  `resolution` tinyint(4) NOT NULL DEFAULT 1 COMMENT '分辨率;1=标清,2=高清,3=超清',
  `status` tinyint(4) NULL DEFAULT 1 COMMENT '1=显示,2=隐藏',
  `pull_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '播流地址',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `web_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用于爬取的网页地址',
  `match_at` datetime(0) NULL DEFAULT NULL COMMENT '比赛时间;只抓取3小时以内的比赛',
  `jump_page` tinyint(1) NULL DEFAULT 1 COMMENT '1=无需跳转,2=需要跳转',
  `video_type` tinyint(1) NULL DEFAULT 1 COMMENT '视频格式;1=flv,2=m3u8',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_index`(`web_url`, `video_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '播流地址' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of live_play
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
