package application

import (
	"fmt"

	"github.com/tillpaid/gitlab-auto-mr/internal/command"
	"github.com/tillpaid/gitlab-auto-mr/internal/config"
	"github.com/tillpaid/gitlab-auto-mr/internal/git"
	"github.com/tillpaid/gitlab-auto-mr/internal/gitlab"
	"github.com/tillpaid/gitlab-auto-mr/internal/jira"
)

type App struct {
	cfg          *config.Config
	gitService   *git.Service
	jiraClient   *jira.Client
	gitlabClient *gitlab.Client
}

func New(cfg *config.Config) *App {
	return &App{
		cfg:          cfg,
		gitService:   git.NewService(command.RunAndTrim),
		jiraClient:   jira.NewClient(cfg.Jira.Url, cfg.Jira.Username, cfg.Jira.Password),
		gitlabClient: gitlab.NewClient(cfg.Gitlab.Url, cfg.Gitlab.Token),
	}
}

func (a *App) Run() error {
	a.processOrigin()
	fmt.Println()

	a.processBranch()
	fmt.Println()

	a.processGitlab()

	return nil
}

func (a *App) processOrigin() {
	originInfo, err := a.gitService.GetOriginInfo()
	if err != nil {
		fmt.Println("Got error during getting origin:", err.Error())
		return
	}

	fmt.Printf("Origin host: %q\n", originInfo.Host)
	fmt.Printf("Origin path: %q\n", originInfo.Path)

	if originInfo.Host != a.cfg.GitConstraints.ExpectedOriginHost {
		fmt.Printf("\tOrigin host %s does not match expected host %s\n", originInfo.Host,
			a.cfg.GitConstraints.ExpectedOriginHost)
	}
}

func (a *App) processBranch() {
	branch, err := a.gitService.GetCurrentBranch()
	if err != nil {
		fmt.Printf("Got error during getting current branch: %q\n", err.Error())
		return
	}

	fmt.Printf("Branch: %q\n", branch)

	issueKey, err := jira.ExtractIssueKey(branch)
	if err != nil {
		fmt.Printf("\tInvalid branch name: %v\n", err)
		return
	}

	fmt.Printf("Issue key: %q\n", issueKey)

	issue, err := a.jiraClient.GetIssueByKey(issueKey)
	if err != nil {
		fmt.Printf("\tError from Jira: %v\n", err)
		return
	}

	fmt.Printf("Issue summary from Jira: %q\n", issue.Fields.Summary)
}

func (a *App) processGitlab() {
	user, err := a.gitlabClient.GetCurrentUser()
	if err != nil {
		fmt.Printf("Got error during getting current user: %v\n", err.Error())
		return
	}

	fmt.Printf("Gitlab user ID: %d\n", user.Id)
	fmt.Printf("Gitlab user username: %q\n", user.Username)
	fmt.Printf("Gitlab user name: %q\n", user.Name)
}
