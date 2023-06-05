package main

import (
	"fmt"

	"github.com/kevinmtanadi/yt-dl/youtube"
)

func main() {
	fmt.Println("Hello")
	yt := youtube.Youtube{
		URL: "https://www.youtube.com/watch?v=akHMQOgGrd8&ab_channel=VickeBlanka",
	}
	err := yt.ExecuteRequest()
	if err != nil {
		fmt.Println(err)
	}

}
