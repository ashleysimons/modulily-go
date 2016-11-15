package routers

import (
	"github.com/ashleysimons/modulily-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/types", &controllers.TypesController{})
}
