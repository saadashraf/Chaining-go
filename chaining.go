package main

import "fmt"

type count struct {
	present_count int
}

func (c count) Add() count {
	c.present_count += 1

	return c
}

func (c count) Remove() count {
	c.present_count -= 1

	return c
}

func main() {
	count_status := count{1}

	count_status.Add().Remove().Add()

	fmt.Println(count_status.present_count)
}
