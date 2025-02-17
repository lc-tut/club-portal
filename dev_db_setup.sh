#!/bin/bash

# データベースのダンプファイルをコンテナにコピー
cp ./dumps/dumps.sql ./cfg/mariadb/data

# MariadbコンテナでBashを実行
docker compose exec -it mariadb /bin/bash -c "mariadb -u root -p club_portal < /var/lib/mysql/dumps.sql"

# ダンプファイルを削除
rm ./cfg/mariadb/data/dumps.sql