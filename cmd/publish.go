package cmd

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ammar-ahmed22/lcgo/fs"
	"github.com/ammar-ahmed22/lcgo/utils"
	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func publishProblem(leetcodeID string, problems map[string]utils.YamlProblem) error {
	problem, exists := problems[leetcodeID]
	if !exists {
		return fmt.Errorf(color.RedString("Problem \"%s\" does not exist\n", leetcodeID))
	}
	if problem.Published {
		return fmt.Errorf(color.RedString("Problem \"%s\" is already published\n", leetcodeID))
	}

	var addTags bool
	var tags []string
	huh.NewConfirm().Title("Would you like to add tags?").Affirmative("Yes").Negative("No").Value(&addTags).Run()
	if addTags {
		var rawTags string
		huh.NewInput().Title("Provide tags (comma-separated)").Value(&rawTags).Run()
		tags = lo.Map(strings.Split(rawTags, ","), func(tag string, _ int) string {
			return strings.TrimSpace(tag)
		})
	}
	// Read the solution file
	solutionFile, err := fs.ReadFileString(fmt.Sprintf("%s/main.go", problem.Directory))
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to read solution file for problem \"%s\"\n", leetcodeID))
	}

	// Extract the solution from the file
	solution, s, e := utils.BetweenLines(solutionFile, "// <-- DO NOT REMOVE: PROBLEM START -->", "// <-- DO NOT REMOVE: PROBLEM END -->")
	if s == -1 || e == -1 {
		return fmt.Errorf(color.RedString("Unable to find solution in file for problem \"%s\"\n", leetcodeID))
	}

	// Read the docs file
	docsFile, err := fs.ReadFileString(fmt.Sprintf("%s/docs.md", problem.Directory))
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to read docs file for problem \"%s\"\n", leetcodeID))
	}
	// Add solution to docs file
	docsFile += fmt.Sprintf("\n## Solution\n\n```go\n%s\n```\n", solution)
	// Write the docs file
	err = fs.WriteFileString(fmt.Sprintf("%s/docs.md", problem.Directory), docsFile)
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to write docs file for problem \"%s\"\n", leetcodeID))
	}
	fmt.Printf("Added %s solution to %s\n", color.CyanString(leetcodeID), color.CyanString("%s/docs.md", problem.Directory))
	fmt.Printf("%s Solution was appended to the docs file. Double-check to ensure it is correct.\n", color.BlueString("NOTE:"))

	// Read the README file
	readme, err := fs.ReadFileString("README.md")
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to read README.md\n"))
	}

	parts := strings.Split(problem.Directory, "-")
	problemName := strings.TrimSpace(parts[1])
	uppercaseDifficulty := strings.ToUpper(problem.Difficulty)
	encodedProblemDirectory := url.PathEscape(problem.Directory)
	readme = utils.TemplateReplace(readme, map[string]string{
		fmt.Sprintf("<!-- %s PROBLEMS -->", uppercaseDifficulty): fmt.Sprintf("- [%s](./%s/docs.md)\n<!-- %s PROBLEMS -->", problemName, encodedProblemDirectory, uppercaseDifficulty),
	})

	easyCount := 0
	medCount := 0
	hardCount := 0
	problem.Published = true
	problem.Date = lo.ToPtr(time.Now().Format("2006-01-02"))
	if addTags {
		problem.Tags = tags
	}
	problems[leetcodeID] = problem
	for _, p := range problems {
		if p.Published {
			if p.Difficulty == "easy" {
				easyCount++
			}
			if p.Difficulty == "medium" {
				medCount++
			}
			if p.Difficulty == "hard" {
				hardCount++
			}
		}
	}
	readme = utils.ReplaceLine(readme, `alt="Easy Badge"`, fmt.Sprintf(`        <img alt="Easy Badge" src="https://img.shields.io/badge/%d-easy-green">`, easyCount))
	readme = utils.ReplaceLine(readme, `alt="Medium Badge"`, fmt.Sprintf(`        <img alt="Medium Badge" src="https://img.shields.io/badge/%d-medium-yellow">`, medCount))
	readme = utils.ReplaceLine(readme, `alt="Hard Badge"`, fmt.Sprintf(`        <img alt="Hard Badge" src="https://img.shields.io/badge/%d-hard-red">`, hardCount))

	err = fs.WriteFileString("README.md", readme)
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to write README.md\n"))
	}
	fmt.Printf("Added %s to %s\n", color.CyanString(leetcodeID), color.CyanString("README.md"))

	err = utils.WriteYamlProblems("problems.yaml", problems)
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to write problems.yaml\n"))
	}
	fmt.Printf("Updated %s status in %s\n", color.CyanString(leetcodeID), color.CyanString("problems.yaml"))
	return nil
}

func filterUnpublishedProblems(problems map[string]utils.YamlProblem) []string {
	return lo.Map(
		lo.Filter(lo.Entries(problems), func(entry lo.Entry[string, utils.YamlProblem], _ int) bool {
			return !entry.Value.Published
		}),
		func(entry lo.Entry[string, utils.YamlProblem], _ int) string {
			return entry.Key
		},
	)

}

var publishCmd = &cobra.Command{
	Use:   "publish <LEETCODE_ID>",
	Short: "Publish the provided problem to the README",
	RunE: func(cmd *cobra.Command, args []string) error {
		problems, err := utils.ReadYamlProblems("problems.yaml")
		if err != nil {
			return fmt.Errorf(color.RedString("Unable to read problems.yaml\n"))
		}
		if len(args) == 1 {
			leetcodeID := args[0]
			return publishProblem(leetcodeID, problems)
		} else if len(args) == 0 {
			unpublished := filterUnpublishedProblems(problems)
			if len(unpublished) == 0 {
				fmt.Println("Nothing to publish!")
				return nil
			}
			var leetcodeID string
			huh.NewSelect[string]().Title("Select a problem to publish").Options(lo.Map(unpublished, func(id string, _ int) huh.Option[string] {
				return huh.NewOption(id, id)
			})...).Value(&leetcodeID).Run()
			return publishProblem(leetcodeID, problems)
		} else {
			return fmt.Errorf(color.RedString("Too many arguments!\n"))
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
