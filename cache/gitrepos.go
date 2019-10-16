package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// GitRepoFileCache contains all information about the Git repository cache.
type GitRepoFileCache struct {
	// Contains path to configuration folder
	configPath string
	// Contains paths to found repositories
	repoPaths []string

	writeFileHandler *os.File
}

// NewGitRepoFileCache creates a new GitRepoCache
func NewGitRepoFileCache(configPath string, regenerateCache bool) *GitRepoFileCache {

	configFilePath := path.Join(configPath, "gitrepos.txt")

	fmt.Println("Using configuration file at " + configFilePath)

	repoPathsByte, err := ioutil.ReadFile(configFilePath)

	var repoPaths []string

	if err == nil {
		repoPaths = strings.Split(string(repoPathsByte), "\n")
	}

	if regenerateCache {
		err = os.Remove(configFilePath)
		if err != nil {
			fmt.Printf("configFilePath could not be deleted: %v", err)
		}
	}
	writeFileHandler, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Printf("writeFileHandler could not be created: %v", err)
	}

	gitRepoCache := GitRepoFileCache{
		configPath:       configPath,
		repoPaths:        repoPaths,
		writeFileHandler: writeFileHandler,
	}

	return &gitRepoCache
}

func (g *GitRepoFileCache) Write(repoPath string) {

	doesExist := false

	for _, rPath := range g.repoPaths {
		if rPath == repoPath {
			doesExist = true
			break
		}
	}

	if doesExist {
		fmt.Printf("%v does already exists in cache.\n", repoPath)
		return
	}

	fmt.Printf("%v will be added in cache.\n", repoPath)

	g.repoPaths = append(g.repoPaths, repoPath)

	if g.writeFileHandler != nil {
		g.writeFileHandler.WriteString(repoPath + "\n")
	}

}

func (g *GitRepoFileCache) ReadAll() []string {
	return g.repoPaths
}
