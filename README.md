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
<br/><br/>
## その他ルール
- 資料系はmasterブランチに移動してから書く！

## TODO
- ヘッダーの文字はVuexで変える？（参照：https://qiita.com/att55/items/91b683c68b5057eaac51）