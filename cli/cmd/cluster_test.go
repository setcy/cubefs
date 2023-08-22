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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cubefs/cubefs/proto"
	"github.com/cubefs/cubefs/util/fake"
)

func TestClusterCmd(t *testing.T) {
	r := newCliTestRunner()
	err := r.testRun("cluster", "help")
	assert.NoError(t, err)
}

func TestClusterInfoCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{},
			expectErr: false,
		},
	}

	successV1 := &proto.AclRsp{
		OK: true,
		List: []*proto.AclIpInfo{
			{
				Ip:    "192.168.0.1",
				CTime: 1689091200,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/aclOp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.JsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient).setCommand("cluster", "info")
	r.runTestCases(t, testCases)
}

func TestClusterStatCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{},
			expectErr: false,
		},
	}

	successV1 := &proto.AclRsp{
		OK: true,
		List: []*proto.AclIpInfo{
			{
				Ip:    "192.168.0.1",
				CTime: 1689091200,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/aclOp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.JsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient).setCommand("cluster", "stat")
	r.runTestCases(t, testCases)
}

func TestClusterFreezeCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments enable",
			args:      []string{"true"},
			expectErr: false,
		},
		{
			name:      "Valid arguments disable",
			args:      []string{"false"},
			expectErr: false,
		},
		{
			name:      "Invalid arguments",
			args:      []string{"invalid"},
			expectErr: true,
		},
	}

	successV1 := &proto.AclRsp{
		OK: true,
		List: []*proto.AclIpInfo{
			{
				Ip:    "192.168.0.1",
				CTime: 1689091200,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/aclOp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.JsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient).setCommand("cluster", "freeze")
	r.runTestCases(t, testCases)
}

func TestClusterSetThresholdCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"0.5"},
			expectErr: false,
		},
		{
			name:      "missing arguments",
			args:      []string{},
			expectErr: true,
		},
		{
			name:      "Invalid arguments",
			args:      []string{"invalid"},
			expectErr: true,
		},
		{
			name:      "too big threshold",
			args:      []string{"1.1"},
			expectErr: true,
		},
	}

	successV1 := &proto.AclRsp{
		OK: true,
		List: []*proto.AclIpInfo{
			{
				Ip:    "192.168.0.1",
				CTime: 1689091200,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/aclOp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.JsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient).setCommand("cluster", "threshold")
	r.runTestCases(t, testCases)
}

func TestClusterSetParasCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{},
			expectErr: false,
		},
	}

	successV1 := &proto.AclRsp{
		OK: true,
		List: []*proto.AclIpInfo{
			{
				Ip:    "192.168.0.1",
				CTime: 1689091200,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/aclOp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.JsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient).setCommand("cluster", "set")
	r.runTestCases(t, testCases)
}

func TestClusterDisableMpDecommissionCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"true"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"forbid-mp-decommission"},
			expectErr: true,
		},
	}

	successV1 := &proto.AclRsp{
		OK: true,
		List: []*proto.AclIpInfo{
			{
				Ip:    "192.168.0.1",
				CTime: 1689091200,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/aclOp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.JsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient).setCommand("cluster", "forbid-mp-decommission")
	r.runTestCases(t, testCases)
}
