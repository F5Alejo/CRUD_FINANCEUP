package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"crud_auth/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// CodigoVerificacionController expone el CRUD usado por el flujo de login.
type CodigoVerificacionController struct{ beego.Controller }

func (c *CodigoVerificacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

func (c *CodigoVerificacionController) Post() {
	var v models.CodigoVerificacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	if _, err := orm.NewOrm().Insert(&v); err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = v
	c.ServeJSON()
}

func (c *CodigoVerificacionController) GetAll() {
	var rows []models.CodigoVerificacion
	qs := orm.NewOrm().QueryTable(new(models.CodigoVerificacion))
	for k, v := range parseAuthQuery(c.GetString("query")) {
		if k == "id_usuario" {
			qs = qs.Filter("id_usuario", v)
		} else if k == "proposito" || k == "canal" || k == "activo" {
			qs = qs.Filter(k, v)
		}
	}
	if sortby := c.GetString("sortby"); sortby != "" {
		order := strings.ToLower(c.GetString("order"))
		if order == "desc" {
			sortby = "-" + sortby
		}
		qs = qs.OrderBy(sortby)
	}
	limit, _ := c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}
	offset, _ := c.GetInt("offset")
	_, err := qs.Limit(limit, offset).All(&rows)
	if err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	c.Data["json"] = rows
	c.ServeJSON()
}

func (c *CodigoVerificacionController) Put() {
	var v models.CodigoVerificacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	v.Id, _ = strconv.Atoi(c.Ctx.Input.Param(":id"))
	if v.Id == 0 {
		writeAuthError(&c.Controller, http.StatusBadRequest, errors.New("id invalido"))
		return
	}
	o := orm.NewOrm()
	old := models.CodigoVerificacion{Id: v.Id}
	if err := o.Read(&old); err != nil {
		writeAuthError(&c.Controller, http.StatusNotFound, err)
		return
	}
	models.ConservarFechas(&v, &old)
	if _, err := o.Update(&v); err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// CuentaSocialController expone las operaciones usadas por login social.
type CuentaSocialController struct{ beego.Controller }

func (c *CuentaSocialController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

func (c *CuentaSocialController) Post() {
	var v models.CuentaSocial
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	if _, err := orm.NewOrm().Insert(&v); err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = v
	c.ServeJSON()
}

func (c *CuentaSocialController) GetAll() {
	var rows []models.CuentaSocial
	qs := orm.NewOrm().QueryTable(new(models.CuentaSocial))
	for k, v := range parseAuthQuery(c.GetString("query")) {
		switch k {
		case "id_usuario":
			qs = qs.Filter("id_usuario", v)
		case "proveedor", "id_externo", "activo":
			qs = qs.Filter(k, v)
		}
	}
	limit, _ := c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}
	offset, _ := c.GetInt("offset")
	_, err := qs.Limit(limit, offset).All(&rows)
	if err != nil {
		writeAuthError(&c.Controller, http.StatusBadRequest, err)
		return
	}
	c.Data["json"] = rows
	c.ServeJSON()
}

func parseAuthQuery(raw string) map[string]string {
	result := make(map[string]string)
	for _, part := range strings.Split(raw, ",") {
		kv := strings.SplitN(part, ":", 2)
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		}
	}
	return result
}

func writeAuthError(c *beego.Controller, status int, err error) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = err.Error()
	c.ServeJSON()
}
