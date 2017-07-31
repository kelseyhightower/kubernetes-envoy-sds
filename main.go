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
	"log"
	"net/http"
)

const (
	endpointsPath = "/api/v1/namespaces/%s/endpoints/%s"
)

var (
	clusterDomain string
	httpAddr      string
)

func main() {
	flag.StringVar(&clusterDomain, "cluster-domain", "svc.cluster.local", "The cluster domain")
	flag.StringVar(&httpAddr, "http", "127.0.0.1:8080", "The HTTP listen address.")
	flag.Parse()

	log.Println("Starting the Kubernetes Envoy SDS Service...")
	log.Printf("Listening on %s...", httpAddr)

	http.Handle("/v1/registration/", registrationServer(clusterDomain))
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
