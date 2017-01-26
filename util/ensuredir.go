package util

import (
	"github.com/majestrate/i2pboted/log"
	"os"
)

// ensure a directory is made
// returns error if it can't be made
func EnsureDir(fpath string) (err error) {
	_, err = os.Stat(fpath)
	if os.IsNotExist(err) {
		log.Debugf("create dir %s", fpath)
		err = os.Mkdir(fpath, 0700)
	}
	return
}
