package mapper

import (
	"fmt"
)

type UnsupportedMapperErr struct {
	mapperId uint8
}

func (err *UnsupportedMapperErr) Error() string {
	return fmt.Sprintf("Unsuppoerted mapper: %d", err.mapperId)
}
