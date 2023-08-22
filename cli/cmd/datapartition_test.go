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

func TestDataPartitionCmd(t *testing.T) {
	r := newCliTestRunner()
	err := r.testRun("datapartition", "help")
	assert.NoError(t, err)
}

func TestDataPartitionGetCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datapartition", "info", "1"},
			expectErr: false,
		},
		{
			name:      "Missing arguments",
			args:      []string{"datapartition", "info"},
			expectErr: true,
		},
		{
			name:      "Invalid arguments",
			args:      []string{"datapartition", "info", "t"},
			expectErr: true,
		},
	}

	successV1 := &proto.DataPartitionInfo{
		PartitionID:              0,
		PartitionTTL:             0,
		PartitionType:            0,
		LastLoadedTime:           0,
		ReplicaNum:               0,
		Status:                   0,
		Recover:                  false,
		Replicas:                 []*proto.DataReplica{},
		Hosts:                    []string{},
		Peers:                    []proto.Peer{},
		Zones:                    []string{},
		MissingNodes:             map[string]int64{},
		VolName:                  "",
		VolID:                    0,
		OfflinePeerID:            0,
		FileInCoreMap:            map[string]*proto.FileInCore{},
		IsRecover:                false,
		FilesWithMissingReplica:  map[string]int64{},
		SingleDecommissionStatus: 0,
		SingleDecommissionAddr:   "",
		RdOnly:                   false,
		IsDiscard:                false,
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataPartition/get":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestListCorruptDataPartitionCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datapartition", "check"},
			expectErr: false,
		},
	}

	diagnosisV1 := &proto.DataPartitionDiagnosis{
		InactiveDataNodes:           []string{"172.16.1.101:17310", "172.16.1.102:17310"},
		CorruptDataPartitionIDs:     []uint64{1, 2},
		LackReplicaDataPartitionIDs: []uint64{1, 2},
		BadDataPartitionIDs: []proto.BadPartitionView{
			{
				Path:         "/test1",
				PartitionIDs: []uint64{1, 2},
			},
			{
				Path:         "/test2",
				PartitionIDs: []uint64{3, 4},
			},
		},
		BadReplicaDataPartitionIDs: []uint64{1, 2},
		RepFileCountDifferDpIDs:    []uint64{1, 2},
		RepUsedSizeDifferDpIDs:     []uint64{1, 2},
		ExcessReplicaDpIDs:         []uint64{1, 2},
	}

	dataNodeV1 := &proto.DataNodeInfo{
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

	dataPartitionV1 := &proto.DataPartitionInfo{
		PartitionID:              0,
		PartitionTTL:             0,
		PartitionType:            0,
		LastLoadedTime:           0,
		ReplicaNum:               0,
		Status:                   0,
		Recover:                  false,
		Replicas:                 []*proto.DataReplica{},
		Hosts:                    []string{},
		Peers:                    []proto.Peer{},
		Zones:                    []string{},
		MissingNodes:             map[string]int64{},
		VolName:                  "",
		VolID:                    0,
		OfflinePeerID:            0,
		FileInCoreMap:            map[string]*proto.FileInCore{},
		IsRecover:                false,
		FilesWithMissingReplica:  map[string]int64{},
		SingleDecommissionStatus: 0,
		SingleDecommissionAddr:   "",
		RdOnly:                   false,
		IsDiscard:                false,
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataPartition/diagnose":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(diagnosisV1)}, nil

		case m == http.MethodGet && p == "/dataNode/get":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(dataNodeV1)}, nil

		case m == http.MethodGet && p == "/dataPartition/get":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(dataPartitionV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataPartitionDecommissionCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datapartition", "decommission", "172.16.1.101:17310", "1"},
			expectErr: false,
		},
		{
			name:      "Missing 1 arguments",
			args:      []string{"datapartition", "decommission", "172.16.1.101:17310"},
			expectErr: true,
		},
		{
			name:      "Missing 2 arguments",
			args:      []string{"datapartition", "decommission"},
			expectErr: true,
		},
		{
			name:      "Invalid arguments",
			args:      []string{"datapartition", "decommission", "172.16.1.101:17310", "t"},
			expectErr: true,
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataPartition/decommission":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(nil)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataPartitionReplicateCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datapartition", "add-replica", "172.16.1.101:17310", "1"},
			expectErr: false,
		},
		{
			name:      "Missing 1 arguments",
			args:      []string{"datapartition", "add-replica", "172.16.1.101:17310"},
			expectErr: true,
		},
		{
			name:      "Missing 2 arguments",
			args:      []string{"datapartition", "add-replica"},
			expectErr: true,
		},
		{
			name:      "Invalid arguments",
			args:      []string{"datapartition", "add-replica", "172.16.1.101:17310", "t"},
			expectErr: true,
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataReplica/add":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(nil)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataPartitionDeleteReplicaCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datapartition", "del-replica", "172.16.1.101:17310", "1"},
			expectErr: false,
		},
		{
			name:      "Missing 1 arguments",
			args:      []string{"datapartition", "del-replica", "172.16.1.101:17310"},
			expectErr: true,
		},
		{
			name:      "Missing 2 arguments",
			args:      []string{"datapartition", "del-replica"},
			expectErr: true,
		},
		{
			name:      "Invalid arguments",
			args:      []string{"datapartition", "del-replica", "172.16.1.101:17310", "t"},
			expectErr: true,
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/dataReplica/delete":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(nil)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})

	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}

func TestDataPartitionGetDiscardCmd(t *testing.T) {
	testCases := []*TestCase{
		{
			name:      "Valid arguments",
			args:      []string{"datapartition", "get-discard"},
			expectErr: false,
		},
	}

	successV1 := &proto.DiscardDataPartitionInfos{
		DiscardDps: []proto.DataPartitionInfo{
			{
				PartitionID:              1,
				PartitionTTL:             0,
				PartitionType:            0,
				LastLoadedTime:           0,
				ReplicaNum:               0,
				Status:                   0,
				Recover:                  false,
				Replicas:                 []*proto.DataReplica{},
				Hosts:                    []string{},
				Peers:                    []proto.Peer{},
				Zones:                    []string{},
				MissingNodes:             map[string]int64{},
				VolName:                  "",
				VolID:                    0,
				OfflinePeerID:            0,
				FileInCoreMap:            map[string]*proto.FileInCore{},
				IsRecover:                false,
				FilesWithMissingReplica:  map[string]int64{},
				SingleDecommissionStatus: 0,
				SingleDecommissionAddr:   "",
				RdOnly:                   false,
				IsDiscard:                false,
			},
		},
	}

	fakeClient := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		switch m, p := req.Method, req.URL.Path; {

		case m == http.MethodGet && p == "/admin/getDiscardDp":
			return &http.Response{StatusCode: http.StatusOK, Header: defaultHeader(), Body: fake.SuccessJsonBody(successV1)}, nil

		default:
			t.Fatalf("unexpected request: %#v\n%#v", req.URL, req)
			return nil, nil
		}
	})
	r := newCliTestRunner().setHttpClient(fakeClient)
	r.runTestCases(t, testCases)
}
