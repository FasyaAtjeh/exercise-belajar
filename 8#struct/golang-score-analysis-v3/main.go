package main

import "fmt"

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	for _, grade := range grades {
		if grade >= 0 && grade <= 100 {
			s.Grades = append(s.Grades, grade)
		}
	}
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	min := s.Grades[0]
	max := s.Grades[0]
	sum := 0
	for _, grade := range s.Grades {
		if grade < min {
			min = grade
		}
		if grade > max {
			max = grade
		}
		sum += grade
	}
	avg := float64(sum) / float64(len(s.Grades))
	return avg, min, max
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{},
	})

	fmt.Println(avg, min, max)
}
