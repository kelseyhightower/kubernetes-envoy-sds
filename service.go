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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Service struct {
	// A list of hosts that make up the service.
	Hosts []Host `json:"hosts"`
}

type Host struct {
	// The IP address of the upstream host.
	IPAddress string `json:"ip_address"`

	// The port of the upstream host.
	Port int32 `json:"port"`

	Tags *Tags `json:"tags,omitempty"`
}

type Tags struct {
	// The optional zone of the upstream host. Envoy uses the zone
	// for various statistics and load balancing tasks.
	AZ string `json:"az,omitempty"`

	// The optional canary status of the upstream host. Envoy uses
	// the canary status for various statistics and load balancing
	// tasks.
	Canary bool `json:"canary,omitempty"`

	// The optional load balancing weight of the upstream host, in
	// the range 1 - 100. Envoy uses the load balancing weight in
	// some of the built in load balancers.
	LoadBalancingWeight int32 `json:"load_balancing_weight,omitempty"`
}

func getService(namespace, name string) (*Service, error) {
	path := fmt.Sprintf(endpointsPath, namespace, name)

	r := &http.Request{
		Header: make(http.Header),
		Method: http.MethodGet,
		URL: &url.URL{
			Host:   "127.0.0.1:8001",
			Path:   path,
			Scheme: "http",
		},
	}

	r.Header.Set("Accept", "application/json, */*")

	ctx := context.Background()
	resp, err := http.DefaultClient.Do(r.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	hosts := make([]Host, 0)

	// If the service does not exist in Kubernetes return an empty
	// list of hosts.
	if resp.StatusCode == 404 {
		return &Service{Hosts: hosts}, nil
	}

	var eps endpoints
	err = json.NewDecoder(resp.Body).Decode(&eps)
	if err != nil {
		return nil, err
	}

	// Envoy backends only support a single port. The backend will use the first
	// port found on the Kubernetes endpoint.
	// Open questions around named ports and services with multiple ports.
	subset := eps.Subsets[0]
	for _, address := range subset.Addresses {
		hosts = append(hosts, Host{IPAddress: address.IP, Port: subset.Ports[0].Port})
	}

	return &Service{Hosts: hosts}, nil
}

type object struct {
	Object endpoints `json:"object"`
	Type   string    `json:"type"`
}

type endpoints struct {
	Kind       string   `json:"kind"`
	ApiVersion string   `json:"apiVersion"`
	Metadata   metadata `json:"metadata"`
	Subsets    []subset `json:"subsets"`
	Message    string   `json:"message"`
}

type metadata struct {
	Name string `json:"name"`
}

type subset struct {
	Addresses []address `json:"addresses"`
	Ports     []port    `json:"ports"`
}

type address struct {
	IP string `json:"ip"`
}

type port struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

type status struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
