package controllers

import(
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

