package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"

	"github.com/fleaz/feierabend/cache"
	"github.com/fleaz/feierabend/gittools"
)

func getTraversalHandler(repoCache cache.Cache) func(path string, f os.FileInfo, err error) error {

	return func(path string, f os.FileInfo, err error) error {

		if !f.IsDir() {
			return filepath.SkipDir
		}

		fmt.Printf("Checking %v", path)

		gittools.CheckRepo(path)

		if repoCache != nil {
			repoCache.Write(path)
		}

		return nil
	}
}

func ensureConfigurationDirectory() string {
	configPath := configdir.LocalConfig("feierabend")
	err := configdir.MakePath(configPath)

	if err != nil {
		panic(err)
	}

	return configPath
}

func main() {
	var workspace = flag.String("workspace", ".", "Path to your GIT repositories")
	var useCache = flag.Bool("use-cache", false, "Uses path cache")

	flag.Parse()

	// Creating configuration directory
	configPath := ensureConfigurationDirectory()
	gitRepoFileCache := cache.NewGitRepoFileCache(configPath, false)

	traversalHandler := getTraversalHandler(gitRepoFileCache)

	if *useCache {
		gitRepos := gitRepoFileCache.ReadAll()

		for _, gitRepo := range gitRepos {
			fmt.Printf("Checking %v\n", gitRepo)
			gittools.CheckRepo(gitRepo)
		}
	} else {
		err := filepath.Walk(*workspace, traversalHandler)
		if err != nil {
			fmt.Println(err)
		}
	}

}
