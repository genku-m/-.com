## スーパー支払い君.com
### 開発環境での使い方
開発環境ではDockerを用いてDBに接続します。
- Dockerを立ち上げる
```
docker-compose up -d
```
- コンテナに入る
```
docker exec -it  db bash
```
- DBに接続
```
mysql -u{user} -p{password} {db_name}
```
initに入ってるCreate Table等をコピーしてテーブル作成
（本来Dockerを立ち上げた時に実行されてほしいができていない）

- サーバーを立ち上げる(もしくはBuildする)
```
go run main.go
```
- ユーザーログインをする
  
事前に用意したユーザーを使用します

ここではcurlを利用した例を挙げます
```
curl -c cookie.txt -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"password"}' localhost:8080/login
```
- 請求書を作成する
```
curl -b cookie.txt -X POST -H "Content-Type: application/json" -d '{"company_guid":"company-1","customer_guid":"customer-1","publish_date":"2024-04-07T17:44:13Z","payment": 10000,"commission_tax_rate":0.4,"tax_rate":0.01,"payment_date":"2024-04-07T17:44:13Z"}' localhost:8080/api/invoices | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   502  100   310  100   192  16131   9991 --:--:-- --:--:-- --:--:-- 26421
{
  "guid": "coalnlon1e4be1vv9eg0",
  "company_guid": "company-1",
  "customer_guid": "customer-1",
  "publish_date": "2024-04-07T17:44:13Z",
  "payment": 10000,
  "commission_tax": 4000,
  "commission_tax_rate": 0.4,
  "consumption_tax": 40,
  "tax_rate": 0.01,
  "billing_amount": 14040,
  "payment_date": "2024-04-07T17:44:13Z",
  "status": "unprocessed"
}
```
- 期間内の請求書一覧を見る
```
curl -b cookie.txt 'localhost:8080/api/invoices?first_payment_date=2024-04-01T00:00:00Z&last_payment_date=2024-04-05T00:00:00Z' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   305  100   305    0     0  61641      0 --:--:-- --:--:-- --:--:-- 76250
[
  {
    "guid": "invoice-2",
    "company_guid": "company-1",
    "customer_guid": "customer-1",
    "publish_date": "2024-04-01T00:00:00Z",
    "payment": 100000,
    "commission_tax": 1000,
    "commission_tax_rate": 0.01,
    "consumption_tax": 1000,
    "tax_rate": 0.01,
    "billing_amount": 102000,
    "payment_date": "2024-04-02T00:00:00Z",
    "status": "processing"
  }
]
```


