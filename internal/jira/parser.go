package jira

import (
	"fmt"
	"regexp"
	"strings"
)

var issueKeyPattern = regexp.MustCompile(`(?i)[A-Z][A-Z0-9]+-\d+`)

func ExtractIssueKey(branchName string) (string, error) {
	match := issueKeyPattern.FindString(branchName)
	if match == "" {
		return "", fmt.Errorf("no jira issue key found in: %q", branchName)
	}

	return strings.ToUpper(match), nil
}
