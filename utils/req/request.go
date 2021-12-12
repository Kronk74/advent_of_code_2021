package request

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Kronk74/advent_of_code_2021/utils/aocg"
	"golang.org/x/net/html"
)

func RequestInput(day int, sessionId string) []byte {
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)

	req, err := http.NewRequest("GET", url, nil)
	aocg.Check(err)

	cookie := &http.Cookie{
		Name:   "session",
		Value:  sessionId,
		MaxAge: 300,
		Domain: ".adventofcode.com",
		Path:   "/",
	}

	req.AddCookie(cookie)

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := http.Client{Transport: tr}
	res, err := client.Do(req)
	aocg.Check(err)

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	aocg.Check(err)

	return body
}

func PushAnswer(answer string, day int, part int, sessionId string) {

	data := url.Values{}
	data.Add("answer", answer)
	data.Add("level", strconv.Itoa(part))

	u, _ := url.ParseRequestURI("https://adventofcode.com")
	u.Path = fmt.Sprintf("/2021/day/%v/answer", day)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(data.Encode()))
	aocg.Check(err)

	cookie := &http.Cookie{
		Name:   "session",
		Value:  sessionId,
		MaxAge: 300,
		Domain: ".adventofcode.com",
		Path:   "/",
	}

	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := http.Client{Transport: tr}
	res, err := client.Do(req)
	aocg.Check(err)

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	aocg.Check(err)

	doc, err := html.Parse(strings.NewReader(string(body)))
	aocg.Check(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "article" {

			fmt.Println(n.FirstChild.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
