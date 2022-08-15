-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql
-- 生成日期： 2022-08-15 15:24:24
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
-- 数据库： `ark_admin_zero`
--

-- --------------------------------------------------------

--
-- 表的结构 `sys_dept`
--

CREATE TABLE `sys_dept` (
  `id` int(11) NOT NULL COMMENT '编号',
  `name` varchar(25) NOT NULL COMMENT '部门简称',
  `full_name` varchar(50) NOT NULL COMMENT '部门全称',
  `unique_key` varchar(25) NOT NULL COMMENT '唯一值',
  `parent_id` int(11) NOT NULL COMMENT '父级id',
  `status` tinyint(1) NOT NULL COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL COMMENT '排序值',
  `remark` varchar(100) NOT NULL DEFAULT '""' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门';

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
  `name` int(11) NOT NULL COMMENT '岗位名称',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启 ',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '开启时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工作岗位';

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
(2, 1, 'routes.permManagement', '/sys/perms', '[]', 0, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-15 14:38:40'),
(3, 2, 'routes.menuList', '/sys/perms/menu/list', '[\"perms1\",\"perms2\"]', 1, '', 0, 'views/system/permission/menu', 1, '', '2022-08-12 02:14:20', '2022-08-15 13:29:46');

-- --------------------------------------------------------

--
-- 表的结构 `sys_profession`
--

CREATE TABLE `sys_profession` (
  `id` int(11) NOT NULL COMMENT '编号',
  `name` varchar(25) NOT NULL COMMENT '职称',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职称';

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
(1, 'arklnk', '596bfe4bb02db60c2a25965598529e7e', '蒋勇', '顾芳', 'http://dummyimage.com/100x100', 1, '1970-01-01 00:00:00', 'm.cqcprwum@qq.com', '18637334616', 1, 1, 1, '[1]', 1, 0, 'nisi', '2022-08-11 06:19:45', '2022-08-15 15:22:46');

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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号';

--
-- 使用表AUTO_INCREMENT `sys_dictionary`
--
ALTER TABLE `sys_dictionary`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号';

--
-- 使用表AUTO_INCREMENT `sys_job`
--
ALTER TABLE `sys_job`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号';

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
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `sys_profession`
--
ALTER TABLE `sys_profession`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号';

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
