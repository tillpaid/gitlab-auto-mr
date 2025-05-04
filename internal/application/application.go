package application

import (
	"fmt"

	"github.com/tillpaid/gitlab-auto-mr/internal/command"
	"github.com/tillpaid/gitlab-auto-mr/internal/config"
	"github.com/tillpaid/gitlab-auto-mr/internal/git"
	"github.com/tillpaid/gitlab-auto-mr/internal/jira"
)

type App struct {
	cfg        *config.Config
	gitService *git.Service
}

func New(cfg *config.Config) *App {
	return &App{
		cfg:        cfg,
		gitService: git.NewService(command.RunAndTrim),
	}
}

func (a *App) Run() error {
	a.processOrigin()
	a.processBranch()

	return nil
}

func (a *App) processOrigin() {
	originInfo, err := a.gitService.GetOriginInfo()
	if err != nil {
		fmt.Println("Got error during getting origin:", err.Error())
		return
	}

	fmt.Println("Origin host:", originInfo.Host)
	fmt.Println("Origin path:", originInfo.Path)

	if originInfo.Host != a.cfg.GitConstraints.ExpectedOriginHost {
		fmt.Printf("\tOrigin host %s does not match expected host %s\n", originInfo.Host,
			a.cfg.GitConstraints.ExpectedOriginHost)
	}
}

func (a *App) processBranch() {
	branch, err := a.gitService.GetCurrentBranch()
	if err != nil {
		fmt.Println("Got error during getting current branch:", err.Error())
		return
	}

	fmt.Println("Branch:", branch)

	issueKey, err := jira.ExtractIssueKey(branch)
	if err != nil {
		fmt.Printf("\tInvalid branch name: %v\n", err)
		return
	}

	fmt.Println("Issue key:", issueKey)
}
