package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"ianmcgraw.com/m/v2/pkg/kitchen"
	"ianmcgraw.com/m/v2/pkg/kitchen/meals"
	"ianmcgraw.com/m/v2/pkg/workflows"
	"os"
)

func main() {
	runCmd()
}

func runCmd() {
	parser := argparse.NewParser("feedme", "Use Argo Workflows to make food!")

	var mealCmds []*argparse.Command
	for _, meal := range meals.All {
		mealCmds = append(mealCmds, parser.NewCommand(meal.Name(), fmt.Sprintf("Cook %s.", meal.Name())))
	}

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	var mealsToBuild []string
	for _, mealCmd := range mealCmds {
		if mealCmd.Happened() {
			mealsToBuild = append(mealsToBuild, mealCmd.GetName())
		}
	}

	chef, err := kitchen.NewChefArgo(mealsToBuild)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = workflows.Client.Create(chef.Cook())
	if err != nil {
		fmt.Println(err)
		return
	}
}
