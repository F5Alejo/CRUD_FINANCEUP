package models

import "github.com/beego/beego/v2/client/orm"

// CuentaSocial vincula una cuenta local con un proveedor de autenticacion social.
type CuentaSocial struct {
	Id             int      `orm:"column(id_cuenta_social);pk;auto"`
	IdUsuario      *Usuario `orm:"column(id_usuario);rel(fk)"`
	Proveedor      string   `orm:"column(proveedor)"`
	IdExterno      string   `orm:"column(id_externo)"`
	EmailProveedor string   `orm:"column(email_proveedor)"`
	Activo         bool     `orm:"column(activo)"`
}

func (t *CuentaSocial) TableName() string { return "cuenta_social" }

func init() { orm.RegisterModel(new(CuentaSocial)) }
