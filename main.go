package main

import (
	p "builder-pattern/phone"
	"fmt"
)

func main() {
	direc := p.Director{}
	iPhone := &p.Iphone{}
	direc.SetBuilder(iPhone)
	phone := direc.BuildPhone()
	
	fmt.Println("Iphone has camera? ", getAns(phone.Camera))
	fmt.Println("Iphone has Dual Sim? ", getAns(phone.DualSim))
	fmt.Println("Iphone has Torch? ", getAns(phone.Torch))
	fmt.Println("Iphone has Color Display? ", getAns(phone.ColorDisplay))

	samsungPhone := &p.Samsung{}
	direc.SetBuilder(samsungPhone)
	phone = direc.BuildPhone()

    fmt.Println("Samsung Galaxy has camera? ", getAns(phone.Camera))
	fmt.Println("Samsung Galaxy has Dual Sim? ", getAns(phone.DualSim))
	fmt.Println("Samsung Galaxy has Torch? ", getAns(phone.Torch))
	fmt.Println("Samsung Galaxy has Color Display? ", getAns(phone.ColorDisplay))
}

func getAns(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}
