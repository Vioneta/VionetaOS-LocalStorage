package v2

import (
	"github.com/Vioneta/VionetaOS-LocalStorage/codegen"
)

type LocalStorage struct{}

func NewLocalStorage() codegen.ServerInterface {
	return &LocalStorage{}
}
