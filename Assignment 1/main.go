package main

import (
	"Assignment1/Administrator"
	"Assignment1/Assistant"
	"Assignment1/Guard"
	"Assignment1/Keeper"
	"Assignment1/Runner"
	"fmt"
)

func main() {
	runner := Runner.NewRunner("Runner", 260000, "Atyrau")

	keeper := Keeper.NewKeeper("Keeper", 370000, "Atyrau")

	guard := Guard.NewGuard("Guard", 280000, "Atyrau")

	assistant := Assistant.NewAssistant("Assistant", 300000, "Atyrau")

	administrator := Administrator.NewAdministrator("Administrator", 170000, "Atyrau")

	fmt.Println("Runner Position:", runner.GetPosition())
	runner.SetPosition("Senior Runner")
	fmt.Println("Runner Position (Updated):", runner.GetPosition())

	fmt.Println("Keeper Salary:", keeper.GetSalary())
	keeper.SetSalary(385000)
	fmt.Println("Keeper Salary (Updated):", keeper.GetSalary())

	fmt.Println("Guard Position:", guard.GetPosition())

	fmt.Println("Assistant Position:", assistant.GetPosition())

	fmt.Println("Administrator Address:", administrator.GetAddress())
}
