package main

import (
	enemy "belajar/enemies"
	"belajar/heros"
	"fmt"
)

func main() {
	// struct
	// users := Person{name: "umar", demage: 12}
	// fmt.Println(users)

	kagura := heros.Hero{Name: "Kagura", Damage: 12, Role: "mage", Hp: 100}
	kagura.ProfileHero()
	kaguraPower := kagura.PowerAttack()
	fmt.Println("Kagura true damge: ", kaguraPower)

	fmt.Println("\n \t VS \n")

	minion := enemy.Enemy{Name: "minion offliner", Damage: 5, Hp: 50, Type: "minion"}
	minion.ProfileStatus()

}
