package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ammar-ahmed22/lcgo/utils"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func runProblem(dirname string) error {
	runCmd := exec.Command("go", "run", ".")
	runCmd.Dir = dirname
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	if err := runCmd.Run(); err != nil {
		return fmt.Errorf(color.RedString("Failed to run problem\n"))
	}
	return nil
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run LeetCode problems",
	RunE: func(cmd *cobra.Command, args []string) error {
		problems, err := utils.ReadYamlProblems("problems.yaml")
		if err != nil {
			return fmt.Errorf(color.RedString("Unable to read problems.yaml\n"))
		}

		if len(args) == 1 {
			leetcodeID := args[0]
			problem, exists := problems[leetcodeID]
			if !exists {
				return fmt.Errorf(color.RedString("Problem \"%s\" does not exist\n", leetcodeID))
			}
			return runProblem(problem.Directory)
		} else if len(args) == 0 {
			prompt := promptui.Select{
				Label: "Select problem to run",
				Items: lo.Keys(problems),
			}
			_, problemID, err := prompt.Run()
			if err != nil {
				return fmt.Errorf(color.RedString("Failed to get selection\n"))
			}
			problem, exists := problems[problemID]
			if !exists {
				return fmt.Errorf(color.RedString("Problem \"%s\" does not exist\n", problemID))
			}
			return runProblem(problem.Directory)
		} else {
			return fmt.Errorf(color.RedString("Too many arguments!\n"))
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
