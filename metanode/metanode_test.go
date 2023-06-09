package metanode

import (
	"github.com/cubefs/cubefs/util/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	data := `
{
  "role": "metanode",
  "listen": "17210",
  "prof": "17220",
  "logLevel": "debug",
  "metadataDir": "/cfs/data/meta",
  "logDir": "/cfs/log",
  "raftDir": "/cfs/data/raft",
  "raftHeartbeatPort": "17230",
  "raftReplicaPort": "17240",
  "consulAddr": "http://192.168.0.101:8500",
  "exporterPort": 9500,
  "warnLogDir": "/cfs/log",
  "totalMem": "536870912",
  "localIP": "192.168.0.21",
  "bindIp": true,
  "zoneName": "zone1",
  "deleteBatchCount": 1000,
  "tickInterval": 10.1,
  "raftRecvBufSize": 1024,
  "masterAddr": [
    "192.168.0.11:17010",
    "192.168.0.12:17010",
    "192.168.0.13:17010"
  ],
  "smuxPortShift": 300,
  "smuxMaxConn": 100,
  "smuxStreamPerConn": 2,
  "smuxMaxBuffer": 32768
}`

	cfg := config.LoadConfigString(data)

	mn := &MetaNode{}
	err := mn.parseConfig(cfg)

	assert.NoError(t, err)
	assert.Equal(t, "17210", mn.listen)
	assert.Equal(t, true, mn.bindIp)
	assert.Equal(t, "17210", serverPort)
	assert.Equal(t, "/cfs/data/meta", mn.metadataDir)
	assert.Equal(t, "/cfs/data/raft", mn.raftDir)
	assert.Equal(t, "192.168.0.21", mn.localAddr)
	assert.Equal(t, "17230", mn.raftHeartbeatPort)
	assert.Equal(t, "17240", mn.raftReplicatePort)
	assert.Equal(t, "zone1", mn.zoneName)
	assert.Equal(t, 10, mn.tickInterval)
	assert.Equal(t, 1024, mn.raftRecvBufSize)
	assert.Equal(t, uint64(1000), DeleteBatchCount())
	assert.NotEqual(t, 0, configTotalMem)
	assert.Equal(t, 300, smuxPortShift)
	assert.Equal(t, 2, smuxPoolCfg.StreamsPerConn)
	assert.Equal(t, 100, smuxPoolCfg.ConnsPerAddr)
	assert.Equal(t, 32768, smuxPoolCfg.MaxReceiveBuffer)
	assert.Equal(t, "192.168.0.11:17010", masterClient.Nodes()[0])
	assert.Equal(t, "192.168.0.12:17010", masterClient.Nodes()[1])
	assert.Equal(t, "192.168.0.13:17010", masterClient.Nodes()[2])
}
