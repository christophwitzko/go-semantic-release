package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/semantic-release/go-semantic-release"
	"github.com/semantic-release/go-semantic-release/condition"
	"github.com/semantic-release/go-semantic-release/update"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var SRVERSION string

func errorHandler(logger *log.Logger) func(error) {
	return func(err error) {
		if err != nil {
			logger.Println(err)
			os.Exit(1)
		}
	}
}

type SemRelConfig struct {
	MaintainedVersion string `json:"maintainedVersion"`
}

func loadConfig() *SemRelConfig {
	f, err := os.OpenFile(".semrelrc", os.O_RDONLY, 0)
	if err != nil {
		return &SemRelConfig{}
	}
	src := &SemRelConfig{}
	json.NewDecoder(f).Decode(src)
	f.Close()
	return src
}

func main() {
	token := flag.String("token", os.Getenv("GITHUB_TOKEN"), "github token")
	slug := flag.String("slug", os.Getenv("TRAVIS_REPO_SLUG"), "slug of the repository")
	ghr := flag.Bool("ghr", false, "create a .ghr file with the parameters for ghr")
	noci := flag.Bool("noci", false, "run semantic-release locally")
	dry := flag.Bool("dry", false, "do not create release")
	vFile := flag.Bool("vf", false, "create a .version file")
	showVersion := flag.Bool("version", false, "outputs the semantic-release version")
	updateFile := flag.String("update", "", "updates the version of a certain file")
	flag.Parse()

	if *showVersion {
		fmt.Printf("semantic-release v%s", SRVERSION)
		return
	}

	logger := log.New(os.Stderr, "[semantic-release]: ", 0)
	exitIfError := errorHandler(logger)

	if val, ok := os.LookupEnv("GH_TOKEN"); *token == "" && ok {
		*token = val
	}

	if *token == "" {
		exitIfError(errors.New("github token missing"))
	}
	if *slug == "" {
		exitIfError(errors.New("slug missing"))
	}

	repo, err := semrel.NewRepository(context.TODO(), *slug, *token)
	exitIfError(err)

	logger.Println("getting default branch...")
	defaultBranch, isPrivate, err := repo.GetInfo()
	exitIfError(err)
	logger.Println("found default branch: " + defaultBranch)

	currentBranch := condition.GetCurrentBranch()
	if currentBranch == "" {
		exitIfError(fmt.Errorf("current branch not found"))
	}
	logger.Println("found current branch: " + currentBranch)

	config := loadConfig()
	if config.MaintainedVersion != "" && currentBranch == defaultBranch {
		exitIfError(fmt.Errorf("maintained version not allowed on default branch"))
	}

	if config.MaintainedVersion != "" {
		logger.Println("found maintained version: " + config.MaintainedVersion)
		defaultBranch = "*"
	}

	if !*noci {
		logger.Println("running CI condition...")
		exitIfError(condition.Travis(*token, defaultBranch, isPrivate))
	}

	logger.Println("getting latest release...")
	release, err := repo.GetLatestRelease(config.MaintainedVersion)
	exitIfError(err)
	logger.Println("found version: " + release.Version.String())

	if strings.Contains(config.MaintainedVersion, "-") && release.Version.Prerelease() == "" {
		exitIfError(fmt.Errorf("no pre-release for this version possible"))
	}

	logger.Println("getting commits...")
	commits, err := repo.GetCommits(currentBranch)
	exitIfError(err)

	logger.Println("calculating new version...")
	newVer := semrel.GetNewVersion(commits, release)
	if newVer == nil {
		exitIfError(errors.New("no change"))
	}
	logger.Println("new version: " + newVer.String())

	if *dry {
		exitIfError(errors.New("DRY RUN: no release was created"))
	}

	logger.Println("creating release...")
	exitIfError(repo.CreateRelease(commits, release, newVer, currentBranch))

	if *ghr {
		exitIfError(ioutil.WriteFile(".ghr", []byte(fmt.Sprintf("-u %s -r %s v%s", repo.Owner, repo.Repo, newVer.String())), 0644))
	}

	if *vFile {
		exitIfError(ioutil.WriteFile(".version", []byte(newVer.String()), 0644))
	}

	if *updateFile != "" {
		exitIfError(update.Apply(*updateFile, newVer.String()))
	}

	logger.Println("done.")
}
