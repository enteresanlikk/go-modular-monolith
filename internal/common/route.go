package common

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type Route struct {
	Method Method
	Path   string
}

func (r *Route) Create(method Method, path string) string {
	return string(method) + " " + r.Path + path
}
