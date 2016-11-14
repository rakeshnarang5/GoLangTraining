package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Lang struct {
	Name  string
	URL   string
	Bytes uint
	Time  string
}

func setNameURLLang(Name string, URL string) Lang {
	temp := Lang{Name: Name, URL: URL}
	return temp
}

func crawl(pfunc func(Lang), lang Lang, wg *sync.WaitGroup, dataChannel *chan int64, timeChannel *chan time.Duration) {
	defer wg.Done()
	initial := time.Now()
	resp, err := http.Get(lang.URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.Copy(ioutil.Discard, resp.Body)
	lang.Bytes = uint(body)
	final := time.Since(initial).String()
	lang.Time = final
	*dataChannel <- body
	*timeChannel <- time.Since(initial)
	pfunc(lang)
}

func pfunc(lang Lang) {
	fmt.Println("\nGolang's default formatting")
	fmt.Printf("%v\n", lang)
	fmt.Println("JSON formatting")
	b, err := json.Marshal(lang)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func main() {
	site1 := setNameURLLang("Python", "http://python.org/")
	site2 := setNameURLLang("Ruby", "http://ruby-lang.org/en/")
	site3 := setNameURLLang("Golang", "http://golang.org/")
	var wg sync.WaitGroup
	var langs = []Lang{
		site1,
		site2,
		site3,
	}
	initial := time.Now()
	dataChannel := make(chan int64, 3)
	timeChannel := make(chan time.Duration, 3)
	for _, lang := range langs {
		wg.Add(1)
		go crawl(pfunc, lang, &wg, &dataChannel, &timeChannel)
	}
	wg.Wait()
	final := time.Since(initial)
	fmt.Println("Total time: \n", final)
	var totalBytes int64
	var totalTime time.Duration
	for i := 1; i <= 3; i++ {
		totalBytes += <-dataChannel
		totalTime += <-timeChannel
	}
	fmt.Printf("Cumulative Bytes: %d\n", totalBytes)
	fmt.Printf("Cumulative Time: %s\n", totalTime.String())
}
