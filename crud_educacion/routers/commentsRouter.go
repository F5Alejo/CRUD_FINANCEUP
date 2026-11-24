package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:CertificadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ContenidoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:EscuelaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:InstructorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:LeccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloDetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ModuloEducativoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoEducativoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"] = append(beego.GlobalControllerRouter["crud_educacion/controllers:ProgresoLeccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
