package models

type ObjectAttributes struct {
	Ref            string   `json:"ref"`
	Sha            string   `json:"sha"`
	BeforeSha      string   `json:"before_sha"`
	ID             float64  `json:"id"`
	DetailedStatus string   `json:"detailed_status"`
	Stages         []string `json:"stages"`
	CreatedAt      string   `json:"created_at"`
	FinishedAt     string   `json:"finished_at"`
	Duration       float64  `json:"duration"`
	Variables      []any    `json:"variables"`
	Tag            bool     `json:"tag"`
	Source         string   `json:"push"`
	Status         string   `json:"status"`
	QueuedDuration float64  `json:"queue_duration"`
}
type User struct {
	ID       float64 `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Avatar   string  `json:"avatar_url"`
	Email    string  `json:"email"`
}
type Project struct {
	ID                float64 `json:"id"`
	Name              string  `json:"name"`
	WebUrl            string  `json:"web_url"`
	GitSshUrl         string  `json:"git_ssh_url"`
	Namespace         string  `json:"namespace"`
	DefaultBranch     string  `json:"default_branch"`
	SshUrl            string  `json:"ssh_url"`
	HttpUrl           string  `json:"http_url"`
	Description       string  `json:"description"`
	Avatar            string  `json:"avatar_url"`
	PathWithNamespace string  `json:"path_with_namespace"`
	VisibilityLevel   float64 `json:"visibility_level"`
	Homepage          string  `json:"homepage"`
	Url               string  `json:"url"`
	GitHttpUrl        string  `json:"git_http_url"`
	CiConfigPath      any     `json:"ci_config_path"`
}
type GitLabEvent struct {
	ObjectKind   string  `json:"object_kind"`
	EventName    string  `json:"event_name"`
	UserId       float64 `json:"user_id"`
	UserName     string  `json:"user_name"`
	UserUsername string  `json:"user_username"`

	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	MergeRequest     string           `json:"merge_request"`
	User             User             `json:"user"`
	ProjectId        float64          `json:"project_id"`
	Project          Project          `json:"project"`
	Commit           map[string]any   `json:"commit"`
	Builds           []map[string]any `json:"builds"`
}
type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}
type Msg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}
