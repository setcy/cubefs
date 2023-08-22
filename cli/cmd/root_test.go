// Copyright 2023 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cmd

import (
	"fmt"
	"github.com/cubefs/cubefs/sdk/master"
	"github.com/cubefs/cubefs/util/fake"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	cliTestAddr     = "127.0.0.1:16517"
	ContentTypeJSON = "application/json"
)

type testErrorRecorder struct {
	errorMessages []string
}

func (t *testErrorRecorder) Write(p []byte) (n int, err error) {
	t.errorMessages = append(t.errorMessages, string(p))
	return len(p), nil
}

func (t *testErrorRecorder) RecordedError() error {
	if len(t.errorMessages) > 0 {
		return fmt.Errorf("%v", t.errorMessages[0])
	}
	return nil
}

func defaultHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", ContentTypeJSON)
	return header
}

func setupMockCommands() *cobra.Command {
	var mc = master.NewMasterClient([]string{cliTestAddr}, false)
	mc.SetTimeout(1)
	cfsRootCmd := NewRootCmd(mc)
	cfsRootCmd.CFSCmd.AddCommand(GenClusterCfgCmd)
	return cfsRootCmd.CFSCmd
}

func setupTestErrorRecorder() *testErrorRecorder {
	recorder := &testErrorRecorder{}
	erroutHandler = func(format string, args ...interface{}) {
		_, _ = fmt.Fprintf(recorder, format, args...)
	}
	return recorder
}

func testRun(args ...string) (err error) {
	recorder := setupTestErrorRecorder()
	cfsCli := setupMockCommands()
	cfsCli.SetArgs(args)
	err = cfsCli.Execute()
	if err != nil {
		return err
	}
	return recorder.RecordedError()
}

type TestCase struct {
	name      string
	args      []string
	expectErr bool
}

func runTestCases(t *testing.T, testCases []*TestCase, fakeClient *http.Client, command ...string) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args := append(command, tc.args...)
			err := testRun(args...)
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRootCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "no command",
			args:      []string{},
			expectErr: false,
		},
		{
			name:      "help command",
			args:      []string{"help"},
			expectErr: false,
		},
		{
			name:      "version command",
			args:      []string{"--version"},
			expectErr: false,
		},
		{
			name:      "wrong command",
			args:      []string{"cluste"},
			expectErr: true,
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch p, m := req.URL.Path, req.Method; {

		case m == http.MethodGet && p == "/apis/certificates.k8s.io/v1/certificatesigningrequests/missing":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader()}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	runTestCases(t, testCases, fakeClient)
}
