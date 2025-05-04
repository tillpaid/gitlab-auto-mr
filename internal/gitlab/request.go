package gitlab

type CreateMergeRequestRequest struct {
	SourceBranch       string `json:"source_branch"`
	TargetBranch       string `json:"target_branch"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	AssigneeId         int    `json:"assignee_id"`
	Squash             bool   `json:"squash"`
	RemoveSourceBranch bool   `json:"remove_source_branch"`
}
