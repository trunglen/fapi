package staff

import (
	"io/ioutil"
	"os"
	xcontroller "revel-modules/xcontroller/app/controllers"

	"github.com/revel/revel"
)

type Staff struct {
	xcontroller.XController
}

func (c Staff) Add() revel.Result {
	return c.Success()
}
func (c Staff) Update() revel.Result {
	return c.Success()
}
func (c Staff) Delete() revel.Result {
	return c.Success()
}

func (c Staff) Upload(avatar []byte) revel.Result {
	upload, _ := revel.Config.String("upload.avatar")
	id := c.Params.Get("id")
	ioutil.WriteFile(upload+id, avatar, os.ModeAppend)
	return c.Success()
}
