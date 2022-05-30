package plugin

import (
	"testing"

	"github.com/Masterminds/semver/v3"

	"github.com/stretchr/testify/require"
)

func parseConstraint(c string) *semver.Constraints {
	constraint, _ := semver.NewConstraint(c)
	return constraint
}

func TestGetPluginInfo(t *testing.T) {
	testCases := []struct {
		t       string
		input   string
		results *PluginInfo
	}{
		{"provider", "git", &PluginInfo{
			Type:                "provider",
			Name:                "git",
			NormalizedName:      "provider-git",
			ShortNormalizedName: "provider-git",
			Resolver:            "default",
		}},
		{"provider", "github:myorg/myplugin", &PluginInfo{
			Type:                "provider",
			Name:                "myplugin",
			NormalizedName:      "provider-github-myorg-myplugin",
			ShortNormalizedName: "provider-myplugin",
			Resolver:            "github",
		}},
		{"ci_condition", "github:myorg/myplugin", &PluginInfo{
			Type:                "ci_condition",
			Name:                "myplugin",
			NormalizedName:      "condition-github-myorg-myplugin",
			ShortNormalizedName: "condition-myplugin",
			Resolver:            "github",
		}},
		{"provider", "github:myorg/myplugin@^1.0.0", &PluginInfo{
			Type:                "provider",
			Name:                "myplugin",
			NormalizedName:      "provider-github-myorg-myplugin",
			ShortNormalizedName: "provider-myplugin",
			Resolver:            "github",
			Constraint:          parseConstraint("^1.0.0"),
		}},
		{"provider", "git@2.0.0", &PluginInfo{
			Type:                "provider",
			Name:                "git",
			NormalizedName:      "provider-git",
			ShortNormalizedName: "provider-git",
			Resolver:            "default",
			Constraint:          parseConstraint("2.0.0"),
		}},
		{"hooks", "registry:logger", &PluginInfo{
			Type:                "hooks",
			Name:                "logger",
			NormalizedName:      "hooks-registry-logger",
			ShortNormalizedName: "hooks-logger",
			Resolver:            "registry",
		}},
	}
	for _, testCase := range testCases {
		results, err := GetPluginInfo(testCase.t, testCase.input)
		require.NoError(t, err)
		require.Equal(t, testCase.results, results)
	}
}
