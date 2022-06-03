package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name         string
	Role         string
	Age          int32
	EmployeeCode int64 `copier:"EmployNum"`
	Salary       int
}

func (u *User) DoubleAge() int32 {
	return 2 * u.Age
}

type Employee struct {
	Name   string `copier:"must"`
	Age    int32  `copier:"must,nopanic"`
	Salary int    `copier:"-"`

	DoubleAge int32
	EmployeId int64 `copier:"EmployeNum"`
	SuperRole string
}

func (e *Employee) Role(role string) {
	e.SuperRole = "Super " + role
}

func main() {
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000}
		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 100000}, {Name: "jinzhu 2", Age: 30, Role: "Dev", Salary: 60000}}
		employee  = Employee{Salary: 150000}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)
	fmt.Printf("%#v \n", employee)

	copier.Copy(&employees, &user)
	fmt.Printf("%#v \n", employees)

	employees = []Employee{}
	copier.Copy(&employees, &users)
	fmt.Printf("%#v \n", employees)

	map1 := map[int]int{3: 6, 4: 8}
	map2 := map[int32]int8{}
	copier.Copy(&map2, map1)

	fmt.Printf("%#v \n", map2)
}
