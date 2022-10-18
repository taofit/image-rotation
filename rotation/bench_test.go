package rotation

import (
	"log"
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	var fileContents = fetchFileContents("../bitmap.pbm")
	pbm, err := fetchPBM(fileContents)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	for i := 0; i < b.N; i++ {
		rotate(pbm, 202)
	}
}