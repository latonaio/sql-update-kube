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

### 複数外部キーへの対応
本マイクロサービスで、複数の外部キーが存在するSQLの登録更新を行うためには、登録更新モジュールとして利用している、[SQLBOILER](https://github.com/volatiletech/sqlboiler)に、以下のソースコードを追加した上で、使用する必要があります。  
SQLBOILERの、boilingcore / boilingcore.goにおける、func Newの配下に、以下のソースコードを配置して、使用してください。  

```
for i, t := range s.Tables {
        var fkeys []drivers.ForeignKey
        var toonerelationships []drivers.ToOneRelationship
        var tomanyrelationships []drivers.ToManyRelationship
        var relalias_foreigns, relalias_tor_locals, relalias_tmr_locals []string

        // eliminate duplicates, if there are multi-column ForeignKeys
        if len(t.FKeys) > 1 {
            for j, fk := range t.FKeys {
                exists := false
                relalias_foreign := s.Config.Aliases.Table(t.Name).Relationship(fk.Name).Foreign
                if j == 0 {
                    relalias_foreigns = append(relalias_foreigns, relalias_foreign)
                    fkeys = append(fkeys, fk)
                    continue
                }
                for _, rf := range relalias_foreigns {
                    if rf == relalias_foreign {
                        exists = true
                        break
                    }
                }
                if !exists {
                    relalias_foreigns = append(relalias_foreigns, relalias_foreign)
                    fkeys = append(fkeys, fk)
                }
            }
            s.Tables[i].FKeys = fkeys
        }

        // eliminate duplicates, if there are multi-column ToOneRelationships
        if len(t.ToOneRelationships) > 1 {
            for j, tor := range t.ToOneRelationships {
                exists := false
                relalias_tor_local := s.Config.Aliases.Table(tor.ForeignTable).Relationship(tor.Name).Local
                if j == 0 {
                    relalias_tor_locals = append(relalias_tor_locals, relalias_tor_local)
                    toonerelationships = append(toonerelationships, tor)
                    continue
                }
                for _, rol := range relalias_tor_locals {
                    if rol == relalias_tor_local {
                        exists = true
                        break
                    }
                }
                if !exists {
                    relalias_tor_locals = append(relalias_tor_locals, relalias_tor_local)
                    toonerelationships = append(toonerelationships, tor)
                }
            }
            s.Tables[i].ToOneRelationships = toonerelationships
        }

        // eliminate duplicates, if there are multi-column ToManyRelationships
        if len(t.ToManyRelationships) > 1 {
            for j, tmr := range t.ToManyRelationships {
                exists := false
                relalias_tmr_local := s.Config.Aliases.ManyRelationship(tmr.ForeignTable, tmr.Name, tmr.JoinTable, tmr.JoinLocalFKeyName).Local
                if j == 0 {
                    relalias_tmr_locals = append(relalias_tmr_locals, relalias_tmr_local)
                    tomanyrelationships = append(tomanyrelationships, tmr)
                    continue
                }
                for _, rml := range relalias_tmr_locals {
                    if rml == relalias_tmr_local {
                        exists = true
                        break
                    }
                }
                if !exists {
                    relalias_tmr_locals = append(relalias_tmr_locals, relalias_tmr_local)
                    tomanyrelationships = append(tomanyrelationships, tmr)
                }
            }
            s.Tables[i].ToManyRelationships = tomanyrelationships
        }
    }
```
