## ベンチマーク

```go
package something

import (
    "testing"
)

func DoSomething() {}

func BenchmarkDoSomething(b *testing.B) {
    for i := 0; i < b.N; i++ {
        DoSomething()
    }
}
```

## ベンチマーク実行

```shell
# 処理時間計測
go test -bench DoSomething

# メモリアロケーション量も計測
go test -benchmem -bench DoSomething

# -benchmem 一処理あたりのメモリアロケーション数が測れる
# B/op 一回のオペレーションで何バイトのm根守が確保が実行されたか
# allocs/op １回のオペレーションで何回メモリ確保が実行されたか
```

## 改善方法

```go
// 計測対象のコード
package something

import "fmt"

func makeSomething(n int) []string {
    var r []string
    for i := 0; i < n; i++ {
        r = append(r, fmt.Sprintf("%05d test", i))
    }
    return r
}

// 計測コード
package something

import "testing"

func BenchmarkMakeSomething(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = makeSomething(1000)
    }
}
```
コマンドを実行して結果を確認
```shell
go test -count 5 -benchmem -bench . 2>&1 | tee old.log
```

## ソースコードの改善
```go
func makeSomething(n int) []string {
    r := make([]string, n, n)
    for i := i < n; i++ {
        r[i] = fmt.Sprintf("%05d test", i)
    }

    return r
}
```

改善後の結果を計測
```shell
go test -count 5 -benchmem -bench . 2>&1 | tee new.log
```

## benchstatを利用して改善前後の結果を比較

```shell
benchstat old.log new.log
```