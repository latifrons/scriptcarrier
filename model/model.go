package model

type AddTaskRequest struct {
	Name            string `json:"name"`
	ScriptType      string `json:"script_type"`
	ScriptFileName  string `json:"script_file_name"`
	ScriptContent   string `json:"script_content"`
	Args            string `json:"args"`
	IntervalSeconds int    `json:"interval_seconds"`
}

type UpdateTaskRequest struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	ScriptType      string `json:"script_type"`
	ScriptFileName  string `json:"script_file_name"`
	ScriptContent   string `json:"script_content"`
	Args            string `json:"args"`
	IntervalSeconds int    `json:"interval_seconds"`
}

type DeleteTaskRequest struct {
	Id int `json:"id"`
}
