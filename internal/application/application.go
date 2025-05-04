package application

import (
	"fmt"
	"time"

	"github.com/tillpaid/gitlab-auto-mr/internal/command"
	"github.com/tillpaid/gitlab-auto-mr/internal/config"
	"github.com/tillpaid/gitlab-auto-mr/internal/git"
	"github.com/tillpaid/gitlab-auto-mr/internal/gitlab"
	"github.com/tillpaid/gitlab-auto-mr/internal/jira"
	"github.com/tillpaid/gitlab-auto-mr/internal/utils/stringutil"
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
	originInfo, err := a.fetchGitOrigin()
	if err != nil {
		return err
	}

	branch, err := a.gitService.GetCurrentBranch()
	if err != nil {
		return fmt.Errorf("error during getting current branch: %v", err)
	}

	issue, err := a.extractIssue(branch)
	if err != nil {
		return err
	}

	user, err := a.getGitlabUser()
	if err != nil {
		return err
	}

	mergeRequest, err := a.createMergeRequest(user.Id, originInfo.Path, branch, issue)
	if err != nil {
		return err
	}

	a.printMergeRequestSummary(mergeRequest)
	return nil
}

func (a *App) fetchGitOrigin() (*git.OriginInfo, error) {
	start := time.Now()
	fmt.Print("🔍 Fetching git origin...")

	originInfo, err := a.gitService.GetOriginInfo()
	if err != nil {
		return nil, fmt.Errorf("error during getting origin: %v", err)
	}

	if originInfo.Host != a.cfg.GitConstraints.ExpectedOriginHost {
		return nil, fmt.Errorf("origin host %q does not match expected host %q", originInfo.Host,
			a.cfg.GitConstraints.ExpectedOriginHost)
	}

	fmt.Printf(" (⏱ %.2fs)\n", time.Since(start).Seconds())
	fmt.Printf("✅ Connected to %s\n", originInfo.Host)
	fmt.Printf("📁 Project: %s\n", originInfo.Path)
	fmt.Println()
	return originInfo, nil
}

func (a *App) extractIssue(branch string) (*jira.Issue, error) {
	start := time.Now()
	fmt.Print("🔗 Extracting Jira issue from branch...")

	issueKey, err := jira.ExtractIssueKey(branch)
	if err != nil {
		return nil, err
	}

	issue, err := a.jiraClient.GetIssueByKey(issueKey)
	if err != nil {
		return nil, fmt.Errorf("error from Jira: %v\n", err)
	}

	fmt.Printf(" (⏱ %.2fs)\n", time.Since(start).Seconds())
	fmt.Println("✅ Found issue:", issue.Key)
	fmt.Println("📝 Title:", stringutil.TruncateWords(issue.Fields.Summary, 30))
	fmt.Println()
	return issue, nil
}

func (a *App) getGitlabUser() (*gitlab.User, error) {
	start := time.Now()
	fmt.Print("👤 Getting GitLab user info...")
	user, err := a.gitlabClient.GetCurrentUser()
	if err != nil {
		return nil, err
	}
	fmt.Printf(" (⏱ %.2fs)\n", time.Since(start).Seconds())
	fmt.Println("✅ Logged in as:", user.Name)
	fmt.Println()
	return user, nil
}

func (a *App) createMergeRequest(assigneeId int, projectPath, branch string, issue *jira.Issue) (*gitlab.MergeRequest, error) {
	start := time.Now()
	fmt.Print("🚀 Creating merge request...")
	title := fmt.Sprintf("Draft: %s - %s", issue.Key, issue.Fields.Summary)
	description := fmt.Sprintf("## %s\n## %s\n\n## Test plan\n\nCloses %s", issue.Key, issue.Fields.Summary, issue.Key)
	mr, err := a.gitlabClient.CreateMergeRequest(assigneeId, projectPath, branch, title, description)
	if err != nil {
		return nil, err
	}
	fmt.Printf(" (⏱ %.2fs)\n", time.Since(start).Seconds())
	fmt.Println("✅ Merge request created successfully!")
	fmt.Println()
	return mr, nil
}

func (a *App) printMergeRequestSummary(mr *gitlab.MergeRequest) {
	fmt.Printf("🔗 %s\n", mr.WebUrl)
	fmt.Printf("📌 Title: %s\n", stringutil.TruncateWords(mr.Title, 30))
	fmt.Printf("👤 Author: %s\n", mr.Author.Name)
	fmt.Printf("🌿 Branch: %s → %s\n", mr.SourceBranch, mr.TargetBranch)
	fmt.Println("👉 You can now open the MR, assign reviewers, and continue in GitLab.")
}
