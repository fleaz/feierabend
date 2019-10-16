package gittools

import (
	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
)

func CheckRepo(repoPath string) (bool, error) {
	r, err := git.PlainOpen(repoPath)
	if err == nil && r != nil {
		// We found valid git repo

		wt, _ := r.Worktree()
		s, _ := wt.Status()
		if s.IsClean() == false {
			color.Red("%q is in a dirty state", repoPath)

			return true, nil
		}
		return false, nil
	}

	return false, err
}
