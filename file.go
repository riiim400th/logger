package logger

import (
	"fmt"
	"os"
	"time"
)

// ログを保持するキュー
var logQueue = make(chan string, 100)

// ログのファイル書き込み
func writeToLog(filename string, logQueue <-chan string) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 現在の日付を取得
			now := time.Now()
			date := now.Format("2006-01-02")

			// ファイルをオープン（既存の場合は追記、存在しない場合は新規作成）
			file, err := os.OpenFile(fmt.Sprintf("log/%s.log", date), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening log file:", err)
				continue
			}
			defer file.Close()

			// キューからログを取り出してファイルに書き込む
			for len(logQueue) > 0 {
				logEntry := <-logQueue
				fmt.Fprintln(file, logEntry)
			}

		case logEntry := <-logQueue:
			// キューからログを取り出してファイルに即時書き込み
			now := time.Now()
			date := now.Format("2006-01-02")

			file, err := os.OpenFile(fmt.Sprintf("log/%s.log", date), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening log file:", err)
				continue
			}
			defer file.Close()

			fmt.Fprintln(file, logEntry)
		}
	}
}

// gorutine
func FileLog(msg string) {
	go writeToLog("log", logQueue)
	time.Sleep(10 * time.Second)
}
