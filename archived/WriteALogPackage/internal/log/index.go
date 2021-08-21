package log

import (
	"io"
	"os"
	"github.com/tysontate/gommap"
)

var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	entWidth =  offWidth + posWidth
)

type index struct {
	file *os.File
	
}