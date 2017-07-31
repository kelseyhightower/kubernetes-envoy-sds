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
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type registrationHandler struct {
	clusterDomain string
}

func registrationServer(clusterDomain string) http.Handler {
	return &registrationHandler{clusterDomain}
}

func extractServiceNamespace(service, clusterDomain string) (string, string) {
	if !strings.HasSuffix(service, clusterDomain) {
		return "", ""
	}
	s := strings.TrimSuffix(service, clusterDomain)
	s = strings.TrimSuffix(s, ".")
	hs := strings.Split(s, ".")

	if len(hs) != 2 {
		return "", ""
	}

	return hs[0], hs[1]
}

func (h *registrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	s := serviceFromURL(r.URL.Path)
	if s == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	service, namespace := extractServiceNamespace(s, h.clusterDomain)
	if namespace == "" {
		log.Printf("Invalid service name: %s. Must be FQDN (service.namespace.%s)", service, h.clusterDomain)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sds, err := getService(namespace, service)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data, err := json.MarshalIndent(sds, "", "  ")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func serviceFromURL(path string) string {
	s := strings.Split(path, "/")
	if len(s) < 3 {
		return ""
	}
	return s[3]
}
