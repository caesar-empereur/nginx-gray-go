SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for service_node
-- ----------------------------
DROP TABLE IF EXISTS `service_node`;
CREATE TABLE `service_node` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(16) NOT NULL DEFAULT '',
  `port` varchar(16) NOT NULL DEFAULT '',
  `weight` int(4) NOT NULL DEFAULT '1' COMMENT '访问权重',
  `service_name` varchar(32) NOT NULL DEFAULT '' COMMENT '服务名称',
  `domain` varchar(32) NOT NULL DEFAULT '' COMMENT '域名',
  `inner_access` int(1) NOT NULL DEFAULT '0' COMMENT '是否为公司内部ip访问，是为1，否为0',
  `zone` varchar(32) NOT NULL DEFAULT '',
  `join` int(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of service_node
-- ----------------------------
INSERT INTO `service_node` VALUES ('5', '127.0.0.1', '8083', '1', '订单对外服务', 'webapi.com', '1', 'zone_for_backends', '1');
INSERT INTO `service_node` VALUES ('6', '127.0.0.1', '8082', '1', '订单对外服务', 'webapi.com', '1', 'zone_for_backends', '1');
INSERT INTO `service_node` VALUES ('7', '127.0.0.1', '8081', '1', '订单对外服务', 'webapi.com', '0', 'zone_for_backends', '0');
