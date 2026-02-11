package options

type ResponseGithubApi struct {
	Body     string      `json:"body"`
	Tag_name string      `json:"tag_name"`
	Assets   []AssetsUrl `json:"assets"`
}

type AssetsUrl struct {
	Url                string `json:"url"`
	Name               string `json:"name"`
	BrowserDownloadUrl string `json:"browser_download_url"`
}
