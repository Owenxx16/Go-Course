package student

import "fmt"

type Student struct {
	Id    int
	Name  string
	Class string
	Score []float64
}

func (s Student) getInfo() string {
	return fmt.Sprintf("Id: %d | Name: %s | Class: %s | DiemTB: %.2f ", s.Id, s.Name, s.Class, s.CalculateAverageScore())
}

func (s Student) CalculateAverageScore() float64 {
	if len(s.Score) == 0 {
		return 0
	}
	var total float64
	for _, score := range s.Score {
		total += score
	}
	return total / float64(len(s.Score))
}

func (s Student) GetId() int {
	return s.Id
}
