package main

import "fmt"

func main() {

	var time, second, minute, hour int

	fmt.Scan(&time)

	second = time % 60
	minute = (time % 3600) / 60
	hour = time / 3600

	fmt.Printf("%d jam %d menit %d detik", hour, minute, second)
}
