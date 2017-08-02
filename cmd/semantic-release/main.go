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
	"github.com/jbcpollak/strcase"
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
	flow := flag.Bool("flow", false, "follow branch naming conventions")
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

	curCommitInfo, err := condition.GetCurCommitInfo()
	exitIfError(err)
	logger.Println("found current branch: " + curCommitInfo.Branch)

	config := loadConfig()
	if config.MaintainedVersion != "" && curCommitInfo.Branch == defaultBranch {
		exitIfError(fmt.Errorf("maintained version not allowed on default branch"))
	}

	if config.MaintainedVersion != "" {
		logger.Println("found maintained version: " + config.MaintainedVersion)
		defaultBranch = "*"
	}

	prerelease := ""
	if *flow && config.MaintainedVersion == "" {

		switch curCommitInfo.Branch {
		// If branch is master -> no pre-latestRelease version
		case "master":
			prerelease = ""
			break
		// If branch is develop -> beta latestRelease
		case "develop":
			prerelease = "beta"
			break
		default:
			branchPath := strings.Split(curCommitInfo.Branch, "/")
			prerelease = branchPath[len(branchPath) - 1]
			prerelease = strcase.ToLowerCamel(prerelease)
		}
	}

	if prerelease != "" {
		logger.Println("Determined prerelease version: " + prerelease)
	}

	if !*noci {
		logger.Println("running CI condition...")
		exitIfError(condition.Travis(*token, defaultBranch, isPrivate))
	}

	logger.Println("getting latest release...")
	latestRelease, err := repo.GetLatestRelease(config.MaintainedVersion, prerelease)
	exitIfError(err)
	logger.Println("found version: " + latestRelease.Version.String())

	if strings.Contains(config.MaintainedVersion, "-") && latestRelease.Version.Prerelease() == "" {
		exitIfError(fmt.Errorf("no pre-release for this version possible"))
	}

	logger.Println("getting commits...")
	commits, err := repo.GetCommits(curCommitInfo.Branch)
	exitIfError(err)

	logger.Println("calculating new version...")
	newVer := semrel.GetNewVersion(commits, latestRelease, prerelease)
	if newVer == nil {
		exitIfError(errors.New("no change"))
	}
	logger.Println("new version: " + newVer.String())

	if *dry {
		exitIfError(errors.New("DRY RUN: no release was created"))
	}

	logger.Println("creating release...")
	exitIfError(repo.CreateRelease(commits, latestRelease, newVer, curCommitInfo.Branch))

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
