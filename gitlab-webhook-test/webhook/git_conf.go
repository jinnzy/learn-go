package webhook
// 解析git post过来的json
type Git_conf struct {
	Object_kind string `json:"object_kind"`
	Event_name string `json:"event_name"`
	Before string `json:"before"`
	After string `json:"after"`
	Ref string `json:"ref"`
	Checkout_sha string `json:"checkout_sha"`
	Message string `json:"message"`
	User_id int `json:"user_id"`
	User_name string `json:"user_name"`
	User_username string `json:"user_username"`
	User_email string `json:"user_email"`
	User_avatar string `json:"user_avatar"`
	Project_id int `json:"project_id"`
	Project ProjectStruct `json:"project"`
	Commits interface{} `json:"commits"`
	Total_commits_count int `json:"total_commits_count"`
	Repository interface{} `json:"repository"`
}
type ProjectStruct struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Web_url string `json:"web_url"`
	Avatar_url interface{} `json:avatar_url"`
	Git_ssh_url string `json:"git_ssh_url"`
	Git_http_url string `json:"git_http_url"`
	Namespace string `json:"namespace"`
	Visibility_level int `json:"visibility_level"`
	Path_with_namespace string `json:"path_with_namespace"`
	Default_branch string `json:"default_branch"`
	Ci_config_path interface{} `json:"ci_config_path"`
	Homepage string `json:"homepage"`
	Url string `json:"url"`
	Ssh_url string `json:"ssh_url"`
	Http_url string `json:"http_url"`
}