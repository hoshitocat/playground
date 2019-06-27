package main

import (
	"errors"
	"fmt"

	"github.com/getsentry/raven-go"
)

func init() {
	//err := sentry.Init(sentry.ClientOptions{Dsn: "https://38abe4c46f224bb78d4b10327eed6357@sentry.io/1490577"})
	//if err != nil {
	//	fmt.Printf("Sentry initialization failed: %v\n", err)
	//}
}

func main() {
	err := raven.SetDSN("https://38abe4c46f224bb78d4b10327eed6357@sentry.io/1490577")
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	raven.CaptureErrorAndWait(errors.New("hogehogehoge"), nil)
	//raven.CaptureMessage("Something bad happened and I would like to know about that")

	//sentry.CaptureException(errors.New("hogehogehoge"))
}
