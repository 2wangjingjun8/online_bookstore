/*
Navicat MySQL Data Transfer

Source Server         : bd
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : bookstore

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-02-20 11:08:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `bookname` varchar(100) NOT NULL,
  `author` varchar(60) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `sales` int(6) NOT NULL DEFAULT '0',
  `stock` int(6) NOT NULL DEFAULT '0',
  `img_path` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO `book` VALUES ('1', 'php书籍', '小小', '50.00', '102', '898', '/static/images/docker.jpg');
INSERT INTO `book` VALUES ('2', 'python书籍', '毛毛', '45.00', '102', '898', '/static/images/docker.jpg');
INSERT INTO `book` VALUES ('3', '三国演义', '罗贯中', '88.80', '102', '898', '/static/images/docker.jpg');
INSERT INTO `book` VALUES ('4', '水浒传', '斯耐庵', '66.88', '1000', '0', '/static/images/docker.jpg');
INSERT INTO `book` VALUES ('6', '西游记', '吴承恩', '25.00', '100', '900', '/static/images/docker.jpg');
INSERT INTO `book` VALUES ('7', '苏东坡传', '林语堂', '20.00', '100', '900', '/static/images/docker.jpg');
INSERT INTO `book` VALUES ('8', '童话', '未知', '50.00', '100', '900', '/static/images/docker.jpg');

-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts` (
  `id` varchar(100) NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11,2) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of carts
-- ----------------------------
INSERT INTO `carts` VALUES ('04f1f4f0-5970-4cd0-4dc8-c530a7be76a7', '20', '950.00', '1');

-- ----------------------------
-- Table structure for cart_items
-- ----------------------------
DROP TABLE IF EXISTS `cart_items`;
CREATE TABLE `cart_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `count` int(11) NOT NULL,
  `amount` double(11,2) NOT NULL,
  `book_id` int(11) NOT NULL,
  `cart_id` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `book_id` (`book_id`),
  KEY `cart_id` (`cart_id`)
) ENGINE=MyISAM AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of cart_items
-- ----------------------------
INSERT INTO `cart_items` VALUES ('43', '10', '450.00', '2', '04f1f4f0-5970-4cd0-4dc8-c530a7be76a7');
INSERT INTO `cart_items` VALUES ('44', '10', '500.00', '1', '04f1f4f0-5970-4cd0-4dc8-c530a7be76a7');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `orderNo` varchar(100) NOT NULL,
  `orderTime` datetime NOT NULL,
  `totalCount` int(10) NOT NULL,
  `totalAmount` double(11,2) NOT NULL,
  `state` tinyint(1) NOT NULL,
  `userId` int(10) NOT NULL,
  PRIMARY KEY (`orderNo`),
  KEY `userId_index` (`userId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES ('8148d835-c1bb-4000-6bd3-4e257fc7bffb', '2020-02-19 12:20:11', '3', '183.80', '2', '1');
INSERT INTO `orders` VALUES ('e7517b4f-3547-404a-5bf1-c37c5bbb2087', '2020-02-19 12:22:20', '3', '183.80', '2', '1');
INSERT INTO `orders` VALUES ('85d8bd22-8f48-45f9-6acb-f34b9ce4099f', '2020-02-19 12:36:33', '10', '668.80', '1', '1');
INSERT INTO `orders` VALUES ('dde9eddd-b4b6-4880-7e07-a1418d94e6aa', '2020-02-19 12:37:39', '100', '6688.00', '0', '3');
INSERT INTO `orders` VALUES ('580a973b-22f3-4536-60e5-95c5ff3bf4ad', '2020-02-19 20:39:12', '150', '10032.00', '0', '3');
INSERT INTO `orders` VALUES ('4ef76038-2c39-47df-63ec-9eb6de300aa5', '2020-02-19 20:48:37', '640', '42803.20', '1', '3');

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `price` decimal(10,2) NOT NULL,
  `count` int(10) NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  `bookName` varchar(100) NOT NULL,
  `author` varchar(60) NOT NULL,
  `imgPath` varchar(255) NOT NULL,
  `orderNo` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `orderNo_index` (`orderNo`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of order_items
-- ----------------------------
INSERT INTO `order_items` VALUES ('7', '50.00', '1', '50.00', 'php书籍', '小小', '/static/images/docker.jpg', '8148d835-c1bb-4000-6bd3-4e257fc7bffb');
INSERT INTO `order_items` VALUES ('8', '45.00', '1', '45.00', 'python书籍', '毛毛', '/static/images/docker.jpg', '8148d835-c1bb-4000-6bd3-4e257fc7bffb');
INSERT INTO `order_items` VALUES ('9', '88.80', '1', '88.80', '三国演义', '罗贯中', '/static/images/docker.jpg', '8148d835-c1bb-4000-6bd3-4e257fc7bffb');
INSERT INTO `order_items` VALUES ('10', '50.00', '1', '50.00', 'php书籍', '小小', '/static/images/docker.jpg', 'e7517b4f-3547-404a-5bf1-c37c5bbb2087');
INSERT INTO `order_items` VALUES ('11', '45.00', '1', '45.00', 'python书籍', '毛毛', '/static/images/docker.jpg', 'e7517b4f-3547-404a-5bf1-c37c5bbb2087');
INSERT INTO `order_items` VALUES ('12', '88.80', '1', '88.80', '三国演义', '罗贯中', '/static/images/docker.jpg', 'e7517b4f-3547-404a-5bf1-c37c5bbb2087');
INSERT INTO `order_items` VALUES ('13', '66.88', '10', '668.80', '水浒传', '斯耐庵', '/static/images/docker.jpg', '85d8bd22-8f48-45f9-6acb-f34b9ce4099f');
INSERT INTO `order_items` VALUES ('14', '66.88', '100', '6688.00', '水浒传', '斯耐庵', '/static/images/docker.jpg', 'dde9eddd-b4b6-4880-7e07-a1418d94e6aa');
INSERT INTO `order_items` VALUES ('15', '66.88', '150', '10032.00', '水浒传', '斯耐庵', '/static/images/docker.jpg', '580a973b-22f3-4536-60e5-95c5ff3bf4ad');
INSERT INTO `order_items` VALUES ('16', '66.88', '640', '42803.20', '水浒传', '斯耐庵', '/static/images/docker.jpg', '4ef76038-2c39-47df-63ec-9eb6de300aa5');

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `uuid` varchar(100) NOT NULL,
  `username` varchar(255) NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`uuid`),
  KEY `user_id_index` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sessions
-- ----------------------------
INSERT INTO `sessions` VALUES ('6a953fda-ffe6-4666-7de1-6ed8d52e4edd', 'admin', '1');
INSERT INTO `sessions` VALUES ('2a808320-9011-4c0e-7f25-bd4f7a2c1efa', 'admin', '1');
INSERT INTO `sessions` VALUES ('322c09e1-ad76-4fe1-4d0c-d8edf220cfd7', 'admin', '1');
INSERT INTO `sessions` VALUES ('bb286dbc-eebc-489b-7050-134b0046b947', 'admin', '1');
INSERT INTO `sessions` VALUES ('e63025f7-37f0-4016-7c1f-622d17bd49bc', 'xiaoxiao', '2');
INSERT INTO `sessions` VALUES ('bc6c93c6-ec7e-41dd-4ba9-e9230e1d4726', 'admin', '1');
INSERT INTO `sessions` VALUES ('cf1da753-5e46-462d-4e0e-daeeecdda36f', 'admin', '1');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `email` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', '123456', 'admin@qq.com');
INSERT INTO `user` VALUES ('2', 'xiaoxiao', '123456', 'xiaoxiao@qq.com');
INSERT INTO `user` VALUES ('3', 'maomao', '123456', 'maomao@qq.com');
