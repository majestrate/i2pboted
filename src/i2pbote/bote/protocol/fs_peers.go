package protocol

import (
	"bufio"
	"fmt"
	"i2pbote/bote/common"
	"i2pbote/i2p/base64"
	"i2pbote/util"
	"os"
	"path/filepath"
	"strings"
)

// filesystem based peer holder
type FsPeerHolder struct {
	File string
	d    []common.Destination
}

func (f *FsPeerHolder) GetPeers(limit int) (p []common.Destination) {
	if len(f.d) <= limit {
		p = append(p, f.d...)
	} else {
		// TODO: randomize
		p = append(p, f.d[:limit]...)
	}
	return
}

func (f *FsPeerHolder) AddPeers(d []common.Destination) {
	f.d = append(f.d, d...)
}

func (f *FsPeerHolder) loadPeerLine(line string) {
	idx := strings.Index(line, "#")
	if idx >= 0 {
		line = line[:idx]
	}
	idx = strings.Index(line, "\t")
	if idx > 0 {
		d, err := base64.Encoding.DecodeString(line[:idx])
		if err == nil {
			if len(d) == common.DestLen {
				var dest common.Destination
				copy(dest[:], d[:])
				f.d = append(f.d, dest)
			}
		}
	}
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
	err = util.EnsureFile(f.File, 0)
	if err == nil {
		var fd *os.File
		fd, err = os.OpenFile(f.File, os.O_WRONLY, 0600)
		if err == nil {
			defer fd.Close()
			_, err = fmt.Fprintf(fd, "# auto generated file\n")
			if err == nil {
				for _, d := range f.d {
					fmt.Fprintf(fd, "%s\ttrue\n", d.String())
				}
			}
		}
	}
	return
}

func NewFsPeerHolder(rootdir string) *FsPeerHolder {
	return &FsPeerHolder{
		File: filepath.Join(rootdir, "relay_peers.txt"),
	}
}
