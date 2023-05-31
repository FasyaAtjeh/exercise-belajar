package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	prorata := float64(j.WorkingMonth) / 12.0
	if prorata > 1.0 {
		prorata = 1.0
	}
	return float64(j.BaseSalary) * 1.0 * prorata
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	prorata := float64(s.WorkingMonth) / 12.0
	if prorata > 1.0 {
		prorata = 1.0
	}
	return float64(s.BaseSalary)*2.0*prorata + float64(s.BaseSalary)*s.PerformanceRate
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	prorata := float64(m.WorkingMonth) / 12.0
	if prorata > 1.0 {
		prorata = 1.0
	}
	return float64(m.BaseSalary)*2.0*prorata + float64(m.BaseSalary)*m.PerformanceRate + float64(m.BaseSalary)*m.BonusManagerRate
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalBonus := 0.0
	for _, employee := range employees {
		totalBonus += employee.GetBonus()
	}
	return totalBonus
}
func main() {
	junior := Junior{Name: "Junior 1", BaseSalary: 100000, WorkingMonth: 12}
	fmt.Println(EmployeeBonus(junior))
}
