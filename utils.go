package cors

type wildcard struct {
	prefix string
	suffix string
}

func (w wildcard) match(s string) bool { _ = "STUB: not implemented"; return false }

// convert converts a list of string using the passed converter function
func convert(s []string, f func(string) string) []string { _ = "STUB: not implemented"; return nil }
