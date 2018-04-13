#!/usr/bin/env bash

pname="jingyong"

echo "start Service........"
curr=$(cd `dirname $0`; pwd)
nohup ${curr}/${pname} >/dev/null 2>&1 &
echo "${pname} started..."
pause