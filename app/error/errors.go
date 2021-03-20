package error

import (
	"database/sql"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

var (
	ErrNotFound = errors.New("Not found")
)

func ParseError(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound
	case strings.Contains(err.Error(), "Validate"):
		return http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
