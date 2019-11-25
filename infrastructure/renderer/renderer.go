package renderer

// TaskRenderer ...
type TaskRenderer struct {
}

// NewTaskRenderer ...
func NewTaskRenderer() *TaskRenderer {
	return &TaskRenderer{}
}

// CreateTask ...
func (r *TaskRenderer) CreateTask() () {
	res, err := i.DataSourceRepository.FindAll()
	if err != nil {
	  return nil, err
	}
	// Output Port の使用
	return i.OutputPort.DownloadDataSources(res)
  }