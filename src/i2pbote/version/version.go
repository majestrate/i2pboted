package version

import "fmt"

const Major = 0
const Minor = 0
const Patch = 0

var Git string

func Version() string {
	return fmt.Sprintf("i2pboted-%d.%d.%d%s", Major, Minor, Patch, Git)
}
