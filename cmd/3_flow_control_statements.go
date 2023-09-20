package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// iを使わない場合は省略可能
	total := 1
	for ; total < 10; { // ;も省略可能 → for total < 10 {
		total += total
	}
	fmt.Println(total)

	fmt.Println(judge(12))

	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os { // switchで使うosという変数をswitch文の定義と同時に作成している
	case "darwin":
		fmt.Println("OS X");
	case "Linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switchの条件を省略することも可能
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer: deferへ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる
	defer fmt.Printf("World")
	fmt.Printf("Hello")

	// deferへ渡す関数が複数ある場合、その呼び出しはスタック（stack）)される。
	// 呼び出し元の関数がreturnするとき、deferへ渡した関数はLIFO(last-in-first-out)の順番で実行される。
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func judge(score int) string {
	if score < 20 { // ifの条件部分の()は省略可能。{}は省略不可。
		return "まだまだだね"
	} else {
		return "いい感じ"
	}
}