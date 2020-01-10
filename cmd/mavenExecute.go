package cmd

import (
	"github.com/SAP/jenkins-library/pkg/command"
)

func mavenExecute(myMavenExecuteOptions mavenExecuteOptions) error {
	parameters := []string{}

	if myMavenExecuteOptions.GlobalSettingsFile != "" {
		globalSettingsFileParameter := " --global-settings " + myMavenExecuteOptions.GlobalSettingsFile
		parameters = append(parameters, globalSettingsFileParameter)
	}

	if myMavenExecuteOptions.ProjectSettingsFile != "" {
		projectSettingsFileParameter := "--settings " + myMavenExecuteOptions.ProjectSettingsFile
		parameters = append(parameters, projectSettingsFileParameter)
	}

	if myMavenExecuteOptions.M2Path != "" {
		m2PathParameter := "-Dmaven.repo.local=" + myMavenExecuteOptions.M2Path
		parameters = append(parameters, m2PathParameter)
	}
	if myMavenExecuteOptions.PomPath != "" {
		pomPathParameter := "--file " + myMavenExecuteOptions.PomPath
		parameters = append(parameters, pomPathParameter)
	}

	if myMavenExecuteOptions.Flags != nil {
		parameters = append(parameters, myMavenExecuteOptions.Flags...)
	}

	parameters = append(parameters, "--batch-mode")

	if myMavenExecuteOptions.LogSuccessfulMavenTransfers {
		parameters = append(parameters, "-Dorg.slf4j.simpleLogger.log.org.apache.maven.cli.transfer.Slf4jMavenTransferListener=warn")
	}

	parameters = append(parameters, myMavenExecuteOptions.Goals...)

	c := command.Command{}
	c.RunExecutable("mvn", parameters...)


	return nil
}
