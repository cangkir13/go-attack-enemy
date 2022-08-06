package enemies

import "fmt"

func (h *Enemy) ProfileStatus() {
	fmt.Println("--------HERO--------")
	fmt.Println("Name \t:", h.Name)
	fmt.Println("Damage \t:", h.Damage)
	fmt.Println("Type \t:", h.Type)
	fmt.Println("HP \t:", h.Hp)
	fmt.Println("--------------------")
}
