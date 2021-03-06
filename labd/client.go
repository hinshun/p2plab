// Copyright 2019 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package labd

import (
	"net/http"

	"github.com/Netflix/p2plab"
)

type client struct {
	httpClient *http.Client
	base       string
}

func NewClient(addr string) (p2plab.LabdAPI, error) {
	return &client{
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy:             http.ProxyFromEnvironment,
				DisableKeepAlives: true,
			},
		},
		base: addr,
	}, nil
}

func (c *client) Cluster() p2plab.ClusterAPI {
	return &clusterAPI{c}
}

func (c *client) Node() p2plab.NodeAPI {
	return &nodeAPI{c}
}

func (c *client) Scenario() p2plab.ScenarioAPI {
	return &scenarioAPI{c}
}

func (c *client) Benchmark() p2plab.BenchmarkAPI {
	return &benchmarkAPI{c}
}
