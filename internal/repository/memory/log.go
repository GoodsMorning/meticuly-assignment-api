package memory

import (
	"meticuly-assignment-api/internal/model"
)

func (m *Memory) GetLogs() []model.Log {
	return m.logs
}

func (m *Memory) WriteLog(log model.Log) model.Log {
	m.logs = append([]model.Log{log}, m.logs...)
	return log
}

func mockLogs() []model.Log {
	return []model.Log{
		{
			Time:    "14:32:08",
			OrderID: "ORD-001",
			Message: "Packaging started",
			Level:   "info",
		},
		{
			Time:    "14:28:41",
			OrderID: "ORD-002",
			Message: "Assembly started",
			Level:   "info",
		},
		{
			Time:    "14:15:20",
			OrderID: "ORD-002",
			Message: "Machining-1 finished",
			Level:   "success",
		},
		{
			Time:    "13:58:02",
			OrderID: "ORD-001",
			Message: "Quality Check finished",
			Level:   "success",
		},
		{
			Time:    "13:45:17",
			OrderID: "ORD-001",
			Message: "Assembly finished",
			Level:   "success",
		},
		{
			Time:    "13:30:05",
			OrderID: "ORD-001",
			Message: "Machining-1 finished",
			Level:   "success",
		},
	}
}