package gobotsdottxt

import "net/http"

type Robots struct {
	URL        string
	Groups     []Group
	SiteMap    []string
	CrawlDelay int
}

type Group struct {
	UserAgent string
	Allow     []string
	Disallow  []string
}

func NewRobots(url string) (Robots, error) {
	robots := Robots{URL: url}
	resp, err := http.Get(url + "/robots.txt")
	if err != nil {
		return robots, err
	}

	//Pretty much no robots.txt
	if resp.StatusCode != 200 {
		return robots, nil
	}

	return robots, nil
}
