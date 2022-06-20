package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "引数に変換対象のファイルのパスを指定してください。")
		os.Exit(1)
		return
	}

	for _, v := range os.Args[1:] {
		err := handleFilepath(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%sは存在しません\n", v)
			os.Exit(1)
			return
		}
	}
}

func handleFilepath(filepath string) error {
	// ファイルの存在チェック
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	// 元のファイル名を元にファイルを作成

	// utf8からsjisに変換

	return nil
}
