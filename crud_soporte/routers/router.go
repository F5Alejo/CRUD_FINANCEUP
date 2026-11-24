// @APIVersion 1.0.0
// @Title FinanceUp - crud_soporte
// @Description CRUD del schema soporte de la base de datos financeup
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"crud_soporte/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/registro_actividad",
			beego.NSInclude(
				&controllers.RegistroActividadController{},
			),
		),

		beego.NSNamespace("/respuesta_pqr",
			beego.NSInclude(
				&controllers.RespuestaPqrController{},
			),
		),

		beego.NSNamespace("/adjunto",
			beego.NSInclude(
				&controllers.AdjuntoController{},
			),
		),

		beego.NSNamespace("/estado_pqr",
			beego.NSInclude(
				&controllers.EstadoPqrController{},
			),
		),

		beego.NSNamespace("/pqr",
			beego.NSInclude(
				&controllers.PqrController{},
			),
		),

		beego.NSNamespace("/pregunta_frecuente",
			beego.NSInclude(
				&controllers.PreguntaFrecuenteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
