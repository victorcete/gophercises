package urlshort

type redirect struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func buildRedirectMap(redirects []redirect) map[string]string {
	redirectMap := make(map[string]string)
	for _, val := range redirects {
		redirectMap[val.Path] = val.URL
	}
	return redirectMap
}
