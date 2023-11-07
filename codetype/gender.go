package codetype

type Gender int

const (
	GenderFemale Gender = 1
	GenderMale   Gender = 2
	GenderOther  Gender = 3
)

func (g Gender) IsValid() bool {
	switch g {
	case GenderFemale, GenderMale, GenderOther:
		return true
	default:
		return false
	}
}
