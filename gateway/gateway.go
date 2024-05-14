package gateway

import "fmt"

type Gateway struct{}

func (g Gateway) Send(userID, message string) error {
	fmt.Printf("user: %s sending message: %s", userID, message)
	return nil
}
