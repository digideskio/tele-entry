package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	var (
		addr           string
		timezone       string
		textNumbersStr string
		textNumbers    []string
	)

	flag.StringVar(&addr, "addr", ":3000", "address to listen to (default is 0.0.0.0:3000)")
	flag.StringVar(&timezone, "tz", "UTC", "timezone (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)")
	flag.StringVar(&textNumbersStr, "text", "", "comma-separated numbers to text")
	flag.Parse()

	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(err)
	}

	if len(textNumbersStr) > 0 {
		textNumbers = strings.Split(textNumbersStr, ",")
	}

	http.HandleFunc("/entry", func(w http.ResponseWriter, r *http.Request) {

		log.Println("entry called")

		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8" ?>
<Response>
    <Say voice="alice">Welcome</Say>
    <Pause length="1" />
    <Play digits="6" />
    <Pause length="1" />
    <Play digits="6" />
    `)
		message := fmt.Sprintf("Someone has arrived at your building at %s.", time.Now().In(location).Format(time.Kitchen))
		for _, number := range textNumbers {
			fmt.Fprintf(w, `<Sms to="+%s">%s</Sms>`, number, message)
		}
		w.Write([]byte("</Response>"))
	})

	http.ListenAndServe(":3000", nil)
}
