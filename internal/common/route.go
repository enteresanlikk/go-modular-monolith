package common

type Route struct {
	Prefix string
}

func (r *Route) Create(pattern string) string {
	return r.Prefix + pattern
}
