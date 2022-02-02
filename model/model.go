package model

type BasicTask struct {
	Name            string `json:"name"`
	ScriptType      string `json:"script_type"`
	ScriptFileName  string `json:"script_file_name"`
	Args            string `json:"args"`
	IntervalSeconds int    `json:"interval_seconds"`
}
type BasicTaskResult struct {
	Time            int64  `json:"time"`
	Code            int    `json:"code"`
	DurationSeconds int    `json:"duration_seconds"`
	LogPath         string `json:"log_path"`
}

type AddTaskRequest struct {
	BasicTask
	ScriptContent string `json:"script_content"`
}

type DeleteTaskRequest struct {
	Id int `json:"id"`
}

type TaskListResponseItem struct {
	BasicTask
	BasicTaskResult
}

type TaskDetailResponseItem struct {
	BasicTask
	BasicTaskResult
	ScriptContent string `json:"script_content"`
}
