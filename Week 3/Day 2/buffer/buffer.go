package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
	} else {
		fmt.Println(url, "Still running")
	}
}
func main() {

	// dataPostRaw()
	dataPostFormData()

	urls := []string{
		"https://github.com/",
		"https://www.kolase.com/",
		"https://www.facebook.com/",
		"https://www.kickstarter.com/",
	}

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			checkUrl(url)
		}(u)
	}

	wg.Wait()

}
func errormessage(err error) {
	if err != nil {
		fmt.Println("there is problem")
	}
}

// func dataPostRaw() {
// 	reqBody, err := json.Marshal(map[string]string{
// 		"name": "arman",
// 	})
// 	errormessage(err)

// 	resp, err := http.Post("https://httpin.org/post"),
// 		"appilication/json",
// 	bytes.NewBuffer(reqBody)),
// 	errormessage(err)

// 	body, err := ioutil.ReadAll(resp.Body)
// 	errormessage(err)
// 	fmt.Println(string(body))
// }

func dataPostFormData() {
	reqBody := url.Values{
		"name": {"arman"},
	}

	resp, err := http.PostForm("https://httpin.org/post", reqBody)
	errormessage(err)

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result["form"])
}
