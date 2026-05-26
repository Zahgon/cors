// adapted from github.com/jub0bs/cors
package internal

// A SortedSet represents a mathematical set of strings sorted in
// lexicographical order.
// Each element has a unique position ranging from 0 (inclusive)
// to the set's cardinality (exclusive).
// The zero value represents an empty set.
type SortedSet struct {
	m      map[string]int
	maxLen int
}

// NewSortedSet returns a SortedSet that contains all of elems,
// but no other elements.
func NewSortedSet(elems ...string) SortedSet { _ = "STUB: not implemented"; return *new(SortedSet) }

// Size returns the cardinality of set.
func (set SortedSet) Size() int {
	_ = "STUB: not implemented"

	// String sorts joins the elements of set (in lexicographical order)
	// with a comma and returns the resulting string.
	return 0
}

func (set SortedSet) String() string { _ = "STUB: not implemented"; return "" }

// safe indexing, by construction of SortedSet

// Accepts reports whether values is a sequence of list-based field values
// whose elements are
//   - all members of set,
//   - sorted in lexicographical order,
//   - unique.
func (set SortedSet) Accepts(values []string) bool {
	_ = "STUB: not implemented"
	// effectively constant
	return false
}

// +1 for comma

// As a defense against maliciously long names in s,
// we process only a small number of s's leading bytes per iteration.

// RFC 9110 requires recipients to tolerate
// "a reasonable number of empty list elements"; see
// https://httpwg.org/specs/rfc9110.html#abnf.extension.recipient.

// We have now exhausted the names in s.

// The names in s are expected to be sorted in lexicographical order
// and to each appear at most once.
// Therefore, the positions (in set) of the names that
// appear in s should form a strictly increasing sequence.
// If that's not actually the case, bail out.

// We have now exhausted the names in s.

const (
	maxOWSBytes      = 1  // number of leading/trailing OWS bytes tolerated
	maxEmptyElements = 16 // number of empty list elements tolerated
)

func cutAtComma(s string, n int) (before, after string, found bool) {
	_ = "STUB: not implemented"
	// Note: this implementation draws inspiration from strings.Cut's.
	return "", "", false
}

// deal with this first to save one bounds check

// TrimOWS trims up to n bytes of [optional whitespace (OWS)]
// from the start of and/or the end of s.
// If no more than n bytes of OWS are found at the start of s
// and no more than n bytes of OWS are found at the end of s,
// it returns the trimmed result and true.
// Otherwise, it returns the original string and false.
//
// [optional whitespace (OWS)]: https://httpwg.org/specs/rfc9110.html#whitespace
func trimOWS(s string, n int) (trimmed string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func trimLeftOWS(s string, n int) (string, bool) { _ = "STUB: not implemented"; return "", false }

func trimRightOWS(s string, n int) (string, bool) { _ = "STUB: not implemented"; return "", false }
