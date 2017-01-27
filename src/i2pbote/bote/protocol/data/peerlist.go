package data

import (
	"i2pbote/bote/common"
	"i2pbote/util"
)

func CreatePeerList(version byte, dests []common.Destination) (buff []byte) {
	l := len(dests)
	buff = make([]byte, 2+(l*common.DestLen))
	util.PutUInt16_i(l, buff[2:])
	for idx, d := range dests {
		copy(buff[4+(idx*common.DestLen):], d[:])
	}
	return New(PeerList, version, buff)
}
