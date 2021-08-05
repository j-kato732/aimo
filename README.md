aimoローカル起動

①コンソール開いてaimoに移動
②docker compose up -d（バックグラウンドで動かす）
③docker exec -ti aimo_web sh
　shコマンドでaimo_webコンテナに入ることができて操作できる
④cd aimo（＝vueプロジェクトに入る）
⑤npm run serve



Git コミット
①変更すると「SOURCE CONTROL」（ブランチのマーク）の
　「Changes」に変更したファイルが上がってくる
②コミットしたいファイルのファイル名にカーソルを合わせると
　＋ボタンが出るのでクリック
　→ファイルが「Staged Changes」に移動する
③「SOURCE CONTROL」にカーソルを合わせるとチェックマークが
　　出るのでクリック（＝コミット）
④「SOURCE CONTROL」にカーソルを合わせて右に出る：をクリック
⑤「Push」をクリック