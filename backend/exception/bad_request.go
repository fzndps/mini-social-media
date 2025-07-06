package exception

type ConflictRequest struct {
	Error string
}

func NewConflictRequestError(error string) ConflictRequest {
	return ConflictRequest{Error: error}
}
