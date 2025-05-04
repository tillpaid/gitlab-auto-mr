package gitlab

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type MergeRequest struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	State        string `json:"state"`
	TargetBranch string `json:"target_branch"`
	SourceBranch string `json:"source_branch"`
	Author       struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
	WebUrl string `json:"web_url"`
}
