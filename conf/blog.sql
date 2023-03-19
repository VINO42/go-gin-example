/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 19/03/2023 22:01:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '标签ID',
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '简述',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_on` int(0) NULL DEFAULT NULL,
  `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(0) UNSIGNED NULL DEFAULT 0,
  `state` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES (2, 5, '坚持严的主基调不动摇 坚持不懈把全面从严治党向纵深推进', '坚持严的主基调不动摇 坚持不懈把全面从严治党向纵深推进1', '一百年来，党外靠发展人民民主、接受人民监督，内靠全面从严治党、推进自我革命，勇于坚持真理、修正错误，勇于刀刃向内、刮骨疗毒，保证了党长盛不衰、不断发展壮大。全面从严治党是新时代党的自我革命的伟大实践，开辟了百年大党自我革命的新境界。必须坚持以党的政治建设为统领，坚守自我革命根本政治方向；必须坚持把思想建设作为党的基础性建设，淬炼自我革命锐利思想武器；必须坚决落实中央八项规定精神、以严明纪律整饬作风，丰富自我革命有效途径；必须坚持以雷霆之势反腐惩恶，打好自我革命攻坚战、持久战；必须坚持增强党组织政治功能和组织力凝聚力，锻造敢于善于斗争、勇于自我革命的干部队伍；必须坚持构建自我净化、自我完善、自我革新、自我提高的制度规范体系，为推进伟大自我革命提供制度保障。', 1642516200, 'vino', 1679233987, '', 0, 0);
INSERT INTO `blog_article` VALUES (3, 5, '坚持严的主基调不动摇 坚持不懈把全面从严治党向纵深推进', '坚持严的主基调不动摇 坚持不懈把全面从严治党向纵深推进1', '一百年来，党外靠发展人民民主、接受人民监督，内靠全面从严治党、推进自我革命，勇于坚持真理、修正错误，勇于刀刃向内、刮骨疗毒，保证了党长盛不衰、不断发展壮大。全面从严治党是新时代党的自我革命的伟大实践，开辟了百年大党自我革命的新境界。必须坚持以党的政治建设为统领，坚守自我革命根本政治方向；必须坚持把思想建设作为党的基础性建设，淬炼自我革命锐利思想武器；必须坚决落实中央八项规定精神、以严明纪律整饬作风，丰富自我革命有效途径；必须坚持以雷霆之势反腐惩恶，打好自我革命攻坚战、持久战；必须坚持增强党组织政治功能和组织力凝聚力，锻造敢于善于斗争、勇于自我革命的干部队伍；必须坚持构建自我净化、自我完善、自我革新、自我提高的制度规范体系，为推进伟大自我革命提供制度保障。', 1648291053, 'vino', 0, '', 0, 1);
INSERT INTO `blog_article` VALUES (4, 5, '坚持严的主基调不动摇 坚持不懈把全面从严治党向纵深推进', '坚持严的主基调不动摇 坚持不懈把全面从严治党向纵深推进1', '一百年来，党外靠发展人民民主、接受人民监督，内靠全面从严治党、推进自我革命，勇于坚持真理、修正错误，勇于刀刃向内、刮骨疗毒，保证了党长盛不衰、不断发展壮大。全面从严治党是新时代党的自我革命的伟大实践，开辟了百年大党自我革命的新境界。必须坚持以党的政治建设为统领，坚守自我革命根本政治方向；必须坚持把思想建设作为党的基础性建设，淬炼自我革命锐利思想武器；必须坚决落实中央八项规定精神、以严明纪律整饬作风，丰富自我革命有效途径；必须坚持以雷霆之势反腐惩恶，打好自我革命攻坚战、持久战；必须坚持增强党组织政治功能和组织力凝聚力，锻造敢于善于斗争、勇于自我革命的干部队伍；必须坚持构建自我净化、自我完善、自我革新、自我提高的制度规范体系，为推进伟大自我革命提供制度保障。', 1679234375, 'vino', 0, '', 0, 1);

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
INSERT INTO `blog_auth` VALUES (1, 'test', 'test123456');

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '标签名称',
  `created_on` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(0) UNSIGNED NULL DEFAULT 0,
  `state` tinyint(0) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章标签管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (3, '体育', 0, 'vino', 0, '', 0, 0);
INSERT INTO `blog_tag` VALUES (4, '文化', 0, 'vino', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (5, '电影哈哈11', 0, 'vino', 0, 'vino', 0, 1);
INSERT INTO `blog_tag` VALUES (6, '电影', 0, 'vino', 0, '', 0, 1);

SET FOREIGN_KEY_CHECKS = 1;
