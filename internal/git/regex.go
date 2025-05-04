package git

import "regexp"

var gitURLPattern = regexp.MustCompile(`(?i)(?:(?:[\w\-]+@)?(?P<host1>[\w\.-]+)[:|/]|(?:(?:ssh|https?)://(?:[\w\-]+@)?)(?P<host2>[\w\.-]+)(?::\d+)?/)(?P<path>[\w\-/]+)(?:\.git)?$`)
