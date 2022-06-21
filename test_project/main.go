package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"encoding/json"
	"net/http"

	"github.com/buildkite/interpolate"
)

func main() {
	fmt.Printf("Hello world\n")

	// 最初の変数への代入と複数行の書き方
	multi_line := `Hey!! we
	are going to
	write multiline strings
	in Go.
	`

	// 2回目以降の変数の代入と複数行の書き方2
	multi_line = "Hey!! we /n" +
		"are going to \n" +
		"write multiline strings \n" +
		"in Go. \n"

	fmt.Printf("%s", multi_line)

	// fmt.Printf(GetHogeByGlue("piyo"))
	fmt.Println(GetHogeByGlue("foo"))        // 文字列結合子（+）を使ったサンプル
	fmt.Println(GetHogeBySprintf("foo"))     // fmt パッケージを使ったサンプル
	fmt.Println(GetHogeByTemplate("foo"))    // template パッケージを使ったサンプル
	fmt.Println(GetHogeByInterpolate("foo")) // interpolate パッケージを使ったサンプル

	// フレームワークを使わずにサーバを立てる
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
		message := map[string]string{
			"message": "hello world",
		}
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			panic(err.Error())
		}
		w.Write(jsonMessage)
	})
	http.ListenAndServe("127.0.0.1:3000", mux)
}

// 関数定義と文字列+変数
func GetHogeByGlue(piyo string) string {
	return "fuga " + piyo + " mogera" + "\n"
}

// フォーマット指定
func GetHogeBySprintf(piyo string) string {
	return fmt.Sprintf("fuga %s mogera\n", piyo)
}

// Template関数

type FieldsToReplace struct {
	Replace1 string
}

func GetHogeByTemplate(piyo string) string {
	var msg_result bytes.Buffer

	msg_tpl, msg_err := template.New("myTemplate").Parse("fuga {{.Replace1}} mogera\n")
	if msg_err != nil {
		log.Fatal(msg_err)
	}
	replace_to := FieldsToReplace{
		Replace1: piyo,
	}
	if msg_err := msg_tpl.Execute(&msg_result, replace_to); msg_err != nil {
		log.Fatal(msg_err)
	}

	return msg_result.String()
}

func GetHogeByInterpolate(piyo string) string {
	env := interpolate.NewSliceEnv([]string{
		"Replace2=" + piyo,
	})

	output, _ := interpolate.Interpolate(env, "fuga ${Replace2} mogera ${Replace3:-🏖}")

	return output
}

// コメント
/*
複数行コメント
何行でも書ける
*/
