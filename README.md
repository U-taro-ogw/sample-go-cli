# cliの作り方を学ぶ
web APIにfetchして出力する

# 準備
下記のようなjson返却APIを立てる（`json-server`を利用）

```bash
$ curl http://localhost:3000/dummy_hoge
{
  "id": 1,
  "text": "this is dummy hoge text",
  "created_at": "2020-05-08T12:00:00Z",
  "updated_at": "2020-05-09T20:00:30Z"
}
```

# 作成手順
1. go mod init hoge
1. go get -u github.com/spf13/cobra/cobra
1. cobra init --pkg-name hoge
