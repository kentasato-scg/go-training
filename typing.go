package main

// Progateでやったタイピングゲームを実際に実装してみる

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	questions := []string{"dog", "cat", "apple", "orange", "lemon", "egg", "pasta"}
	const q_count = 7

	// valid
	if q_count > len(questions) {
		fmt.Println("[ERROR]質問数がクイズ数を超えています")
		return
	}

	// 本処理
	q_len := len(questions)
	score := 0
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= q_count; i++ {
		var q int
		if q_len-1 == 0 {
			q = 0
		} else {
			q = rand.Intn(q_len - 1)
		}
		ask(questions[q], &score)

		questions = unset(questions, q)
		q_len = len(questions)
	}

}

// 質問→回答のメソッド
func ask(q string, scorePtr *int) {
	fmt.Printf("次の単語を入力してください: %s\n", q)
	var input string
	fmt.Printf("入力待ち...: ")
	fmt.Scan(&input)
	fmt.Printf("入力された値: %s\n", input)

	if q == input {
		*scorePtr += 10
	} else {
		*scorePtr += 0
	}
	fmt.Printf("スコア: %d\n", *scorePtr)
}

// スライスから要素を削除
func unset(s []string, i int) []string {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
