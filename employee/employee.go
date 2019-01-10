// Package employee 提供员工数据结构及相关方法
package employee

import "time"

// Employee 员工
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// EmployeeByID 根据员工ID返回对应的员工
func EmployeeByID(id int) *Employee {
	dilbert := new(Employee)
	if id != 1 {
		return dilbert
	}
	dilbert.ID = 1
	dilbert.Name = "Dilbert"
	dilbert.Position = "Engineer"
	dilbert.Salary = 290000
	return dilbert
}
