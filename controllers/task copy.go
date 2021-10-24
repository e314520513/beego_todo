package bk

import(
	"encoding/json"
	"strconv"
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

func (this *TaskController) GetTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	this.Data["json"] = t
	this.ServeJSON()
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (this *TaskController) UpdateTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &t); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
}