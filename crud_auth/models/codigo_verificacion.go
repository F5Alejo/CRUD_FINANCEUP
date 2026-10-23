package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// CodigoVerificacion guarda los codigos de un solo uso enviados durante el login.
type CodigoVerificacion struct {
	Id                int       `orm:"column(id_codigo);pk;auto"`
	IdUsuario         *Usuario  `orm:"column(id_usuario);rel(fk)"`
	Proposito         string    `orm:"column(proposito)"`
	Canal             string    `orm:"column(canal)"`
	CodigoHash        string    `orm:"column(codigo_hash)"`
	FechaExpiracion   time.Time `orm:"column(fecha_expiracion);type(timestamp without time zone)"`
	FechaUso          time.Time `orm:"column(fecha_uso);type(timestamp without time zone);null"`
	Intentos          int       `orm:"column(intentos)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *CodigoVerificacion) TableName() string { return "codigo_verificacion" }

func init() { orm.RegisterModel(new(CodigoVerificacion)) }
