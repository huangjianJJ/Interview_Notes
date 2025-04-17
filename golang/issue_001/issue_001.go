package issue001

import (
	"fmt"
	"time"
)

// 一个goroutine只打印一个字母，写出一个程序交替打印ABCD

func PrintABCD() {
	signal := make(chan string, 1)

	go func() {
		for {
			v := <-signal
			if v == "A" {
				fmt.Println("A")
				signal <- "B"
				continue
			}
			signal <- v
		}
	}()

	go func() {
		for {
			v := <-signal
			if v == "B" {
				fmt.Println("B")
				signal <- "C"
				continue
			}
			signal <- v
		}
	}()

	go func() {
		for {
			v := <-signal
			if v == "C" {
				fmt.Println("C")
				signal <- "D"
				continue
			}
			signal <- v
		}
	}()

	go func() {
		for {
			v := <-signal
			if v == "D" {
				fmt.Println("D")
				signal <- "A"
				continue
			}
			signal <- v
		}
	}()

	signal <- "A"

	time.Sleep(1 * time.Minute)
}

//Deepseek

// func main() {
//     // 创建四个无缓冲的channel用于同步
//     a := make(chan struct{})
//     b := make(chan struct{})
//     c := make(chan struct{})
//     d := make(chan struct{})

//     // 启动goroutine打印A
//     go func() {
//         for {
//             <-a       // 等待a的信号
//             fmt.Print("A")
//             b <- struct{}{} // 触发B
//         }
//     }()

//     // 启动goroutine打印B
//     go func() {
//         for {
//             <-b       // 等待b的信号
//             fmt.Print("B")
//             c <- struct{}{} // 触发C
//         }
//     }()

//     // 启动goroutine打印C
//     go func() {
//         for {
//             <-c       // 等待c的信号
//             fmt.Print("C")
//             d <- struct{}{} // 触发D
//         }
//     }()

//     // 启动goroutine打印D
//     go func() {
//         for {
//             <-d       // 等待d的信号
//             fmt.Print("D")
//             a <- struct{}{} // 循环触发A
//         }
//     }()

//     // 初始化，触发A开始执行
//     a <- struct{}{}

//     // 阻塞主goroutine，防止程序退出
//     select {}
// }
