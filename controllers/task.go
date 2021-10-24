package controllers

import(
	"encoding/json"
	"github.com/astaxie/beego"
	"beego_todo/models"
)

type TaskController struct{
	beego.Controller
}

func (this *TaskController) ListTasks(){
	res := struct{tasks []*models.Task}{models.DefaultTaskList.All()}
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *TaskController) NewTask(){
	req  := struct{Title string}{}
	if err:= json.Unmarshal(this.Ctx.Input.RequestBody,&req); err != nil {
		this.Ctx.Output.SetStatus(4000)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	t,err := models.NewTask(req.Title)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return 
	}
	models.DefaultTaskList.Save(t)
}

