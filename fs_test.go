package fs

import . "github.com/pkg4go/assert"
import "testing"

func TestCopy(t *testing.T) {
	a := A{t}
	err := Copy("fs.go", "fs.go.tmp")
	// TODO: stat, remove *.tmp
	a.Equal(err, nil)
}
