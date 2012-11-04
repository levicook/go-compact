package compact

var zeroString string

func Strings(in []string) []string {
	if in == nil {
		return in
	}

	for i, s := range in {
		if s == zeroString {
			in = append(in[:i], in[i+1:]...)
		}
	}
	return in
}
