// Copyright 2022 The etcd Authors
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

package e2e

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"go.etcd.io/etcd/tests/v3/framework/e2e"
)

func TestEtcdDurationFlags(t *testing.T) {
	e2e.SkipInShortMode(t)

	timeDur := 60 * time.Second
	timeDurStr := timeDur.String()
	inputFlags := []string{
		"backend-batch-interval",
		"discovery-dial-timeout",
		"discovery-request-timeout",
		"discovery-keepalive-time",
		"discovery-keepalive-timeout",
		"experimental-corrupt-check-time",
		"experimental-downgrade-check-time",
		"experimental-wait-cluster-ready-timeout",
		"grpc-keepalive-min-time",
		"grpc-keepalive-interval",
		"grpc-keepalive-timeout",
		"raft-read-timeout",
		"raft-write-timeout",
		"experimental-compact-hash-check-time",
		"experimental-compaction-sleep-interval",
		"experimental-watch-progress-notify-interval",
		"experimental-warning-apply-duration",
		"experimental-warning-unary-request-duration",
	}

	etcdCmd := []string{e2e.BinPath.Etcd, "--log-level", "debug"}
	outExpectedStrs := []string{}
	for _, inFlag := range inputFlags {
		flagWithHyphens := fmt.Sprintf("--%s", inFlag)
		etcdCmd = append(etcdCmd, flagWithHyphens, timeDurStr)
		outExpectedStrs = append(outExpectedStrs, fmt.Sprintf(`"%s":"%s"`, inFlag, timeDur))
	}

	fmt.Println(etcdCmd, outExpectedStrs)
	proc, err := e2e.SpawnCmd(etcdCmd, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = e2e.WaitReadyExpectProc(context.TODO(), proc, outExpectedStrs); err != nil {
		t.Fatal(err)
	}
	if err = proc.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestEtcdDurationFlags_IntegerInputValues(t *testing.T) {
	e2e.SkipInShortMode(t)

	timeDurIntFormat := 88
	timeDurIntFormatStr := strconv.Itoa(timeDurIntFormat) // This is the actual check: input should be an (duration) integer format string
	outTimeDur := time.Duration(timeDurIntFormat) * time.Second
	inputFlags := []string{
		"backend-batch-interval",
		"discovery-dial-timeout",
		"discovery-request-timeout",
		"discovery-keepalive-time",
		"discovery-keepalive-timeout",
		"experimental-corrupt-check-time",
		"experimental-downgrade-check-time",
		"experimental-wait-cluster-ready-timeout",
		"grpc-keepalive-min-time",
		"grpc-keepalive-interval",
		"grpc-keepalive-timeout",
		"raft-read-timeout",
		"raft-write-timeout",
		"experimental-compact-hash-check-time",
		"experimental-compaction-sleep-interval",
		"experimental-watch-progress-notify-interval",
		"experimental-warning-apply-duration",
		"experimental-warning-unary-request-duration",
	}

	etcdCmd := []string{e2e.BinPath.Etcd, "--log-level", "debug"}
	outExpectedStrs := []string{}
	for _, inFlag := range inputFlags {
		flagWithHyphens := fmt.Sprintf("--%s", inFlag)
		etcdCmd = append(etcdCmd, flagWithHyphens, timeDurIntFormatStr)
		outExpectedStrs = append(outExpectedStrs, fmt.Sprintf(`"%s":"%s"`, inFlag, outTimeDur))
	}

	proc, err := e2e.SpawnCmd(etcdCmd, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = e2e.WaitReadyExpectProc(context.TODO(), proc, outExpectedStrs); err != nil {
		t.Fatal(err)
	}
	if err = proc.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestEtcdGatewayDurationFlags(t *testing.T) {
	e2e.SkipInShortMode(t)

	timeDur := 60 * time.Second
	timeDurStr := timeDur.String()
	inputFlags := []string{
		"retry-delay",
	}

	etcdCmd := []string{e2e.BinPath.Etcd, "gateway", "start", "--log-level", "debug"}
	outExpectedStrs := []string{}
	for _, inFlag := range inputFlags {
		flagWithHyphens := fmt.Sprintf("--%s", inFlag)
		etcdCmd = append(etcdCmd, flagWithHyphens, timeDurStr)
		outExpectedStrs = append(outExpectedStrs, fmt.Sprintf(`"%s":"%s"`, inFlag, timeDur))
	}

	proc, err := e2e.SpawnCmd(etcdCmd, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = e2e.WaitReadyExpectProc(context.TODO(), proc, outExpectedStrs); err != nil {
		t.Fatal(err)
	}
	if err = proc.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestEtcdGatewayDurationFlags_IntegerInputValues(t *testing.T) {
	e2e.SkipInShortMode(t)

	timeDurIntFormat := 88
	timeDurIntFormatStr := strconv.Itoa(timeDurIntFormat) // This is the actual check: input should be an (duration) integer format string
	outTimeDur := time.Duration(timeDurIntFormat) * time.Second
	inputFlags := []string{
		"retry-delay",
	}

	etcdCmd := []string{e2e.BinPath.Etcd, "gateway", "start", "--log-level", "debug"}
	outExpectedStrs := []string{}
	for _, inFlag := range inputFlags {
		flagWithHyphens := fmt.Sprintf("--%s", inFlag)
		etcdCmd = append(etcdCmd, flagWithHyphens, timeDurIntFormatStr)
		outExpectedStrs = append(outExpectedStrs, fmt.Sprintf(`"%s":"%s"`, inFlag, outTimeDur))
	}

	proc, err := e2e.SpawnCmd(etcdCmd, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = e2e.WaitReadyExpectProc(context.TODO(), proc, outExpectedStrs); err != nil {
		t.Fatal(err)
	}
	if err = proc.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestEtcdGRPCProxyDurationFlags(t *testing.T) {
	e2e.SkipInShortMode(t)

	timeDur := 60 * time.Second
	timeDurStr := timeDur.String()
	inputFlags := []string{
		"endpoints-auto-sync-interval",
		"grpc-keepalive-min-time",
		"grpc-keepalive-interval",
		"grpc-keepalive-timeout",
	}

	etcdCmd := []string{e2e.BinPath.Etcd, "grpc-proxy", "start", "--debug"}
	outExpectedStrs := []string{}
	for _, inFlag := range inputFlags {
		flagWithHyphens := fmt.Sprintf("--%s", inFlag)
		etcdCmd = append(etcdCmd, flagWithHyphens, timeDurStr)
		outExpectedStrs = append(outExpectedStrs, fmt.Sprintf(`"%s":"%s"`, inFlag, timeDur))
	}

	proc, err := e2e.SpawnCmd(etcdCmd, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = e2e.WaitReadyExpectProc(context.TODO(), proc, outExpectedStrs); err != nil {
		t.Fatal(err)
	}
	if err = proc.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestEtcdGRPCProxyDurationFlags_IntegerInputValues(t *testing.T) {
	e2e.SkipInShortMode(t)

	timeDurIntFormat := 88
	timeDurIntFormatStr := strconv.Itoa(timeDurIntFormat) // This is the actual check: input should be an (duration) integer format string
	outTimeDur := time.Duration(timeDurIntFormat) * time.Second
	inputFlags := []string{
		"endpoints-auto-sync-interval",
		"grpc-keepalive-min-time",
		"grpc-keepalive-interval",
		"grpc-keepalive-timeout",
	}

	etcdCmd := []string{e2e.BinPath.Etcd, "grpc-proxy", "start", "--debug"}
	outExpectedStrs := []string{}
	for _, inFlag := range inputFlags {
		flagWithHyphens := fmt.Sprintf("--%s", inFlag)
		etcdCmd = append(etcdCmd, flagWithHyphens, timeDurIntFormatStr)
		outExpectedStrs = append(outExpectedStrs, fmt.Sprintf(`"%s":"%s"`, inFlag, outTimeDur))
	}

	proc, err := e2e.SpawnCmd(etcdCmd, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = e2e.WaitReadyExpectProc(context.TODO(), proc, outExpectedStrs); err != nil {
		t.Fatal(err)
	}
	if err = proc.Stop(); err != nil {
		t.Fatal(err)
	}
}
