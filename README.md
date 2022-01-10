# sql-update-kube  

sql-update-kube は、主にエッジコンピューティング環境において 入力した JSONデータ に SQL で保持された静的データとの差分がある場合に、SQLを更新するマイクロサービスです。  
本マイクロサービスには、SQLテーブルへのレコードの登録機能、更新機能が含まれます
本マイクロサービスは、JSONデータを RabbitMQ のキューで受け取ります。  

## 動作環境  

* OS: Linux OS  
* CPU: ARM/AMD/Intel  
* Kubernetes  

## クラウド環境での利用  
sql-update-kube は、クラウド環境においても、利用可能なように設計されています。  

## SQLの登録更新 モジュール SQLBOILER の利用
sql-update-kube は、SQLの登録更新モジュールとして、[SQLBOILER](https://github.com/volatiletech/sqlboiler)を利用しています。  

## 事前準備（SQLデータベースからのソースコードの生成）
sql-update-kube のランタイム環境を利用するための事前準備として、  

```
/database/model_generate.sh
```
  
の実行をします。  
これにより、対象のSQLデータベースを読み込み、当該SQLデータベースの登録更新モジュールとして必要なソースコードが、/database/models/ 内に自動生成されます。  
（同一環境内での２回目以降の実行ではソースコードが上書きされます）  

model_generate.sh の 対象データベース接続時の情報は、/database/sqlboiler.toml に書かれる必要があります。    

## RabbitMQ からの JSON Input

sql-update-kube は、Inputとして、RabbitMQ からのメッセージをJSON形式で受け取ります。 

## RabbitMQ からのメッセージ受信による イベントドリヴン の ランタイム実行

sql-update-kube は、RabbitMQ からのメッセージを受け取ると、イベントドリヴンでランタイムを実行します。  
AION の仕様では、Kubernetes 上 の 当該マイクロサービスPod は 立ち上がったまま待機状態で当該メッセージを受け取り、（コンテナ起動などの段取時間をカットして）即座にランタイムを実行します。　

## RabbitMQ の マスタサーバ環境

sql-update-kube が利用する RabbitMQ のマスタサーバ環境は、[rabbitmq-on-kubernetes](https://github.com/latonaio/rabbitmq-on-kubernetes) です。  
当該マスタサーバ環境は、同じエッジコンピューティングデバイスに配置されても、別の物理(仮想)サーバ内に配置されても、どちらでも構いません。

## RabbitMQ の Golang Runtime ライブラリ
sql-update-kube は、RabbitMQ の Golang Runtime ライブラリ として、[rabbitmq-golang-client](https://github.com/latonaio/rabbitmq-golang-client)を利用しています。

## デプロイ・稼働
sql-update-kube の デプロイ・稼働 を行うためには、aion-service-definitions の services.yml に、本レポジトリの services.yml を設定する必要があります。

kubectl apply - f 等で Deployment作成後、以下のコマンドで Pod が正しく生成されていることを確認してください。
```
$ kubectl get pods
```
