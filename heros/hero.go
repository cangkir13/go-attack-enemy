package heros

import "fmt"

func (p *Hero) PowerAttack() int {
	roles := p.Role
	switch roles {
	case "mage":
		return p.Damage + 2
	default:
		return p.Damage + 10
	}
}

func (h *Hero) ProfileHero() {
	fmt.Println("--------HERO--------")
	fmt.Println("Name \t:", h.Name)
	fmt.Println("Damage \t:", h.Damage)
	fmt.Println("Role \t:", h.Role)
	fmt.Println("HP \t:", h.Hp)
	fmt.Println("--------------------")
}
