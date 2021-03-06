package git

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestURLParser_ParseURL(t *testing.T) {
	c := make(SSHConfig)
	c["github.com"] = "ssh.github.com"
	c["git.company.com"] = "ssh.git.company.com"

	p := &URLParser{c}

	u, err := p.Parse("https://github.com/octokit/go-octokit.git")
	assert.Equal(t, nil, err)
	assert.Equal(t, "github.com", u.Host)
	assert.Equal(t, "https", u.Scheme)
	assert.Equal(t, "/octokit/go-octokit.git", u.Path)

	u, err = p.Parse("git://github.com/octokit/go-octokit.git")
	assert.Equal(t, nil, err)
	assert.Equal(t, "github.com", u.Host)
	assert.Equal(t, "git", u.Scheme)
	assert.Equal(t, "/octokit/go-octokit.git", u.Path)

	u, err = p.Parse("git@github.com:lostisland/go-sawyer.git")
	assert.Equal(t, nil, err)
	assert.Equal(t, "github.com", u.Host)
	assert.Equal(t, "ssh", u.Scheme)
	assert.Equal(t, "git", u.User.Username())
	assert.Equal(t, "/lostisland/go-sawyer.git", u.Path)

	u, err = p.Parse("https://git.company.com/octokit/go-octokit.git")
	assert.Equal(t, nil, err)
	assert.Equal(t, "ssh.git.company.com", u.Host)
	assert.Equal(t, "https", u.Scheme)
	assert.Equal(t, "/octokit/go-octokit.git", u.Path)
}
