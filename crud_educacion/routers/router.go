// @APIVersion 1.0.0
// @Title FinanceUp - crud_educacion
// @Description CRUD del schema educacion de la base de datos financeup
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"crud_educacion/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/escuela",
			beego.NSInclude(
				&controllers.EscuelaController{},
			),
		),

		beego.NSNamespace("/instructor",
			beego.NSInclude(
				&controllers.InstructorController{},
			),
		),

		beego.NSNamespace("/leccion",
			beego.NSInclude(
				&controllers.LeccionController{},
			),
		),

		beego.NSNamespace("/modulo_detalle",
			beego.NSInclude(
				&controllers.ModuloDetalleController{},
			),
		),

		beego.NSNamespace("/certificado",
			beego.NSInclude(
				&controllers.CertificadoController{},
			),
		),

		beego.NSNamespace("/modulo_educativo",
			beego.NSInclude(
				&controllers.ModuloEducativoController{},
			),
		),

		beego.NSNamespace("/progreso_educativo",
			beego.NSInclude(
				&controllers.ProgresoEducativoController{},
			),
		),

		beego.NSNamespace("/progreso_leccion",
			beego.NSInclude(
				&controllers.ProgresoLeccionController{},
			),
		),

		beego.NSNamespace("/contenido",
			beego.NSInclude(
				&controllers.ContenidoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
