package main

import (
	"fmt"
	"modak/notification"
)

func main() {
	service := notification.NewService(notification.NewRateLimiter())
	err := service.Send("News", "user1", "news 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = service.Send("News", "user1", "news 2")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = service.Send("News", "user1", "news 3")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = service.Send("News", "user2", "news 1")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = service.Send("Marketing", "user1", "update 1")
	if err != nil {
		fmt.Println(err.Error())
	}
}
