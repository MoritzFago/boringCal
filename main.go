// webuntistoics project main.go
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"time"

	"code.google.com/p/go.net/publicsuffix"
)

func main() {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{Jar: jar, Transport: tr}
	resp, err := client.Get("https://poly.webuntis.com/WebUntis/?school=RBZ-Kiel")
	if err != nil {
		log.Fatal(err)
	}
	//	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	w, w2, w3 := getDates()
	//	getics(jar, "2015-05-21")
	var wo = stripend(getics(jar, w))
	_ = "breakpoint"

	var wo2 = stripbeginning(stripend(getics(jar, w2)))
	var wo3 = stripbeginning(getics(jar, w3))
	_ = "breakpoint"

	var resul = wo + wo2 + wo3
	fmt.Print(resul)

}

func stripall(input string) string {
	return stripbeginning(stripend(input))
}

func stripbeginning(input string) string {
	_ = "breakpoint"
	re := regexp.MustCompile("BEGIN:VCALENDAR")
	a := re.ReplaceAllString(input, "")
	_ = "breakpoint"
	re2 := regexp.MustCompile("PRODID:-//Ben Fortuna//iCal4j 1.0//EN")
	a = re2.ReplaceAllString(a, "")
	re3 := regexp.MustCompile("VERSION:2.0")
	a = re3.ReplaceAllString(a, "")
	re4 := regexp.MustCompile("CALSCALE:GREGORIAN")
	a = re4.ReplaceAllString(a, "")
	return a
}

/*
BEGIN:VCALENDAR
PRODID:-//Ben Fortuna//iCal4j 1.0//EN
VERSION:2.0
CALSCALE:GREGORIAN

[...]

END:VCALENDAR
*/

func stripend(input string) string {
	re := regexp.MustCompile("END:VCALENDAR")
	a := re.ReplaceAllString(input, "")
	return a
}

func getDates() (string, string, string) {
	now := time.Now()
	const layout = "2006-01-2"

	t := now.Format(layout)
	t2 := now.AddDate(0, 0, 7).Format(layout)
	t3 := now.AddDate(0, 0, 14).Format(layout)
	return t, t2, t3
	//7 Tage == 168 Stunden
	//14 Tage == 336 Stunden
}
func getics(sharedjar http.CookieJar, datum string) string {
	client := http.Client{Jar: sharedjar}
	url := "https://poly.webuntis.com/WebUntis/Ical.do?elemType=1&elemId=569&rpt_sd=" + datum
	resp, err := client.Get(url)
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
func sum(a, b int) int {
	return a + b
}
