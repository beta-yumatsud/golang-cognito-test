# 概要
Amazon Cognito周りの検証をGo言語アプリケーションでやってみる。

# アプリケーション
[go-chi/chi](https://github.com/go-chi/chi) でリクエストを受け付けるのみ。
`/header` でCognito経由の際のリクエスト内容（特にHeader情報）をみてみる目的。
これだけだったらLambdaでもよかったが、実際の想定使い所はECS on Fargateのため用意。

# 使い方
* アプリケーション単体を動かすのは下記
```shell
$ go run -race main.go
```
* `Docker` で動かす場合は下記
```shell
$ ./build
$ docker build -t golang-cognito-test .
$ docker run -p 8080:8080 -d golang-cognito-test
```

# Reference
- [Notion（閲覧制限あり）](https://www.notion.so/Cognito-22e861414ba041f897823a165e259c92)
