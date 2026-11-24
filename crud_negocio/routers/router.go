// @APIVersion 1.0.0
// @Title FinanceUp - crud_negocio
// @Description CRUD del schema negocio de la base de datos financeup
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"crud_negocio/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/asesor_bancario",
			beego.NSInclude(
				&controllers.AsesorBancarioController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/conversacion_usuario_asesor",
			beego.NSInclude(
				&controllers.ConversacionUsuarioAsesorController{},
			),
		),

		beego.NSNamespace("/credito_desembolsado",
			beego.NSInclude(
				&controllers.CreditoDesembolsadoController{},
			),
		),

		beego.NSNamespace("/lead",
			beego.NSInclude(
				&controllers.LeadController{},
			),
		),

		beego.NSNamespace("/producto_crediticio",
			beego.NSInclude(
				&controllers.ProductoCrediticioController{},
			),
		),

		beego.NSNamespace("/producto_detalle",
			beego.NSInclude(
				&controllers.ProductoDetalleController{},
			),
		),

		beego.NSNamespace("/contacto_asesor",
			beego.NSInclude(
				&controllers.ContactoAsesorController{},
			),
		),

		beego.NSNamespace("/producto_pregunta",
			beego.NSInclude(
				&controllers.ProductoPreguntaController{},
			),
		),

		beego.NSNamespace("/producto_tarifa",
			beego.NSInclude(
				&controllers.ProductoTarifaController{},
			),
		),

		beego.NSNamespace("/transaccion_comision",
			beego.NSInclude(
				&controllers.TransaccionComisionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
