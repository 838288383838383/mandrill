package github

import (
	"os/exec"
	"strings"
)

type GhClient struct {
	Authenticated bool
	Username      string
	Token         string
}

func DetectGhCli() (*GhClient, error) {
	client := &GhClient{}

	statusCmd := exec.Command("gh", "auth", "status")
	output, err := statusCmd.CombinedOutput()
	if err != nil {
		return client, nil
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Logged in as") {
			client.Authenticated = true
			parts := strings.SplitN(line, "as ", 2)
			if len(parts) == 2 {
				client.Username = strings.TrimSpace(parts[1])
			}
		}
	}

	if client.Authenticated {
		tokenCmd := exec.Command("gh", "auth", "token")
		tokenOut, err := tokenCmd.Output()
		if err == nil {
			client.Token = strings.TrimSpace(string(tokenOut))
		}
	}

	return client, nil
}

func (c *GhClient) IsAvailable() bool {
	_, err := exec.LookPath("gh")
	return err == nil
}
