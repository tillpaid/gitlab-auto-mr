package git

import (
	"errors"
)

type OriginInfo struct {
	Host string
	Path string
}

type Service struct {
	runCommand func(name string, arg ...string) (string, error)
}

func NewService(run func(name string, arg ...string) (string, error)) *Service {
	return &Service{runCommand: run}
}

func (s *Service) GetOriginInfo() (*OriginInfo, error) {
	raw, err := s.runCommand("git", "config", "--get", "remote.origin.url")
	if err != nil {
		return nil, err
	}

	match := gitURLPattern.FindStringSubmatch(raw)
	if match == nil {
		return nil, errors.New("failed to parse git origin")
	}

	var host string
	if match[1] != "" {
		host = match[1]
	} else {
		host = match[2]
	}

	path := match[3]

	return &OriginInfo{
		Host: host,
		Path: path,
	}, nil
}

func (s *Service) GetCurrentBranch() (string, error) {
	return s.runCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
}
