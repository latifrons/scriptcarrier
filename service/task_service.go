package service

import (
	"database/sql"
	"github.com/latifrons/scriptcarrier/db"
	"github.com/latifrons/scriptcarrier/model"
	"github.com/latifrons/scriptcarrier/tools"
	"gorm.io/gorm"
	"os"
	"path"
)

type TaskService struct {
	Db       *gorm.DB
	RootPath string
}

func (s *TaskService) AddTask(req model.AddTaskRequest) error {
	projectFolder := path.Join(s.RootPath, req.Name)

	task := db.Task{
		Name:            req.Name,
		ScriptType:      req.ScriptType,
		ScriptPath:      path.Join(projectFolder, req.ScriptFileName),
		Args:            req.Args,
		IntervalSeconds: req.IntervalSeconds,
		NextRunTime:     sql.NullTime{}, // run instantly
		LastRunTime:     sql.NullTime{},
		LastRunCode:     0,
		LastRunDuration: 0,
		LastRunLogPath:  "",
	}
	err := s.Db.Create(&task).Error
	if err != nil {
		return err
	}

	tools.EnsureFolder(projectFolder, 0755)

	err = os.WriteFile(task.ScriptPath, []byte(req.ScriptContent), 0755)
	if err != nil {
		return err
	}
	return nil

}

func (s *TaskService) ListTask() (list []db.Task, err error) {
	err = s.Db.Model(db.Task{}).Order("last_run_time desc").Find(&list).Error
	if err != nil {
		return
	}
	return

}
