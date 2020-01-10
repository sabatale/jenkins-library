package cmd

import (
	"fmt"

	"github.com/SAP/jenkins-library/pkg/command"
)

func mavenExecute(MavenExecuteOptions mavenExecuteOptions) error {

	c := command.Command{}
	c.RunShell("bash", "mvn clean package")
	_, err := fmt.Printf("mvn")

	return err
}
