package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/JohannesKaufmann/html-to-markdown"
	"github.com/ammar-ahmed22/lcgo/fs"
	"github.com/ammar-ahmed22/lcgo/utils"
	"github.com/briandowns/spinner"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func createProblemDirectory(dirname string) error {
	// Check if the directory already exists
	if _, err := os.Stat(dirname); !os.IsNotExist(err) {
		return fmt.Errorf(color.RedString("Directory \"%s\" already exists\n", dirname))
	}
	err := os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to create directory")
	}
	return nil
}

func createDocsFile(dirname, problemTitle, problemDescription string) error {
	docsFile, err := os.Create(fmt.Sprintf("%s/docs.md", dirname))
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to create docs.md file\n"))
	}
	docsString := fmt.Sprintf("# %s\n\n## Problem\n\n%s\n\n## Approach\n\n", problemTitle, problemDescription)
	docsFile.WriteString(docsString)
	docsFile.Close()
	return nil
}

type LeetcodeProblem struct {
	DescriptionHtml string
	Difficulty      string
	ProblemTitle    string
	CodeSnippet     string
}

func getLeetcodeProblemData(leetcodeID string) (*LeetcodeProblem, error) {
	url := fmt.Sprintf("https://leetcode.com/problems/%s", leetcodeID)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),     // Disable headless mode
		chromedp.Flag("disable-gpu", false), // (optional) Enable GPU for rendering
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // hides headless
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
			"(KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var (
		descriptionHtml string
		difficulty      string
		problemTitle    string
		codeSnippet     string
	)

	err := chromedp.Run(ctx,
		// Navigate to the URL
		chromedp.Navigate(url),
		// Wait for the description to load
		chromedp.WaitVisible(`div[data-track-load="description_content"]`, chromedp.ByQuery),
		// Extract the description HTML
		chromedp.InnerHTML(`div[data-track-load="description_content"]`, &descriptionHtml, chromedp.ByQuery),
		// Extract the difficulty
		chromedp.Text(`[class*="text-difficulty"]`, &difficulty, chromedp.ByQuery),
		// Extract the problem title
		chromedp.Text(fmt.Sprintf(`[href="/problems/%s/"]`, leetcodeID), &problemTitle, chromedp.ByQuery),

		// Wait for the language selector to load
		chromedp.WaitVisible(`//button[text()='C++']`, chromedp.BySearch),
		// Click on the language selector
		chromedp.Click(`//button[text()='C++']`, chromedp.BySearch),
		// Wait for the language selection to load
		chromedp.WaitVisible(`//div[text()='Go']`, chromedp.BySearch),
		// Click on the language selection
		chromedp.Click(`//div[text()='Go']`, chromedp.BySearch),
		// Wait for the code editor to load
		chromedp.WaitVisible(`//span[text()='func']`, chromedp.BySearch),
		// Extract the code snippet
		chromedp.Text(`.view-lines`, &codeSnippet, chromedp.ByQuery),
	)
	if err != nil {
		return nil, fmt.Errorf(color.RedString("Unable to fetch leetcode problem data\n"))
	}

	return &LeetcodeProblem{
		DescriptionHtml: descriptionHtml,
		Difficulty:      strings.ToLower(difficulty),
		ProblemTitle:    problemTitle,
		CodeSnippet:     codeSnippet,
	}, nil
}

func createGoModule(dirname string, moduleName string) error {
	modCmd := exec.Command("go", "mod", "init", moduleName)
	modCmd.Dir = dirname
	if err := modCmd.Run(); err != nil {
		return fmt.Errorf(color.RedString("Unable to create go module\n"))
	}
	return nil
}

func tidyGoModule(dirname string) error {
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = dirname
	if err := tidyCmd.Run(); err != nil {
		return fmt.Errorf(color.RedString("Unable to tidy go module\n"))
	}
	return nil
}

func createGoFiles(dirname, moduleName, codeSnippet string) error {
	templateFile, err := fs.ReadFileString("tpl/main.tpl")
	if err != nil {
		return err
	}
	codeSnippet = utils.NormalizeSpaces(codeSnippet)
	functionName, argsTypes, returnType := utils.ExtractFunctionMetadata(codeSnippet)
	templateFile = utils.TemplateReplace(templateFile, map[string]string{
		"// INSERT PROBLEM": codeSnippet,
		"module_name":    moduleName,
		"RETURN_TYPE":    returnType,
		"FUNCTION_CALL": fmt.Sprintf("%s(%s)", functionName, strings.Join(lo.Map(argsTypes, func(typ string, i int) string {
			return fmt.Sprintf("args[%d].(%s)", i, typ)
		}), ",")),
	})

	// Create the main.go file
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", dirname))
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to create main.go file\n"))
	}
	mainFile.WriteString(templateFile)
	mainFile.Close()

	// Create the testutils directory
	err = os.MkdirAll(fmt.Sprintf("%s/testutils", dirname), os.ModePerm)
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to create testutils directory\n"))
	}

	// Create the testutils.go file
	testutilsFile, err := os.Create(fmt.Sprintf("%s/testutils/testutils.go", dirname))
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to create testutils.go file\n"))
	}
	testutilsFileContent, err := fs.ReadFileString("tpl/testutils/testutils.tpl")
	if err != nil {
		return fmt.Errorf(color.RedString("Unable to read testutils.tpl file\n"))
	}
	testutilsFile.WriteString(testutilsFileContent)
	testutilsFile.Close()
	return nil
}

func cleanDirectory(dirname string) {
	err := os.RemoveAll(dirname)
	if err != nil {
		fmt.Printf("%s Failed to delete %s", color.YellowString("WARNING:"), color.CyanString(dirname))
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <LEETCODE_ID>",
	Short: "Add a new leetcode problem",
	Long: `Add a new leetcode problem to the current directory. 

The leetcode id can be found in the URL of the problem page. For example, for the URL https://leetcode.com/problems/two-sum/, the leetcode ID is "two-sum".

The command will create a new directory with the leetcode ID and generate a template testing file for the problem. It will also include a docs.md file containing the problem description and section for solution notes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Leetcode ID is required")
		}
		start := time.Now()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Fetching leetcode problem data"
		leetcodeID := args[0]

		problems, err := utils.ReadYamlProblems("problems.yaml")
		if err != nil {
			return fmt.Errorf(color.RedString("Unable to read problems.yaml\n"))
		}

		if problem, exists := problems[leetcodeID]; exists {
			return fmt.Errorf(color.RedString("Problem \"%s\" already exists in directory \"%s\"", leetcodeID, problem.Directory))
		}

		attempts := 0
		s.Start()
		data, err := getLeetcodeProblemData(leetcodeID)
		if err != nil {
			// Retry fetching the problem data
			s.Stop()
			attempts++
			for attempts < 3 {
				fmt.Printf("Retrying leetcode problem fetch (%d/3)\n", attempts)
				s.Start()
				data, err = getLeetcodeProblemData(leetcodeID)
				if err == nil {
					s.Stop()
					break
				}
				attempts++
			}
			if err != nil {
				return fmt.Errorf("Unable to fetch leetcode problem")
			}
		}
		s.Stop()
		fmt.Printf("%s Fetched leetcode problem data\n", color.GreenString("\u2713"))

		converter := md.NewConverter("", true, nil)
		markdown, err := converter.ConvertString(data.DescriptionHtml)
		if err != nil {
			return fmt.Errorf(color.RedString("Unable to convert leetcode problem description to markdown\n"))
		}

		// Create the directory
		dirname := fmt.Sprintf("%s - %s", lo.Capitalize(data.Difficulty), data.ProblemTitle)
		fmt.Printf("Creating new problem directory \"%s\"\n", color.CyanString(dirname))
		if err = createProblemDirectory(dirname); err != nil {
			return err
		}

		// Create the docs.md file
		if err = createDocsFile(dirname, data.ProblemTitle, markdown); err != nil {
			cleanDirectory(dirname)
			return err
		}
		fmt.Printf("Created %s\n", color.CyanString("%s/docs.md", dirname))

		// Create the go module
		if err = createGoModule(dirname, leetcodeID); err != nil {
			cleanDirectory(dirname)
			return err
		}
		fmt.Printf("Created %s\n", color.CyanString("%s/go.mod", dirname))

		// Create the go files
		if err = createGoFiles(dirname, leetcodeID, data.CodeSnippet); err != nil {
			cleanDirectory(dirname)
			return err
		}
		fmt.Printf("Created %s\n", color.CyanString("%s/main.go", dirname))
		fmt.Printf("Created %s\n", color.CyanString("%s/testutils/testutils.go", dirname))

		// Tidy the go module
		if err = tidyGoModule(dirname); err != nil {
			cleanDirectory(dirname)
			return err
		}
		fmt.Printf("Created %s\n", color.CyanString("%s/go.sum", dirname))

		// Add to the yaml file
		problems[leetcodeID] = utils.YamlProblem{
			Difficulty: data.Difficulty,
			Directory: dirname,
			Published: false,
			Date: nil,
			Tags: nil,
		}
		err = utils.WriteYamlProblems("problems.yaml", problems)
		if err != nil {
			cleanDirectory(dirname)
			return fmt.Errorf(color.RedString("Unable to write problems.yaml\n"))
		}
		fmt.Printf("%s Updated %s\n",color.GreenString("\u2713"), color.CyanString("problems.yaml"))

		duration := time.Since(start)

		fmt.Printf("✨ Done in %s\n", utils.FormatDuration(duration))
		fmt.Printf("View the problem description and write solution notes in %s\n", color.CyanString("%s/docs.md", dirname))
		fmt.Printf("Write and test your solution in %s\n", color.CyanString("%s/main.go", dirname))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
