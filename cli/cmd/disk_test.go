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
	"github.com/cubefs/cubefs/util/fake"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskCmd(t *testing.T) {
	err := testRun("disk", "help")
	assert.NoError(t, err)
}

func TestListBadDiskCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"disk", "check"},
			expectErr: false,
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
