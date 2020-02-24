package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
)

func callTelephone(telephoneNumber string) error {
	sourcePhoneNumber := "+XXXXXXXXXXXX"
	instanceID := "XXXXXXXXXXXXXXXX"
	contactFlowID := "XXXXXXXXXXXXXXXXXXXXXX"
	attributeTemp := "Something Message."
	callAttribute := map[string]*string{"message": &attributeTemp}

	mySession := session.Must(session.NewSession())
	svc := connect.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))
	callInput := connect.StartOutboundVoiceContactInput{}
	callInput.SourcePhoneNumber = &sourcePhoneNumber
	callInput.InstanceId = &instanceID
	callInput.ContactFlowId = &contactFlowID
	callInput.Attributes = callAttribute
	callInput.DestinationPhoneNumber = &telephoneNumber
	_, error := svc.StartOutboundVoiceContact(&callInput)
	return error
}

func startCall() error {

	destinationPhoneNumbers := []string{}
	destinationPhoneNumbers = append(destinationPhoneNumbers, "+XXXXXXXXXXX") // anything
	destinationPhoneNumbers = append(destinationPhoneNumbers, "+XXXXXXXXXXX") // anything

	for i := range destinationPhoneNumbers {
		error := callTelephone(destinationPhoneNumbers[i])
		if error != nil {
			return error
		}
	}
	return nil
}

func main() {
	lambda.Start(startCall)
}
