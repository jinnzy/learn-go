#!/bi/bash

echo 123456
sleep 10
echo "`ip a|grep -A 5 "2: "|egrep "1[0-9][0-9]+?.* brd"|awk '{print $2}'|awk -F '/' '{print $1}'` 完成"

