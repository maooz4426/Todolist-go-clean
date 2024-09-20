# infrastructure層

ここでは、Frameworks & Driversレイヤーとして実装しています。

## databases
ここではdbの接続処理を記述しています。
dbはmysqlやpostgresql、Nosql等で書き方が変わってくると思うので、repositoryとは別で実装しています。

## router
goのフレームワークであるechoを使ったルーティングの設定を行なっています。