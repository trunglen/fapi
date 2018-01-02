package controllers

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/revel/revel"
)

type App struct {
	revel.Controller
}

func (c App) Index() revel.Result {
	return c.RenderJSON(map[string]interface{}{
		"Count":  "len(files)",
		"Status": "Successfully uploaded",
	})
}

func (c App) HandleUpload(avatar []byte) revel.Result {
	// Validation rules.
	fmt.Println(c.Params.Files["avatar"][0].Filename)
	ioutil.WriteFile("icrating", avatar, os.ModeAppend)
	return c.RenderJSON(map[string]interface{}{
		"dsad": "dsadasd",
	})
}
