package main

import (
	"fmt"
	"time"
)

// Функция для реализации ожидания
func CustomSleep(duration time.Duration) {
	done := make(chan bool)
	go func() {
		time.AfterFunc(duration, func() {
			done <- true
		})
	}()

	<-done
}

// Суть программы заключается в точ то мы запискаем горутин на определенное время и не даем выполняться программе далбше пока не закончится горутин
func main() {
	fmt.Println("Начало")
	CustomSleep(6 * time.Second)
	fmt.Println("Прошло 6 секунд")
}
