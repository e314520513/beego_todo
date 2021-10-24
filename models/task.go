package models

import(
	"fmt"
)

var DefaultTaskList *TaskManager

type Task struct{
	ID int64
	Title string
	Done bool
}

type TaskManager struct{
	tasks []*Task
	lastID int64
}
func (this *TaskManager) Save(task *Task) error{
	if task.ID == 0 {
		this.lastID++
		task.ID = this.lastID
	    this.tasks = append(this.tasks,cloneTask(task))
		return nil
	}
	for i,t :=range this.tasks{
		if t.ID==task.ID{
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}

	return fmt.Errorf("unknown task")
}

func cloneTask(t *Task) *Task{
	c := *t
	return  &c
}
func NewTaskManager() *TaskManager{
	return &TaskManager{}
}

func (this *TaskManager) All() []*Task{
	return this.tasks
}

func (this * TaskManager) Find(ID int64)(*Task, bool){
	for _,t := range this.tasks{
		if t.ID == ID {
			return t,true
		}
	}
	return nil,false
}


func init(){
	DefaultTaskList = NewTaskManager()
}