package object

import "fmt"

type cat struct {
	name string
}

func (c cat) setName(s string) {
	c.name = s
}

func (c *cat) seriousSetName(s string) {
	c.name = s
}

func myTest() {
	c := cat{}
	c.setName("lucy")
	fmt.Println("cat 1 name is ", c.name)
	c.seriousSetName("lawrence")
	fmt.Println("seriously cat 1 name is ", c.name)

	c2 := &cat{}
	c2.setName("lu ben wei")
	fmt.Println("cat 2 name is ", c2.name)
	c2.seriousSetName("Zz1tail")
	fmt.Println("seriously cat 2 name is ", c2.name)

}
