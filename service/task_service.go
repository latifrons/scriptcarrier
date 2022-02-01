package service

import (
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
	}
	err := s.Db.Create(&task).Error
	if err != nil {
		return err
	}

	tools.EnsureFolder(projectFolder, 0755)

	err = os.WriteFile(path.Join(task.ScriptPath, req.ScriptFileName), []byte(req.ScriptContent), 0755)
	if err != nil {
		return err
	}
	return nil

}
