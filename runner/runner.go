package runner

import (
	"fmt"
	"os/exec"
)

// External is interface for rule external processes
type External interface {
	Start(name, args string) (string, error)
	Stop(name string) error
	Restart(name string) (string, error)
	Check(name string) (string, error)
	GetStd(name string, stdType int) (string, error)
}

// New return new Config
func New() Config {
	return Config{}
}

type Config struct {
}

// Start running new process
func (c *Config) Start(name, args string) (string, error) {
	cmd := exec.Command(name, args)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Config) Stop(name string) error {
	cmd := exec.Command("pkill", name)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Restart(name, args string) (string, error) {
	cmd := exec.Command(name, args)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Config) Check(name string) (string, error) {
	cmd := exec.Command("pgrep", name)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Config) GetStd(id string, stdType int) (string, error) {
	cmd := exec.Command("tail", fmt.Sprintf("/proc/%s/fd/%d", id, stdType))
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
