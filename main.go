package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/tjarratt/babble"
	"io"
	"os"
	"time"
)

// タイムリミットのデフォルト10秒
const timeLimit = 10

// スコア初期値
var score = 0

var (
	babbler  babble.Babbler
	question string
	d        time.Duration
)

func init() {
	flag.DurationVar(&d, "d", timeLimit*time.Second, "duration flag")
	flag.Parse()

	babbler = babble.NewBabbler()
	babbler.Count = 1
}

func main() {
	startGame()
}

// こうすることで_main関数が終わると同時のキャンセル処理が走る
func startGame() {
	// ゲーム開始前のカウントダウン表示
	countdown()

	// コンテキストのキャンセル処理
	// ゴルーチンをまたいだ処理のキャンセルに使う
	// done := time.After(timeLimit * time.Second)
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, d)
	defer cancel()

	start := time.Now()
	questions()

	ch := input(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n\nTime UP!!")
			fmt.Println("score:", score)
			return // この行がないと無限ループになる
		case v := <-ch:
			if v == question {
				score++
				fmt.Println("Good!")
			} else {
				fmt.Println("oops...")
			}
			end := time.Now()
			fmt.Printf("%d秒経過\n", int((end.Sub(start)).Seconds()))
			questions()
		}
	}

}

func countdown() {
	for i := 3; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Go!")
}

func questions() {
	question = babbler.Babble()
	fmt.Println("\ntype this:", question)
	fmt.Print("> ")
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		// 標準入力から１行ずつ文字をよみこむ
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
