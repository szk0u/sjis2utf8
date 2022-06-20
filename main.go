package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "引数に変換対象のファイルのパスを指定してください。")
		os.Exit(1)
		return
	}

	for _, v := range os.Args[1:] {
		// ファイルの存在チェック
		_, err := os.Stat(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%sは存在しません\n", v)
			os.Exit(1)
			return
		}

		err = handleFilepath(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%sの変換に失敗しました。%v\n", v, err)
			os.Exit(1)
			return
		}
	}
}

func handleFilepath(path string) error {
	// 元のファイル名を元にファイルを作成
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	ext := filepath.Ext(path)

	pathWithoutExtension := strings.TrimSuffix(absPath, ext)

	newPath := pathWithoutExtension + "_converted" + ext

	existingFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer existingFile.Close()

	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := &runeWriter{transform.NewWriter(newFile, japanese.ShiftJIS.NewEncoder())}
	tee := io.TeeReader(existingFile, writer)
	s := bufio.NewScanner(tee)
	for s.Scan() {
	}

	return nil
}

type runeWriter struct {
	w io.Writer
}

func (rw *runeWriter) Write(b []byte) (int, error) {

	var err error

	l := 0

loop:

	for len(b) > 0 {

		_, n := utf8.DecodeRune(b)

		if n == 0 {

			break loop

		}

		_, err = rw.w.Write(b[:n])

		if err != nil {

			_, err = rw.w.Write([]byte{'?'})

			if err != nil {

				break loop

			}

		}

		l += n

		b = b[n:]

	}

	return l, err

}
