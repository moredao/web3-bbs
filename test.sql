/*
Navicat MySQL Data Transfer

Source Server         : 8.218.119.251
Source Server Version : 50650
Source Host           : 8.218.119.251:3306
Source Database       : myzx

Target Server Type    : MYSQL
Target Server Version : 50650
File Encoding         : 65001

Date: 2022-05-02 19:41:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for f_binance
-- ----------------------------
DROP TABLE IF EXISTS `f_binance`;
CREATE TABLE `f_binance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `time` datetime DEFAULT NULL,
  `strategy_name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `f_pnl_p` float DEFAULT NULL,
  `equity_change` float DEFAULT NULL,
  `trade_usdt_simulate` float DEFAULT NULL,
  `account` float DEFAULT NULL,
  `pnl_p` float DEFAULT NULL,
  `d_trade_usdt_simulate` float DEFAULT NULL,
  `f_type` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5913 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
