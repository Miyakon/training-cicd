package main

import "fmt"

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
}

// 関数定義と文字列+変数
func GetHogeByGlue(piyo string) string {
	return "fuga " + piyo + " mogera" + "\n"
}

// コメント
/*
複数行コメント
何行でも書ける
*/
