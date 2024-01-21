package keycloakservice

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"

	"github.com/halflifeviper/keycloak/config"
)

func CreatePostgres(cf *config.Config) error {
	session, err := CreateSessionAws(cf)
	if err != nil {
		return err
	}
	rdsClient := rds.New(session)
	// Describe the RDS instance
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(cf.Database.Postgres.Name)}

	existingInstance, _ := rdsClient.DescribeDBInstances(input)

	if len(existingInstance.DBInstances) > 0 {
		return nil
	}

	params := &rds.CreateDBInstanceInput{
		AllocatedStorage:        aws.Int64(cf.Database.Postgres.AllocatedStorage),
		DBInstanceIdentifier:    aws.String(cf.Database.Postgres.Name),
		DBInstanceClass:         aws.String(cf.Database.Postgres.InstanceClass),
		Engine:                  aws.String(cf.Database.Postgres.Engine),
		MasterUsername:          aws.String(cf.Database.Postgres.UserName),
		MasterUserPassword:      aws.String(cf.Database.Postgres.Password),
		DBName:                  aws.String(cf.Database.Postgres.Name),
		AutoMinorVersionUpgrade: aws.Bool(true),
		PubliclyAccessible:      aws.Bool(false),
	}

	result, err := rdsClient.CreateDBInstance(params)
	if err != nil {
		return err
	}

	log.Println(result.DBInstance.Endpoint)
	return nil
}

func DestroyPostgres(cf *config.Config) error {
	session, err := CreateSessionAws(cf)
	if err != nil {
		return err
	}
	rdsClient := rds.New(session)

	params := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(cf.Database.Postgres.Name),
		SkipFinalSnapshot:    aws.Bool(true),
	}

	// Delete the RDS instance
	_, err = rdsClient.DeleteDBInstance(params)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
