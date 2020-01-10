package cmd

import (
	"os"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/spf13/cobra"
)

type mavenExecuteOptions struct {
	PomPath                     string   `json:"pomPath,omitempty"`
	ProjectSettingsFile         string   `json:"projectSettingsFile,omitempty"`
	GlobalSettingsFile          string   `json:"globalSettingsFile,omitempty"`
	M2Path                      string   `json:"m2Path,omitempty"`
	Goals                       []string `json:"goals,omitempty"`
	Defines                     []string `json:"defines,omitempty"`
	Flags                       []string `json:"flags,omitempty"`
	LogSuccessfulMavenTransfers bool     `json:"logSuccessfulMavenTransfers,omitempty"`
}

var myMavenExecuteOptions mavenExecuteOptions
var mavenExecuteStepConfigJSON string

// MavenExecuteCommand Ths step allows to run maven commands
func MavenExecuteCommand() *cobra.Command {
	metadata := mavenExecuteMetadata()
	var createMavenExecuteCmd = &cobra.Command{
		Use:   "mavenExecute",
		Short: "Ths step allows to run maven commands",
		Long:  `Ths step allows to run maven commands`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetStepName("mavenExecute")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "mavenExecute", &myMavenExecuteOptions, config.OpenPiperFile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return mavenExecute(myMavenExecuteOptions)
		},
	}

	addMavenExecuteFlags(createMavenExecuteCmd)
	return createMavenExecuteCmd
}

func addMavenExecuteFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&myMavenExecuteOptions.PomPath, "pomPath", os.Getenv("PIPER_pomPath"), "Path to the pom file that should be used.")
	cmd.Flags().StringVar(&myMavenExecuteOptions.ProjectSettingsFile, "projectSettingsFile", os.Getenv("PIPER_projectSettingsFile"), "Path to the mvn settings file that should be used as project settings file.")
	cmd.Flags().StringVar(&myMavenExecuteOptions.GlobalSettingsFile, "globalSettingsFile", os.Getenv("PIPER_globalSettingsFile"), "Path to the mvn settings file that should be used as global settings file.")
	cmd.Flags().StringVar(&myMavenExecuteOptions.M2Path, "m2Path", os.Getenv("PIPER_m2Path"), "Path to the location of the local repository that should be used.*/")
	cmd.Flags().StringSliceVar(&myMavenExecuteOptions.Goals, "goals", []string{}, "Maven goals that should be executed.")
	cmd.Flags().StringSliceVar(&myMavenExecuteOptions.Defines, "defines", []string{}, "Additional properties in form of -Dkey=value.")
	cmd.Flags().StringSliceVar(&myMavenExecuteOptions.Flags, "flags", []string{}, "Flags to provide when running mvn.")
	cmd.Flags().BoolVar(&myMavenExecuteOptions.LogSuccessfulMavenTransfers, "logSuccessfulMavenTransfers", false, "Configures maven to log successful downloads. This is set to `false` by default to reduce the noise in build logs.")

	cmd.MarkFlagRequired("goals")
}

// retrieve step metadata
func mavenExecuteMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:      "pomPath",
						Scope:     []string{"PARAMETERS", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "projectSettingsFile",
						Scope:     []string{"PARAMETERS", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "globalSettingsFile",
						Scope:     []string{"PARAMETERS", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "m2Path",
						Scope:     []string{"PARAMETERS", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "goals",
						Scope:     []string{"PARAMETERS"},
						Type:      "[]string",
						Mandatory: true,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "defines",
						Scope:     []string{"PARAMETERS"},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "flags",
						Scope:     []string{"PARAMETERS", "STEPS"},
						Type:      "[]string",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
					{
						Name:      "logSuccessfulMavenTransfers",
						Scope:     []string{"PARAMETERS"},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
