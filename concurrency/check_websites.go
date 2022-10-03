package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	// synchronously
	// for _, url := range urls {
	// 	results[url] = wc(url)
	// }

	// asynchronously
	/*The body of the anonymous function is just the same as the
	  as the loop body was before. The only difference is that each
	  iteration of the loop will start a new goroutine, concurrent
	  with the current process (the WebsiteChecker function) each of
	  which will add its result to the results map.
	*/
	// var wg sync.WaitGroup
	for _, url := range urls {
		go func(u string) {
			resultChan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.string] = r.bool
	}

	return results
}

// By giving each anonymous function a parameter for the url - u - and
// then calling the anonymous function with the url as the argument,
// we make sure that the value of u is fixed as the value of url
// for the iteration of the loop that we're
// launching the goroutine in. u is a copy of the value
// of url, and so can't be changed.
