# 数据库文档

<a name="返回顶部"></a>

## 数据表列表

* [sys_dept(部门)](#sys_dept_pointer)

* [sys_dictionary(字典)](#sys_dictionary_pointer)

* [sys_job(工作岗位)](#sys_job_pointer)

* [sys_log_action(操作日志)](#sys_log_action_pointer)

* [sys_log_login(登录日志)](#sys_log_login_pointer)

* [sys_param(系统参数)](#sys_param_pointer)

* [sys_perm_menu(权限&菜单)](#sys_perm_menu_pointer)

* [sys_profession(职称)](#sys_profession_pointer)

* [sys_role(角色)](#sys_role_pointer)

* [sys_user(用户)](#sys_user_pointer)



## 数据表说明

<a name="sys_dept_pointer"></a>

* sys_dept表(部门)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(11)|编号|
|name|varchar(25)|部门简称|
|full_name|varchar(50)|部门全称|
|unique_key|varchar(25)|唯一值|
|parent_id|int(11)|父级id|
|status|tinyint(1)|0=禁用 1=开启|
|order_num|int(11)|排序值|
|remark|varchar(100)|备注|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_dictionary_pointer"></a>

* sys_dictionary表(字典)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(11)|编号|
|parent_id|int(11)|父级id|
|name|varchar(25)|名称|
|type|tinyint(2)|1文本 2数字 3数组 4单选 5多选 6下拉 7日期 8时间 9单文件 10多文件  |
|unique_key|varchar(25)|唯一值|
|value|int(11)|配置值|
|order_num|int(11)|排序值|
|remark|varchar(100)|备注|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_job_pointer"></a>

* sys_job表(工作岗位)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(11)|编号|
|name|varchar(50)|岗位名称|
|status|tinyint(1)|0=禁用 1=开启 |
|order_num|int(11)|排序值|
|create_time|timestamp|创建时间|
|update_time|timestamp|开启时间|

<a name="sys_log_action_pointer"></a>

* sys_log_action表(操作日志)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(10)|编号|
|user_id|int(11)|操作账号|
|ip|varchar(100)|ip|
|os|varchar(50)|系统|
|browser|varchar(50)|浏览器|
|uri|varchar(200)|请求路径|
|request|json|请求数据|
|response|json|响应数据|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_log_login_pointer"></a>

* sys_log_login表(登录日志)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(10)|编号|
|user_id|int(11)|登录账号|
|ip|varchar(100)|ip|
|os|varchar(50)|操作系统|
|browser|varchar(50)|浏览器|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_param_pointer"></a>

* sys_param表(系统参数)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(11)|编号|
|parent_id|int(11)|父级id|
|name|varchar(25)|名称|
|unique_key|varchar(25)|唯一值|
|value|int(11)|配置值|
|status|tinyint(1)|0=禁用 1=开启|
|order_num|int(11)|排序值|
|remark|varchar(100)|备注|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_perm_menu_pointer"></a>

* sys_perm_menu表(权限&菜单)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(11)|编号|
|parent_id|int(11)|父级id|
|name|varchar(255)|名称|
|router|varchar(255)|路由|
|perms|varchar(255)|权限|
|type|tinyint(1)|0=目录 1=菜单 2=权限|
|icon|varchar(255)|图标|
|order_num|int(11)|排序值|
|view_path|varchar(255)|页面路径|
|is_show|tinyint(1)|0=隐藏 1=显示|
|active_router|varchar(255)|当前激活的菜单|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_profession_pointer"></a>

* sys_profession表(职称)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(11)|编号|
|name|varchar(25)|职称|
|status|tinyint(1)|0=禁用 1=开启|
|order_num|int(11)|排序值|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_role_pointer"></a>

* sys_role表(角色)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(10)|编号|
|parent_id|int(11)|父级id|
|name|varchar(25)|名称|
|unique_key|varchar(25)|唯一标识|
|remark|varchar(100)|备注|
|perm_menu_ids|json|权限集|
|status|tinyint(1)|0=禁用 1=开启|
|order_num|int(11)|排序值|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

<a name="sys_user_pointer"></a>

* sys_user表(用户)[↑](#返回顶部)

|字段名称|字段类型|字段含义|
|:---:|:---:|:---:|
|id|int(10)|编号|
|account|varchar(25)|账号|
|password|char(32)|密码|
|username|varchar(25)|姓名|
|nickname|varchar(25)|昵称|
|avatar|varchar(200)|头像|
|gender|tinyint(1)|0=保密 1=女 2=男|
|birthday|datetime|生日|
|email|varchar(50)|邮件|
|mobile|char(11)|手机号|
|profession_id|int(11)|职称|
|job_id|int(11)|岗位|
|dept_id|int(11)|部门|
|role_ids|json|角色集|
|status|tinyint(1)|0=禁用 1=开启|
|order_num|int(11)|排序值|
|remark|varchar(100)|备注|
|create_time|timestamp|创建时间|
|update_time|timestamp|更新时间|

