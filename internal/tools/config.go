package config

import (
	"fmt"

	"github.com/xUnholy/go-proxy/pkg/execute"
)

func EnableGITProxyConfiguration(port string) error {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "http.proxy", port}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "https.proxy", port}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to enable git http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to enable git https command %q", err)
	}
	return nil
}

func DisableGITProxyConfiguration() error {
	http := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "http.proxy"}}
	https := execute.Command{Cmd: "git", Args: []string{"config", "--global", "--unset", "https.proxy"}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to disable git http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to disable git https command %q", err)
	}
	return nil
}

func EnableNPMProxyConfiguration(port string) error {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "set", "proxy", port}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "set", "https-proxy", port}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to enable npm http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to enable npm https command %q", err)
	}
	return nil
}

func DisableNPMProxyConfiguration() error {
	http := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "proxy"}}
	https := execute.Command{Cmd: "npm", Args: []string{"config", "delete", "https-proxy"}}
	_, err := execute.RunCommand(http)
	if err != nil {
		return fmt.Errorf("failed to disable npm http command %q", err)
	}
	_, err = execute.RunCommand(https)
	if err != nil {
		return fmt.Errorf("failed to disable npm https command %q", err)
	}
	return nil
}
