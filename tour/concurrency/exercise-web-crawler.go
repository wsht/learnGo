package concurrency

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var Logs *log.Logger
var FileName = "./test.log"

func MyLog() (*log.Logger, error) {
	if Logs == nil {
		file, error := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, 0755)
		fmt.Println("new logger")
		if error != nil {
			return Logs, error
		}
		Logs = log.New(file, "", log.LstdFlags)
	}

	return Logs, nil
}

var urlMap = make(map[string]bool)

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	mylog, error := MyLog()
	if error != nil {
		fmt.Println(error)
		return
	}

	if urlMap[url] {
		return
	}

	if depth <= 0 {
		return
	}
	//set the url is fetch
	urlMap[url] = true

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		// log.Println(err)
		mylog.Println(err)
		return
	}
	mylog.Printf("found: %s %q\n", url, body)
	// log.Printf("found: %s %q\n", url, body)
	time.Sleep(100 * time.Millisecond)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}

	return
}

func ExerciseWebCrawler() {

	log1, error := MyLog()
	if error != nil {
		fmt.Println(error)
	}
	log2, error := MyLog()
	if error != nil {
		fmt.Println(error)
	}
	log1.Output(2, "123")
	log2.Output(2, "456")
	go Crawl("http://golang.org/", 4, fetcher)
	log.Println(123)

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
