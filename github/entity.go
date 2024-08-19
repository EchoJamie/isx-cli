package github

type Repository struct {
	Id          int    `json:"id"`
	NodeId      string `json:"node_id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Private     bool   `json:"private"`
	HtmlUrl     string `json:"html_url"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
	Url         string `json:"url"`
}

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	State  string `json:"state"`

	Labels []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
	} `json:"labels"`
}

type Branch struct {
	Name string `json:"name"`
}
