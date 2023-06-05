package youtube

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Youtube struct {
	URL      string
	Title    string
	Duration time.Time
}

func (yt *Youtube) ExecuteRequest() error {
	// base_header := map[string]interface{}{
	// 	"User-Agent":      "Mozila/5.0",
	// 	"accept-language": "en-US,en",
	// }

	client := &http.Client{}
	if strings.HasPrefix(strings.ToLower(yt.URL), "http") {
		req, err := http.NewRequest("GET", yt.URL, nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))
	}

	return nil
}

func (yt *Youtube) Scrape() {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
