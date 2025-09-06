package teacher

import "fmt"

type Teacher struct {
	Id         int
	Name       string
	Subject    string
	BaseSalary float64
	Bonus      float64
}

func (t Teacher) GetInfo() string {
	return fmt.Sprintf("Id: %d | Name: %s | Subject: %s | BaseSalary: %.2f | Bonus: %.2f", t.Id, t.Name, t.Subject, t.BaseSalary, t.Bonus)
}

func (t Teacher) CalculateTotalSalary() float64 {
	return t.BaseSalary + t.Bonus
}

func (t Teacher) GetId() int {
	return t.Id
}
