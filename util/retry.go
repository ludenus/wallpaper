package util

import (
	"time"
	"log"
	"fmt"
)

// shamelessly borrowed from https://blog.abourget.net/en/2016/01/04/my-favorite-golang-retry-function/
//
// USAGE:
//
//var signedContent []byte
//err := retry(5, 2*time.Second, func ()(err error) {
//	signedContent, err = signFile(unsignedFile, contents)
//	return
//})
//if err != nil {
//	log.Println(err)
//	http.Error(w, err.Error(), 500)
//	return
//}

func retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func retryDuring(duration time.Duration, sleep time.Duration, callback func() error) (err error) {
	t0 := time.Now()
	i := 0
	for {
		i++

		err = callback()
		if err == nil {
			return
		}

		delta := time.Now().Sub(t0)
		if delta > duration {
			return fmt.Errorf("after %d attempts (during %s), last error: %s", i, delta, err)
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
}
