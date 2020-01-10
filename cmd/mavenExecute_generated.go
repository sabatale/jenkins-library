package cmd

import (
	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/spf13/cobra"
)

type mavenExecuteOptions struct {
}

var myMavenExecuteOptions mavenExecuteOptions
var mavenExecuteStepConfigJSON string

// MavenExecuteCommand Maven
func MavenExecuteCommand() *cobra.Command {
	metadata := mavenExecuteMetadata()
	var createMavenExecuteCmd = &cobra.Command{
		Use:   "mavenExecute",
		Short: "Maven",
		Long:  `Maven`,
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

}

// retrieve step metadata
func mavenExecuteMetadata() config.StepData {
	var theMetaData = config.StepData{
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{},
			},
		},
	}
	return theMetaData
}
