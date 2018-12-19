package readfq

import (
	"github.com/brentp/xopen"
	"github.com/sstadick/readfq"
)

// ExampleIter shows how to use the iter function
func ExampleIter() {

	fp, r := xopen.Ropen("-")
	defer fp.Close()
	n, sLen, qLen := 0, int64(0), int64(0)
	var fqr readfq.FqReader
	fqr.Reader = r.Reader // Any reader you want ..
	for r, done := fqr.iter(); !done; r, done = fqr.iter() {
		n++
		sLen += int64(len(r.seq))
		qLen += int64(len(r.qual))
	}
}
