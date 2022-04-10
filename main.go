package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
  "math/rand"
)

// TypingGameを扱う情報
type GameInfo struct {
	timeLimit time.Duration
	score 		int
	name			string
}

func main() {
	// 初期化
 	g := GameInfo{time.Duration(10 * time.Second), 0, "defalut",}
	stdin := os.Stdin

	// 名前入力
	fmt.Print("名前を入力してね！(何も入力しない場合は defalut になります)\n > ")
	name := infoInput(stdin)
	if name != "" {
		g.name = name
	}
	fmt.Printf("%s さんようこそ！\n", g.name)

	// タイムリミットの設定
	fmt.Print("タイムリミットを設定してね！(正しい値が入力されない場合は 10s になります) \n > ")
	timeLimit := infoInput(stdin)
	if v, err := strconv.Atoi(timeLimit); err == nil && v > 0 {
		g.timeLimit = time.Duration(time.Duration(v) * time.Second)
	}
	fmt.Printf("タイムリミットは %v です！\n", g.timeLimit)

	// カウントダウン開始
	countdown()
	
	// ゲーム開始
	g.score = game(g.timeLimit)

	// 結果発表
	time.Sleep(time.Second * 2)
	fmt.Printf("\n %s さんのスコアは。。%d点でした！お疲れ様でした！ \n", g.name, g.score)
}

// 3秒間カウントダウンを行うコード
func countdown() {
	for i := 3; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Go!")
}

func game(t time.Duration) int {
	words, err := questions()
	if err != nil {
		fmt.Println("エラーが起きたため、ゲームを終了します。")
		os.Exit(1)
	}

	ch := gameInput(os.Stdin)
	score := 0
	timer := time.NewTimer(t)
	defer timer.Stop()
	LOOP:
	for {
		word := words[rand.Intn(len(words))]
		fmt.Printf("入力する単語：%s \n > ", word)
		select {
		case res := <-ch:
			if res == word {
				score++
				fmt.Println("+1点！！")
			} else {
				fmt.Println("不正解！！")
			}
		case <- timer.C:
			fmt.Println("\n 終了！！！")
			break LOOP
		}
	}
	return score
}

// 自分で考えたお題を出すための関数
func questions() ([]string, error) {
	var words []string
	f, err := os.Open("question.txt")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	return words, nil
}

// 初期情報を入力するための関数
func infoInput(r io.Reader) string {
	s := bufio.NewScanner(r)
	s.Scan()
	return s.Text()
}

// タイピングを入力するための関数
func gameInput(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		defer close(ch) // ここでチャネル閉じても意味なくないか？
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}
