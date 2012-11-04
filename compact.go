package compact

var zeroString string

func Strings(in []string) []string {
	if in == nil {
		return in
	}

	lastI := len(in) - 1
	for i, s := range in {
		if s == zeroString { // remove it from in
			switch {
			case i == 0:
				in = in[1:lastI]
			case i == lastI:
				continue
			default:
				in = append(in[:i-1], in[i:]...)
			}
		}
	}

	return in
}
