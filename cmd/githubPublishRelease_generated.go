package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type githubPublishReleaseOptions struct {
	AddClosedIssues       bool     `json:"addClosedIssues,omitempty"`
	AddDeltaToLastRelease bool     `json:"addDeltaToLastRelease,omitempty"`
	APIURL                string   `json:"apiUrl,omitempty"`
	AssetPath             string   `json:"assetPath,omitempty"`
	Commitish             string   `json:"commitish,omitempty"`
	ExcludeLabels         []string `json:"excludeLabels,omitempty"`
	Labels                []string `json:"labels,omitempty"`
	Owner                 string   `json:"owner,omitempty"`
	ReleaseBodyHeader     string   `json:"releaseBodyHeader,omitempty"`
	Repository            string   `json:"repository,omitempty"`
	ServerURL             string   `json:"serverUrl,omitempty"`
	Token                 string   `json:"token,omitempty"`
	UploadURL             string   `json:"uploadUrl,omitempty"`
	Version               string   `json:"version,omitempty"`
}

// GithubPublishReleaseCommand Publish a release in GitHub
func GithubPublishReleaseCommand() *cobra.Command {
	metadata := githubPublishReleaseMetadata()
	var stepConfig githubPublishReleaseOptions
	var startTime time.Time

	var createGithubPublishReleaseCmd = &cobra.Command{
		Use:   "githubPublishRelease",
		Short: "Publish a release in GitHub",
		Long: `This step creates a tag in your GitHub repository together with a release.
The release can be filled with text plus additional information like:

* Closed pull request since last release
* Closed issues since last release
* Link to delta information showing all commits since last release

The result looks like

![Example release](../images/githubRelease.png)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("githubPublishRelease")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "githubPublishRelease", &stepConfig, config.OpenPiperFile)
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, "githubPublishRelease")
			githubPublishRelease(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addGithubPublishReleaseFlags(createGithubPublishReleaseCmd, &stepConfig)
	return createGithubPublishReleaseCmd
}

func addGithubPublishReleaseFlags(cmd *cobra.Command, stepConfig *githubPublishReleaseOptions) {
	cmd.Flags().BoolVar(&stepConfig.AddClosedIssues, "addClosedIssues", false, "If set to `true`, closed issues and merged pull-requests since the last release will added below the `releaseBodyHeader`")
	cmd.Flags().BoolVar(&stepConfig.AddDeltaToLastRelease, "addDeltaToLastRelease", false, "If set to `true`, a link will be added to the relese information that brings up all commits since the last release.")
	cmd.Flags().StringVar(&stepConfig.APIURL, "apiUrl", "https://api.github.com", "Set the GitHub API url.")
	cmd.Flags().StringVar(&stepConfig.AssetPath, "assetPath", os.Getenv("PIPER_assetPath"), "Path to a release asset which should be uploaded to the list of release assets.")
	cmd.Flags().StringVar(&stepConfig.Commitish, "commitish", "master", "Target git commitish for the release")
	cmd.Flags().StringSliceVar(&stepConfig.ExcludeLabels, "excludeLabels", []string{}, "Allows to exclude issues with dedicated list of labels.")
	cmd.Flags().StringSliceVar(&stepConfig.Labels, "labels", []string{}, "Labels to include in issue search.")
	cmd.Flags().StringVar(&stepConfig.Owner, "owner", os.Getenv("PIPER_owner"), "Set the GitHub organization.")
	cmd.Flags().StringVar(&stepConfig.ReleaseBodyHeader, "releaseBodyHeader", os.Getenv("PIPER_releaseBodyHeader"), "Content which will appear for the release.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Set the GitHub repository.")
	cmd.Flags().StringVar(&stepConfig.ServerURL, "serverUrl", "https://github.com", "GitHub server url for end-user access.")
	cmd.Flags().StringVar(&stepConfig.Token, "token", os.Getenv("PIPER_token"), "GitHub personal access token as per https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
	cmd.Flags().StringVar(&stepConfig.UploadURL, "uploadUrl", "https://uploads.github.com", "Set the GitHub API url.")
	cmd.Flags().StringVar(&stepConfig.Version, "version", os.Getenv("PIPER_version"), "Define the version number which will be written as tag as well as release name.")

	cmd.MarkFlagRequired("apiUrl")
	cmd.MarkFlagRequired("owner")
	cmd.MarkFlagRequired("repository")
	cmd.MarkFlagRequired("serverUrl")
	cmd.MarkFlagRequired("token")
	cmd.MarkFlagRequired("uploadUrl")
	cmd.MarkFlagRequired("version")
}

// retrieve step metadata
func githubPublishReleaseMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "addClosedIssues",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "addDeltaToLastRelease",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "apiUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubApiUrl"}},
					},
					{
						Name:        "assetPath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "commitish",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "excludeLabels",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "labels",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "owner",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "github/owner"}},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubOrg"}},
					},
					{
						Name:        "releaseBodyHeader",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "repository",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "github/repository"}},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubRepo"}},
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubServerUrl"}},
					},
					{
						Name:        "token",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubToken"}},
					},
					{
						Name:        "uploadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "githubUploadUrl"}},
					},
					{
						Name:        "version",
						ResourceRef: []config.ResourceReference{{Name: "commonPipelineEnvironment", Param: "artifactVersion"}},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
