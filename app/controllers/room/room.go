package room

import (
	xcontroller "github.com/trunglen/revel-modules/controller/app/controllers"

	"github.com/revel/revel"
)

type Room struct {
	*xcontroller.XController
}

func (c Room) Join(ws revel.ServerWebSocket) {

}
