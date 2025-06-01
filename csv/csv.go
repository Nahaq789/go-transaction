package csv

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"time"
)

const (
	minLength  int32  = 3
	maxLength  int32  = 5
	fileSize   int32  = 1024 * 1024
	collection string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func createFIle() string {
	year, month, day := time.Now().Date()
	filePath := fmt.Sprintf("./csvfile/%v-%v-%v.csv", year, month, day)

	_, err := os.Create(filePath)
	if err != nil {
		slog.Warn("ディレクトリが存在しません。", "Warn", err)
		log.Println("ディレクトリを作成します。")
		if err = os.Mkdir("./csvfile", 0750); err != nil {
			slog.Warn("ディレクトリ作成に失敗しました。", "error", err)
			return ""
		}
	}
	return filePath
}

func writeFile(filePath string, content []byte) error {
	if err := os.WriteFile(filePath, content, 0666); err != nil {
		slog.Error("ファイル書き込みに失敗しました。", "error", err)
		return err

	}
	return nil
}

func generateContent() []byte {
	b := make([]byte, random())
	return b
}

func random() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31()*(maxLength-minLength) + minLength
}

func CreateCsvFIle() {
	filePath := createFIle()
	if len(filePath) < 0 {
		return
	}

	content := []byte("hoge")
	if err := writeFile(filePath, content); err != nil {
		log.Printf("%v", err)
		return
	}
}
