## プロジェクト作成時
- DDLを実行すること(.read ddl-file)
- configの作成（.emvとauth_config.json）
- npm install
  - ERR LIFECYCLEの場合
  - `rm -rm node_modules`
  - `rm package-lock.json`
  - npm install

## aimoローカル起動
1. コンソール開いてaimoに移動
2. ```docker compose up -d```<br/>
  バックグラウンドで動かす
3. ```docker exec -ti aimo_web sh```<br/>
   shコマンドでaimo_webコンテナに入ることができて操作できる
4. ```cd aimo```<br/>
  vueプロジェクト（aimo）に入る
5. ```npm run serve```
<br/><br/><br/>

# SERVER起動
ターミナル開く
## serverの起動
```
docker exec -ti aimo_api bash
go run grpc/server/main.go
```
別ウィンドウ
## gatewayプロキシの起動
```
docker exec -ti aimo_api bash
go run grpc/gateway/main.go
```
## Gitコミット
1. 変更すると「SOURCE CONTROL」（ブランチのマーク）の<br/>
  「Changes」に変更したファイルが上がってくる
2. コミットしたいファイルのファイル名にカーソルを合わせると<br/>
    ＋ボタンが出るのでクリック<br/>
    →ファイルが「Staged Changes」に移動する
3. 「SOURCE CONTROL」にカーソルを合わせるとチェックマークが<br/>
    出るのでクリック（＝コミット）
4. 「SOURCE CONTROL」にカーソルを合わせて右に出る：をクリック
5. 「Push」をクリック
<br/><br/><br/>
## aimo_frontにmasterを反映
1. 「SOURCE CONTROL」で「Pull from…」選択
2. 候補に「origin」出てくるので選択
3. 候補に「origin/master」出てくるので選択

（4. pushすると他の人のaimo_frontにもmaster反映されてる状態になる）
<br/><br/><br/>
## インストールしたよ
- router
- vuex (npm install vuex)<br/>
（storeも作成）
（"ストア" は、基本的にアプリケーションの 状態（state） を保持するコンテナ）
- vuetify
- $ npm install chart.js@2.8
- $ npm install vue-chartjs@3.4.2
- npm install @auth0/auth0-spa-js
<br/><br/><br/>
## その他ルール
- 資料系はmasterブランチに移動してから書く！
<br/><br/><br/>
## TODO
- axiosで取得した値の中を自由自在に使いたい 済
- 目標設定シート記入画面作成
  - 方針系はコンポーネント作成　済
  - タブの中身は１つのコンポーネントで作る　済
    - 配置はfloatでやってみたけどだめだ、やっぱinline-blockにしよ、、　済
  - スケジュールはとりあえずチェックボックスで 済
    - 将来的にボタンはいかが？
  - タブに制御つける
  - ManagementPolicy.vueで「¥n」が入ってたら改行するロジック作る
  <br/>
- postもputも改行のところは「¥n」か何かしら入れて返す
- 9/24に大石さんのアポ取って目標設定シートのputとpostやる
  - 9/28 続編

<br/>

- curlコマンドくらい叩けるようにせんかい！
- /front/Dockerfileにnpm install npm を追記
- opneapi作成
- docker-compose up時にdocker networkのアドレスが変更される問題(参考：http://ajisaba.net/develop/docker/docker-compose_ip.html)
- ER図作成
- 部署によって目標が役割が変わるとしたら、部署は札幌エイジア含めめちゃ細かくする必要あり？いろいろマスターを変更した時に順序はどうやって担保すんの？マスター変更のオペレーションは？
- ユーザー側
  - 総合評価表示、
- 管理者側
  - 期末面談者コメント入力、総合評価入力
- 目標テーブル、管理側入力用評価テーブルと分離
- go getの内容がgo.modに記述されない問題
- aimo_open.yml、IDとキャメルフィールドに変更する

  - ManagementPolicy.vueで「¥n」が入ってたら改行するロジック作るtお
  - ManagementPolicy.vueで「¥n」が入ってたら改行するロジック作る

# aimo_openapi.yml変更点
- Post personalEva requestBody personal_evaluation_id削除
- aimApi全て achievement_means削除
- achievementMeanAPI追加
- achievementMeanModelのフィールド名変更　mean_number -> achievement_mean_number
- 全getメソッドレスポンスのフィールド名変更　API名_id -> id
- achievementMeans　req,res複数対応
- GETパラメータ記述修正（対象api: evaluationBefore, evaluation, comprehensiveComment)
- evaluation GETリクエストパラメータ追加（追加項目：evaluatorNumber)
- userInfoをuserに変更
- userのuserIDをauthIdに変更j
- period, department, jobのGETレスポンスフィールド名を複数形に変更
