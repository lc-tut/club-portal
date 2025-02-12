# Club Portal | 東京工科大学サークルポータル

## 概要
東京工科大学公認サークルである [LinuxClub](https://www2.linux.it.teu.ac.jp) が開発, 運用 (予定) である課外活動用ポータル HP のバックエンド実装となります.

フロントエンドの実装は -> https://github.com/lc-tut/club-portal-frontend

## 言語, フレームワーク等
作成には以下の言語, フレームワーク等が使われています.
- Golang
- [gin](https://gin-gonic.com)
- MariaDB
- Redis
- Docker

## ディレクトリ構造
- `cfg` -> 認証用情報用のコンフィグファイル, 初期化用 SQL
- `consts` -> プロジェクトで使われるグローバルな定数や変数
- `models` -> データベース用のモデル情報
- `repos` -> モデルとデータベースを繋げる Repository
- `router` -> API 用 URL ルータ
- `utils` -> 汎用関数など

## 環境開発(2025/02/12追記)
### 現在はClub-Portal-FrontendにあるWikiをご覧ください。
リンクはこちら
https://github.com/lc-tut/club-portal-frontend/wiki

### データベースのセットアップ手順(LC部員向け)
この方法は、実際に稼働しているclub-portalの情報を使用します。情報の取り扱いには十分注意してください。
1. こちらから、dumps.sqlをダウンロードしてください(dumps{日付}となっているディレクトリ内にあります。)<br>
    https://github.com/lc-tut/club-portal-dumps
2. ダウンロードしたdumps.sqlを、手元環境のdumpsディレクトリへ格納してください。
3. Linuxの方は`dev_db_setup.sh`を実行してください。尚、mariadbのパスワードを問われますので、db_passに設定したパスワードを入力してください。(一応Windows用のdev_db_setup.ps1も作っておきました)
4. データベースがリストアされたことを確認してください。

## 開発
基本的には, 以下のようにブランチを切って作業を行い, Pull Request (PR) を `dev` ブランチに送ってください.
`dev` ブランチがメインブランチとなります.
```shell
$ git checkout -b <branch_name>
```

## ビルド & 実行
`docker-compose up -d --build`

## ライセンス
このプロジェクトは MIT License 下で作成されています.
