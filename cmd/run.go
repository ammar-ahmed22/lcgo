package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ammar-ahmed22/lcgo/utils"
	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
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
				filtered := lo.Filter(lo.Keys(problems), func (id string, _ int) bool {
					return strings.Contains(id, leetcodeID)
				})
				if len(filtered) == 0 {
					color.Yellow("No matches for search \"%s\"", leetcodeID)
					return nil
				}
				var problemID string
				huh.NewSelect[string]().Title(fmt.Sprintf("Select from problems matching \"%s\"", leetcodeID)).Options(lo.Map(filtered, func(id string, _ int) huh.Option[string] {
					return huh.NewOption(id, id)
				})...).Value(&problemID).Run()
				problem, _ := problems[problemID]
				return runProblem(problem.Directory)
			}
			return runProblem(problem.Directory)
		} else if len(args) == 0 {
			var problemID string
			huh.NewSelect[string]().Title("Select a problem to run").Options(lo.Map(lo.Keys(problems), func(id string, _ int) huh.Option[string] {
				return huh.NewOption(id, id)
			})...).Value(&problemID).Run()
			problem, _ := problems[problemID]
			return runProblem(problem.Directory)
		} else {
			return fmt.Errorf(color.RedString("Too many arguments!\n"))
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
