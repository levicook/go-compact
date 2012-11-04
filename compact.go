package compact

var zeroString string

func Strings(in []string) []string {
	if in == nil {
		return in
	}

	out := []string{}
	for _, s := range in {
		if s == zeroString {
			continue
		}
		out = append(out, s)
	}

	return out
}
