-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql
-- 生成日期： 2022-08-23 03:15:55
-- 服务器版本： 5.7.36
-- PHP 版本： 7.4.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `ark_admin`
--

-- --------------------------------------------------------

--
-- 表的结构 `sys_dept`
--

CREATE TABLE `sys_dept` (
  `id` int(11) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL COMMENT '父级id',
  `name` varchar(25) NOT NULL COMMENT '部门简称',
  `full_name` varchar(50) NOT NULL COMMENT '部门全称',
  `unique_key` varchar(25) NOT NULL COMMENT '唯一值',
  `type` tinyint(1) NOT NULL COMMENT '1=公司 2=子公司 3=部门',
  `status` tinyint(1) NOT NULL COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门';

--
-- 转存表中的数据 `sys_dept`
--

INSERT INTO `sys_dept` (`id`, `parent_id`, `name`, `full_name`, `unique_key`, `type`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 0, '方舟', '方舟互联', 'arklnk', 1, 1, 0, '', '2022-08-17 02:09:17', '2022-08-22 02:13:54'),
(2, 0, '思忆', '思忆技术', 'siyee', 1, 1, 0, '', '2022-08-19 06:40:10', '2022-08-22 02:13:39');

-- --------------------------------------------------------

--
-- 表的结构 `sys_dictionary`
--

CREATE TABLE `sys_dictionary` (
  `id` int(11) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '0=配置集 !0=父级id',
  `name` varchar(25) NOT NULL COMMENT '名称',
  `type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0字典 1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单文件 10多文件	',
  `unique_key` varchar(25) NOT NULL COMMENT '唯一值',
  `value` varchar(2048) NOT NULL COMMENT '配置值',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统参数';

--
-- 转存表中的数据 `sys_dictionary`
--

INSERT INTO `sys_dictionary` (`id`, `parent_id`, `name`, `type`, `unique_key`, `value`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 0, '系统配置', 0, 'sys', '', 1, 0, '', '2022-08-22 10:03:58', '2022-08-23 01:25:31'),
(2, 1, '站点名称', 1, 'sys_title', '方舟管理后台', 1, 0, '', '2022-08-22 10:03:58', '2022-08-23 03:02:15'),
(3, 1, '站点logo', 1, 'sys_logo', '', 1, 0, '', '2022-08-22 10:03:58', '2022-08-23 03:02:19'),
(4, 1, '系统语言', 4, 'sys_language', 'cn', 1, 0, '', '2022-08-22 10:03:58', '2022-08-23 03:02:27'),
(5, 1, '默认密码', 1, 'sys_pwd', '123456', 1, 0, '', '2022-08-22 10:03:58', '2022-08-23 03:02:31');

-- --------------------------------------------------------

--
-- 表的结构 `sys_job`
--

CREATE TABLE `sys_job` (
  `id` int(11) NOT NULL COMMENT '编号',
  `name` varchar(50) NOT NULL COMMENT '岗位名称',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启 ',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '开启时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工作岗位';

--
-- 转存表中的数据 `sys_job`
--

INSERT INTO `sys_job` (`id`, `name`, `status`, `order_num`, `create_time`, `update_time`) VALUES
(1, '前端', 1, 0, '2022-08-17 03:15:56', '2022-08-17 05:27:26'),
(2, '后端', 1, 0, '2022-08-17 03:15:56', '2022-08-17 05:32:50'),
(3, '设计', 1, 0, '2022-08-17 03:15:56', '2022-08-17 05:32:55');

-- --------------------------------------------------------

--
-- 表的结构 `sys_log_action`
--

CREATE TABLE `sys_log_action` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '编号',
  `user_id` int(11) NOT NULL COMMENT '操作账号',
  `ip` varchar(100) NOT NULL COMMENT 'ip',
  `os` varchar(50) NOT NULL COMMENT '系统',
  `browser` varchar(50) NOT NULL COMMENT '浏览器',
  `uri` varchar(200) NOT NULL COMMENT '请求路径',
  `request` json NOT NULL COMMENT '请求数据',
  `response` json NOT NULL COMMENT '响应数据',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志';

-- --------------------------------------------------------

--
-- 表的结构 `sys_log_login`
--

CREATE TABLE `sys_log_login` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '编号',
  `user_id` int(11) NOT NULL COMMENT '登录账号',
  `ip` varchar(100) NOT NULL COMMENT 'ip',
  `os` varchar(50) NOT NULL COMMENT '操作系统',
  `browser` varchar(50) NOT NULL COMMENT '浏览器',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='登录日志';

-- --------------------------------------------------------

--
-- 表的结构 `sys_perm_menu`
--

CREATE TABLE `sys_perm_menu` (
  `id` int(11) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `router` varchar(255) NOT NULL DEFAULT '' COMMENT '路由',
  `perms` varchar(255) NOT NULL DEFAULT '' COMMENT '权限',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0=目录 1=菜单 2=权限',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `order_num` int(11) DEFAULT '0' COMMENT '排序值',
  `view_path` varchar(255) NOT NULL DEFAULT '' COMMENT '页面路径',
  `is_show` tinyint(1) DEFAULT '1' COMMENT '0=隐藏 1=显示',
  `active_router` varchar(255) NOT NULL DEFAULT '' COMMENT '当前激活的菜单',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限&菜单';

--
-- 转存表中的数据 `sys_perm_menu`
--

INSERT INTO `sys_perm_menu` (`id`, `parent_id`, `name`, `router`, `perms`, `type`, `icon`, `order_num`, `view_path`, `is_show`, `active_router`, `create_time`, `update_time`) VALUES
(1, 0, 'routes.system.name', '/sys', '[]', 0, 'system', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:51:41'),
(2, 1, 'routes.system.menu.name', '/sys/menu', '[]', 1, '', 0, 'views/system/menu', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:27:53'),
(3, 2, 'common.basic.query', '', '[\"sys/perm/menu/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 04:55:04'),
(4, 2, 'common.basic.add', '', '[\"sys/perm/menu/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 04:55:06'),
(5, 2, 'common.basic.delete', '', '[\"sys/perm/menu/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 04:55:08'),
(6, 2, 'common.basic.update', '', '[\"sys/perm/menu/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 04:55:10'),
(7, 1, 'routes.system.role.name', '/sys/role', '[]', 1, '', 0, 'views/system/role', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:28:15'),
(8, 7, 'common.basic.query', '', '[\"sys/role/list\",\"sys/perm/menu/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 04:55:36'),
(9, 7, 'common.basic.add', '', '[\"sys/role/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:18:12'),
(10, 7, 'common.basic.delete', '', '[\"sys/role/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:18:54'),
(11, 7, 'common.basic.update', '', '[\"sys/role/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:19:15'),
(13, 1, 'routes.system.department.name', '/sys/dept', '[]', 1, '', 0, 'views/system/dept', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:28:33'),
(14, 13, 'common.basic.query', '', '[\"sys/dept/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:19:42'),
(15, 13, 'common.basic.add', '', '[\"sys/dept/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:19:53'),
(16, 13, 'common.basic.delete', '', '[\"sys/dept/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:20:04'),
(17, 13, 'common.basic.update', '', '[\"sys/dept/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:20:13'),
(18, 1, 'routes.system.job.name', '/sys/job', '[]', 1, '', 0, 'views/system/job', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:28:47'),
(19, 18, 'common.basic.query', '', '[\"sys/job/page\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:22:57'),
(20, 18, 'common.basic.add', '', '[\"sys/job/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:20:49'),
(21, 18, 'common.basic.delete', '', '[\"sys/job/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:20:59'),
(22, 18, 'common.basic.update', '', '[\"sys/job/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:21:12'),
(24, 1, 'routes.system.profession.name', '/sys/profession', '[]', 1, '', 0, 'views/system/profession', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:29:07'),
(25, 24, 'common.basic.query', '', '[\"sys/profession/page\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:23:08'),
(26, 24, 'common.basic.add', '', '[\"sys/profession/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:21:50'),
(27, 24, 'common.basic.delete', '', '[\"sys/profession/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:22:02'),
(28, 24, 'common.basic.update', '', '[\"sys/profession/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:22:12'),
(30, 1, 'routes.system.user.name', '/sys/user', '[]', 1, '', 0, 'views/system/user', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:29:23'),
(31, 30, 'common.basic.query', '', '[\"sys/user/list\",\"sys/dept/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 08:01:27'),
(32, 30, 'common.basic.add', '', '[\"sys/user/add\",\"sys/dept/list\",\"sys/job/list\",\"sys/role/list\",\"sys/profession/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:59:16'),
(33, 30, 'common.basic.delete', '', '[\"sys/user/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:58:21'),
(34, 30, 'common.basic.update', '', '[\"sys/user/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:58:12'),
(35, 30, 'permission.changepasswd', '', '[\"sys/user/password/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:58:01'),
(36, 30, 'permission.transfer', '', '[\"sys/user/transfer\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-22 07:58:46'),
(37, 0, 'routes.config.name', '/config', '[]', 0, 'documentation', 0, '', 1, '', '2022-08-22 03:33:42', '2022-08-23 02:54:22'),
(38, 37, 'routes.config.dictionary.name', '/config/dict', '[]', 1, '', 0, 'views/config/dict', 1, '', '2022-08-22 03:39:21', '2022-08-23 02:54:50'),
(39, 38, 'common.basic.query', '', '[\"config/dict/list\",\"config/dict/data/page\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 02:35:06'),
(40, 38, 'common.basic.add', '', '[\"config/dict/add\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 02:36:06'),
(41, 38, 'common.basic.delete', '', '[\"config/dict/delete\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 02:36:17'),
(42, 38, 'common.basic.update', '', '[\"config/dict/update\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 02:36:29');

-- --------------------------------------------------------

--
-- 表的结构 `sys_profession`
--

CREATE TABLE `sys_profession` (
  `id` int(11) NOT NULL COMMENT '编号',
  `name` varchar(50) NOT NULL COMMENT '职称',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职称';

--
-- 转存表中的数据 `sys_profession`
--

INSERT INTO `sys_profession` (`id`, `name`, `status`, `order_num`, `create_time`, `update_time`) VALUES
(1, 'CEO', 1, 0, '2022-08-17 05:09:26', '2022-08-17 05:09:26'),
(2, 'CTO', 1, 0, '2022-08-17 05:09:26', '2022-08-17 05:09:26');

-- --------------------------------------------------------

--
-- 表的结构 `sys_role`
--

CREATE TABLE `sys_role` (
  `id` int(10) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(25) NOT NULL COMMENT '名称',
  `unique_key` varchar(25) NOT NULL COMMENT '唯一标识',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `perm_menu_ids` json NOT NULL COMMENT '权限集',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';

--
-- 转存表中的数据 `sys_role`
--

INSERT INTO `sys_role` (`id`, `parent_id`, `name`, `unique_key`, `remark`, `perm_menu_ids`, `status`, `order_num`, `create_time`, `update_time`) VALUES
(1, 0, '超级管理员', 'superadmin', '超级管理员', '[]', 1, 0, '2022-08-19 02:38:19', '2022-08-19 02:38:19');

-- --------------------------------------------------------

--
-- 表的结构 `sys_user`
--

CREATE TABLE `sys_user` (
  `id` int(10) NOT NULL COMMENT '编号',
  `account` varchar(25) NOT NULL COMMENT '账号',
  `password` char(32) NOT NULL COMMENT '密码',
  `username` varchar(25) NOT NULL COMMENT '姓名',
  `nickname` varchar(25) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0=保密 1=女 2=男',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮件',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `profession_id` int(11) NOT NULL COMMENT '职称',
  `job_id` int(11) NOT NULL COMMENT '岗位',
  `dept_id` int(11) NOT NULL COMMENT '部门',
  `role_ids` json NOT NULL COMMENT '角色集',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

--
-- 转存表中的数据 `sys_user`
--

INSERT INTO `sys_user` (`id`, `account`, `password`, `username`, `nickname`, `avatar`, `gender`, `email`, `mobile`, `profession_id`, `job_id`, `dept_id`, `role_ids`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 'arklnk', '596bfe4bb02db60c2a25965598529e7e', 'arklnk', 'arklnk', '', 1, 'arklnk@163.com', '12000000000', 0, 0, 0, '[1]', 1, 0, 'arklnk', '2022-08-11 06:19:45', '2022-08-22 09:42:42');

--
-- 转储表的索引
--

--
-- 表的索引 `sys_dept`
--
ALTER TABLE `sys_dept`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_key` (`unique_key`);

--
-- 表的索引 `sys_dictionary`
--
ALTER TABLE `sys_dictionary`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_key` (`unique_key`);

--
-- 表的索引 `sys_job`
--
ALTER TABLE `sys_job`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- 表的索引 `sys_log_action`
--
ALTER TABLE `sys_log_action`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `sys_log_login`
--
ALTER TABLE `sys_log_login`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `sys_perm_menu`
--
ALTER TABLE `sys_perm_menu`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `sys_profession`
--
ALTER TABLE `sys_profession`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- 表的索引 `sys_role`
--
ALTER TABLE `sys_role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_key` (`unique_key`);

--
-- 表的索引 `sys_user`
--
ALTER TABLE `sys_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `account` (`account`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `sys_dept`
--
ALTER TABLE `sys_dept`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `sys_dictionary`
--
ALTER TABLE `sys_dictionary`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `sys_job`
--
ALTER TABLE `sys_job`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `sys_log_action`
--
ALTER TABLE `sys_log_action`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号';

--
-- 使用表AUTO_INCREMENT `sys_log_login`
--
ALTER TABLE `sys_log_login`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号';

--
-- 使用表AUTO_INCREMENT `sys_perm_menu`
--
ALTER TABLE `sys_perm_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=43;

--
-- 使用表AUTO_INCREMENT `sys_profession`
--
ALTER TABLE `sys_profession`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `sys_role`
--
ALTER TABLE `sys_role`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `sys_user`
--
ALTER TABLE `sys_user`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
