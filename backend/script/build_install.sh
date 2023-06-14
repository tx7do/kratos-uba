cd ..
make build

mkdir -p ~/app/service

mkdir -p ~/app/service/agent/service/bin/
mkdir -p ~/app/service/admin/service/bin/
mkdir -p ~/app/service/core/service/bin/
mkdir -p ~/app/service/logger/service/bin/
mkdir -p ~/app/service/report/service/bin/

mkdir -p ~/app/service/agent/service/configs/
mkdir -p ~/app/service/admin/service/configs/
mkdir -p ~/app/service/core/service/configs/
mkdir -p ~/app/service/logger/service/configs/
mkdir -p ~/app/service/report/service/configs/

mv -f ./app/agent/service/bin/server ~/app/service/agent/service/bin/server
mv -f ./app/admin/service/bin/server ~/app/service/admin/service/bin/server
mv -f ./app/core/service/bin/server ~/app/service/core/service/bin/server
mv -f ./app/logger/service/bin/server ~/app/service/logger/service/bin/server
mv -f ./app/report/service/bin/server ~/app/service/report/service/bin/server

/bin/cp -rf ./app/agent/service/configs/*.yaml ~/app/service/agent/service/configs
/bin/cp -rf ./app/admin/service/configs/*.yaml ~/app/service/admin/service/configs
/bin/cp -rf ./app/core/service/configs/*.yaml ~/app/service/core/service/configs
/bin/cp -rf ./app/logger/service/configs/*.yaml ~/app/service/logger/service/configs
/bin/cp -rf ./app/report/service/configs/*.yaml ~/app/service/report/service/configs

cp -rf ./script/supervisor/*.ini /etc/supervisord.d/

mkdir -p /data/logs

sudo supervisorctl reload
sudo supervisorctl restart all
