package controllers

import (
	"crud_soporte/models"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type RespuestaPqrController struct{ beego.Controller }

func (c *RespuestaPqrController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetByPqr", c.GetByPqr)
}

func (c *RespuestaPqrController) Post() {
	var v models.RespuestaPqr
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	if v.IdPqr == nil || v.IdPqr.Id == 0 || v.Autor == "" || v.Mensaje == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "id_pqr, autor y mensaje son obligatorios"}
		c.ServeJSON()
		return
	}
	if _, err := models.AddRespuestaPqr(&v); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = v
	c.ServeJSON()
}

func (c *RespuestaPqrController) GetByPqr() {
	id, err := c.GetInt("id_pqr")
	if err != nil || id < 1 {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "id_pqr invalido"}
		c.ServeJSON()
		return
	}
	activo := c.GetString("activo") != "false"
	respuestas, err := models.GetRespuestasPqr(id, activo)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = respuestas
	}
	c.ServeJSON()
}
