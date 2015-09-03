package gobotsdottxt

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

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

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return robots, err
	}

	var group Group
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		options := strings.SplitN(line, ":", 2)
		if len(options) == 1 {
			robots.Groups = append(robots.Groups, group)
			group = Group{}
			continue
		}

		command := strings.ToLower(strings.Trim(options[0], " "))
		value := strings.Trim(options[1], " ")
		switch command {
		case "user-agent":
			group.UserAgent = value
		case "disallow":
			group.Disallow = append(group.Disallow, value)
		case "allow":
			group.Allow = append(group.Allow, value)
		case "sitemap":
			robots.SiteMap = append(robots.SiteMap, value)
		case "crawl-delay":
			delay, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			robots.CrawlDelay = delay
		}
	}

	return robots, nil
}
