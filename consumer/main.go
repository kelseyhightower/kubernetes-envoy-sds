// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	proxy string
)

func main() {
	flag.StringVar(&proxy, "proxy", "http://127.0.0.1:80", "Envoy proxy")
	flag.Parse()

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100

	for {
		resp, err := http.Get(proxy)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.StatusCode != 200 {
			log.Println("Non 200 response code")
			continue
		}
		resp.Body.Close()

		time.Sleep(10 * time.Millisecond)
		fmt.Println(resp)
	}
}
