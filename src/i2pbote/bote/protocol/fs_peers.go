package protocol

import (
	"bufio"
	"i2pbote/bote/common"
	"os"
	"path/filepath"
)

type FsPeerHolder struct {
	File string
	d    []common.Destination
}

func (f *FsPeerHolder) GetPeers(limit int) (p []common.Destination) {
	return
}

func (f *FsPeerHolder) AddPeers(d []common.Destination) {
	f.d = append(f.d, d...)
}

func (f *FsPeerHolder) loadPeerLine(line string) {

}

func (f *FsPeerHolder) Load() (err error) {
	var fd *os.File
	fd, err = os.Open(f.File)
	if err == nil {
		defer fd.Close()
		r := bufio.NewReader(fd)
		var line string
		for err == nil {
			line, err = r.ReadString(10)
			if len(line) > 0 {
				f.loadPeerLine(line)
			}
		}
	}
	return
}

func (f *FsPeerHolder) Store() (err error) {

	return
}

func NewFsPeerHolder(rootdir string) *FsPeerHolder {
	return &FsPeerHolder{
		File: filepath.Join(rootdir, "relay_peers.txt"),
	}
}
