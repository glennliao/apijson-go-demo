/*
 Navicat Premium Data Transfer

 Source Server         : pi
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : 192.168.31.70:3306
 Source Schema         : apijson_go_demo_hello

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 07/01/2023 16:40:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
                           `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                           `user_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                           `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                           `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                           `created_at` datetime NULL DEFAULT NULL,
                           PRIMARY KEY (`id`) USING BTREE,
                           UNIQUE INDEX `User_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (2, '10001', 'wangmiao', '汪淼', '2022-10-24 17:04:11');
INSERT INTO `t_user` VALUES (4, '10002', 'shiqiang', '史强', '2022-10-24 17:06:09');
INSERT INTO `t_user` VALUES (6, '10003', 'dingyi', '丁仪', '2022-10-24 17:06:57');
INSERT INTO `t_user` VALUES (8, '10004', 'linyun', '林云', '2022-10-24 17:07:23');

SET FOREIGN_KEY_CHECKS = 1;
