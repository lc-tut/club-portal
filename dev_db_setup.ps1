# filepath: /d:/create/_lc/club-portal/setup.ps1

# データベースのダンプファイルをコンテナにコピー
Copy-Item ./dumps/dumps.sql ./cfg/mariadb/data

# MariadbコンテナでPowerShellを実行
docker compose exec -it mariadb /bin/bash -c "mariadb -u root -p club_portal < /var/lib/mysql/dumps.sql"

# 削除
Remove-Item ./cfg/mariadb/data/dumps.sql