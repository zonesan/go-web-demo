package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

var oauthUrl string = `https://aws-lab.datafoundry.cn/oauth/authorize?response_type=code&redirect_uri=http%3A%2F%2Fweb-demo.chaizs.app.dataos.io%2Fcallback&scope=user%3Ainfo%20user%3Acheck-access%20user%3Alist-projects&client_id=oauth-test`

func httpsrv() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)
		r.Header["CLIENT-INFO"] = []string{r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto}
		r.Header["hostname"] = []string{os.Getenv("HOSTNAME")}
		fmt.Printf("%#v", r.Header)
		resp, _ := json.MarshalIndent(r.Header, "", "  ")
		fmt.Fprintf(w, string(resp))
	})

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)
		http.Redirect(w, r, oauthUrl, http.StatusFound)
	})

	//https://aws-lab.datafoundry.cn/oauth/authorize?response_type=code&redirect_uri=http%3A%2F%2Fweb-demo.chaizs.app.dataos.io%2Fcallback&scope=user%3Ainfo%20user%3Acheck-access%20user%3Alist-projects&client_id=oauth-test

	log.Println("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	go httpsrv()

	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val

		}
		return items

	}
	cnt := 0
	for {
		environment := getenvironment(os.Environ(), func(item string) (key, val string) {
			splits := strings.Split(item, "=")
			key = splits[0]
			val = splits[1]
			return

		})
		keys := []string{}
		for k := range environment {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key, "=", environment[key])
		}
		cnt += 1
		fmt.Println("hello#", cnt)
		/*
			for k, v := range environment {
				fmt.Println(k, "=", v)
			}
		*/
		time.Sleep(time.Second * 3000)
	}
}
