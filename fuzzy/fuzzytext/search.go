package fuzzytext

const (
	SequentialBonus  = 15 // bonus for adjacent matches
	SeparatorBonus   = 30 // bonus if match occurs after a separator
	CamelBonus       = 30 // bonus if match is uppercase and prev is lower
	FirstLetterBonus = 15 // bonus if the first letter is matched

	MaxLeadingLetterPenalty = -15 // penalty applied for every letter in str before the first match
	LeadingLetterPenalty    = -5  // maximum penalty for leading letters
	UnmatchedLetterPenalty  = -1
)
