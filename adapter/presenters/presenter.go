package presenters

type Presenter interface {
	ToErrorResponse(err error, code int) (mapResponse map[string]interface{})
}
