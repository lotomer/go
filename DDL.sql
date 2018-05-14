/*
Navicat MySQL Data Transfer

Source Server         : vps2
Source Server Version : 50173
Source Host           : vps2.tomstools.org:3306
Source Database       : of

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2018-05-11 23:06:41
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for T_OF_DS_CATEGORY
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_DS_CATEGORY`;
CREATE TABLE `T_OF_DS_CATEGORY` (
  `CATEGORY_CODE` varchar(64) NOT NULL COMMENT '分类编码',
  `CATEGORY_NAME` varchar(64) DEFAULT NULL COMMENT '分类名称',
  `PARENT_CODE` varchar(64) DEFAULT NULL COMMENT '父节点编码',
  `ALL_PATH` varchar(1000) DEFAULT NULL COMMENT '全路径',
  `IS_LEAF` tinyint(4) DEFAULT '1' COMMENT '是否有效标识 0非叶子节点  非0是叶子节点  默认1',
  `CREATOR` varchar(64) DEFAULT NULL COMMENT '创建者',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `LAST_UPDATE_BY` varchar(64) DEFAULT NULL COMMENT '最后修改人',
  `LAST_UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `ENABLE_FLAG` tinyint(4) DEFAULT '1' COMMENT '是否有效标识 0无效  非0有效 默认1',
  PRIMARY KEY (`CATEGORY_CODE`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据服务类别信息表';

-- ----------------------------
-- Table structure for T_OF_DS_DATASOURCE
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_DS_DATASOURCE`;
CREATE TABLE `T_OF_DS_DATASOURCE` (
  `DS_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '数据源编号',
  `DS_TYPE_ID` int(11) NOT NULL COMMENT '数据源类型编号',
  `NAME` varchar(128) NOT NULL COMMENT '数据源名称',
  `OPTIONS` varchar(4000) DEFAULT NULL COMMENT '数据源配置选项(JSON格式）。参照数据源类型中定义的数据源配置',
  `COMMENT` varchar(500) DEFAULT NULL COMMENT '备注',
  `CREATOR` varchar(64) DEFAULT NULL COMMENT '创建者',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `IS_VALID` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效。0 无效；非0 有效。默认1',
  PRIMARY KEY (`DS_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='数据源配置表';

-- ----------------------------
-- Table structure for T_OF_DS_DATASOURCE_TYPE
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_DS_DATASOURCE_TYPE`;
CREATE TABLE `T_OF_DS_DATASOURCE_TYPE` (
  `DS_TYPE_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '数据源类型编号',
  `NAME` varchar(128) NOT NULL COMMENT '数据源类型名称',
  `DS_DEFINE` varchar(2000) NOT NULL COMMENT '数据源配置选项定义（JSON）。在配置该类型的数据源时需要填写的属性定义，如JDBC方式需要填写驱动、url、用户名和密码等信息',
  `DS_DATA_DEFINE` varchar(1000) NOT NULL COMMENT '数据源数据配置选项定义（JSON）。在数据封装时需要填写的内容定义，比如JDBC数据需要填SQL',
  `COMMENT` varchar(500) DEFAULT NULL COMMENT '备注',
  `CREATOR` varchar(64) DEFAULT NULL COMMENT '创建者',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `IS_VALID` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效。0 无效；非0 有效。默认1',
  `PARSE_CLASS` varchar(512) DEFAULT NULL COMMENT '解析类。该类型数据源对应的解析类',
  PRIMARY KEY (`DS_TYPE_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='数据源类型配置表';

-- ----------------------------
-- Table structure for T_OF_DS_SERVICE_CONFIG
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_DS_SERVICE_CONFIG`;
CREATE TABLE `T_OF_DS_SERVICE_CONFIG` (
  `ID` varchar(128) NOT NULL COMMENT '服务编码',
  `DS_ID` int(11) NOT NULL COMMENT '数据源编号',
  `CATEGORY_CODE` varchar(64) NOT NULL COMMENT '分类编码',
  `NAME` varchar(64) NOT NULL COMMENT '服务名称',
  `QUERY_PARAMS` varchar(4000) DEFAULT NULL COMMENT '查询条件列属性集合(JSON格式)',
  `RETURNS` varchar(4000) DEFAULT NULL COMMENT '返回结果配置',
  `OPTIONS` varchar(4000) DEFAULT NULL COMMENT '配置选项(JSON格式)。参照该类数据源的数据配置',
  `STATUS` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态（1表示激活，0表示未激活)。默认 1',
  `CACHE_SECONDS` int(11) DEFAULT '0' COMMENT '缓存时长。单位：秒。小于等于0 表示不缓存；大于0 表示具体缓存的秒数。默认0不缓存',
  `ORDER_NUM` int(11) DEFAULT '1' COMMENT '显示顺序',
  `MEMO` varchar(500) DEFAULT NULL COMMENT '备注',
  `CREATOR` varchar(64) DEFAULT NULL COMMENT '创建者',
  `IN_TIME` datetime NOT NULL  COMMENT '创建时间',
  `MODIFY_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `IS_VALID` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效。0 无效；非0 有效；默认1',
  `IS_HASH` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'rowkey是否采用hash算法处理。0-不处理，1-处理，默认为1',
  PRIMARY KEY (`ID`),
  KEY `AK_UNI_SERVICE_NAME` (`NAME`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据服务配置表';

-- ----------------------------
-- Table structure for T_OF_DS_SERVICE_DIY
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_DS_SERVICE_DIY`;
CREATE TABLE `T_OF_DS_SERVICE_DIY` (
  `ID` varchar(128) NOT NULL COMMENT '服务编码',
  `RETURNS` varchar(8000) DEFAULT NULL COMMENT '返回结果配置',
  `PROCESS_TYPE` varchar(255) DEFAULT 'CLASS' COMMENT '执行类型。CLASS 指定类；CODE 代码片段。默认CLASS',
  `PROCESSOR` varchar(8000) NOT NULL COMMENT '执行器。当PROCESS_TYPE为CLASS时，存放具体类名；为CODE时，存放代码片段',
  `CREATOR` varchar(64) DEFAULT NULL COMMENT '创建者',
  `IN_TIME` datetime NOT NULL  COMMENT '创建时间',
  `MODIFY_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `IS_VALID` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效。0 无效；非0 有效；默认1',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='服务编排配置表';

-- ----------------------------
-- Table structure for T_OF_DS_SERVICE_DIY_DETAIL
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_DS_SERVICE_DIY_DETAIL`;
CREATE TABLE `T_OF_DS_SERVICE_DIY_DETAIL` (
  `ID` varchar(128) NOT NULL COMMENT '服务编码',
  `SUB_ID` varchar(128) NOT NULL COMMENT '子服务编码',
  `CREATOR` varchar(64) DEFAULT NULL COMMENT '创建者',
  `IN_TIME` datetime NOT NULL COMMENT '创建时间',
  `MODIFY_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `IS_VALID` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效。0 无效；非0 有效；默认1',
  PRIMARY KEY (`ID`,`SUB_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='服务编排配置明细表';

-- ----------------------------
-- Table structure for T_OF_PRI_ROLE_RES
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_PRI_ROLE_RES`;
CREATE TABLE `T_OF_PRI_ROLE_RES` (
  `PRIVILEGE_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `RES_TYPE_CODE` varchar(128) NOT NULL COMMENT '资源类型编码',
  `RES_CODE` varchar(128) NOT NULL COMMENT '资源编码',
  `ROLE_ID` varchar(64) NOT NULL COMMENT '角色编号',
  `CONTAIN_CHILD` char(1) NOT NULL DEFAULT '0' COMMENT '是否包含子资源。1 包含子资源； 0 不包含子资源。默认0',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `CREATE_USER` varchar(128) DEFAULT NULL COMMENT '创建者',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `UPDATE_USER` varchar(128) DEFAULT NULL COMMENT '更新者',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效； 0 无效；默认1',
  PRIMARY KEY (`PRIVILEGE_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8 COMMENT='角色资源配置表';

-- ----------------------------
-- Table structure for T_OF_PRI_ROLE_URL
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_PRI_ROLE_URL`;
CREATE TABLE `T_OF_PRI_ROLE_URL` (
  `PRIVILEGE_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `URL_CODE` varchar(128) NOT NULL COMMENT '资源编码',
  `URL_TYPE` varchar(128) NOT NULL DEFAULT '' COMMENT '是否包含子资源。1 包含子资源； 0 不包含子资源。默认0',
  `ROLE_ID` varchar(64) NOT NULL COMMENT '角色编号',
  `CREATE_TIME` datetime DEFAULT NULL COMMENT '创建时间',
  `CREATE_USER` varchar(128) DEFAULT NULL COMMENT '创建者',
  `UPDATE_TIME` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `UPDATE_USER` varchar(128) DEFAULT NULL COMMENT '更新者',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效； 0 无效；默认1',
  PRIMARY KEY (`PRIVILEGE_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色URL配置表';

-- ----------------------------
-- Table structure for T_OF_PRI_USER_RES
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_PRI_USER_RES`;
CREATE TABLE `T_OF_PRI_USER_RES` (
  `PRIVILEGE_ID` int(128) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `RES_TYPE_CODE` varchar(128) NOT NULL COMMENT '资源类型编码',
  `RES_CODE` varchar(128) NOT NULL COMMENT '资源编码',
  `USER_ID` varchar(128) NOT NULL COMMENT '用户编号',
  `CONTAIN_CHILD` char(1) NOT NULL DEFAULT '0' COMMENT '是否包含子资源。1 包含子资源； 0 不包含子资源。默认0',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `CREATE_USER` varchar(128) DEFAULT NULL COMMENT '创建者',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `UPDATE_USER` varchar(128) DEFAULT NULL COMMENT '更新者',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效； 0 无效；默认1',
  PRIMARY KEY (`PRIVILEGE_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户资源配置表';

-- ----------------------------
-- Table structure for T_OF_PRI_USER_URL
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_PRI_USER_URL`;
CREATE TABLE `T_OF_PRI_USER_URL` (
  `PRIVILEGE_ID` int(128) NOT NULL AUTO_INCREMENT COMMENT '权限编号',
  `URL_CODE` varchar(128) NOT NULL COMMENT 'URL编码',
  `URL_TYPE` varchar(2) DEFAULT NULL COMMENT 'URL类型  a：api的url   p：页面的url',
  `USER_ID` varchar(128) NOT NULL COMMENT '用户编号',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `CREATE_USER` varchar(128) DEFAULT NULL COMMENT '创建者',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `UPDATE_USER` varchar(128) DEFAULT NULL COMMENT '更新者',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效； 0 无效；默认1',
  PRIMARY KEY (`PRIVILEGE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户URL配置表';

-- ----------------------------
-- Records of T_OF_PRI_USER_URL
-- ----------------------------

-- ----------------------------
-- Table structure for T_OF_REQUEST_LOG
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_REQUEST_LOG`;
CREATE TABLE `T_OF_REQUEST_LOG` (
  `LOG_ID` varchar(64) DEFAULT NULL,
  `REQUEST_URL` varchar(64) DEFAULT NULL,
  `CLIENT_IP` varchar(32) DEFAULT NULL,
  `TOKEN` varchar(64) DEFAULT NULL,
  `REQUEST_TIME` decimal(32,0) NOT NULL,
  `RESPONSE_TIME` decimal(32,0) NOT NULL,
  `SUCCESS` tinyint(1) DEFAULT NULL,
  `RESPONSE_CONTENT_LEN` decimal(10,0) DEFAULT NULL,
  `API_NAME` varchar(32) DEFAULT NULL,
  `FAULT_MESSAGE` varchar(128) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of T_OF_REQUEST_LOG
-- ----------------------------

-- ----------------------------
-- Table structure for T_OF_RES_API
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_RES_API`;
CREATE TABLE `T_OF_RES_API` (
  `RES_CODE` varchar(128) NOT NULL COMMENT '资源编码',
  `RES_NAME` varchar(128) NOT NULL COMMENT '资源名称',
  `URL` varchar(1000) DEFAULT NULL COMMENT '请求URL',
  `CREATE_TIME` datetime DEFAULT NULL COMMENT '创建时间',
  `CREATE_USER` varchar(128) DEFAULT NULL COMMENT '创建者',
  `UPDATE_TIME` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `UPDATE_USER` varchar(128) DEFAULT NULL COMMENT '更新者',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效； 0 无效；默认1',
  PRIMARY KEY (`RES_CODE`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='API资源定义表';

-- ----------------------------
-- Table structure for T_OF_RES_TYPE_4_DATA
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_RES_TYPE_4_DATA`;
CREATE TABLE `T_OF_RES_TYPE_4_DATA` (
  `RES_TYPE_CODE` varchar(128) NOT NULL COMMENT '资源类型编码',
  `RES_TYPE_NAME` varchar(128) NOT NULL COMMENT '资源类型名称',
  `RES_TYPE_PARENT` varchar(128) NOT NULL DEFAULT '-1' COMMENT '资源类型父类编码。顶级类别的PID固定为-1',
  `RES_TYPE_VALUE` varchar(4000) DEFAULT NULL COMMENT '资源类型。一般对应SQL。SQL查询',
  `PRIMARY_KEY` varchar(128) NOT NULL COMMENT '数据主键',
  `CREATE_TIME` datetime NOT NULL  COMMENT '创建时间',
  `CREATE_USER` varchar(128) DEFAULT NULL COMMENT '创建者',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `UPDATE_USER` varchar(128) DEFAULT NULL COMMENT '更新者',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效； 0 无效；默认1',
  PRIMARY KEY (`RES_TYPE_CODE`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据资源类型表';

-- ----------------------------
-- Table structure for T_OF_SYS_MENUS
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_SYS_MENUS`;
CREATE TABLE `T_OF_SYS_MENUS` (
  `MENU_ID` varchar(64) NOT NULL COMMENT '菜单编号',
  `PAGE_ID` varchar(64) DEFAULT NULL COMMENT '页面编号',
  `MENU_NAME` varchar(64) NOT NULL COMMENT '菜单名称',
  `PARENT_ID` varchar(64) NOT NULL DEFAULT '-1' COMMENT '父编号。默认为-1，表示一级模块。',
  `IN_TIME` datetime DEFAULT NULL COMMENT '生成时间',
  `ICON_CLASS` varchar(32) DEFAULT NULL COMMENT '图标样式类。为空表示不指定图标',
  `ORDER_NUM` int(11) NOT NULL DEFAULT '0' COMMENT '显示顺序',
  `IS_SHOW` char(1) NOT NULL DEFAULT '0' COMMENT '是否显示。可选值：0 不显示；1 显示',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效；0 无效。默认1',
  `UPDATE_TIME` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`MENU_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单配置表';

-- ----------------------------
-- Table structure for T_OF_SYS_PAGES
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_SYS_PAGES`;
CREATE TABLE `T_OF_SYS_PAGES` (
  `PAGE_ID` varchar(64) NOT NULL COMMENT '页面编号',
  `PAGE_NAME` varchar(64) NOT NULL COMMENT '页面名称',
  `CONTENT_URL` varchar(512) DEFAULT NULL COMMENT '内容对应的URL',
  `PARAMS` varchar(512) DEFAULT NULL COMMENT '页面参数',
  `IN_TIME` datetime DEFAULT NULL COMMENT '生成时间',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效；0 无效。默认1',
  `WIDTH` int(11) DEFAULT '1200' COMMENT '页面宽度。0表示自适应',
  `HEIGHT` int(11) DEFAULT '500' COMMENT '页面高度。0表示自适应',
  `ICON_CLASS` varchar(32) DEFAULT NULL COMMENT '图标样式类。为空表示不指定图标',
  `AUTO_FRESH_TIME` int(11) NOT NULL DEFAULT '0' COMMENT '自动刷新时间。单位：秒。默认0，表示不自动刷新',
  `UPDATE_TIME` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`PAGE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='页面配置表';

-- ----------------------------
-- Table structure for T_OF_SYS_ROLES
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_SYS_ROLES`;
CREATE TABLE `T_OF_SYS_ROLES` (
  `ROLE_ID` varchar(64) NOT NULL COMMENT '角色编号',
  `ROLE_NAME` varchar(64) NOT NULL COMMENT '角色名称',
  `ALIVE_TIME` int(11) DEFAULT NULL COMMENT '密钥有效时间（单位：秒）。小于等于0表示长效；大于0表示密钥有效时间',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效；0 无效。默认1',
  PRIMARY KEY (`ROLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色信息表';

-- ----------------------------
-- Table structure for T_OF_SYS_USERS
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_SYS_USERS`;
CREATE TABLE `T_OF_SYS_USERS` (
  `USER_ID` int(64) NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `USER_NAME` varchar(64) NOT NULL COMMENT '用户名称',
  `USER_PASSWD` varchar(128) NOT NULL COMMENT '用户密码',
  `NICK_NAME` varchar(64) NOT NULL COMMENT '昵称',
  `EMAIL` varchar(128) DEFAULT NULL COMMENT '邮箱地址',
  `PHONE_NUMBER` varchar(32) DEFAULT NULL COMMENT '手机号码',
  `CLIENT_IP` varchar(1000) DEFAULT NULL COMMENT '客户端限制',
  `NEED_CHANGE_PASSWORD` char(1) DEFAULT '1' COMMENT '是否需要修改密码。1 需要修改密码；0 不需要修改密码。默认1',
  `CREATE_TIME` datetime DEFAULT NULL COMMENT '创建时间',
  `UPDATE_TIME` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `IS_VALID` char(1) NOT NULL DEFAULT '1' COMMENT '是否有效。1 有效；0 无效。默认1',
  `ALIVE_TIME` int(11) NOT NULL,
  PRIMARY KEY (`USER_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Table structure for T_OF_U_KEY
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_U_KEY`;
CREATE TABLE `T_OF_U_KEY` (
  `USER_ID` int(11) NOT NULL COMMENT '用户编号',
  `KEY` varchar(128) NOT NULL COMMENT '密钥',
  `UPDATE_TIME` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `INVALID_TIME` datetime NOT NULL COMMENT '失效时间',
  PRIMARY KEY (`KEY`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户密钥信息表';

-- ----------------------------
-- Table structure for T_OF_U_ROLE
-- ----------------------------
DROP TABLE IF EXISTS `T_OF_U_ROLE`;
CREATE TABLE `T_OF_U_ROLE` (
  `USER_ID` int(11) NOT NULL COMMENT '用户编号',
  `ROLE_ID` varchar(64) NOT NULL COMMENT '角色编号',
  PRIMARY KEY (`ROLE_ID`,`USER_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色关联表';

-- ----------------------------
-- Table structure for T_U_CONFIG
-- ----------------------------
DROP TABLE IF EXISTS `T_U_CONFIG`;
CREATE TABLE `T_U_CONFIG` (
  `USER_ID` int(11) NOT NULL COMMENT '用户编号',
  `CONFIG_NAME` varchar(64) NOT NULL COMMENT '配置名',
  `CONFIG_VALUE` varchar(20000) DEFAULT NULL COMMENT '配置值',
  `CONFIG_TITLE` varchar(64) DEFAULT NULL COMMENT '配置说明',
  `UPDATE_TIME` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  UNIQUE KEY `UNI_U_CONFIG` (`USER_ID`,`CONFIG_NAME`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户个性化配置表';
