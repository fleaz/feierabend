package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
)

func handler(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		r, err := git.PlainOpen(path)
		if err == nil && r != nil {
			// We found valid git repo
			wt, _ := r.Worktree()
			s, _ := wt.Status()
			if s.IsClean() == false {
				color.Red("%q is in a dirty state", path)
			}
			return filepath.SkipDir
		}
	}

	return nil
}

func main() {
	var workspace = flag.String("workspace", ".", "Path to your GIT repositories")
	flag.Parse()

	err := filepath.Walk(*workspace, handler)
	if err != nil {
		fmt.Println(err)
	}

}
