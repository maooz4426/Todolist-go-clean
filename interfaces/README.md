# interface層

ここではFrameworks & DriversレイヤーとApplication Business Rulesの橋渡しに当たる、Interface Adaptersレイヤーを実装しています。

## gateway/repository/datasource
dbの抽象メソッドを定義しています。
usecases層で定義したrepositoryインターフェースを使っています。

## controllers
ここでは、外部のリクエストに対する処理を書いています。