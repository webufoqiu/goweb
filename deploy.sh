#! /bin/sh

kill -9 $(pgrep goweb)
cd ~/goweb/
git pull https://github.com/webufoqiu/goweb

./goweb &
