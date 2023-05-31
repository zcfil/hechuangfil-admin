SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 开始初始化表结构 ;
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` int(11) DEFAULT NULL COMMENT '角色ID',
  `dept_id` int(11) DEFAULT NULL COMMENT '部门ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int(11) DEFAULT NULL COMMENT '角色id',
  `menu_id` int(11) DEFAULT NULL COMMENT '权限id',
  `rule_type` char(1) DEFAULT 'p',
  `role_name` varchar(64) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `action` varchar(8) DEFAULT NULL COMMENT '请求方式 GET  POST '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `salt` varchar(255) DEFAULT NULL COMMENT '盐',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `nick_name` varchar(64) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `role_id` int(11) DEFAULT NULL COMMENT '角色id',
  `dept_id` int(11) DEFAULT NULL COMMENT '部门编码',
  `post_id` int(11) DEFAULT NULL COMMENT '职位编码',
  `status` varchar(255) DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_by` varchar(255) DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(255) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `dataScope` varchar(255) DEFAULT NULL,
  `params` varchar(255) DEFAULT NULL,
  `is_del` int(1) DEFAULT '0' COMMENT '是否删除',
  `dept_path` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `configId` int(11) NOT NULL AUTO_INCREMENT,
  `configName` varchar(128) DEFAULT NULL COMMENT '参数名称',
  `configKey` varchar(128) DEFAULT NULL COMMENT '参数键名',
  `configValue` varchar(255) DEFAULT NULL COMMENT '参数键值',
  `configType` varchar(64) DEFAULT NULL COMMENT '是否系统内置',
  `create_by` varchar(64) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(64) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `dataScope` varchar(255) DEFAULT NULL,
  `params` varchar(255) DEFAULT NULL,
  `is_del` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`configId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `deptId` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门编码',
  `parent_id` int(11) DEFAULT NULL COMMENT '上级部门',
  `dept_path` varchar(255) DEFAULT NULL COMMENT '上级路径',
  `dept_name` varchar(255) DEFAULT NULL COMMENT '部门名称',
  `sort` int(4) DEFAULT NULL COMMENT '排序',
  `leader` varchar(255) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `status` int(1) DEFAULT '0' COMMENT '状态',
  `is_del` char(1) DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `create_by` varchar(64) DEFAULT NULL,
  `update_by` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`deptId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dictCode` int(128) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dictSort` int(4) DEFAULT NULL COMMENT '显示顺序',
  `dictLabel` varchar(128) DEFAULT NULL COMMENT '数据标签',
  `dictValue` varchar(255) DEFAULT NULL COMMENT '数据键值',
  `dictType` varchar(64) DEFAULT NULL COMMENT '字典类型',
  `cssClass` varchar(255) DEFAULT NULL,
  `listClass` varchar(255) DEFAULT NULL,
  `isDefault` varchar(8) DEFAULT NULL,
  `status` varchar(8) DEFAULT NULL COMMENT '状态',
  `default` varchar(8) DEFAULT NULL,
  `create_by` varchar(64) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(64) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `params` varchar(255) DEFAULT NULL,
  `dataScope` varchar(255) DEFAULT NULL,
  `is_del` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`dictCode`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dictId` int(11) NOT NULL AUTO_INCREMENT COMMENT '字典编号',
  `dictName` varchar(128) DEFAULT NULL COMMENT '字典名称',
  `dictType` varchar(128) DEFAULT NULL COMMENT '字典类型',
  `status` varchar(8) DEFAULT NULL COMMENT '状态',
  `dataScope` varchar(255) DEFAULT NULL,
  `params` varchar(255) DEFAULT NULL,
  `create_by` varchar(64) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `is_del` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`dictId`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_loginlog`;
CREATE TABLE `sys_loginlog` (
  `infoId` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dataScope` varchar(255) DEFAULT NULL COMMENT '数据',
  `params` varchar(255) DEFAULT NULL COMMENT '参数',
  `userName` varchar(255) DEFAULT NULL COMMENT '用户名',
  `status` varchar(255) DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) DEFAULT NULL COMMENT 'ip地址',
  `loginLocation` varchar(255) DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) DEFAULT NULL COMMENT '系统',
  `loginTime` datetime DEFAULT NULL COMMENT '登录时间',
  `create_by` varchar(255) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(255) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '书签',
  `is_del` char(2) DEFAULT NULL,
  `platform` varchar(255) DEFAULT NULL COMMENT '系统版本',
  `msg` varchar(255) DEFAULT NULL COMMENT '信息',
  PRIMARY KEY (`infoId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menuId` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(64) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL COMMENT '路径',
  `paths` varchar(128) DEFAULT NULL,
  `menu_type_path` varchar(255) DEFAULT NULL,
  `action` varchar(16) DEFAULT '无' COMMENT '请求方式',
  `permission` varchar(32) DEFAULT NULL COMMENT '菜单权限标识',
  `menuType` varchar(1) DEFAULT NULL COMMENT '菜单类型0:页面；5:接口',
  `parentId` int(11) DEFAULT NULL COMMENT '父菜单ID',
  `noCache` varchar(255) DEFAULT NULL COMMENT '如果设置为true，则不会被 <keep-alive> 缓存(默认 false)',
  `breadcrumb` varchar(255) DEFAULT NULL COMMENT '如果设置为false，则不会在breadcrumb面包屑中显示',
  `menuName` varchar(255) DEFAULT NULL COMMENT '设定路由的名字，一定要填写不然使用<keep-alive>时会出现各种问题',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `component` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL COMMENT '创建人',
  `sort` int(4) NOT NULL DEFAULT '0' COMMENT '排序',
  `visible` char(1) DEFAULT NULL COMMENT '是否显示 0 显示；1 删；',
  `update_time` datetime DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL COMMENT '更新人',
  `is_del` char(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否；1 删；',
  PRIMARY KEY (`menuId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=258 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_operlog`;
CREATE TABLE `sys_operlog` (
  `operId` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `data_scope` varchar(255) DEFAULT NULL COMMENT '数据',
  `params` varchar(2048) DEFAULT NULL COMMENT '参数',
  `is_del` varchar(255) DEFAULT NULL,
  `latency_time` varchar(128) DEFAULT NULL COMMENT '耗时',
  `title` varchar(255) DEFAULT NULL COMMENT '操作模块',
  `businessType` varchar(255) DEFAULT NULL COMMENT '操作类型',
  `businessTypes` varchar(255) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL COMMENT '函数',
  `requestMethod` varchar(255) DEFAULT NULL COMMENT '请求方式',
  `operatorType` varchar(255) DEFAULT NULL COMMENT '操作类型',
  `operName` varchar(255) DEFAULT NULL COMMENT '操作者',
  `deptName` varchar(255) DEFAULT NULL COMMENT '部门名称',
  `operUrl` varchar(255) DEFAULT NULL COMMENT '访问地址',
  `operIp` varchar(255) DEFAULT NULL COMMENT '客户端ip',
  `operLocation` varchar(255) DEFAULT NULL COMMENT '访问位置',
  `operParam` varchar(2048) DEFAULT NULL COMMENT '请求参数',
  `status` varchar(255) DEFAULT NULL COMMENT '操作状态',
  `operTime` datetime DEFAULT NULL COMMENT '操作时间',
  `jsonResult` varchar(255) DEFAULT NULL COMMENT '返回参数',
  `create_by` varchar(255) DEFAULT NULL COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(255) DEFAULT NULL COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '书签',
  `user_agent` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`operId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `postId` int(11) NOT NULL AUTO_INCREMENT COMMENT '岗位编号',
  `postName` varchar(255) DEFAULT NULL COMMENT '岗位名称',
  `postCode` varchar(255) DEFAULT NULL COMMENT '岗位代码',
  `sort` int(4) DEFAULT '0' COMMENT '岗位排序',
  `status` varchar(255) DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '描述',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '最后修改时间',
  `is_del` int(1) DEFAULT '0' COMMENT '是否删除',
  `create_by` varchar(255) DEFAULT NULL,
  `update_by` varchar(255) DEFAULT NULL,
  `dataScope` varchar(255) DEFAULT NULL,
  `params` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`postId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '角色编码',
  `name` varchar(64) DEFAULT NULL COMMENT '角色名称',
  `status` varchar(255) DEFAULT NULL COMMENT '状态',
  `role_key` varchar(255) DEFAULT NULL COMMENT '角色代码',
  `role_sort` int(255) DEFAULT NULL COMMENT '角色排序',
  `is_del` varchar(255) DEFAULT NULL COMMENT '是否删除',
  `flag` varchar(255) DEFAULT NULL,
  `create_by` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_by` varchar(255) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `data_scope` varchar(255) DEFAULT NULL,
  `params` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
-- 表结构完成 ;
-- 开始初始化数据 ;
BEGIN;
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menulist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/databytype/', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUserList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser/', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/user/profile', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolelist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/configList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/config/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/configKey/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/config', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/config/', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/config/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menurole', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/roleMenuTreeselect/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menuTreeselect', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolemenu', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolemenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolemenu/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/doctor', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/doctor', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/doctor/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/doctor/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/doctor/pic/', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/deptList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/datalist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/data/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/databytype/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/data', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/data/', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/data/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/typelist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/type/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/type', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/type', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/type/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/postlist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/calendar', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/calendar', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/calendar/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/calendar/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/member', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/member', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/member/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/member/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/orders', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/orders/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/orders', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/orders/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/orders/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menuids', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/loginloglist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/loginlog/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/operloglist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/getinfo', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menulist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/databytype/', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUserList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser/', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user/profile', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolelist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/configList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menurole', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/roleMenuTreeselect/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menuTreeselect', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolemenu', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolemenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolemenu/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/doctor', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/doctor', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/doctor/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/doctor/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/doctor/pic/', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/deptList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/datalist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/databytype/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data/', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/typelist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/postlist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/calendar', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/calendar', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/calendar/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/calendar/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/member', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/member', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/member/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/member/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/orders', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/orders/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/orders', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/orders/:id', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/orders/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menuids', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/loginloglist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/loginlog/:id', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/operloglist', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/getinfo', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/roledatascope', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/roleDeptTreeselect/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/deptTree', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/configKey/*', 'GET', NULL, NULL, NULL);
COMMIT;
BEGIN;
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys_index_skinName', 'skin-blue', 'Y', '1', '2020-02-29 10:37:48', '1', '2020-03-11 00:04:51', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', '', '', '');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', '1', '2020-02-29 10:38:23', '1', '2020-03-11 00:35:28', '初始化密码 123456', '', '', '');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys_index_sideTheme', 'theme-dark', 'Y', '1', '2020-02-29 10:39:01', '1', '2020-03-11 00:08:52', '深色主题theme-dark，浅色主题theme-light', '', '', '');
COMMIT;
BEGIN;
INSERT INTO `sys_dept` VALUES (1, 0, '/0/1', '云构科技', 0, 'aituo', '13782218188', 'atuo@aituo.com', 0, '0', '2020-02-27 15:30:19', '2020-03-10 21:09:21', '1', '1');
INSERT INTO `sys_dept` VALUES (7, 1, '/0/1/7', '研发部', 0, '', '', '', 0, '0', '2020-03-08 23:10:59', '2020-03-10 21:06:30', '1', NULL);
INSERT INTO `sys_dept` VALUES (8, 1, '/0/1/8', '运维部', 0, '', '', '', 0, '0', '2020-03-08 23:11:08', '2020-03-10 16:50:27', '1', NULL);
INSERT INTO `sys_dept` VALUES (9, 1, '/0/1/9', '客服部', 0, '', '', '', 0, '0', '2020-03-08 23:11:15', '2020-03-08 23:11:15', '1', NULL);
COMMIT;
BEGIN;
INSERT INTO `sys_dict_data` VALUES (1, 0, '正常', '0', 'sys_normal_disable', '', '', '', '0', '', '1', '2020-02-28 20:55:34', '', '2020-03-08 23:11:15', '系统正常', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (2, 0, '停用', '1', 'sys_normal_disable', '', '', '', '0', '', '1', '2020-02-28 21:10:41', '', '2020-03-08 23:11:15', '系统停用', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (3, 0, '男', '0', 'sys_user_sex', '', '', '', '0', '', '1', '2020-02-28 21:37:28', '', '2020-03-08 23:11:15', '性别男', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (4, 0, '女', '1', 'sys_user_sex', '', '', '', '0', '', '1', '2020-02-28 21:37:40', '', '2020-03-08 23:11:15', '性别女', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (5, 0, '未知', '2', 'sys_user_sex', '', '', '', '0', '', '1', '2020-02-28 21:38:05', '', '2020-03-08 23:11:15', '性别未知', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (6, 0, '显示', '0', 'sys_show_hide', '', '', '', '0', '', '1', '2020-02-28 21:38:36', '', '2020-03-08 23:11:15', '显示菜单', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (7, 0, '隐藏', '1', 'sys_show_hide', '', '', '', '0', '', '1', '2020-02-28 21:38:50', '', '2020-03-08 23:11:15', '隐藏菜单', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (8, 0, '是', 'Y', 'sys_yes_no', '', '', '', '0', '', '1', '2020-02-28 21:39:40', '', '2020-03-08 23:11:15', '系统默认是', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (9, 0, '否', 'N', 'sys_yes_no', '', '', '', '0', '', '1', '2020-02-28 21:40:06', '', '2020-03-08 23:11:15', '系统默认否', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (10, 0, '正常', '0', 'sys_job_status', '', '', '', '0', '', '1', '2020-02-28 21:41:02', '', '2020-03-08 23:11:15', '正常状态', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (11, 0, '停用', '1', 'sys_job_status', '', '', '', '0', '', '1', '2020-02-28 21:41:15', '', '2020-03-08 23:11:15', '停用状态', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (12, 0, '默认', 'DEFAULT', 'sys_job_group', '', '', '', '0', '', '1', '2020-02-28 21:41:48', '', '2020-03-08 23:11:15', '默认分组', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (13, 0, '系统', 'SYSTEM', 'sys_job_group', '', '', '', '0', '', '1', '2020-02-28 21:42:02', '', '2020-03-08 23:11:15', '系统分组', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (14, 0, '通知', '1', 'sys_notice_type', '', '', '', '0', '', '1', '2020-02-28 21:42:43', '', '2020-03-08 23:11:15', '通知', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (15, 0, '公告', '2', 'sys_notice_type', '', '', '', '0', '', '1', '2020-02-28 21:42:53', '', '2020-03-08 23:11:15', '公告', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (16, 0, '正常', '0', 'sys_common_status', '', '', '', '0', '', '1', '2020-02-28 21:43:21', '', '2020-03-08 23:11:15', '正常状态', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (17, 0, '关闭', '1', 'sys_common_status', '', '', '', '0', '', '1', '2020-02-28 21:43:31', '', '2020-03-08 23:11:15', '关闭状态', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (18, 0, '新增', '1', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:44:14', '', '2020-02-28 22:00:22', '新增操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (19, 0, '修改', '2', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:44:34', '', '2020-03-08 23:11:15', '修改操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (20, 0, '删除', '3', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:44:52', '', '2020-03-08 23:11:15', '删除操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (21, 0, '授权', '4', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:45:18', '', '2020-03-08 23:11:15', '授权操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (22, 0, '导出', '5', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:45:44', '', '2020-03-08 23:11:15', '导出操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (23, 0, '导入', '6', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:46:02', '', '2020-03-08 23:11:15', '导入操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (24, 0, '强退', '7', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:46:25', '', '2020-03-08 23:11:15', '强退操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (25, 0, '生成代码', '8', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:46:53', '', '2020-03-08 23:11:15', '生成操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (26, 0, '清空数据', '9', 'sys_oper_type', '', '', '', '0', '', '1', '2020-02-28 21:47:15', '', '2020-03-08 23:11:15', '清空操作', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (27, 0, '成功', '0', 'sys_notice_status', '', '', '', '0', '', '1', '2020-02-28 22:03:24', '', '2020-03-08 23:11:15', '成功状态', '', '', '0');
INSERT INTO `sys_dict_data` VALUES (28, 0, '失败', '1', 'sys_notice_status', '', '', '', '0', '', '1', '2020-02-28 22:03:39', '', '2020-03-08 23:11:15', '失败状态', '', '', '0');
COMMIT;
BEGIN;
INSERT INTO `sys_dict_type` VALUES (1, '系统开关', 'sys_normal_disable', '0', '', '', '1', '2020-02-28 19:44:30', '', '2020-03-08 23:11:15', '系统开关列表', '');
INSERT INTO `sys_dict_type` VALUES (2, '用户性别', 'sys_user_sex', '0', '', '', '1', '2020-02-28 21:12:04', '', '2020-03-08 23:11:15', '用户性别列表', '');
INSERT INTO `sys_dict_type` VALUES (3, '菜单状态', 'sys_show_hide', '0', '', '', '1', '2020-02-28 21:13:08', '', '2020-03-08 23:11:15', '菜单状态列表', '');
INSERT INTO `sys_dict_type` VALUES (4, '系统是否', 'sys_yes_no', '0', '', '', '1', '2020-02-28 21:13:34', '', '2020-03-08 23:11:15', '系统是否列表', '');
INSERT INTO `sys_dict_type` VALUES (5, '任务状态', 'sys_job_status', '0', '', '', '1', '2020-02-28 21:13:58', '', '2020-03-08 23:11:15', '任务状态列表', '');
INSERT INTO `sys_dict_type` VALUES (6, '任务分组', 'sys_job_group', '0', '', '', '1', '2020-02-28 21:14:20', '', '2020-03-08 23:11:15', '任务分组列表', '');
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', '', '', '1', '2020-02-28 21:14:48', '', '2020-03-08 23:11:15', '通知类型列表', '');
INSERT INTO `sys_dict_type` VALUES (8, '系统状态', 'sys_common_status', '0', '', '', '1', '2020-02-28 21:15:35', '', '2020-03-08 23:11:15', '登录状态列表', '');
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', '', '', '1', '2020-02-28 21:16:00', '', '2020-03-08 23:11:15', '操作类型列表', '');
INSERT INTO `sys_dict_type` VALUES (10, '通知状态', 'sys_notice_status', '0', '', '', '1', '2020-02-28 21:16:31', '', '2020-03-08 23:11:15', '通知状态列表', '');
COMMIT;
BEGIN;
INSERT INTO `sys_menu` VALUES (2, '系统管理', '/upms', '/0/2', NULL, '无', '', 'M', 0, '1', '', 'Upms', 'example', 'Layout', '2019-11-26 21:22:09', '1', 2, '0', '2020-03-10 20:42:19', '1', '0');
INSERT INTO `sys_menu` VALUES (3, '用户管理', 'sysuser', '/0/2/3', NULL, '无', 'system:sysuser:list', 'C', 2, NULL, NULL, 'Sysuser', 'user', '/sysuser/index', '2019-09-10 13:55:17', '1', 0, '0', '2020-03-09 20:31:45', '1', '0');
INSERT INTO `sys_menu` VALUES (43, '新增用户', '/api/v1/sysuser', '/0/2/3/43', NULL, 'POST', 'system:sysuser:add', 'F', 3, NULL, NULL, NULL, NULL, NULL, '2019-11-25 10:36:34', '1', 0, '0', '2020-03-09 20:31:54', '1', '0');
INSERT INTO `sys_menu` VALUES (44, '查询用户', '/api/v1/sysuser', '/0/2/3/44', NULL, 'GET', 'system:sysuser:query', 'F', 3, NULL, NULL, NULL, NULL, NULL, '2019-11-25 10:37:02', '1', 0, '0', '2020-03-09 20:31:56', '1', '0');
INSERT INTO `sys_menu` VALUES (45, '修改用户', '/api/v1/sysuser/', '/0/2/3/45', NULL, 'PUT', 'system:sysuser:edit', 'F', 3, NULL, NULL, NULL, NULL, NULL, '2019-11-25 10:37:25', '1', 0, '0', '2020-03-09 20:31:59', '1', '0');
INSERT INTO `sys_menu` VALUES (46, '删除用户', '/api/v1/sysuser/', '/0/2/3/46', NULL, 'DELETE', 'system:sysuser:remove', 'F', 3, NULL, NULL, NULL, NULL, NULL, '2019-11-25 10:37:38', '1', 0, '0', '2020-03-09 20:32:01', '1', '0');
INSERT INTO `sys_menu` VALUES (50, '基础信息', '/mangent', '/0/50', NULL, '无', '', 'M', 0, '1', '', 'Mangent', 'network', 'Layout', '2019-11-26 23:27:47', '1', 1, '0', '2020-03-09 19:46:35', '1', '0');
INSERT INTO `sys_menu` VALUES (51, '菜单管理', 'menu', '/0/2/51', NULL, '无', 'system:sysmenu:list', 'C', 2, '1', '', 'Menu', 'tree-table', '/menu/index', '2019-11-26 23:35:47', '1', 0, '0', '2020-03-09 19:45:53', '1', '0');
INSERT INTO `sys_menu` VALUES (52, '角色管理', 'role', '/0/2/52', NULL, '无', 'system:sysrole:list', 'C', 2, '1', '', 'Role', 'peoples', '/role/index', '2019-11-26 23:40:59', '1', 0, '0', '2020-03-09 19:45:57', '1', '0');
INSERT INTO `sys_menu` VALUES (53, '医生管理', 'doctor', '/0/50/53', NULL, '无', '', 'C', 50, '1', '', 'Doctor', 'pass', '/doctor/index', '2019-11-26 23:41:40', '1', 0, '0', '2020-03-09 20:34:43', '1', '0');
INSERT INTO `sys_menu` VALUES (54, '排班管理', 'calendar', '/0/50/54', NULL, '无', '', 'C', 50, '1', '', 'Calendar', 'calendar', '/calendar/index', '2019-11-26 23:42:10', '1', 0, '0', '2020-03-09 20:34:47', '1', '0');
INSERT INTO `sys_menu` VALUES (55, '会员管理', 'member', '/0/50/55', NULL, '无', '', 'C', 50, '1', '', 'Menber', 'vip', '/member/index', '2019-11-26 23:43:01', '1', 0, '0', '2020-03-09 20:34:49', '1', '0');
INSERT INTO `sys_menu` VALUES (56, '部门管理', 'dept', '/0/2/56', NULL, '无', 'system:sysdept:list', 'C', 2, '0', '', 'Dept', 'tree', '/dept/index', '2020-02-27 10:19:49', '1', 0, '0', '2020-03-09 19:46:00', '1', '0');
INSERT INTO `sys_menu` VALUES (57, '岗位管理', 'post', '/0/2/57', NULL, '无', 'system:syspost:list', 'C', 2, '0', '', 'post', 'pass', '/post/index', '2020-02-27 21:39:02', '1', 0, '0', '2020-03-09 19:46:03', '1', '0');
INSERT INTO `sys_menu` VALUES (58, '字典管理', 'dict', '/0/2/58', NULL, '无', 'system:sysdicttype:list', 'C', 2, '0', '', 'Dict', 'education', '/dict/index', '2020-02-28 17:51:22', '1', 0, '0', '2020-03-09 19:46:06', '1', '0');
INSERT INTO `sys_menu` VALUES (59, '字典数据', 'dict/data/:dictId', '/0/2/59', NULL, '无', 'system:sysdictdata:list', 'C', 2, '0', '', 'DictData', 'education', '/dict/data', '2020-02-28 20:02:36', '1', 0, '1', '2020-03-09 19:46:24', '1', '0');
INSERT INTO `sys_menu` VALUES (60, '系统工具', '/tools', '/0/60', NULL, '无', '', 'M', 0, '0', '', 'Tools', 'component', 'Layout', '2020-02-28 23:36:21', '1', 1, '0', '2020-03-09 19:46:38', '1', '0');
INSERT INTO `sys_menu` VALUES (61, '系统接口', 'swagger', '/0/60/61', NULL, '无', '', 'C', 60, '0', '', 'Swagger', 'guide', '/tools/swagger/index', '2020-02-28 23:41:07', '1', 1, '0', '2020-03-09 20:35:02', '1', '0');
INSERT INTO `sys_menu` VALUES (62, '参数设置', '/config', '/0/2/62', NULL, '无', 'system:sysconfig:list', 'C', 2, '0', '', 'Config', 'list', '/config/index', '2020-02-29 10:32:02', '1', 1, '0', '2020-03-09 19:46:27', '1', '0');
INSERT INTO `sys_menu` VALUES (63, '接口权限', '', '/0/63', NULL, '', '', 'M', 0, '0', '', '', 'bug', '', '2020-03-03 22:08:43', '1', 1, '1', '2020-03-09 19:46:42', '1', '0');
INSERT INTO `sys_menu` VALUES (64, '用户管理', '', '/0/63/64', NULL, '', '', 'M', 63, '0', '', '', 'user', '', '2020-03-03 22:10:27', '1', 1, '1', '2020-03-09 20:35:12', '1', '0');
INSERT INTO `sys_menu` VALUES (65, '用户列表', '/api/v1/sysUserList', '', NULL, 'GET', NULL, 'A', 64, NULL, NULL, NULL, NULL, NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-06 15:03:14', '1', '1');
INSERT INTO `sys_menu` VALUES (66, '菜单管理', '', '/0/63/66', NULL, '', '', 'C', 63, '0', '', '', 'tree-table', '', '2020-03-04 23:52:22', '1', 1, '1', '2020-03-09 20:35:38', '1', '0');
INSERT INTO `sys_menu` VALUES (67, '菜单列表', '/api/v1/menulist', '/0/63/66/67', NULL, 'GET', '', 'A', 66, '0', '', '', 'tree-table', '', '2020-03-04 23:52:55', '1', 1, '1', '2020-03-09 20:35:44', '1', '0');
INSERT INTO `sys_menu` VALUES (68, '新建菜单', '/api/v1/menu', '/0/63/66/68', NULL, 'POST', '', 'A', 66, '0', '', '', 'tree', '', '2020-03-04 23:54:29', '1', 1, '1', '2020-03-09 20:35:48', '1', '0');
INSERT INTO `sys_menu` VALUES (69, '字典', '', '/0/63/69', NULL, '', '', 'M', 63, '0', '', '', 'dict', '', '2020-03-04 23:57:43', '1', 1, '1', '2020-03-09 19:47:19', '1', '0');
INSERT INTO `sys_menu` VALUES (70, '类型', '', '/0/63/69/70', NULL, '', '', 'C', 69, '0', '', '', 'dict', '', '2020-03-04 23:59:26', '1', 1, '1', '2020-03-09 20:37:22', '1', '0');
INSERT INTO `sys_menu` VALUES (71, '字典类型获取', '/api/v1/dict/databytype/', '/0/63/256/71', NULL, 'GET', '', 'A', 256, '0', '', '', 'tree', '', '2020-03-05 00:00:41', '1', 1, '1', '2020-03-10 20:33:13', '1', '0');
INSERT INTO `sys_menu` VALUES (72, '修改菜单', '/api/v1/menu', '/0/63/66/72', NULL, 'PUT', '', 'A', 66, '0', '', '', 'bug', '', '2020-03-05 00:32:09', '1', 1, '1', '2020-03-10 20:51:05', '1', '0');
INSERT INTO `sys_menu` VALUES (73, '删除菜单', '/api/v1/menu/:id', '/0/63/66/73', NULL, 'DELETE', '', 'A', 66, '0', '', '', 'bug', '', '2020-03-05 00:33:03', '1', 1, '1', '2020-03-10 20:43:09', '1', '0');
INSERT INTO `sys_menu` VALUES (74, '管理员列表', '/api/v1/sysUserList', '/0/63/64/74', NULL, 'GET', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:17', '1', '0');
INSERT INTO `sys_menu` VALUES (75, '根据id获取管理员', '/api/v1/sysUser/:id', '/0/63/64/75', NULL, 'GET', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:21', '1', '0');
INSERT INTO `sys_menu` VALUES (76, '获取管理员', '/api/v1/sysUser/', '/0/63/64/76', NULL, 'GET', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:23', '1', '0');
INSERT INTO `sys_menu` VALUES (77, '创建管理员', '/api/v1/sysUser', '/0/63/64/77', NULL, 'POST', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:26', '1', '0');
INSERT INTO `sys_menu` VALUES (78, '修改管理员', '/api/v1/sysUser', '/0/63/64/78', NULL, 'PUT', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:29', '1', '0');
INSERT INTO `sys_menu` VALUES (79, '删除管理员', '/api/v1/sysUser/:id', '/0/63/64/79', NULL, 'DELETE', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:31', '1', '0');
INSERT INTO `sys_menu` VALUES (80, '当前用户个人信息', '/api/v1/user/profile', '/0/63/64/80', NULL, 'GET', NULL, 'A', 64, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:34', '1', '0');
INSERT INTO `sys_menu` VALUES (81, '角色列表', '/api/v1/rolelist', '/0/63/201/81', NULL, 'GET', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:18', '1', '0');
INSERT INTO `sys_menu` VALUES (82, '获取角色信息', '/api/v1/role/:id', '/0/63/201/82', NULL, 'GET', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:38', '1', '0');
INSERT INTO `sys_menu` VALUES (83, '创建角色', '/api/v1/role', '/0/63/201/83', NULL, 'POST', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:44', '1', '0');
INSERT INTO `sys_menu` VALUES (84, '修改角色', '/api/v1/role', '/0/63/201/84', NULL, 'PUT', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:47', '1', '0');
INSERT INTO `sys_menu` VALUES (85, '删除角色', '/api/v1/role/:id', '/0/63/201/85', NULL, 'DELETE', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:50', '1', '0');
INSERT INTO `sys_menu` VALUES (86, '参数列表', '/api/v1/configList', '/0/63/202/86', NULL, 'GET', NULL, 'A', 202, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:06', '1', '0');
INSERT INTO `sys_menu` VALUES (87, '根据id获取参数', '/api/v1/config/:id', '/0/63/202/87', NULL, 'GET', NULL, 'A', 202, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:09', '1', '0');
INSERT INTO `sys_menu` VALUES (88, '根据key获取参数', '/api/v1/configKey/:id', '/0/63/202/88', NULL, 'GET', NULL, 'A', 202, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-10 23:37:11', '1', '1');
INSERT INTO `sys_menu` VALUES (89, '创建参数', '/api/v1/config', '/0/63/202/89', NULL, 'POST', NULL, 'A', 202, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:14', '1', '0');
INSERT INTO `sys_menu` VALUES (90, '修改参数', '/api/v1/config/', '/0/63/202/90', NULL, 'PUT', NULL, 'A', 202, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:17', '1', '0');
INSERT INTO `sys_menu` VALUES (91, '删除参数', '/api/v1/config/:id', '/0/63/202/91', NULL, 'DELETE', NULL, 'A', 202, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:20', '1', '0');
INSERT INTO `sys_menu` VALUES (92, '获取角色菜单', '/api/v1/menurole', '/0/63/201/92', NULL, 'GET', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:53', '1', '0');
INSERT INTO `sys_menu` VALUES (93, '根据角色id获取角色', '/api/v1/roleMenuTreeselect/:id', '/0/63/201/93', NULL, 'GET', NULL, 'A', 201, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:55', '1', '0');
INSERT INTO `sys_menu` VALUES (94, '获取菜单树', '/api/v1/menuTreeselect', '/0/63/205/94', NULL, 'GET', NULL, 'A', 205, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:08', '1', '0');
INSERT INTO `sys_menu` VALUES (95, '获取角色菜单', '/api/v1/rolemenu', '/0/63/205/95', NULL, 'GET', NULL, 'A', 205, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:11', '1', '0');
INSERT INTO `sys_menu` VALUES (96, '创建角色菜单', '/api/v1/rolemenu', '/0/63/205/96', NULL, 'POST', NULL, 'A', 205, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:15', '1', '0');
INSERT INTO `sys_menu` VALUES (97, '删除用户菜单数据', '/api/v1/rolemenu/:id', '/0/63/205/97', NULL, 'DELETE', NULL, 'A', 205, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:18', '1', '0');
INSERT INTO `sys_menu` VALUES (98, '医生获取', '/api/v1/doctor', '/0/63/208/98', NULL, 'GET', NULL, 'A', 208, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 1, '1', '2020-03-09 20:40:40', '1', '0');
INSERT INTO `sys_menu` VALUES (99, '创建医生', '/api/v1/doctor', '/0/63/208/99', NULL, 'POST', NULL, 'A', 208, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:43', '1', '0');
INSERT INTO `sys_menu` VALUES (100, '修改医生', '/api/v1/doctor/:id', '/0/63/208/100', NULL, 'PUT', NULL, 'A', 208, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:46', '1', '0');
INSERT INTO `sys_menu` VALUES (101, '删除医生', '/api/v1/doctor/:id', '/0/63/208/101', NULL, 'DELETE', NULL, 'A', 208, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:49', '1', '0');
INSERT INTO `sys_menu` VALUES (102, '添加医生头像', '/api/v1/doctor/pic/', '/0/63/208/102', NULL, 'POST', NULL, 'A', 208, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:52', '1', '0');
INSERT INTO `sys_menu` VALUES (103, '部门菜单列表', '/api/v1/deptList', '/0/63/203/103', NULL, 'GET', NULL, 'A', 203, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:25', '1', '0');
INSERT INTO `sys_menu` VALUES (104, '根据id获取部门信息', '/api/v1/dept/:id', '/0/63/203/104', NULL, 'GET', NULL, 'A', 203, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:27', '1', '0');
INSERT INTO `sys_menu` VALUES (105, '创建部门', '/api/v1/dept', '/0/63/203/105', NULL, 'POST', NULL, 'A', 203, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:30', '1', '0');
INSERT INTO `sys_menu` VALUES (106, '修改部门', '/api/v1/dept', '/0/63/203/106', NULL, 'PUT', NULL, 'A', 203, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:35', '1', '0');
INSERT INTO `sys_menu` VALUES (107, '删除部门', '/api/v1/dept/:id', '/0/63/203/107', NULL, 'DELETE', NULL, 'A', 203, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:37', '1', '0');
INSERT INTO `sys_menu` VALUES (108, '字典数据列表', '/api/v1/dict/datalist', '/0/63/69/206/108', NULL, 'GET', NULL, 'A', 206, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:37:57', '1', '0');
INSERT INTO `sys_menu` VALUES (109, '通过编码获取字典数据', '/api/v1/dict/data/:id', '/0/63/69/206/109', NULL, 'GET', NULL, 'A', 206, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:00', '1', '0');
INSERT INTO `sys_menu` VALUES (110, '通过字典类型获取字典数据', '/api/v1/dict/databytype/:id', '/0/63/256/110', NULL, 'GET', NULL, 'A', 256, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-10 20:32:40', '1', '0');
INSERT INTO `sys_menu` VALUES (111, '创建字典数据', '/api/v1/dict/data', '/0/63/69/206/111', NULL, 'POST', NULL, 'A', 206, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:06', '1', '0');
INSERT INTO `sys_menu` VALUES (112, '修改字典数据', '/api/v1/dict/data/', '/0/63/69/206/112', NULL, 'PUT', NULL, 'A', 206, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:08', '1', '0');
INSERT INTO `sys_menu` VALUES (113, '删除字典数据', '/api/v1/dict/data/:id', '/0/63/69/206/113', NULL, 'DELETE', NULL, 'A', 206, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:38:11', '1', '0');
INSERT INTO `sys_menu` VALUES (114, '字典类型列表', '/api/v1/dict/typelist', '/0/63/69/70/114', NULL, 'GET', NULL, 'A', 70, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:37:39', '1', '0');
INSERT INTO `sys_menu` VALUES (115, '通过id获取字典类型', '/api/v1/dict/type/:id', '/0/63/69/70/115', NULL, 'GET', NULL, 'A', 70, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:37:42', '1', '0');
INSERT INTO `sys_menu` VALUES (116, '创建字典类型', '/api/v1/dict/type', '/0/63/69/70/116', NULL, 'POST', NULL, 'A', 70, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:37:45', '1', '0');
INSERT INTO `sys_menu` VALUES (117, '修改字典类型', '/api/v1/dict/type', '/0/63/69/70/117', NULL, 'PUT', NULL, 'A', 70, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:37:48', '1', '0');
INSERT INTO `sys_menu` VALUES (118, '删除字典类型', '/api/v1/dict/type/:id', '/0/63/69/70/118', NULL, 'DELETE', NULL, 'A', 70, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:37:51', '1', '0');
INSERT INTO `sys_menu` VALUES (119, '获取岗位列表', '/api/v1/postlist', '/0/63/204/119', NULL, 'GET', NULL, 'A', 204, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:49', '1', '0');
INSERT INTO `sys_menu` VALUES (120, '通过id获取岗位信息', '/api/v1/post/:id', '/0/63/204/120', NULL, 'GET', NULL, 'A', 204, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:52', '1', '0');
INSERT INTO `sys_menu` VALUES (121, '创建岗位', '/api/v1/post', '/0/63/204/121', NULL, 'POST', NULL, 'A', 204, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:55', '1', '0');
INSERT INTO `sys_menu` VALUES (122, '修改岗位', '/api/v1/post', '/0/63/204/122', NULL, 'PUT', NULL, 'A', 204, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:39:58', '1', '0');
INSERT INTO `sys_menu` VALUES (123, '删除岗位', '/api/v1/post/:id', '/0/63/204/123', NULL, 'DELETE', NULL, 'A', 204, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:01', '1', '0');
INSERT INTO `sys_menu` VALUES (124, '获取排班信息', '/api/v1/calendar', '/0/63/210/124', NULL, 'GET', NULL, 'A', 210, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:23', '1', '0');
INSERT INTO `sys_menu` VALUES (125, '创建排班', '/api/v1/calendar', '/0/63/210/125', NULL, 'POST', NULL, 'A', 210, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:26', '1', '0');
INSERT INTO `sys_menu` VALUES (126, '修改排班', '/api/v1/calendar/:id', '/0/63/210/126', NULL, 'PUT', NULL, 'A', 210, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:29', '1', '0');
INSERT INTO `sys_menu` VALUES (127, '删除排班', '/api/v1/calendar/:id', '/0/63/210/127', NULL, 'DELETE', NULL, 'A', 210, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:31', '1', '0');
INSERT INTO `sys_menu` VALUES (128, '获取会员', '/api/v1/member', '/0/63/207/128', NULL, 'GET', NULL, 'A', 207, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:26', '1', '0');
INSERT INTO `sys_menu` VALUES (129, '创建会员', '/api/v1/member', '/0/63/207/129', NULL, 'POST', NULL, 'A', 207, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:29', '1', '0');
INSERT INTO `sys_menu` VALUES (130, '修改会员', '/api/v1/member/:id', '/0/63/207/130', NULL, 'PUT', NULL, 'A', 207, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:32', '1', '0');
INSERT INTO `sys_menu` VALUES (131, '删除会员', '/api/v1/member/:id', '/0/63/207/131', NULL, 'DELETE', NULL, 'A', 207, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:35', '1', '0');
INSERT INTO `sys_menu` VALUES (132, '获取订单', '/api/v1/orders', '/0/63/209/132', NULL, 'GET', NULL, 'A', 209, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:40:59', '1', '0');
INSERT INTO `sys_menu` VALUES (133, '通过id获取订单', '/api/v1/orders/:id', '/0/63/209/133', NULL, 'GET', NULL, 'A', 209, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:02', '1', '0');
INSERT INTO `sys_menu` VALUES (134, '新建订单', '/api/v1/orders', '/0/63/209/134', NULL, 'POST', NULL, 'A', 209, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:06', '1', '0');
INSERT INTO `sys_menu` VALUES (135, '修改订单', '/api/v1/orders/:id', '/0/63/209/135', NULL, 'PUT', NULL, 'A', 209, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:13', '1', '0');
INSERT INTO `sys_menu` VALUES (136, '删除订单', '/api/v1/orders/:id', '/0/63/209/136', NULL, 'DELETE', NULL, 'A', 209, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:17', '1', '0');
INSERT INTO `sys_menu` VALUES (137, '菜单列表', '/api/v1/menulist', '/0/63/66/137', NULL, 'GET', NULL, 'A', 66, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:35:57', '1', '1');
INSERT INTO `sys_menu` VALUES (138, '获取根据id菜单信息', '/api/v1/menu/:id', '/0/63/66/138', NULL, 'GET', NULL, 'A', 66, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:36:16', '1', '0');
INSERT INTO `sys_menu` VALUES (139, '创建菜单', '/api/v1/menu', '/0/63/66/139', NULL, 'POST', NULL, 'A', 66, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:36:19', '1', '1');
INSERT INTO `sys_menu` VALUES (140, '修改菜单', '/api/v1/menu/:id', '/0/63/66/140', NULL, 'PUT', NULL, 'A', 66, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-10 20:46:20', '1', '1');
INSERT INTO `sys_menu` VALUES (141, '删除菜单', '/api/v1/menu/:id', '', NULL, 'DELETE', NULL, 'A', 66, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 19:14:55', '1', '1');
INSERT INTO `sys_menu` VALUES (142, '获取角色对应的菜单id数组', '/api/v1/menuids', '/0/63/256/142', NULL, 'GET', NULL, 'A', 256, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-10 20:33:35', '1', '0');
INSERT INTO `sys_menu` VALUES (201, '角色管理', '', '/0/63/201', NULL, 'GET', '', 'C', 63, '0', '', '', 'peoples', '', '2020-03-06 18:53:19', '1', 1, '1', '2020-03-09 19:47:22', '1', '0');
INSERT INTO `sys_menu` VALUES (202, '参数设置', '', '/0/63/202', NULL, 'DELETE', '', 'C', 63, '0', '', '', 'list', '', '2020-03-06 18:56:13', '1', 1, '1', '2020-03-09 20:39:00', '1', '0');
INSERT INTO `sys_menu` VALUES (203, '部门管理', '', '/0/63/203', NULL, 'POST', '', 'C', 63, '0', '', '', 'tree', '', '2020-03-06 19:36:52', '1', 1, '1', '2020-03-09 19:47:34', '1', '0');
INSERT INTO `sys_menu` VALUES (204, '岗位管理', '', '/0/63/204', NULL, '', '', 'C', 63, '0', '', '', 'pass', '', '2020-03-06 19:37:10', '1', 1, '1', '2020-03-09 19:47:41', '1', '0');
INSERT INTO `sys_menu` VALUES (205, '角色菜单管理', '', '/0/63/205', NULL, '', '', 'C', 63, '0', '', '', 'nested', '', '2020-03-06 19:37:35', '1', 1, '1', '2020-03-09 19:47:44', '1', '0');
INSERT INTO `sys_menu` VALUES (206, '数据', '', '/0/63/69/206', NULL, 'PUT', '', 'C', 69, '0', '', '', '', '', '2020-03-06 19:43:15', '1', 2, '1', '2020-03-09 20:09:37', '1', '0');
INSERT INTO `sys_menu` VALUES (207, '会员管理', '', '/0/63/207', NULL, 'DELETE', '', 'C', 63, '0', '', '', 'vip', '', '2020-03-06 19:51:59', '1', 1, '1', '2020-03-09 19:47:48', '1', '0');
INSERT INTO `sys_menu` VALUES (208, '医生管理', '', '/0/63/208', NULL, '', '', 'C', 63, '0', '', '', 'theme', '', '2020-03-06 19:52:35', '1', 1, '1', '2020-03-09 19:47:52', '1', '0');
INSERT INTO `sys_menu` VALUES (209, '订单管理', '', '/0/63/209', NULL, '', '', 'M', 63, '0', '', '', 'tab', '', '2020-03-06 19:56:47', '1', 1, '1', '2020-03-09 19:47:55', '1', '0');
INSERT INTO `sys_menu` VALUES (210, '排班管理', '', '/0/63/210', NULL, 'DELETE', '', 'C', 63, '0', '', '', 'calendar', '', '2020-03-06 19:58:35', '1', 1, '1', '2020-03-09 19:47:59', '1', '0');
INSERT INTO `sys_menu` VALUES (211, '日志管理', '/log', '/0/2/211', NULL, '', '', 'M', 2, '0', '', 'Log', 'log', '/log/index', '2020-03-06 21:50:57', '1', 1, '0', '2020-03-09 19:46:31', '1', '0');
INSERT INTO `sys_menu` VALUES (212, '登录日志', '/loginlog', '/0/2/211/212', NULL, '', 'system:sysloginlog:list', 'C', 211, '0', '', 'LoginLog', 'logininfor', '/loginlog/index', '2020-03-06 21:53:03', '1', 1, '0', '2020-03-09 20:34:11', '1', '0');
INSERT INTO `sys_menu` VALUES (213, '获取登录日志', '/api/v1/loginloglist', '/0/63/214/213', NULL, 'GET', NULL, 'A', 214, NULL, NULL, NULL, 'bug', NULL, '2020-03-09 20:43:15', '1', 0, '1', '2020-03-09 20:41:37', '1', '0');
INSERT INTO `sys_menu` VALUES (214, '日志管理', '', '/0/63/214', NULL, 'GET', '', 'M', 63, '0', '', '', 'log', '', '2020-03-06 23:12:37', '1', 1, '1', '2020-03-09 19:48:07', '1', '0');
INSERT INTO `sys_menu` VALUES (215, '删除日志', '/api/v1/loginlog/:id', '/0/63/214/215', NULL, 'DELETE', '', 'A', 214, '0', '', '', 'bug', '', '2020-03-07 01:09:21', '1', 1, '1', '2020-03-09 20:41:39', '1', '0');
INSERT INTO `sys_menu` VALUES (216, '操作日志', '/operlog', '/0/2/211/216', NULL, '', 'system:sysoperlog:list', 'C', 211, '0', '', 'OperLog', 'skill', '/operlog/index', '2020-03-08 00:56:43', '1', 1, '0', '2020-03-09 20:34:30', '1', '0');
INSERT INTO `sys_menu` VALUES (217, '获取操作日志', '/api/v1/operloglist', '/0/63/214/217', NULL, 'GET', '', 'A', 214, '0', '', '', 'bug', '', '2020-03-08 00:59:41', '1', 1, '1', '2020-03-09 20:41:42', '1', '0');
INSERT INTO `sys_menu` VALUES (218, '日历', 'calendar', '/0/60/218', NULL, '', '', 'C', 60, '0', '', 'Calendars', 'calendar', '/calendar/index', '2020-03-08 16:09:53', '1', 1, '0', '2020-03-10 21:30:57', '1', '0');
INSERT INTO `sys_menu` VALUES (219, 'Excel导入', '/excel', '/0/50/219', NULL, '', '', 'C', 50, '0', '', 'Excel', 'excel', '/excel/upload-excel', '2020-03-08 17:55:46', '1', 1, '0', '2020-03-09 20:34:52', '1', '0');
INSERT INTO `sys_menu` VALUES (220, '新增菜单', '', '/0/2/51/220', NULL, '', 'system:sysmenu:add', 'F', 51, '0', '', '', '', '', '2020-03-08 18:53:36', '1', 1, '0', '2020-03-09 20:32:08', '1', '0');
INSERT INTO `sys_menu` VALUES (221, '修改菜单', '', '/0/2/51/221', NULL, '', 'system:sysmenu:edit', 'F', 51, '0', '', '', 'edit', '', '2020-03-08 18:54:04', '1', 1, '0', '2020-03-09 20:32:11', '1', '0');
INSERT INTO `sys_menu` VALUES (222, '查询菜单', '', '/0/2/51/222', NULL, '', 'system:sysmenu:query', 'F', 51, '0', '', '', 'search', '', '2020-03-08 18:56:47', '1', 1, '0', '2020-03-09 20:32:15', '1', '0');
INSERT INTO `sys_menu` VALUES (223, '删除菜单', '', '/0/2/51/223', NULL, '', 'system:sysmenu:remove', 'F', 51, '0', '', '', '', '', '2020-03-08 18:57:14', '1', 1, '0', '2020-03-09 20:32:18', '1', '0');
INSERT INTO `sys_menu` VALUES (224, '新增角色', '', '/0/2/52/224', NULL, '', 'system:sysrole:add', 'F', 52, '0', '', '', '', '', '2020-03-08 19:54:18', '1', 1, '0', '2020-03-09 20:32:23', '1', '0');
INSERT INTO `sys_menu` VALUES (225, '查询角色', '', '/0/2/52/225', NULL, '', 'system:sysrole:query', 'F', 52, '0', '', '', '', '', '2020-03-08 19:54:46', '1', 1, '0', '2020-03-09 20:32:26', '1', '0');
INSERT INTO `sys_menu` VALUES (226, '修改角色', '', '/0/2/52/226', NULL, '', 'system:sysrole:edit', 'F', 52, '0', '', '', '', '', '2020-03-08 19:55:19', '1', 1, '0', '2020-03-09 20:32:29', '1', '0');
INSERT INTO `sys_menu` VALUES (227, '删除角色', '', '/0/2/52/227', NULL, '', 'system:sysrole:remove', 'F', 52, '0', '', '', '', '', '2020-03-08 19:55:44', '1', 1, '0', '2020-03-09 20:32:32', '1', '0');
INSERT INTO `sys_menu` VALUES (228, '查询部门', '', '/0/2/56/228', NULL, '', 'system:sysdept:query', 'F', 56, '0', '', '', '', '', '2020-03-08 19:56:23', '1', 1, '0', '2020-03-09 20:32:57', '1', '0');
INSERT INTO `sys_menu` VALUES (229, '新增部门', '', '/0/2/56/229', NULL, '', 'system:sysdept:add', 'F', 56, '0', '', '', '', '', '2020-03-08 19:56:43', '1', 1, '0', '2020-03-09 20:33:00', '1', '0');
INSERT INTO `sys_menu` VALUES (230, '修改部门', '', '/0/2/56/230', NULL, '', 'system:sysdept:edit', 'F', 56, '0', '', '', '', '', '2020-03-08 19:58:21', '1', 0, '0', '2020-03-09 20:33:04', '1', '0');
INSERT INTO `sys_menu` VALUES (231, '删除部门', '', '/0/2/56/231', NULL, '', 'system:sysdept:remove', 'F', 56, '0', '', '', '', '', '2020-03-08 19:58:35', '1', 0, '0', '2020-03-09 20:33:07', '1', '0');
INSERT INTO `sys_menu` VALUES (232, '查询岗位', '', '/0/2/57/232', NULL, '', 'system:syspost:query', 'F', 57, '0', '', '', '', '', '2020-03-08 19:59:10', '1', 0, '0', '2020-03-09 20:33:13', '1', '0');
INSERT INTO `sys_menu` VALUES (233, '新增岗位', '', '/0/2/57/233', NULL, '', 'system:syspost:add', 'F', 57, '0', '', '', '', '', '2020-03-08 19:59:26', '1', 0, '0', '2020-03-09 20:33:15', '1', '0');
INSERT INTO `sys_menu` VALUES (234, '修改岗位', '', '/0/2/57/234', NULL, '', 'system:syspost:edit', 'F', 57, '0', '', '', '', '', '2020-03-08 19:59:45', '1', 0, '0', '2020-03-09 20:33:18', '1', '0');
INSERT INTO `sys_menu` VALUES (235, '删除岗位', '', '/0/2/57/235', NULL, '', 'system:syspost:remove', 'F', 57, '0', '', '', '', '', '2020-03-08 19:59:59', '1', 0, '0', '2020-03-09 20:33:21', '1', '0');
INSERT INTO `sys_menu` VALUES (236, '字典查询', '', '/0/2/58/236', NULL, '', 'system:sysdicttype:query', 'F', 58, '0', '', '', '', '', '2020-03-08 20:01:14', '1', 0, '0', '2020-03-09 20:33:27', '1', '0');
INSERT INTO `sys_menu` VALUES (237, '新增类型', '', '/0/2/58/237', NULL, '', 'system:sysdicttype:add', 'F', 58, '0', '', '', '', '', '2020-03-08 20:01:51', '1', 0, '0', '2020-03-09 20:33:30', '1', '0');
INSERT INTO `sys_menu` VALUES (238, '修改类型', '', '/0/2/58/238', NULL, '', 'system:sysdicttype:edit', 'F', 58, '0', '', '', '', '', '2020-03-08 20:02:07', '1', 0, '0', '2020-03-09 20:33:32', '1', '0');
INSERT INTO `sys_menu` VALUES (239, '删除类型', '', '/0/2/58/239', NULL, '', 'system:sysdicttype:remove', 'F', 58, '0', '', '', '', '', '2020-03-08 20:02:29', '1', 0, '0', '2020-03-09 20:33:35', '1', '0');
INSERT INTO `sys_menu` VALUES (240, '查询数据', '', '/0/2/59/240', NULL, '', 'system:sysdictdata:query', 'F', 59, '0', '', '', '', '', '2020-03-08 20:03:24', '1', 0, '0', '2020-03-09 20:33:40', '1', '0');
INSERT INTO `sys_menu` VALUES (241, '新增数据', '', '/0/2/59/241', NULL, '', 'system:sysdictdata:add', 'F', 59, '0', '', '', '', '', '2020-03-08 20:04:07', '1', 0, '0', '2020-03-09 20:33:43', '1', '0');
INSERT INTO `sys_menu` VALUES (242, '修改数据', '', '/0/2/59/242', NULL, '', 'system:sysdictdata:edit', 'F', 59, '0', '', '', '', '', '2020-03-08 20:04:19', '1', 0, '0', '2020-03-09 20:33:45', '1', '0');
INSERT INTO `sys_menu` VALUES (243, '删除数据', '', '/0/2/59/243', NULL, '', 'system:sysdictdata:remove', 'F', 59, '0', '', '', '', '', '2020-03-08 20:04:36', '1', 0, '0', '2020-03-09 20:33:48', '1', '0');
INSERT INTO `sys_menu` VALUES (244, '查询参数', '', '/0/2/62/244', NULL, '', 'system:sysconfig:query', 'F', 62, '0', '', '', '', '', '2020-03-08 20:05:19', '1', 0, '0', '2020-03-09 20:33:55', '1', '0');
INSERT INTO `sys_menu` VALUES (245, '新增参数', '', '/0/2/62/245', NULL, '', 'system:sysconfig:add', 'F', 62, '0', '', '', '', '', '2020-03-08 20:05:35', '1', 0, '0', '2020-03-09 20:33:59', '1', '0');
INSERT INTO `sys_menu` VALUES (246, '修改参数', '', '/0/2/62/246', NULL, '', 'system:sysconfig:edit', 'F', 62, '0', '', '', '', '', '2020-03-08 20:05:49', '1', 0, '0', '2020-03-09 20:34:02', '1', '0');
INSERT INTO `sys_menu` VALUES (247, '删除参数', '', '/0/2/62/247', NULL, '', 'system:sysconfig:remove', 'F', 62, '0', '', '', '', '', '2020-03-08 20:06:04', '1', 0, '0', '2020-03-09 20:34:05', '1', '0');
INSERT INTO `sys_menu` VALUES (248, '查询登录日志', '', '/0/2/211/212/248', NULL, '', 'system:sysloginlog:query', 'F', 212, '0', '', '', '', '', '2020-03-08 20:07:28', '1', 0, '0', '2020-03-09 20:34:16', '1', '0');
INSERT INTO `sys_menu` VALUES (249, '删除登录日志', '', '/0/2/211/212/249', NULL, '', 'system:sysloginlog:remove', 'F', 212, '0', '', '', '', '', '2020-03-08 20:08:18', '1', 0, '0', '2020-03-09 20:34:19', '1', '0');
INSERT INTO `sys_menu` VALUES (250, '查询操作日志', '', '/0/2/211/216/250', NULL, '', 'system:sysoperlog:query', 'F', 216, '0', '', '', '', '', '2020-03-08 20:09:51', '1', 0, '0', '2020-03-09 20:34:33', '1', '0');
INSERT INTO `sys_menu` VALUES (251, '删除操作日志', '', '/0/2/211/216/251', NULL, '', 'system:sysoperlog:remove', 'F', 216, '0', '', '', '', '', '2020-03-08 20:10:08', '1', 0, '0', '2020-03-09 20:34:36', '1', '0');
INSERT INTO `sys_menu` VALUES (252, '获取登录用户信息', '/api/v1/getinfo', '/0/63/256/252', NULL, 'GET', '', 'A', 256, '0', '', '', '', '', '2020-03-08 20:47:56', '1', 0, '1', '2020-03-10 20:31:15', '1', '');
INSERT INTO `sys_menu` VALUES (253, '角色数据权限', '/api/v1/roledatascope', '/0/63/201/253', NULL, 'PUT', '', 'A', 201, '0', '', '', '', '', '2020-03-09 23:37:10', '1', 0, '1', '2020-03-10 20:32:07', '1', '');
INSERT INTO `sys_menu` VALUES (254, '部门树接口【数据权限】', '/api/v1/roleDeptTreeselect/:id', '/0/63/256/254', NULL, 'GET', '', 'A', 256, '0', '', '', '', '', '2020-03-09 23:38:36', '1', 0, '1', '2020-03-10 20:31:33', '1', '');
INSERT INTO `sys_menu` VALUES (255, '部门树【用户列表】', '/api/v1/deptTree', '/0/63/256/255', NULL, 'GET', '', 'A', 256, '0', '', '', '', '', '2020-03-10 20:30:18', '1', 0, '1', '2020-03-10 20:52:37', '1', '');
INSERT INTO `sys_menu` VALUES (256, '必开接口', '', '/0/63/256', NULL, 'GET', '', 'M', 63, '0', '', '', '', '', '2020-03-10 20:31:00', '1', 0, '1', '2020-03-08 23:11:15', '', '');
INSERT INTO `sys_menu` VALUES (257, '通过key获取参数', '/api/v1/configKey/:id', '/0/63/256/257', NULL, 'GET', '', 'A', 256, '0', '', '', 'bug', '', '2020-03-10 23:15:59', '1', 1, '1', '2020-03-10 23:37:32', '1', '0');
COMMIT;
BEGIN;
INSERT INTO `sys_post` VALUES (1, '首席执行官', 'CEO', 0, '0', '首席执行官', '2020-02-27 21:45:58', '2020-03-08 23:11:15', 0, '1', '2020-03-08 23:11:15', '', '');
INSERT INTO `sys_post` VALUES (2, '开发工程师', 'Development ', 1, '0', '开发工程师', '2020-02-29 18:49:22', '2020-03-08 23:11:15', 0, '1', '2020-03-08 23:11:15', '', '');
INSERT INTO `sys_post` VALUES (3, '测试工程师', 'Test', 2, '0', '测试工程师', '2020-02-29 18:50:04', '2020-03-08 23:11:15', 0, '1', '2020-03-08 23:11:15', '', '');
INSERT INTO `sys_post` VALUES (4, '产品经理', 'Product', 1, '0', '产品经理', '2020-02-29 18:50:31', '2020-03-08 23:11:15', 0, '1', '2020-03-08 23:11:15', '', '');
INSERT INTO `sys_post` VALUES (5, '运维工程师', 'Opetion&Maintenance', 1, '0', '', '2020-02-29 18:53:00', '2020-03-06 23:45:48', 0, '1', '2020-03-08 23:11:15', '', '');
COMMIT;
BEGIN;
INSERT INTO `sys_role` VALUES (1, '系统管理员', '0', 'admin', 1, '0', NULL, '1', '2020-03-09 21:21:54', NULL, '2020-03-10 23:38:09', NULL, '3', NULL, 0);
INSERT INTO `sys_role` VALUES (2, '普通角色', '0', 'common', 0, '0', NULL, '1', '2020-03-09 21:21:54', NULL, '2020-03-10 20:12:00', NULL, '2', NULL, 0);
COMMIT;
BEGIN;
INSERT INTO `sys_role_menu` VALUES (1, 2, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 3, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 43, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 44, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 45, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 46, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 50, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 51, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 52, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 56, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 57, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 58, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 59, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 60, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 61, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 62, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 63, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 64, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 66, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 67, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 68, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 69, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 70, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 71, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 72, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 73, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 74, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 75, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 76, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 77, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 78, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 79, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 80, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 81, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 82, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 83, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 84, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 85, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 86, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 87, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 89, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 90, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 91, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 92, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 93, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 94, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 95, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 96, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 97, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 98, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 99, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 100, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 101, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 102, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 103, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 104, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 105, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 106, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 107, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 108, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 109, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 110, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 111, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 112, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 113, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 114, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 115, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 116, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 117, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 118, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 119, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 120, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 121, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 122, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 123, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 124, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 125, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 126, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 127, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 128, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 129, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 130, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 131, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 132, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 133, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 134, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 135, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 136, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 138, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 142, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 201, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 202, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 203, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 204, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 205, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 206, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 207, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 208, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 209, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 210, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 211, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 212, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 213, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 214, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 215, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 216, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 217, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 218, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 219, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 220, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 221, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 222, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 223, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 224, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 225, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 226, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 227, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 228, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 229, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 230, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 231, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 232, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 233, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 234, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 235, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 236, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 237, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 238, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 239, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 240, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 241, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 242, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 243, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 244, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 245, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 246, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 247, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 248, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 249, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 250, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 251, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 252, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 253, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 254, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 255, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 256, 'p', 'admin', NULL, NULL);
INSERT INTO `sys_role_menu` VALUES (1, 257, 'p', 'admin', NULL, NULL);
COMMIT;
BEGIN;
INSERT INTO `sys_user` VALUES (1, 'admin', NULL, '$2a$10$fWkAKmNCntimDhPhHoa9luB6qGCQKVi4S73OVJVd1UujiQVK2Q9qy', 'zhangwj', '13818888888', NULL, '1@1.com', 1, 1, 2, '0', NULL, '1', '0', '2019-11-10 14:05:55', '1', '2020-03-11 20:26:15', NULL, NULL, 0, NULL);
INSERT INTO `sys_user` VALUES (2, 'zhangwj', NULL, '$2a$10$CqMwHahA3cNrNv16CoSxmeD4XMPU.BiKHPEAeaG5oXMavOKrjInXi', 'zhangwj', '13211111111', NULL, 'q@q.com', 3, 8, 2, '0', NULL, '1', '0', '2019-11-12 18:28:27', '1', '2020-03-11 20:27:39', NULL, NULL, 0, NULL);
COMMIT;
SET FOREIGN_KEY_CHECKS = 1;
-- 数据完成 ;
