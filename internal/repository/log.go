
package repository

import "meticuly-assignment-api/internal/model"

func (r *Repository) GetLogs() []model.Log {
	return r.mockMemory.GetLogs()
}

func (r *Repository) WriteLog(log model.Log) model.Log {
	return r.mockMemory.WriteLog(log)
}