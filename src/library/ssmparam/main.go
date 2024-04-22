package ssmparam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"strconv"
)

func getParameter(parameterName string) (*string, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	ssmClient := ssm.NewFromConfig(cfg)

	input := &ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(false),
	}

	result, err := ssmClient.GetParameter(context.Background(), input)
	if err != nil {
		return nil, err
	}

	return result.Parameter.Value, nil
}

func updateParameter(parameterName, parameterValue string) (*ssm.PutParameterOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	ssmClient := ssm.NewFromConfig(cfg)

	overwrite := true
	input := &ssm.PutParameterInput{
		Name:      aws.String(parameterName),
		Value:     aws.String(parameterValue),
		Type:      "String",
		Overwrite: &overwrite,
	}

	result, err := ssmClient.PutParameter(context.Background(), input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// IncreaseCounter is used when new firewall rule is added
func IncreaseCounter(parameterName string) error {
	s, err := getParameter(parameterName)
	if err != nil {
		return err
	}

	num, err := strconv.Atoi(*s)
	if err != nil {
		return err
	}

	num = num + 1 // increase number

	_, err = updateParameter(parameterName, fmt.Sprint(num))
	if err != nil {
		return err
	}

	return nil
}
