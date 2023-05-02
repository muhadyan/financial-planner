package model

var (
	DefaultPage = 1
	DefaultLimit = 10
)

type BasedFilter struct {
	Limit   int
	Offset  int
	Page    int
	OrderBy string
	SortBy  string
}

func (c *BasedFilter) DefaultQuery() BasedFilter {
	if c.Limit <= 0 {
		c.Limit = 10
	}

	if c.Page > 0 {
		c.Offset = (c.Page - 1) * c.Limit
	}

	return *c
}
