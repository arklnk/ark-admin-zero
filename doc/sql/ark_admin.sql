-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql
-- 生成日期： 2022-08-17 14:19:52
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
  `status` tinyint(1) NOT NULL COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '""' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门';

--
-- 转存表中的数据 `sys_dept`
--

INSERT INTO `sys_dept` (`id`, `parent_id`, `name`, `full_name`, `unique_key`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 0, '方舟', '方舟互联', 'arklnk', 1, 1, '', '2022-08-17 02:09:17', '2022-08-17 03:16:32');

-- --------------------------------------------------------

--
-- 表的结构 `sys_dictionary`
--

CREATE TABLE `sys_dictionary` (
  `id` int(11) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(25) NOT NULL COMMENT '名称',
  `type` tinyint(2) NOT NULL COMMENT '1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单文件 10多文件  ',
  `unique_key` varchar(25) NOT NULL COMMENT '唯一值',
  `value` int(11) NOT NULL COMMENT '配置值',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '""' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典';

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
-- 表的结构 `sys_param`
--

CREATE TABLE `sys_param` (
  `id` int(11) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(25) NOT NULL COMMENT '名称',
  `unique_key` varchar(25) NOT NULL COMMENT '唯一值',
  `value` int(11) NOT NULL COMMENT '配置值',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '""' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统参数';

-- --------------------------------------------------------

--
-- 表的结构 `sys_perm_menu`
--

CREATE TABLE `sys_perm_menu` (
  `id` int(11) NOT NULL COMMENT '编号',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `router` varchar(255) NOT NULL DEFAULT '""' COMMENT '路由',
  `perms` varchar(255) NOT NULL DEFAULT '""' COMMENT '权限',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0=目录 1=菜单 2=权限',
  `icon` varchar(255) NOT NULL DEFAULT '""' COMMENT '图标',
  `order_num` int(11) DEFAULT '0' COMMENT '排序值',
  `view_path` varchar(255) NOT NULL DEFAULT '""' COMMENT '页面路径',
  `is_show` tinyint(1) DEFAULT '1' COMMENT '0=隐藏 1=显示',
  `active_router` varchar(255) NOT NULL DEFAULT '""' COMMENT '当前激活的菜单',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限&菜单';

--
-- 转存表中的数据 `sys_perm_menu`
--

INSERT INTO `sys_perm_menu` (`id`, `parent_id`, `name`, `router`, `perms`, `type`, `icon`, `order_num`, `view_path`, `is_show`, `active_router`, `create_time`, `update_time`) VALUES
(1, 0, 'routes.systemManagement', '/sys', '[]', 0, 'system', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-15 13:29:53'),
(2, 1, 'routes.permManagement', '/sys/perms', '[]', 0, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-16 06:26:42'),
(3, 2, 'routes.menuList', '/sys/perms/menu/list', '[]', 1, '', 0, 'views/system/permission/menu', 1, '', '2022-08-12 02:14:20', '2022-08-16 06:57:24'),
(4, 3, 'common.basic.query', '', '[\"sys/perm/menu/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 07:31:41'),
(5, 3, 'common.basic.add', '', '[\"sys/perm/menu/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 06:21:43'),
(6, 3, 'common.basic.delete', '', '[\"sys/perm/menu/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 07:32:09'),
(7, 3, 'common.basic.update', '', '[\"sys/perm/menu/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 07:32:18'),
(8, 2, '角色列表', '/sys/perms/role/list', '[]', 1, '', 0, 'views/system/permission/role', 1, '', '2022-08-12 02:14:20', '2022-08-16 13:12:14'),
(9, 8, '列表', '', '[\"sys/role/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-16 13:12:54'),
(10, 8, '新增', '', '[\"sys/role/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-16 13:13:29'),
(11, 8, '删除', '', '[\"sys/role/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-16 13:13:51'),
(12, 8, '更新', '', '[\"sys/role/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-16 13:14:10'),
(13, 8, '权限', '', '[\"sys/role/perm/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:53:24'),
(14, 2, '部门列表', '/sys/perms/dept/list', '[]', 1, '', 0, 'views/system/permission/dept', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:51:04'),
(15, 14, '列表', '', '[\"sys/dept/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:51:48'),
(16, 14, '新增', '', '[\"sys/dept/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:52:17'),
(17, 14, '删除', '', '[\"sys/dept/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:52:44'),
(18, 14, '更新', '', '[\"sys/dept/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:53:04'),
(19, 14, '转移', '', '[\"sys/dept/transfer\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 01:53:53'),
(20, 2, '岗位列表', '/sys/perms/job/list', '[]', 1, '', 0, 'views/system/permission/job', 1, '', '2022-08-12 02:14:20', '2022-08-17 02:59:03'),
(21, 20, '列表', '', '[\"sys/job/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 02:59:25'),
(22, 20, '新增', '', '[\"sys/job/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 02:59:54'),
(23, 20, '删除', '', '[\"sys/job/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 03:00:11'),
(24, 20, '更新', '', '[\"sys/job/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 03:00:26'),
(25, 2, '职称列表', '/sys/perms/profession/list', '[]', 1, '', 0, 'views/system/profession/job', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:01:11'),
(26, 25, '列表', '', '[\"sys/profession/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(27, 25, '新增', '', '[\"sys/profession/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(28, 25, '删除', '', '[\"sys/profession/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(29, 25, '更新', '', '[\"sys/profession/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(30, 2, '用户列表', '/sys/perms/user/list', '[]', 1, '', 0, 'views/system/user/job', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:01:11'),
(31, 30, '列表', '', '[\"sys/user/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(32, 30, '新增', '', '[\"sys/user/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(33, 30, '删除', '', '[\"sys/user/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(34, 30, '更新', '', '[\"sys/user/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22'),
(35, 30, '改密', '', '[\"sys/user/password/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-17 05:02:22');

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
  `remark` varchar(100) NOT NULL DEFAULT '""' COMMENT '备注',
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
(1, 0, '超级管理员', 'superadmin', '', '[0]', 1, 0, '2022-08-11 09:18:21', '2022-08-15 14:31:00');

-- --------------------------------------------------------

--
-- 表的结构 `sys_user`
--

CREATE TABLE `sys_user` (
  `id` int(10) NOT NULL COMMENT '编号',
  `account` varchar(25) NOT NULL COMMENT '账号',
  `password` char(32) NOT NULL COMMENT '密码',
  `username` varchar(25) NOT NULL COMMENT '姓名',
  `nickname` varchar(25) NOT NULL DEFAULT '""' COMMENT '昵称',
  `avatar` varchar(200) NOT NULL DEFAULT '""' COMMENT '头像',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0=保密 1=女 2=男',
  `birthday` datetime NOT NULL DEFAULT '1970-01-01 00:00:00' COMMENT '生日',
  `email` varchar(50) NOT NULL DEFAULT '""' COMMENT '邮件',
  `mobile` char(11) NOT NULL DEFAULT '""' COMMENT '手机号',
  `profession_id` int(11) NOT NULL COMMENT '职称',
  `job_id` int(11) NOT NULL COMMENT '岗位',
  `dept_id` int(11) NOT NULL COMMENT '部门',
  `role_ids` json NOT NULL COMMENT '角色集',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '""' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

--
-- 转存表中的数据 `sys_user`
--

INSERT INTO `sys_user` (`id`, `account`, `password`, `username`, `nickname`, `avatar`, `gender`, `birthday`, `email`, `mobile`, `profession_id`, `job_id`, `dept_id`, `role_ids`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 'arklnk', '596bfe4bb02db60c2a25965598529e7e', 'arklnk', 'arklnk', 'http://dummyimage.com/100x100', 1, '1970-01-01 00:00:00', 'arklnk@163.com', '12000000000', 1, 1, 1, '[1]', 1, 0, 'arklnk', '2022-08-11 06:19:45', '2022-08-17 14:15:02');

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
-- 表的索引 `sys_param`
--
ALTER TABLE `sys_param`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_key` (`unique_key`);

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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `sys_dictionary`
--
ALTER TABLE `sys_dictionary`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号';

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
-- 使用表AUTO_INCREMENT `sys_param`
--
ALTER TABLE `sys_param`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号';

--
-- 使用表AUTO_INCREMENT `sys_perm_menu`
--
ALTER TABLE `sys_perm_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=36;

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
