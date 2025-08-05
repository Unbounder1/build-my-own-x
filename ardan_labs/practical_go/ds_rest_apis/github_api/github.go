package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println(UserInfo("ardanlabs"))
}

func demo() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
	if err != nil {
		fmt.Println("error:", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error bad status - %s \n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)

	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}
	// io.Copy(os.Stdout, resp.Body)

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		return
	}
	fmt.Println(reply.Name, reply.NumRepos)
}

func UserInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}
	// io.Copy(os.Stdout, resp.Body)

	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		return "", 0, err
	}

	return reply.Name, reply.NumRepos, nil

}
