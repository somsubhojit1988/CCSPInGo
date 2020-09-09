package compression

import "testing"

func TestCompression(t *testing.T) {
	tt := []string{"ATCG", "ACGTG", "AG", "ACT", "ACTGAACTG", ""}
	for _, tc := range tt {
		cg := Compress(tc)
		res := cg.Decompress()
		if tc != res {
			t.Errorf("Expected: %s, decompressed: %s", tc, res)
			t.FailNow()
		}
	}
}
