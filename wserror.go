package nfe

import "fmt"

type WSError struct {
	Url           string
	StatusCode    int
	StatusMessage string
	Body          string
}

func (e *WSError) Error() string {
	return fmt.Sprintf("Falha na consulta à receita (url '%s'): %d - %s", e.Url, e.StatusCode, e.StatusMessage)
}
