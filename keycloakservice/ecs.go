package keycloakservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"

	"github.com/halflifeviper/keycloak/config"
)

func RegisterTask(cf *config.Config) (string, error) {
	session, err := CreateSessionAws(cf)
	if err != nil {
		return "", err
	}
	ecsClient := ecs.New(session)

	// Specify the container definition for the task
	containerDefinition := &ecs.ContainerDefinition{
		Name:   aws.String(cf.ECS.KeyCloak.Name),
		Image:  aws.String(cf.ECS.KeyCloak.Image),
		Memory: aws.Int64(cf.ECS.KeyCloak.Memory),
	}

	// Specify the task definition
	taskDefinitionInput := &ecs.RegisterTaskDefinitionInput{
		Family: aws.String(cf.ECS.KeyCloak.Family),
		ContainerDefinitions: []*ecs.ContainerDefinition{
			containerDefinition,
		},
	}

	// Register the task definition
	result, err := ecsClient.RegisterTaskDefinition(taskDefinitionInput)
	if err != nil {
		return "", err
	}

	return *result.TaskDefinition.TaskDefinitionArn, nil
}

func CreateECS(cf *config.Config) (string, error) {
	session, err := CreateSessionAws(cf)
	if err != nil {
		return "", err
	}
	ecsClient := ecs.New(session)

	// Specify the name of the ECS cluster
	clusterName := cf.ECS.Cluster.Name

	// Create ECS cluster input
	clusterInput := &ecs.CreateClusterInput{
		ClusterName: aws.String(clusterName),
		// Add any other cluster parameters as needed
	}

	// Create the ECS cluster
	result, err := ecsClient.CreateCluster(clusterInput)
	if err != nil {
		return "", err
	}

	return *result.Cluster.ClusterArn, nil
}
