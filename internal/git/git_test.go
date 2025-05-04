package git

import (
	"errors"
	"testing"
)

func TestService_GetOriginInfo(t *testing.T) {
	tests := []struct {
		name        string
		mockOutput  string
		mockError   error
		expectHost  string
		expectPath  string
		expectError bool
	}{
		{
			name:       "Classic git@ with single-level path",
			mockOutput: "git@github.com:org/project.git",
			expectHost: "github.com",
			expectPath: "org/project",
		},
		{
			name:       "git@ with deeply nested path",
			mockOutput: "git@git.example.com:org/team/component/module/project.git",
			expectHost: "git.example.com",
			expectPath: "org/team/component/module/project",
		},
		{
			name:       "ssh:// with user and port, nested path",
			mockOutput: "ssh://dev@git.enterprise.io:2222/org/team/project.git",
			expectHost: "git.enterprise.io",
			expectPath: "org/team/project",
		},
		{
			name:       "ssh:// with no user, nested path",
			mockOutput: "ssh://git.enterprise.io:2222/org/team/subsystem/project.git",
			expectHost: "git.enterprise.io",
			expectPath: "org/team/subsystem/project",
		},
		{
			name:       "https:// with user and port",
			mockOutput: "https://user@git.enterprise.io:8443/org/infra/api.git",
			expectHost: "git.enterprise.io",
			expectPath: "org/infra/api",
		},
		{
			name:       "https:// with no user and no port",
			mockOutput: "https://git.enterprise.io/org/infra/api.git",
			expectHost: "git.enterprise.io",
			expectPath: "org/infra/api",
		},
		{
			name:       "HTTPS with long nested path and no .git",
			mockOutput: "https://git.enterprise.io/org/infra/platform/deploy/api",
			expectHost: "git.enterprise.io",
			expectPath: "org/infra/platform/deploy/api",
		},
		{
			name:       "Mixed-case host and path",
			mockOutput: "git@GitHub.COM:Org/Repo-Name.git",
			expectHost: "GitHub.COM",
			expectPath: "Org/Repo-Name",
		},
		{
			name:       "Extra long nested subgroup",
			mockOutput: "git@gitlab.example.net:division/alpha/beta/gamma/delta/project.git",
			expectHost: "gitlab.example.net",
			expectPath: "division/alpha/beta/gamma/delta/project",
		},
		{
			name:       "URL with no .git extension",
			mockOutput: "git@gitlab.example.net:team/repo",
			expectHost: "gitlab.example.net",
			expectPath: "team/repo",
		},
		{
			name:        "invalid origin",
			mockOutput:  "not-a-valid-origin",
			expectError: true,
		},
		{
			name:        "command fails",
			mockError:   errors.New("git error"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewService(func(name string, args ...string) (string, error) {
				return tc.mockOutput, tc.mockError
			})

			result, err := svc.GetOriginInfo()
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result.Host != tc.expectHost {
				t.Errorf("expected host %q, got %q", tc.expectHost, result.Host)
			}

			if result.Path != tc.expectPath {
				t.Errorf("expected path %q, got %q", tc.expectPath, result.Path)
			}
		})
	}
}

func TestService_GetCurrentBranch(t *testing.T) {
	tests := []struct {
		name       string
		mockOutput string
		mockError  error
		want       string
		wantErr    bool
	}{
		{
			name:       "Main",
			mockOutput: "Main",
			want:       "Main",
		},
		{
			name:       "PROJ-1231",
			mockOutput: "PROJ-1231",
			want:       "PROJ-1231",
		},
		{
			name:      "Error",
			mockError: errors.New("git failed"),
			wantErr:   true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewService(func(name string, args ...string) (string, error) {
				return tc.mockOutput, tc.mockError
			})

			result, err := svc.GetCurrentBranch()
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result != tc.mockOutput {
				t.Errorf("expected branch %q, got %q", tc.mockOutput, result)
			}
		})
	}
}
