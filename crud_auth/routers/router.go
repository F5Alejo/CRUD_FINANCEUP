// @APIVersion 1.0.0
// @Title FinanceUp - crud_auth
// @Description CRUD del schema auth de la base de datos financeup
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"crud_auth/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/usuario_rol",
			beego.NSInclude(
				&controllers.UsuarioRolController{},
			),
		),

		beego.NSNamespace("/auditoria_login",
			beego.NSInclude(
				&controllers.AuditoriaLoginController{},
			),
		),

		beego.NSNamespace("/codigo_verificacion",
			beego.NSInclude(
				&controllers.CodigoVerificacionController{},
			),
		),

		beego.NSNamespace("/credencial",
			beego.NSInclude(
				&controllers.CredencialController{},
			),
		),

		beego.NSNamespace("/cuenta_social",
			beego.NSInclude(
				&controllers.CuentaSocialController{},
			),
		),

		beego.NSNamespace("/rol",
			beego.NSInclude(
				&controllers.RolController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
