```shell
# カバレッジ出力
go test -v -cover
# テスト未実施の箇所をファイルに出力
go test -v -cover -coverprofile=cover.out
# カバレッジを可視化ツール
go tool cover -html=cover.out -o cover.html
```

[CodeCov](https://about.codecov.io/)