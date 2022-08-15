set mysql=root:root@tcp(127.0.0.1:3306)/ark_admin_zero
set table=%1
set dir=%2
goctl model mysql datasource -url="%mysql%" -table="%table%" -c -dir %dir% -home ./doc/goctl