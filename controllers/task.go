package controllers

import(
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"gihub.com/beego/samples/todo/models"
)

type TaskController struct{
	beego.Controller
}

func (this *TaskController) ListTasks(){
	res := struct{tasks []*model.Task}{models.DefaultTaskList.All()}
	this.Data["json"] = res
	this.ServeJSON()
}

