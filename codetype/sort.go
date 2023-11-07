package codetype

import "strings"

type SortType string

const (
	SortTypeASC  SortType = "ASC"
	SortTypeDESC SortType = "DESC"
)

func (s *SortType) IsValid() bool {
	if s != nil {
		switch SortType(strings.ToUpper(string(*s))) {
		case SortTypeASC, SortTypeDESC:
			return true
		default:
			return false
		}
	}

	return false
}

func (s *SortType) Format() {
	if !s.IsValid() {
		*s = SortTypeASC
	}
}
