package gohosts

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	domains = map[string]bool{}
	dLock   = &sync.RWMutex{}
)

func Write(fName string) {
	out, err := os.Create(fName)
	defer out.Close()
	if err != nil {
		fmt.Println(err)
	}
	checkErr := func(err error) {
		if err != nil {
			fmt.Println(err)
		}
	}
	for k, _ := range domains {
		_, err = out.WriteString("0.0.0.0 ")
		checkErr(err)
		_, err = out.WriteString(k)
		checkErr(err)
		_, err = out.WriteString("\n")
		checkErr(err)
	}

}
func Merge(sources []string) {
	var wg sync.WaitGroup
	wg.Add(len(sources))

	for _, s := range sources {
		go func() {
			mergeOne(s)
			wg.Done()
		}()
	}
	wg.Wait()
}
func mergeOne(s string) {
	var err error
	var f io.ReadCloser
	respCode := 0
	if strings.HasPrefix(s, "http") {
		client := http.Client{}
		var resp *http.Response
		resp, err = client.Get(s)

		if resp != nil {
			f = resp.Body
			respCode = resp.StatusCode
		}

	} else {
		f, err = os.Open(s)
	}
	defer f.Close()

	if err != nil || respCode != http.StatusOK {
		return
	}

	reader := bufio.NewReader(f)

	for {
		hostLine, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if !strings.HasPrefix(hostLine, "1") && !strings.HasPrefix(hostLine, "0") {
			continue
		}
		if strings.HasPrefix(hostLine, "1") {
			hostLine = strings.TrimPrefix(hostLine, "127.0.0.1")
		} else {
			hostLine = strings.TrimPrefix(hostLine, "0.0.0.0")
		}
		hostLine = strings.TrimSpace(hostLine)
		if strings.EqualFold(hostLine, "localhost") ||
			strings.EqualFold(hostLine, "127.0.0.1") ||
			strings.EqualFold(hostLine, "0.0.0.0") {
			continue
		}
		dLock.RLock()
		_, ok := domains[hostLine]
		dLock.RUnlock()
		if !ok {
			dLock.Lock()
			domains[hostLine] = true
			dLock.Unlock()
		}
	}

}
