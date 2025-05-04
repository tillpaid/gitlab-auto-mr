

# gitlab-auto-mr

🚀 **CLI tool to automatically create GitLab Merge Requests based on your current Git branch and Jira issue.**

---

## 📦 Features

- 🔍 Detects Git remote and project
- 🔗 Extracts Jira issue key from branch name
- 📝 Fetches issue title from Jira
- 👤 Authenticates user via GitLab API
- 🧠 Auto-generates MR title and description
- ⏱ Pretty terminal output with duration timers

---

## 🧑‍💻 Installation

### Via Go

```bash
go install github.com/tillpaid/gitlab-auto-mr@latest
```

### Or clone and build manually

```bash
git clone https://github.com/tillpaid/gitlab-auto-mr
cd gitlab-auto-mr
go build -o gitlab-auto-mr ./cmd
```

---

## ⚙️ Configuration

By default, config is expected at:

```text
~/.config/gitlab-auto-mr/config.yml
```

### Example config file

```yaml
gitConstraints:
  expectedOriginHost: gitlab.example.com

jira:
  url: https://jira.example.com
  username: your.email@example.com
  password: your-jira-api-token

gitlab:
  url: https://gitlab.example.com
  token: your-gitlab-token
```

---

## 🚀 Usage

Just run the command inside your Git repository:

```bash
gitlab-auto-mr
```

It will:
1. Detect the remote GitLab project and current branch
2. Extract the Jira issue key from the branch name (e.g., `feature/ACME-123`)
3. Fetch the Jira issue summary
4. Create a draft merge request in GitLab

---

## 🧪 Demo Output

```bash
🔍 Fetching git origin... (⏱ 0.12s)
✅ Connected to gitlab.example.com
📁 Project: dir/project

🔗 Extracting Jira issue from branch... (⏱ 0.05s)
✅ Issue: ACME-123
📝 Title: Create command for customers

👤 Getting GitLab user info... (⏱ 0.10s)
✅ Logged in as: Oleksandr Maiboroda

🚀 Creating merge request... (⏱ 0.80s)
✅ Merge request created successfully!

🔗 https://gitlab.example.com/...
📌 Title: Draft: ACME-123 - Create command...
👤 Author: Oleksandr Maiboroda
🌿 Branch: feature/ACME-123 → main
👉 You can now open the MR, assign reviewers, and continue in GitLab.
```

---

## 🧰 Development

Build:

```bash
go build -o gitlab-auto-mr ./cmd
```

Run directly:

```bash
go run ./cmd
```

Test:

```bash
go test ./internal/...
```

---

## 🧩 Project Structure

```
internal/
├── application/  # Orchestration logic (Run)
├── git/          # Git remote and branch utils
├── gitlab/       # GitLab API client
├── jira/         # Jira API client
├── httpclient/   # Shared HTTP wrapper
├── config/       # Config loader & validation
├── utils/        # Helpers like TruncateWords
```

---

## 🪪 License

This project is licensed under the MIT License — see the [LICENSE](./LICENSE) file for details.

---

Made with 💙 by Oleksandr Maiboroda.
