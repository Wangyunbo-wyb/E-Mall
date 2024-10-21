package filter

type Result struct {
	Word    string // the sensitive word
	Matched string // the string that matched
	Start   int    // the start position matched in the original string
	End     int    // the end position matched in the original string
}
