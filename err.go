package imgur_go

import "fmt"

type ServerError ErrorResponse

func (e *ServerError) Error() string {
	return fmt.Sprintf("cannot request Imgur, error message: %v, error code: %d\n", e.Data.Error, e.Status)
}
