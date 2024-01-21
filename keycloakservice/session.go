package keycloakservice

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/halflifeviper/keycloak/config"
)

func CreateSessionAws(config *config.Config) (*session.Session, error) {

	// Set up an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Owner.Region),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return sess, nil
}

func CreateKeyCloak(config *config.Config) (err error) {
	err = CreatePostgres(config)
	if err != nil {
		return err
	}

	taskArn, err := RegisterTask(config)
	if err != nil {
		return err
	}
	log.Println(taskArn)

	ecs, err := CreateECS(config)
	if err != nil {
		return err
	}
	log.Println(ecs)

	return nil
}

func DestroyKeyCloak(config *config.Config) (err error) {
	err = DestroyPostgres(config)
	if err != nil {
		return err
	}
	return nil
}
