package utils

import (
	"regexp"
	"strings"
)

func ExtractFunctionMetadata(snippet string) (string, []string, string) {
	functionNamePart := snippet[:strings.Index(snippet, "(")]
	functionName := strings.TrimSpace(strings.TrimPrefix(functionNamePart, "func "))
	argsPart := snippet[strings.Index(snippet, "(")+1 : strings.Index(snippet, ")")]
	returnPart := strings.TrimSpace(snippet[strings.Index(snippet, ")")+1:])
	returnType := ""
	if strings.HasPrefix(returnPart, "(") {
		// Handle multiple return values (optional extension)
		end := strings.Index(returnPart, ")")
		returnType = returnPart[1:end]
	} else {
		// Single return type
		re := regexp.MustCompile(`^\s*([^\s{]+)`)
		matches := re.FindStringSubmatch(returnPart)
		if len(matches) > 1 {
			returnType = matches[1]
		}
	}
	var argsTypes []string
	for _, arg := range strings.Split(argsPart, ",") {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}
		parts := strings.Fields(arg)
		if len(parts) == 2 {
			argsTypes = append(argsTypes, parts[1])
		}
	}

	return functionName, argsTypes, returnType
}
