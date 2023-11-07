package codetype

const PageSizeDefault = 20

type Paginator struct {
	Page  int `json:"page,omitempty" query:"page"`
	Limit int `json:"limit,omitempty" query:"limit"`
}

func (p *Paginator) Format() {
	if p.Limit == 0 {
		p.Limit = PageSizeDefault
	}

	if p.Limit < 0 {
		p.Limit = -1
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}
