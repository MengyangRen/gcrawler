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

 Date: 24/11/2020 14:22:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for matchs
-- ----------------------------
DROP TABLE IF EXISTS `matchs`;
CREATE TABLE `matchs`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_id` tinyint(4) NOT NULL COMMENT 'live_class表的id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '赛事名称',
  `nick_name` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '别名集合,以,分割',
  `icon` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `show_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '简称,防止名称太长影响UI',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39537 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '赛事表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of matchs
-- ----------------------------
INSERT INTO `matchs` VALUES (39363, 1, '欧罗巴杯', '', '', '');
INSERT INTO `matchs` VALUES (39364, 1, '墨西乙春', '', '', '');
INSERT INTO `matchs` VALUES (39365, 1, '俄青联', '', '', '');
INSERT INTO `matchs` VALUES (39366, 1, '白俄甲', '', '', '');
INSERT INTO `matchs` VALUES (39367, 2, 'CBA', '', '', '');
INSERT INTO `matchs` VALUES (39368, 2, '菲篮杯', '', '', '');
INSERT INTO `matchs` VALUES (39369, 2, 'NBL', '', '', '');
INSERT INTO `matchs` VALUES (39370, 1, '泰乙', '', '', '');
INSERT INTO `matchs` VALUES (39371, 1, '冰岛超', '', '', '');
INSERT INTO `matchs` VALUES (39372, 1, '中超', '', '', '');
INSERT INTO `matchs` VALUES (39373, 1, '日职乙', '', '', '');
INSERT INTO `matchs` VALUES (39374, 2, '罗女杯', '', '', '');
INSERT INTO `matchs` VALUES (39375, 2, '土女篮', '', '', '');
INSERT INTO `matchs` VALUES (39376, 1, '挪甲', '', '', '');
INSERT INTO `matchs` VALUES (39377, 1, '越南联', '', '', '');
INSERT INTO `matchs` VALUES (39378, 1, '欧女杯', '', '', '');
INSERT INTO `matchs` VALUES (39379, 1, '拉脱超', '', '', '');
INSERT INTO `matchs` VALUES (39380, 1, '爱超', '', '', '');
INSERT INTO `matchs` VALUES (39381, 1, '日职丙', '', '', '');
INSERT INTO `matchs` VALUES (39382, 1, '格鲁甲', '', '', '');
INSERT INTO `matchs` VALUES (39383, 1, '港菁杯', '', '', '');
INSERT INTO `matchs` VALUES (39384, 1, '日职联', '', '', '');
INSERT INTO `matchs` VALUES (39385, 1, '南球杯', '', '', '');
INSERT INTO `matchs` VALUES (39386, 1, '白俄超', '', '', '');
INSERT INTO `matchs` VALUES (39387, 1, '中甲', '', '', '');
INSERT INTO `matchs` VALUES (39388, 1, '中北美联', '', '', '');
INSERT INTO `matchs` VALUES (39389, 1, '不丹廷联', '', '', '');
INSERT INTO `matchs` VALUES (39390, 1, '澳昆超附', '', '', '');
INSERT INTO `matchs` VALUES (39391, 1, '非洲女杯', '', '', '');
INSERT INTO `matchs` VALUES (39392, 2, '俄女超2', '', '', '');
INSERT INTO `matchs` VALUES (39393, 2, '韩篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39394, 1, '中乙', '', '', '');
INSERT INTO `matchs` VALUES (39395, 1, '赞比亚超', '', '', '');
INSERT INTO `matchs` VALUES (39396, 2, '欧篮联', '', '', '');
INSERT INTO `matchs` VALUES (39397, 2, '墨联', '', '', '');
INSERT INTO `matchs` VALUES (39398, 1, '俄杯', '', '', '');
INSERT INTO `matchs` VALUES (39399, 1, '新加坡联', '', '', '');
INSERT INTO `matchs` VALUES (39400, 1, '卡塔杯', '', '', '');
INSERT INTO `matchs` VALUES (39401, 1, '保杯', '', '', '');
INSERT INTO `matchs` VALUES (39402, 1, '哈萨超', '', '', '');
INSERT INTO `matchs` VALUES (39403, 2, '菲律宾PBA', '', '', '');
INSERT INTO `matchs` VALUES (39404, 1, '哥伦甲秋', '', '', '');
INSERT INTO `matchs` VALUES (39405, 1, '约超联', '', '', '');
INSERT INTO `matchs` VALUES (39406, 1, '俄女杯', '', '', '');
INSERT INTO `matchs` VALUES (39407, 1, '阿尔巴甲A', '', '', '');
INSERT INTO `matchs` VALUES (39408, 1, '厄瓜甲秋', '', '', '');
INSERT INTO `matchs` VALUES (39409, 1, '巴圣青联', '', '', '');
INSERT INTO `matchs` VALUES (39410, 1, '泰超', '', '', '');
INSERT INTO `matchs` VALUES (39411, 1, '阿美超', '', '', '');
INSERT INTO `matchs` VALUES (39412, 1, '黑山甲', '', '', '');
INSERT INTO `matchs` VALUES (39413, 1, '巴西甲', '', '', '');
INSERT INTO `matchs` VALUES (39414, 2, '俄女超1', '', '', '');
INSERT INTO `matchs` VALUES (39415, 1, '智利甲', '', '', '');
INSERT INTO `matchs` VALUES (39416, 1, '巴西杯', '', '', '');
INSERT INTO `matchs` VALUES (39417, 2, '俄篮超', '', '', '');
INSERT INTO `matchs` VALUES (39418, 1, '土杯', '', '', '');
INSERT INTO `matchs` VALUES (39419, 1, '坦桑超', '', '', '');
INSERT INTO `matchs` VALUES (39420, 1, '韩足总', '', '', '');
INSERT INTO `matchs` VALUES (39421, 1, '保超', '', '', '');
INSERT INTO `matchs` VALUES (39422, 1, '俄乙南', '', '', '');
INSERT INTO `matchs` VALUES (39423, 1, '波兰丁', '', '', '');
INSERT INTO `matchs` VALUES (39424, 1, '西乙', '', '', '');
INSERT INTO `matchs` VALUES (39425, 1, '意丁', '', '', '');
INSERT INTO `matchs` VALUES (39426, 1, '委內超秋', '', '', '');
INSERT INTO `matchs` VALUES (39427, 1, '芬兰超', '', '', '');
INSERT INTO `matchs` VALUES (39428, 1, '奥地利杯', '', '', '');
INSERT INTO `matchs` VALUES (39429, 1, '德国男排联赛', '', '', '');
INSERT INTO `matchs` VALUES (39430, 2, '阿根廷篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39431, 2, '芬冰青SML', '', '', '');
INSERT INTO `matchs` VALUES (39432, 1, '南俱杯', '', '', '');
INSERT INTO `matchs` VALUES (39433, 1, '中华职棒', '', '', '');
INSERT INTO `matchs` VALUES (39434, 1, 'MHL', '', '', '');
INSERT INTO `matchs` VALUES (39435, 1, '巴圣甲', '', '', '');
INSERT INTO `matchs` VALUES (39436, 1, '欧冠杯', '', '', '');
INSERT INTO `matchs` VALUES (39437, 1, '综合', '', '', '');
INSERT INTO `matchs` VALUES (39438, 1, '体育频道', '', '', '');
INSERT INTO `matchs` VALUES (39439, 2, '瑞士篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39440, 1, '美职联', '', '', '');
INSERT INTO `matchs` VALUES (39441, 2, '哈萨克斯坦篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39442, 1, '俄罗斯乙', '', '', '');
INSERT INTO `matchs` VALUES (39443, 2, '欧篮联盟杯', '', '', '');
INSERT INTO `matchs` VALUES (39444, 2, '俄女篮超', '', '', '');
INSERT INTO `matchs` VALUES (39445, 1, '斯诺克', '', '', '');
INSERT INTO `matchs` VALUES (39446, 2, 'VHL', '', '', '');
INSERT INTO `matchs` VALUES (39447, 1, '电子竞技', '', '', '');
INSERT INTO `matchs` VALUES (39448, 1, '德西联', '', '', '');
INSERT INTO `matchs` VALUES (39449, 1, '非洲冠军杯', '', '', '');
INSERT INTO `matchs` VALUES (39450, 2, '西篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39451, 1, '中北美冠联', '', '', '');
INSERT INTO `matchs` VALUES (39452, 2, '欧篮冠联', '', '', '');
INSERT INTO `matchs` VALUES (39453, 2, '德国女排联赛', '', '', '');
INSERT INTO `matchs` VALUES (39454, 1, '意甲', '', '', '');
INSERT INTO `matchs` VALUES (39455, 1, '秘鲁甲', '', '', '');
INSERT INTO `matchs` VALUES (39456, 1, '智利乙', '', '', '');
INSERT INTO `matchs` VALUES (39457, 1, '瑞士超', '', '', '');
INSERT INTO `matchs` VALUES (39458, 1, '爱尔兰超', '', '', '');
INSERT INTO `matchs` VALUES (39459, 2, '斯洛伐克篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39460, 2, '俄篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39461, 2, '芬兰篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39462, 1, '乌兹别克超', '', '', '');
INSERT INTO `matchs` VALUES (39463, 1, '爱沙杯', '', '', '');
INSERT INTO `matchs` VALUES (39464, 1, '卡塔尔杯', '', '', '');
INSERT INTO `matchs` VALUES (39465, 2, '韩国篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39466, 1, '挪威超', '', '', '');
INSERT INTO `matchs` VALUES (39467, 1, '法罗群岛甲', '', '', '');
INSERT INTO `matchs` VALUES (39468, 1, '体育节目', '', '', '');
INSERT INTO `matchs` VALUES (39469, 1, '阿联酋超', '', '', '');
INSERT INTO `matchs` VALUES (39470, 1, '马耳他甲', '', '', '');
INSERT INTO `matchs` VALUES (39471, 2, '爱拉篮球联合联赛', '', '', '');
INSERT INTO `matchs` VALUES (39472, 2, '俄女篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39473, 2, '韩国职棒', '', '', '');
INSERT INTO `matchs` VALUES (39474, 2, '韩国排球联赛', '', '', '');
INSERT INTO `matchs` VALUES (39475, 1, '阿甲', '', '', '');
INSERT INTO `matchs` VALUES (39476, 1, '瑞典超', '', '', '');
INSERT INTO `matchs` VALUES (39477, 1, '意乙', '', '', '');
INSERT INTO `matchs` VALUES (39478, 1, '西甲', '', '', '');
INSERT INTO `matchs` VALUES (39479, 1, '德甲', '', '', '');
INSERT INTO `matchs` VALUES (39480, 2, '土篮超', '', '', '');
INSERT INTO `matchs` VALUES (39481, 1, '俄甲', '', '', '');
INSERT INTO `matchs` VALUES (39482, 1, '法乙', '', '', '');
INSERT INTO `matchs` VALUES (39483, 2, '篮冠联', '', '', '');
INSERT INTO `matchs` VALUES (39484, 1, '葡超', '', '', '');
INSERT INTO `matchs` VALUES (39485, 1, '阿曼苏杯', '', '', '');
INSERT INTO `matchs` VALUES (39486, 1, '奥乙', '', '', '');
INSERT INTO `matchs` VALUES (39487, 1, '德乙', '', '', '');
INSERT INTO `matchs` VALUES (39488, 1, '巴西乙', '', '', '');
INSERT INTO `matchs` VALUES (39489, 1, '韩挑K联', '', '', '');
INSERT INTO `matchs` VALUES (39490, 1, '土超', '', '', '');
INSERT INTO `matchs` VALUES (39491, 1, '俄超', '', '', '');
INSERT INTO `matchs` VALUES (39492, 1, '英超', '', '', '');
INSERT INTO `matchs` VALUES (39493, 1, '荷甲', '', '', '');
INSERT INTO `matchs` VALUES (39494, 1, '罗甲', '', '', '');
INSERT INTO `matchs` VALUES (39495, 1, '克亚甲', '', '', '');
INSERT INTO `matchs` VALUES (39496, 1, '丹麦超', '', '', '');
INSERT INTO `matchs` VALUES (39497, 1, '比甲', '', '', '');
INSERT INTO `matchs` VALUES (39498, 1, '墨西联春', '', '', '');
INSERT INTO `matchs` VALUES (39499, 1, '波兰超', '', '', '');
INSERT INTO `matchs` VALUES (39500, 1, '法甲', '', '', '');
INSERT INTO `matchs` VALUES (39501, 2, 'WCBA', '', '', '');
INSERT INTO `matchs` VALUES (39502, 1, '巴拉甲秋', '', '', '');
INSERT INTO `matchs` VALUES (39503, 1, '美职业', '', '', '');
INSERT INTO `matchs` VALUES (39504, 2, '俄女超', '', '', '');
INSERT INTO `matchs` VALUES (39505, 2, '乌克杯', '', '', '');
INSERT INTO `matchs` VALUES (39506, 1, '乌兹甲', '', '', '');
INSERT INTO `matchs` VALUES (39507, 1, '乌克U21', '', '', '');
INSERT INTO `matchs` VALUES (39508, 1, '乌兹别克甲', '', '', '');
INSERT INTO `matchs` VALUES (39509, 2, '欧协杯', '', '', '');
INSERT INTO `matchs` VALUES (39510, 2, '乌拉联', '', '', '');
INSERT INTO `matchs` VALUES (39511, 1, '哥斯女甲', '', '', '');
INSERT INTO `matchs` VALUES (39512, 2, 'KHL', '', '', '');
INSERT INTO `matchs` VALUES (39513, 1, '哥伦乙附', '', '', '');
INSERT INTO `matchs` VALUES (39514, 1, '球会友谊', '', '', '');
INSERT INTO `matchs` VALUES (39515, 1, '秘鲁甲秋', '', '', '');
INSERT INTO `matchs` VALUES (39516, 1, '不丹超', '', '', '');
INSERT INTO `matchs` VALUES (39517, 1, '墨女超', '', '', '');
INSERT INTO `matchs` VALUES (39518, 2, 'NBL全男联', '', '', '');
INSERT INTO `matchs` VALUES (39519, 1, '国际友谊', '', '', '');
INSERT INTO `matchs` VALUES (39520, 1, '欧青U21外', '', '', '');
INSERT INTO `matchs` VALUES (39521, 1, '亚冠杯', '', '', '');
INSERT INTO `matchs` VALUES (39522, 1, '摩尔多瓦甲', '', '', '');
INSERT INTO `matchs` VALUES (39523, 2, '澳洲女篮甲', '', '', '');
INSERT INTO `matchs` VALUES (39524, 1, '印度超', '', '', '');
INSERT INTO `matchs` VALUES (39525, 2, '委内瑞拉篮联', '', '', '');
INSERT INTO `matchs` VALUES (39526, 1, '科威特超', '', '', '');
INSERT INTO `matchs` VALUES (39527, 1, '中协杯', '', '', '');
INSERT INTO `matchs` VALUES (39528, 1, '意杯', '', '', '');
INSERT INTO `matchs` VALUES (39529, 1, '中冠联赛', '', '', '');
INSERT INTO `matchs` VALUES (39530, 1, '哈萨克斯坦超', '', '', '');
INSERT INTO `matchs` VALUES (39531, 1, '自由杯', '', '', '');
INSERT INTO `matchs` VALUES (39532, 2, 'CHBL', '', '', '');
INSERT INTO `matchs` VALUES (39533, 2, 'VTB杯', '', '', '');
INSERT INTO `matchs` VALUES (39534, 2, '韩女甲', '', '', '');
INSERT INTO `matchs` VALUES (39535, 2, '世欧预', '', '', '');
INSERT INTO `matchs` VALUES (39536, 2, '澳女篮', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
