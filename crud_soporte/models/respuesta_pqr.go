package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type RespuestaPqr struct {
	Id             int       `orm:"column(id_respuesta);pk;auto"`
	IdPqr          *Pqr      `orm:"column(id_pqr);rel(fk)"`
	IdUsuario      *int      `orm:"column(id_usuario);null"`
	Autor          string    `orm:"column(autor)"`
	Mensaje        string    `orm:"column(mensaje)"`
	FechaRespuesta time.Time `orm:"column(fecha_respuesta);type(timestamp without time zone);auto_now_add"`
	Activo         bool      `orm:"column(activo)"`
}

func (t *RespuestaPqr) TableName() string { return "respuesta_pqr" }

func init() { orm.RegisterModel(new(RespuestaPqr)) }

func AddRespuestaPqr(m *RespuestaPqr) (int64, error) {
	return orm.NewOrm().Insert(m)
}

func GetRespuestasPqr(idPqr int, activo bool) ([]RespuestaPqr, error) {
	var respuestas []RespuestaPqr
	_, err := orm.NewOrm().QueryTable(new(RespuestaPqr)).
		Filter("id_pqr", idPqr).
		Filter("activo", activo).
		OrderBy("id_respuesta").All(&respuestas)
	return respuestas, err
}
