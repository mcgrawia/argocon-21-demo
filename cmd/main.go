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

	mealCmds := make(map[string]*bool, len(meals.All))
	for _, meal := range meals.All {
		mealCmds[meal.Name()] = parser.Flag("", meal.Name(), &argparse.Options{Help: fmt.Sprintf("Cook %s.", meal.Name())})
	}

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	var mealsToBuild []string
	for mealName, mealFlag := range mealCmds {
		if *mealFlag {
			mealsToBuild = append(mealsToBuild, mealName)
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
