package controllers

import (
	"github.com/revel/modules/orm/gorp/app/controllers"
	"github.com/revel/revel"
)

type Boards struct {
	//*revel.Controller
	gorpController.Controller
}

func (c Boards) Index() revel.Result {
	return c.Render()
}


