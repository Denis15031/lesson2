package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func main() {
	cache := NewCache()
	defer cache.Close()

	cache.Set("user", User{Name: "Alice"}, time.Hour)

	cache.Set("temp", 42, 2*time.Second)

	if val, ok := cache.Get("user"); ok {
		fmt.Printf("User: %+v\n", val)
	}

	if user, err := GetAs[User](cache, "user"); err == nil {
		fmt.Printf("Typed: %+v\n", user)
	}

	time.Sleep(3 * time.Second)

	fmt.Println("Exists(temp):", cache.Exists("temp"))

	jsonData, err := cache.ToJSON()
	if err != nil {
		fmt.Printf("ToJSON error: %v\n", err)
	} else {
		fmt.Println("JSON:", string(jsonData))
	}

	cache.Clear()

	fmt.Println("After clear, Exists(user):", cache.Exists("user"))
}
