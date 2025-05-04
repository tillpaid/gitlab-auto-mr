package jira

import "testing"

func TestExtractIssueKey(t *testing.T) {
	tests := []struct {
		branch  string
		want    string
		wantErr bool
	}{
		{branch: "ACME-1", want: "ACME-1"},
		{branch: "PROJ-123", want: "PROJ-123"},
		{branch: "TST999-42", want: "TST999-42"},
		{branch: "FOO2-8888", want: "FOO2-8888"},
		{branch: "feature/ACME-123", want: "ACME-123"},
		{branch: "hotfix/PROJ-5-something", want: "PROJ-5"},
		{branch: "bugfix/TST999-7_fix", want: "TST999-7"},
		{branch: "release/FOO2-8888-version-bump", want: "FOO2-8888"},
		{branch: "feature/some-thing-ACME-123", want: "ACME-123"},
		{branch: "ACME-123-some-desc", want: "ACME-123"},
		{branch: "dev/ACME-1234-xyz-567", want: "ACME-1234"},
		{branch: "task/anything-ACME-123-rest", want: "ACME-123"},
		{branch: "abc-abc-TST999-99-final", want: "TST999-99"},
		{branch: "main", wantErr: true},
		{branch: "develop", wantErr: true},
		{branch: "feature/no-key-here", wantErr: true},
		{branch: "hotfix/something_else", wantErr: true},
		{branch: "acme-123", want: "ACME-123"},
		{branch: "aCmE-123", want: "ACME-123"},
		{branch: "ACME 123", wantErr: true},
		{branch: "PROJ123", wantErr: true},
		{branch: "FOO--123", wantErr: true},
		{branch: "FOO/123", wantErr: true},
		{branch: "ACME-1-FOO2-2", want: "ACME-1"},
		{branch: "ACME-1/FOO2-2", want: "ACME-1"},
		{branch: "ACME-1_FOO2-2", want: "ACME-1"},
		{branch: "ACME-123 fix something", want: "ACME-123"},
		{branch: "feature/ACME-123/fix/bug", want: "ACME-123"},
	}
	for _, tt := range tests {
		t.Run(tt.branch, func(t *testing.T) {
			got, err := ExtractIssueKey(tt.branch)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("ExtractIssueKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
