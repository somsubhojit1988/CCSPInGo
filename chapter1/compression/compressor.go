package compression

type CompressedGene struct {
	seq []byte
	N   int
}

// A : 00
// C : 01
// T : 10
// G : 11

func Compress(strSeq string) *CompressedGene {
	// byte -> 8 bits ->4 nucleotides

	n := func(n int) int {
		ret := int(n / 4)
		if n%4 > 0 {
			ret++
		}
		return ret
	}(len(strSeq))
	data := make([]byte, n)

	func(s string) {
		buffIndx := -1
		for i, c := range s {

			if i%4 == 0 {
				buffIndx++
			}

			switch string(c) {
			case "A":
				data[buffIndx] = (data[buffIndx] << 2) | 0b00
			case "C":
				data[buffIndx] = (data[buffIndx] << 2) | 0b01
			case "T":
				data[buffIndx] = (data[buffIndx] << 2) | 0b10
			case "G":
				data[buffIndx] = (data[buffIndx] << 2) | 0b11
			}
		}

	}(strSeq)

	return &CompressedGene{seq: data, N: len(strSeq)}
}

func (g *CompressedGene) Decompress() string {
	s, cntr := "", g.N

	decode := func(b byte, n int) string {
		s := ""
		for i := 0; i < n; i++ {
			switch b & 0b11 {
			case 0b00:
				s = "A" + s
			case 0b01:
				s = "C" + s
			case 0b10:
				s = "T" + s
			case 0b11:
				s = "G" + s
			}
			b = b >> 2
		}
		return s
	}

	for i := 0; i < (len(g.seq) - 1); i++ {
		b := g.seq[i]
		s += decode(b, 4)
		cntr -= 4
	}

	if cntr > 0 {
		b := g.seq[len(g.seq)-1]
		s += decode(b, cntr)
		cntr = 0
	}

	return s
}
