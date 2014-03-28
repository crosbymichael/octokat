package octokat

import (
	"fmt"
)

type Label struct {
	URL   string `json:"url,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

func (c *Client) Labels(repo Repo) (labels []*Label, err error) {
	path := fmt.Sprintf("repos/%s/labels", repo)
	err = c.jsonGet(path, &Options{}, &labels)
	return
}
