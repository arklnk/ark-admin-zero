
--
-- 更新status、is_show、type等字段，修改逻辑，
-- 1.状态原先0和1的，将0修改为2，1不做修改
-- 2.状态原先0到2的，所有参数均+1

ALTER TABLE `sys_perm_menu`
    MODIFY COLUMN `type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1=目录 2=菜单 3=权限' AFTER `perms`,
    MODIFY COLUMN `is_show` tinyint(1) UNSIGNED NULL DEFAULT 1 COMMENT '2=隐藏 1=显示' AFTER `view_path`;
UPDATE sys_perm_menu SET `is_show` = 2 WHERE  `is_show`=0;
UPDATE sys_perm_menu SET `type` = 3 WHERE  `type`=2;
UPDATE sys_perm_menu SET `type` = 2 WHERE  `type`=1;
UPDATE sys_perm_menu SET `type` = 1 WHERE  `type`=0;


ALTER TABLE `sys_user`
    MODIFY COLUMN `gender` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1=保密 2=女 3=男' AFTER `avatar`,
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 2 COMMENT '2=禁用 1=开启' AFTER `role_ids`;
UPDATE sys_perm_menu SET `gender` = 3 WHERE  `gender`=2;
UPDATE sys_perm_menu SET `gender` = 2 WHERE  `gender`=1;
UPDATE sys_perm_menu SET `gender` = 1 WHERE  `gender`=0;


ALTER TABLE `sys_dept`
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 2 COMMENT '2=禁用 1=开启' AFTER `type`;
ALTER TABLE `sys_dictionary`
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '2=禁用 1=开启' AFTER `value`;

ALTER TABLE `sys_job`
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '2=禁用 1=开启 ' AFTER `name`;
ALTER TABLE `sys_log`
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '2=失败 1=成功' AFTER `request`;

ALTER TABLE `sys_profession`
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '2=禁用 1=开启' AFTER `name`;

ALTER TABLE `sys_role`
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '2=禁用 1=开启' AFTER `perm_menu_ids`;

ALTER TABLE `sys_user`
    MODIFY COLUMN `gender` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '1=保密 2=女 3=男' AFTER `avatar`,
    MODIFY COLUMN `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '2=禁用 1=开启' AFTER `role_ids`;


UPDATE sys_dept SET `status` = 2 WHERE  `status`=0;
UPDATE sys_dictionary SET `status` = 2 WHERE  `status`=0;
UPDATE sys_job SET `status` = 2 WHERE  `status`=0;
UPDATE sys_log SET `status` = 2 WHERE  `status`=0;
UPDATE sys_profession SET `status` = 2 WHERE  `status`=0;
UPDATE sys_role SET `status` = 2 WHERE  `status`=0;
UPDATE sys_user SET `status` = 2 WHERE  `status`=0;