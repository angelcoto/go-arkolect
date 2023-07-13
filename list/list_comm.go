package list

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/angelcoto/go-artamiz/hash"
)

type file struct {
	bytes    int64
	modtime  time.Time
	hash     []byte
	pathname string
}

func (f file) getFileProp(fi fs.FileInfo, dir, algo string) (file, error) {
	f.bytes = fi.Size()
	f.modtime = fi.ModTime()
	f.pathname = filepath.Join(dir, fi.Name())
	var err error
	f.hash, err = hash.SumArchivo(f.pathname, algo)
	if err != nil {
		return file{}, err

	}
	return f, nil
}

// printLine prints the ouput line
func printLine(lineObj file) {
	fmt.Printf("%x\t%d\t%s\t%s\n",
		lineObj.hash,
		lineObj.bytes,
		lineObj.modtime.Format(time.RFC3339),
		lineObj.pathname)
}

func PrintError(err error) {
	fmt.Printf("* Error: %s\n", err)
}
