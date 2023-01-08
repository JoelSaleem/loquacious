package internal

import (
	"fmt"
	"net/http"
)

func Login(serverAddress, userID string) error {
	req, err := http.NewRequest("POST", serverAddress+"/login/"+userID, nil)
	if err != nil {
		fmt.Println("err", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("login: %v", err)
	}
	defer resp.Body.Close()

	return nil
}

func Logout(serverAddress, userID string) error {
	req, err := http.NewRequest("POST", serverAddress+"/logout/"+userID, nil)
	if err != nil {
		fmt.Println("err", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("logout: %v", err)
	}
	resp.Body.Close()
	return nil
}
