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

package p2plab

import (
	"context"
	"io"

	"github.com/Netflix/p2plab/metadata"
)

// BenchmarkAPI defines API for benchmark operations.
type BenchmarkAPI interface {
	// Create starts benchmarking a scenario on a cluster.
	Start(ctx context.Context, cluster, scenario string, opts ...StartBenchmarkOption) (Benchmark, error)

	// Get returns a benchmark.
	Get(ctx context.Context, id string) (Benchmark, error)

	// List returns available benchmarks.
	List(ctx context.Context) ([]Benchmark, error)
}

// Benchmark is an execution of a scenario on a cluster.
type Benchmark interface {
	Metadata() metadata.Benchmark

	// Cancel cancels a running benchmark.
	Cancel(ctx context.Context) error

	// Report returns statistics on how the P2P application behaved during the
	// benchmark.
	Report(ctx context.Context) (Report, error)

	// Logs returns a streaming log of the benchmark operation.
	Logs(ctx context.Context, opt ...LogsOption) (io.ReadCloser, error)
}

type StartBenchmarkOption func(*StartBenchmarkSettings) error

type StartBenchmarkSettings struct {
	NoReset bool
}

func WithBenchmarkNoReset() StartBenchmarkOption {
	return func(s *StartBenchmarkSettings) error {
		s.NoReset = true
		return nil
	}
}

// Report is a benchmark summary on how the P2P application behaved during the
// benchmark.
type Report interface {
}

// LogsOption is an option to modify logging settings.
type LogsOption func(LogsSettings) error

// LogsSettings specify logging settings.
type LogsSettings struct {
	// Follow specify that the log should be followed.
	Follow bool
}
