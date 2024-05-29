package dftest

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/rocketlaunchr/dataframe-go/exports"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"github.com/tobgu/qframe"
)

// ３種類のベンチマークコードを取得する方法

func benchmarkQFrame(b *testing.B) {
	bs, err := os.ReadFile("iris.csv")
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		qf := qframe.ReadCSV(bytes.NewReader(bs))
		qf = qf.Select("sepal_length", "species")
		err = qf.ToCSV(&out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGota(b *testing.B) {
	bs, err := os.ReadFile("iris.csv")
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		df := dataframe.ReadCSV(bytes.NewReader(bs))
		df = df.Select([]string{"sepal_length", "species"})
		err = df.WriteCSV(&out, dataframe.WriteHeader(true))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataframeGo(b *testing.B) {
	bs, err := os.ReadFile("iris.csv")
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var out bytes.Buffer
		df, err := imports.LoadFromCSV(context.Background(), bytes.NewReader(bs))
		if err != nil {
			b.Fatal(err)
		}

		df.RemoveSeries("sepal_width")
		df.RemoveSeries("petal_length")
		df.RemoveSeries("petal_width")
		err = exports.ExportToCSV(context.Background(), &out, df)
		if err != nil {
			b.Fatal(err)
		}
	}
}
