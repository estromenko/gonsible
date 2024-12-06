package ssh

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/melbahja/goph"
)

func Execute(rawCommand string, host string) ([]byte, error) {
	command := strings.TrimSpace(rawCommand)

	if host == "localhost" || host == "127.0.0.1" {
		cmd := exec.Command("bash", "-c", command)

		return cmd.CombinedOutput()
	}

	uhd, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	privateKeyPath := path.Join(uhd, ".ssh/id_rsa")
	if _, err := os.Stat(privateKeyPath); errors.Is(err, os.ErrNotExist) {
		privateKeyPath = path.Join(uhd, ".ssh/id_ed25519")
		if _, err := os.Stat(privateKeyPath); errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	auth, err := goph.Key(privateKeyPath, "")
	if err != nil {
		return nil, err
	}

	client, err := goph.New("estromenko", host, auth)
	if err != nil {
		return nil, err
	}

	return client.Run(command)
}
