package main

import (
	"fmt"
	"time"
)

// Employee as test struct
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

// EmployeeByID as demo function
func EmployeeByID(id int, employees []Employee) (*Employee, error) {
	for _, employ := range employees {
		if id == employ.ID {
			return &employ, nil
		}
	}
	return nil, fmt.Errorf("No employee found by given ID: %d", id)
}

func main() {
	e1 := Employee{
		ID:        0,
		Name:      "Shawn",
		Address:   "14 Mitchell CT, NY",
		Salary:    1000,
		DoB:       time.Now(),
		ManagerID: 0,
	}
	e2 := Employee{
		ID:        1,
		Name:      "Yadong",
		Address:   "516 60th ST, NY",
		Salary:    1500,
		ManagerID: 1,
	}
	e3 := Employee{
		ID:        2,
		Name:      "Ryan",
		Address:   "City Point Plaza, NY",
		Salary:    1000,
		ManagerID: 1,
	}
	// input employees and target id
	source := []Employee{e1, e2, e3}
	// tester := &source[0]
	// tester.Name = "Tyler"
	// fmt.Printf("Employee Name: %s\n", tester.Name)
	// fmt.Printf("Employee Name: %s\n", source[0].Name)
	id := 0
	// find target employee
	employee, err := EmployeeByID(id, source)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Found empolyee, ID: %d, Name: %s \n", employee.ID, employee.Name)
	fmt.Printf("Start to change employee \n")
	employee.Name = "Tyler"
	index := employee.ID
	fmt.Printf("Changed empolyee, ID: %d, Name: %s \n", source[index].ID, source[index].Name)
}
