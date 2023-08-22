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
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cubefs/cubefs/proto"
	"github.com/cubefs/cubefs/util/fake"
)

func TestDataNodeCmd(t *testing.T) {
	r := newCliTestRunner()
	err := r.testRun("datanode", "help")
	assert.NoError(t, err)
}

func TestDataNodeListCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datanode", "list"},
			expectErr: false,
		},
	}

	successV1 := &proto.ClusterView{
		DataNodes: []proto.NodeView{
			{
				Addr:       "172.16.1.101:17310",
				Status:     false,
				DomainAddr: "",
				ID:         2,
				IsWritable: false,
			},
			{
				Addr:       "172.16.1.102:17310",
				Status:     false,
				DomainAddr: "",
				ID:         3,
				IsWritable: false,
			},
			{
				Addr:       "172.16.1.103:17310",
				Status:     false,
				DomainAddr: "",
				ID:         4,
				IsWritable: false,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/getCluster":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataNodeInfoCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datanode", "info", "172.16.1.110:17310"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"datanode", "info"},
			expectErr: true,
		},
	}

	successV1 := &proto.DataNodeInfo{
		Total:                     0,
		Used:                      0,
		AvailableSpace:            0,
		ID:                        0,
		ZoneName:                  "",
		Addr:                      "",
		DomainAddr:                "",
		ReportTime:                time.Time{},
		IsActive:                  false,
		IsWriteAble:               false,
		UsageRatio:                0,
		SelectedTimes:             0,
		Carry:                     0,
		DataPartitionReports:      []*proto.PartitionReport{},
		DataPartitionCount:        0,
		NodeSetID:                 0,
		PersistenceDataPartitions: []uint64{},
		BadDisks:                  []string{},
		RdOnly:                    false,
		MaxDpCntLimit:             0,
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataNode/get":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataNodeDecommissionCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datanode", "decommission", "172.16.1.110:17310", "--count", "1"},
			expectErr: false,
		},
		{
			name:      "Missing node address",
			args:      []string{"datanode", "decommission"},
			expectErr: true,
		},
		{
			name:      "Invalid migrate dp count",
			args:      []string{"datanode", "decommission", "172.16.1.110:17310", "--count", "-1"},
			expectErr: true,
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataNode/decommission":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(nil)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataNodeMigrateCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datanode", "migrate", "172.16.1.110:17310", "172.16.1.110:17310"},
			expectErr: false,
		},
		{
			name:      "Missing 1 node address",
			args:      []string{"datanode", "migrate", "172.16.1.110:17310"},
			expectErr: true,
		},
		{
			name:      "Missing 2 node address",
			args:      []string{"datanode", "migrate"},
			expectErr: true,
		},
		{
			name:      "invalid migrate dp count",
			args:      []string{"datanode", "migrate", "172.16.1.101:17310", "172.16.1.102:17310", "--count", "-1"},
			expectErr: true,
		},
		{
			name:      "too much migrate dp count",
			args:      []string{"datanode", "migrate", "172.16.1.101:17310", "172.16.1.102:17310", "--count", "500"},
			expectErr: true,
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataNode/migrate":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(nil)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}
