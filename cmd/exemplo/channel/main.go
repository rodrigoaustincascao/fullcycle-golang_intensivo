package main

import "time"

/*
Primeira forma
*/

// func main() {
// 	canal := make(chan string)

// 	go func() {
// 		canal <- "Olá Mundo"
// 	}()

// 	result := <-canal
// 	println(result)
// }

/*
Segunda forma
*/
// func x(ch chan string) {
// 	time.Sleep(time.Second * 5)
// 	ch <- "Olá Mundo!"
// }

// func main() {
// 	canal := make(chan string)
// 	go x(canal)

// 	println(<-canal)
// }

/*
Terceira forma
*/
// func contador(ch chan int) {
// 	i := 0
// 	for {
// 		time.Sleep(time.Second * 2)
// 		ch <- i
// 		i++
// 	}
// }

// func main() {
// 	canal := make(chan int)

// 	go contador(canal)
// 	for x := range canal {
// 		println(x)
// 		if x == 10 {
// 			close(canal)
// 			println("Canal fechado!")
// 		}
// 	}

// }

/*
Quarta Forma
*/
func worker(workerId int, msg chan int) {
	for res := range msg {
		println("worker: ", workerId, "recebeu: ", res)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	for i := 0; i < 5; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 10; i++ {
		canal <- i
	}
}
