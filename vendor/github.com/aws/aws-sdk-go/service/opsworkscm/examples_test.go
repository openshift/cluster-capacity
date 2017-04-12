// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package opsworkscm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleOpsWorksCM_AssociateNode() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.AssociateNodeInput{
		NodeName:   aws.String("NodeName"),   // Required
		ServerName: aws.String("ServerName"), // Required
		EngineAttributes: []*opsworkscm.EngineAttribute{
			{ // Required
				Name:  aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.AssociateNode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_CreateBackup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.CreateBackupInput{
		ServerName:  aws.String("ServerName"), // Required
		Description: aws.String("String"),
	}
	resp, err := svc.CreateBackup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_CreateServer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.CreateServerInput{
		InstanceProfileArn:     aws.String("InstanceProfileArn"), // Required
		ServerName:             aws.String("ServerName"),         // Required
		ServiceRoleArn:         aws.String("ServiceRoleArn"),     // Required
		BackupId:               aws.String("BackupId"),
		BackupRetentionCount:   aws.Int64(1),
		DisableAutomatedBackup: aws.Bool(true),
		Engine:                 aws.String("String"),
		EngineAttributes: []*opsworkscm.EngineAttribute{
			{ // Required
				Name:  aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		EngineModel:                aws.String("String"),
		EngineVersion:              aws.String("String"),
		InstanceType:               aws.String("String"),
		KeyPair:                    aws.String("KeyPair"),
		PreferredBackupWindow:      aws.String("TimeWindowDefinition"),
		PreferredMaintenanceWindow: aws.String("TimeWindowDefinition"),
		SecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SubnetIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateServer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DeleteBackup() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DeleteBackupInput{
		BackupId: aws.String("BackupId"), // Required
	}
	resp, err := svc.DeleteBackup(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DeleteServer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DeleteServerInput{
		ServerName: aws.String("ServerName"), // Required
	}
	resp, err := svc.DeleteServer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DescribeAccountAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	var params *opsworkscm.DescribeAccountAttributesInput
	resp, err := svc.DescribeAccountAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DescribeBackups() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DescribeBackupsInput{
		BackupId:   aws.String("BackupId"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
		ServerName: aws.String("ServerName"),
	}
	resp, err := svc.DescribeBackups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DescribeEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DescribeEventsInput{
		ServerName: aws.String("ServerName"), // Required
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DescribeNodeAssociationStatus() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DescribeNodeAssociationStatusInput{
		NodeAssociationStatusToken: aws.String("NodeAssociationStatusToken"), // Required
		ServerName:                 aws.String("ServerName"),                 // Required
	}
	resp, err := svc.DescribeNodeAssociationStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DescribeServers() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DescribeServersInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
		ServerName: aws.String("ServerName"),
	}
	resp, err := svc.DescribeServers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_DisassociateNode() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.DisassociateNodeInput{
		NodeName:   aws.String("NodeName"),   // Required
		ServerName: aws.String("ServerName"), // Required
		EngineAttributes: []*opsworkscm.EngineAttribute{
			{ // Required
				Name:  aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.DisassociateNode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_RestoreServer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.RestoreServerInput{
		BackupId:     aws.String("BackupId"),   // Required
		ServerName:   aws.String("ServerName"), // Required
		InstanceType: aws.String("String"),
		KeyPair:      aws.String("KeyPair"),
	}
	resp, err := svc.RestoreServer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_StartMaintenance() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.StartMaintenanceInput{
		ServerName: aws.String("ServerName"), // Required
	}
	resp, err := svc.StartMaintenance(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_UpdateServer() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.UpdateServerInput{
		ServerName:                 aws.String("ServerName"), // Required
		BackupRetentionCount:       aws.Int64(1),
		DisableAutomatedBackup:     aws.Bool(true),
		PreferredBackupWindow:      aws.String("TimeWindowDefinition"),
		PreferredMaintenanceWindow: aws.String("TimeWindowDefinition"),
	}
	resp, err := svc.UpdateServer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleOpsWorksCM_UpdateServerEngineAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := opsworkscm.New(sess)

	params := &opsworkscm.UpdateServerEngineAttributesInput{
		AttributeName:  aws.String("AttributeName"), // Required
		ServerName:     aws.String("ServerName"),    // Required
		AttributeValue: aws.String("AttributeValue"),
	}
	resp, err := svc.UpdateServerEngineAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
