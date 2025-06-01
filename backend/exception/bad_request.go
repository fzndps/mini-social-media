package exception

type BadRequest struct {
	Error string
}

func NewBadRequestError(error string) BadRequest {
	return BadRequest{Error: error}
}
