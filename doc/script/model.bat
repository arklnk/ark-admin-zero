set mysql=root:k9kzCqGft0@tcp(192.168.1.200:31746)/go_zero_admin
set table=%1
set dir=%2
goctl model mysql datasource -url="%mysql%" -table="%table%" -dir %dir% -home ./doc/template