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

func TestUidCmd(t *testing.T) {
	r := newCliTestRunner()
	err := r.testRun("uid", "help")
	assert.NoError(t, err)
}

func TestUidAddCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"uid", "add", "testVol", "1", "20"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"uid", "add"},
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

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestUidListCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"uid", "list", "testVol"},
			expectErr: false,
		},
		{
			name:      "List all volumes",
			args:      []string{"uid", "list", "testVol", "all"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"uid", "list"},
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

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestUidDelCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"uid", "del", "testVol", "1"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"uid", "del"},
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

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}
func TestUidCheckCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"uid", "check", "testVol", "1"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"uid", "check"},
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

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}
