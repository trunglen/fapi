package smiley

import (
	"g/x/math"
	"g/x/web"
	"io/ioutil"
	"os"
	"tamson/app/models"

	xcontroller "github.com/trunglen/revel-modules/xcontroller/app/controllers"

	"github.com/revel/revel"
)

type Smiley struct {
	xcontroller.XController
}

func (c *Smiley) Add() revel.Result {
	var smiley *models.Smiley
	c.MustDecodeBody(&smiley)
	smiley.Validate(c.Validation)
	web.AssertValidation(c.Validation.Errors)
	web.AssertNil(smiley.Create())
	return c.Success()
}

func (c Smiley) Upload(icon []byte) revel.Result {
	var fileName = math.RandString("smiley", 4)
	var folder, _ = revel.Config.String("upload.smiley")
	web.AssertNil(ioutil.WriteFile(folder+fileName, icon, os.ModeAppend))
	return c.Success()
}

func (c Smiley) List(icon []byte) revel.Result {
	var listIcon = []string{}
	files, err := ioutil.ReadDir("./public/smiley")
	web.AssertNil(err)
	for _, item := range files {
		listIcon = append(listIcon, item.Name())
	}
	return c.SendData(listIcon)
}
