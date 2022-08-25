-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql
-- 生成日期： 2022-08-25 02:41:25
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
  `id` int(11) UNSIGNED NOT NULL COMMENT '编号',
  `parent_id` int(11) UNSIGNED NOT NULL COMMENT '父级id',
  `name` varchar(50) NOT NULL COMMENT '部门简称',
  `full_name` varchar(50) NOT NULL COMMENT '部门全称',
  `unique_key` varchar(50) NOT NULL COMMENT '唯一值',
  `type` tinyint(1) UNSIGNED NOT NULL COMMENT '1=公司 2=子公司 3=部门',
  `status` tinyint(1) UNSIGNED NOT NULL COMMENT '0=禁用 1=开启',
  `order_num` int(11) UNSIGNED NOT NULL COMMENT '排序值',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门';

--
-- 转存表中的数据 `sys_dept`
--

INSERT INTO `sys_dept` (`id`, `parent_id`, `name`, `full_name`, `unique_key`, `type`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 0, '方舟', '方舟互联', 'arklnk', 1, 1, 0, '', '2022-08-17 02:09:17', '2022-08-22 02:13:54'),
(2, 0, '思忆', '思忆技术', 'siyee', 1, 1, 0, '', '2022-08-19 06:40:10', '2022-08-22 02:13:39'),
(3, 0, '演示', '演示部门', 'demo', 1, 1, 0, '', '2022-08-23 14:02:27', '2022-08-25 02:27:51');

-- --------------------------------------------------------

--
-- 表的结构 `sys_dictionary`
--

CREATE TABLE `sys_dictionary` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '编号',
  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '0=配置集 !0=父级id',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `type` tinyint(2) UNSIGNED NOT NULL DEFAULT '1' COMMENT '1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单图 10多图 11单文件 12多文件',
  `unique_key` varchar(50) NOT NULL COMMENT '唯一值',
  `value` varchar(2048) NOT NULL DEFAULT '' COMMENT '配置值',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统参数';

--
-- 转存表中的数据 `sys_dictionary`
--

INSERT INTO `sys_dictionary` (`id`, `parent_id`, `name`, `type`, `unique_key`, `value`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 0, '系统配置', 0, 'sys', '', 1, 0, '', '2022-08-22 10:03:58', '2022-08-23 01:25:31'),
(2, 1, '默认密码', 1, 'sys_pwd', '123456', 1, 0, '新建用户默认密码', '2022-08-22 10:03:58', '2022-08-24 05:28:06');

-- --------------------------------------------------------

--
-- 表的结构 `sys_job`
--

CREATE TABLE `sys_job` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '编号',
  `name` varchar(50) NOT NULL COMMENT '岗位名称',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启 ',
  `order_num` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '开启时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='工作岗位';

--
-- 转存表中的数据 `sys_job`
--

INSERT INTO `sys_job` (`id`, `name`, `status`, `order_num`, `create_time`, `update_time`) VALUES
(1, '前端', 1, 0, '2022-08-17 03:15:56', '2022-08-17 05:27:26'),
(2, '后端', 1, 0, '2022-08-17 03:15:56', '2022-08-17 05:32:50'),
(3, '设计', 1, 0, '2022-08-17 03:15:56', '2022-08-17 05:32:55'),
(4, '演示', 1, 0, '2022-08-23 14:01:23', '2022-08-23 14:01:23');

-- --------------------------------------------------------

--
-- 表的结构 `sys_log`
--

CREATE TABLE `sys_log` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '编号',
  `user_id` int(11) UNSIGNED NOT NULL COMMENT '操作账号',
  `ip` varchar(100) NOT NULL COMMENT 'ip',
  `uri` varchar(200) NOT NULL COMMENT '请求路径',
  `type` tinyint(1) UNSIGNED NOT NULL COMMENT '1=登录日志 2=操作日志',
  `request` varchar(2048) NOT NULL DEFAULT '' COMMENT '请求数据',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '0=失败 1=成功',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统日志';

-- --------------------------------------------------------

--
-- 表的结构 `sys_perm_menu`
--

CREATE TABLE `sys_perm_menu` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '编号',
  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `router` varchar(200) NOT NULL DEFAULT '' COMMENT '路由',
  `perms` varchar(200) NOT NULL DEFAULT '' COMMENT '权限',
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '0=目录 1=菜单 2=权限',
  `icon` varchar(50) NOT NULL DEFAULT '' COMMENT '图标',
  `order_num` int(11) UNSIGNED DEFAULT '0' COMMENT '排序值',
  `view_path` varchar(200) NOT NULL DEFAULT '' COMMENT '页面路径',
  `is_show` tinyint(1) UNSIGNED DEFAULT '1' COMMENT '0=隐藏 1=显示',
  `active_router` varchar(200) NOT NULL DEFAULT '' COMMENT '当前激活的菜单',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限&菜单';

--
-- 转存表中的数据 `sys_perm_menu`
--

INSERT INTO `sys_perm_menu` (`id`, `parent_id`, `name`, `router`, `perms`, `type`, `icon`, `order_num`, `view_path`, `is_show`, `active_router`, `create_time`, `update_time`) VALUES
(1, 0, '系统管理', '/sys', '[]', 0, 'system', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:32:28'),
(2, 1, '菜单管理', '/sys/menu', '[]', 1, '', 0, 'views/system/menu', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:32:44'),
(3, 2, '查询', '', '[\"sys/perm/menu/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:35:53'),
(4, 2, '新增', '', '[\"sys/perm/menu/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:07'),
(5, 2, '删除', '', '[\"sys/perm/menu/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:20'),
(6, 2, '更新', '', '[\"sys/perm/menu/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:34'),
(7, 1, '角色管理', '/sys/role', '[]', 1, '', 0, 'views/system/role', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:33:26'),
(8, 7, '查询', '', '[\"sys/role/list\",\"sys/perm/menu/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:35:53'),
(9, 7, '新增', '', '[\"sys/role/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:07'),
(10, 7, '删除', '', '[\"sys/role/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:20'),
(11, 7, '更新', '', '[\"sys/role/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:34'),
(13, 1, '部门管理', '/sys/dept', '[]', 1, '', 0, 'views/system/dept', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:33:40'),
(14, 13, '查询', '', '[\"sys/dept/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:35:53'),
(15, 13, '新增', '', '[\"sys/dept/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:07'),
(16, 13, '删除', '', '[\"sys/dept/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:20'),
(17, 13, '更新', '', '[\"sys/dept/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:34'),
(18, 1, '岗位管理', '/sys/job', '[]', 1, '', 0, 'views/system/job', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:33:55'),
(19, 18, '查询', '', '[\"sys/job/page\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:35:53'),
(20, 18, '新增', '', '[\"sys/job/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:07'),
(21, 18, '删除', '', '[\"sys/job/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:20'),
(22, 18, '更新', '', '[\"sys/job/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:34'),
(24, 1, '职称管理', '/sys/profession', '[]', 1, '', 0, 'views/system/profession', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:34:09'),
(25, 24, '查询', '', '[\"sys/profession/page\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:35:53'),
(26, 24, '新增', '', '[\"sys/profession/add\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:07'),
(27, 24, '删除', '', '[\"sys/profession/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:20'),
(28, 24, '更新', '', '[\"sys/profession/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:34'),
(29, 1, '用户管理', '/sys/user', '[]', 1, '', 0, 'views/system/user', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:34:20'),
(30, 29, '查询', '', '[\"sys/user/page\",\"sys/dept/list\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-24 03:46:56'),
(31, 29, '新增', '', '[\"sys/user/add\",\"sys/user/rdpj/info\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-24 03:17:19'),
(32, 29, '删除', '', '[\"sys/user/delete\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:37:20'),
(33, 29, '更新', '', '[\"sys/user/update\",\"sys/user/rdpj/info\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-24 03:08:07'),
(34, 29, '更改密码', '', '[\"sys/user/password/update\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:34:28'),
(35, 29, '转移', '', '[\"sys/user/transfer\"]', 2, '', 0, '', 1, '', '2022-08-12 02:14:20', '2022-08-23 09:34:36'),
(36, 0, '配置管理', '/config', '[]', 0, 'config', 0, '', 1, '', '2022-08-22 03:33:42', '2022-08-24 03:41:35'),
(37, 36, '字典管理', '/config/dict', '[]', 1, '', 0, 'views/config/dict', 1, '', '2022-08-22 03:39:21', '2022-08-23 09:33:47'),
(38, 37, '查询', '', '[\"config/dict/list\",\"config/dict/data/page\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 09:35:53'),
(39, 37, '新增', '', '[\"config/dict/add\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 09:37:07'),
(40, 37, '删除', '', '[\"config/dict/delete\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 09:37:20'),
(41, 37, '更新', '', '[\"config/dict/update\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 09:37:34'),
(42, 0, '日志管理', '/log', '[]', 0, 'log', 0, '', 1, '', '2022-08-23 04:47:23', '2022-08-24 03:52:57'),
(43, 42, '登录日志', '/log/login', '[]', 1, '', 0, 'views/log/login', 1, '', '2022-08-23 04:47:51', '2022-08-23 09:42:43'),
(44, 43, '查询', '', '[\"log/login/page\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 09:35:53'),
(45, 43, '删除', '', '[\"log/login/delete\"]', 2, '', 0, '', 1, '', '2022-08-22 03:42:07', '2022-08-23 09:37:20');

-- --------------------------------------------------------

--
-- 表的结构 `sys_profession`
--

CREATE TABLE `sys_profession` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '编号',
  `name` varchar(50) NOT NULL COMMENT '职称',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职称';

--
-- 转存表中的数据 `sys_profession`
--

INSERT INTO `sys_profession` (`id`, `name`, `status`, `order_num`, `create_time`, `update_time`) VALUES
(1, 'CEO', 1, 0, '2022-08-17 05:09:26', '2022-08-17 05:09:26'),
(2, 'CTO', 1, 0, '2022-08-17 05:09:26', '2022-08-17 05:09:26'),
(3, '演示', 1, 0, '2022-08-23 14:01:43', '2022-08-23 14:01:43');

-- --------------------------------------------------------

--
-- 表的结构 `sys_role`
--

CREATE TABLE `sys_role` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '编号',
  `parent_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `unique_key` varchar(50) NOT NULL COMMENT '唯一标识',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `perm_menu_ids` json NOT NULL COMMENT '权限集',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色';

--
-- 转存表中的数据 `sys_role`
--

INSERT INTO `sys_role` (`id`, `parent_id`, `name`, `unique_key`, `remark`, `perm_menu_ids`, `status`, `order_num`, `create_time`, `update_time`) VALUES
(1, 0, '超级管理员', 'superadmin', '超级管理员', '[]', 1, 0, '2022-08-19 02:38:19', '2022-08-19 02:38:19'),
(2, 0, '演示', 'demo', '', '[3, 8, 14, 19, 25, 30, 38, 44, 1, 2, 7, 13, 18, 24, 29, 36, 37, 42, 43]', 0, 0, '2022-08-23 13:13:05', '2022-08-25 02:29:13'),
(3, 0, '测试', 'testing', '', '[]', 1, 0, '2022-08-24 05:01:46', '2022-08-24 07:59:33');

-- --------------------------------------------------------

--
-- 表的结构 `sys_user`
--

CREATE TABLE `sys_user` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '编号',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `password` char(32) NOT NULL COMMENT '密码',
  `username` varchar(50) NOT NULL COMMENT '姓名',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(400) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '0=保密 1=女 2=男',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮件',
  `mobile` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `profession_id` int(11) UNSIGNED NOT NULL COMMENT '职称',
  `job_id` int(11) UNSIGNED NOT NULL COMMENT '岗位',
  `dept_id` int(11) UNSIGNED NOT NULL COMMENT '部门',
  `role_ids` json NOT NULL COMMENT '角色集',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '0=禁用 1=开启',
  `order_num` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序值',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

--
-- 转存表中的数据 `sys_user`
--

INSERT INTO `sys_user` (`id`, `account`, `password`, `username`, `nickname`, `avatar`, `gender`, `email`, `mobile`, `profession_id`, `job_id`, `dept_id`, `role_ids`, `status`, `order_num`, `remark`, `create_time`, `update_time`) VALUES
(1, 'arklnk', '596bfe4bb02db60c2a25965598529e7e', 'arklnk2', 'arklnk', 'https://avataaars.io/?clotheColor=Blue02&accessoriesType=Prescription01&avatarStyle=Circle&clotheType=BlazerSweater&eyeType=Hearts&eyebrowType=RaisedExcitedNatural&facialHairColor=Brown&facialHairType=MoustacheFancy&hairColor=PastelPink&hatColor=Pink&mouthType=Default&skinColor=Light&topType=Hijab', 0, 'arklnk@163.com', '12000000000', 0, 0, 0, '[1]', 1, 0, 'arklnk', '2022-08-11 06:19:45', '2022-08-25 02:40:01'),
(2, 'demo', '596bfe4bb02db60c2a25965598529e7e', 'demo', '', 'https://avataaars.io/?avatarStyle=Circle&topType=Hat&accessoriesType=Sunglasses&facialHairType=Blank&clotheType=Hoodie&clotheColor=Heather&eyeType=Hearts&eyebrowType=UpDown&mouthType=Tongue&skinColor=DarkBrown', 0, '', '', 3, 4, 3, '[2]', 1, 0, '', '2022-08-23 14:04:24', '2022-08-24 08:44:59');

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
-- 表的索引 `sys_log`
--
ALTER TABLE `sys_log`
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
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `sys_dictionary`
--
ALTER TABLE `sys_dictionary`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `sys_job`
--
ALTER TABLE `sys_job`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `sys_log`
--
ALTER TABLE `sys_log`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号';

--
-- 使用表AUTO_INCREMENT `sys_perm_menu`
--
ALTER TABLE `sys_perm_menu`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=46;

--
-- 使用表AUTO_INCREMENT `sys_profession`
--
ALTER TABLE `sys_profession`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `sys_role`
--
ALTER TABLE `sys_role`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `sys_user`
--
ALTER TABLE `sys_user`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '编号', AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
