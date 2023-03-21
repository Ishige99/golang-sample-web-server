package main

import (
	"bufio"
	"os"
)

type WardList struct {
	Wards []string
}

func New(wards []string) WardList {
	return WardList{Wards: wards}
}

func FileRead(fileName string) []string {
	var wardList []string
	file, err := os.Open(fileName) // 引数で受け取った該当ファイルを開く

	if os.IsNotExist(err) {
		return nil
	}
	defer file.Close() // ファイルは関数終了時に閉じる

	scanner := bufio.NewScanner(file) // 引数で受け取ったファイルを1行ずつ読み込みます

	// 1行ずつ読み取ったファイルを1行ずつappendしていく
	// .Scan()は次の行にデータがあるまでtrueを返すので、先ほど取得した行分だけ実行されます。
	for scanner.Scan() {
		wardList = append(wardList, scanner.Text()) // .Text()は現在の行のテキストを全て取得します。
	}

	return wardList
}
