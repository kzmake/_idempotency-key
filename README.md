# Idempotency-Key のサイドカーパターン実装
## インフラ構成

もともと下記のようなサービスがあったとして、

```
client --request--> gateway --grpc--> service
```

そこに、Idempotency-Key を処理するproxyをサイドカーとして追加し、

```
client --request--> sidecar:"Idempotency-Key を処理するproxy" --proxy--> gateway --grpc--> service
```

とアプリケーションの実装なしで対応する例。

## 結果

```bash
$ make http
: --- 1st/2nd requests ---
curl -i -XPOST -H "Idempotency-Key: 8e03978e-40d5-43e8-bc93-6894a57f9324" localhost:58081/v1/now&
curl -i -XPOST -H "Idempotency-Key: 8e03978e-40d5-43e8-bc93-6894a57f9324" localhost:58081/v1/now

HTTP/1.1 409 Conflict
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 20 Feb 2022 17:19:35 GMT
Content-Length: 9

conflict

HTTP/1.1 200 OK
Content-Length: 66
Content-Type: application/json
Date: Sun, 20 Feb 2022 17:19:38 GMT
Grpc-Metadata-Content-Type: application/grpc

{"now":"2022-02-20 17:19:35.842370254 +0000 UTC m=+324.852386024"}

: --- delay 3 sec ---

: --- 3rd request ---
curl -s -XPOST -H "Idempotency-Key: 8e03978e-40d5-43e8-bc93-6894a57f9324" localhost:58081/v1/now
{"now":"2022-02-20 17:19:35.842370254 +0000 UTC m=+324.852386024"}

:  --- 4th request ---
curl -s -XPOST -H "Idempotency-Key: 12345678-1234-4321-1234-123456789abc" localhost:58081/v1/now
{"now":"2022-02-20 17:19:41.920781674 +0000 UTC m=+330.930797443"}
```

## Run

```
skaffold dev
```

```
make http
```
