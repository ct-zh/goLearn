package chrome

import (
	"net/url"
	"os/exec"
)

type Chrome struct {
	Path   string
	Args   map[string]string
	Target *url.URL
}

func NewChrome(path string) (*Chrome, error) {
	shell := `which ` + path
	cmd := exec.Command("bash", "-c", shell)
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return &Chrome{Path: path}, nil
}

func (c Chrome) HtmlToPdf() (string, error) {
	s := c.Path + " "
	for name, value := range c.Args {
		s += name + value + " "
	}
	s += c.Target.String()

	cmd := exec.Command("bash", "-c", s)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
