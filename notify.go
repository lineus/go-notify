package notify

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// SNSPublishAPI words
type SNSPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

// Notify notifies via sms via aws
func Notify(msg string) error {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", config)

	client := sns.NewFromConfig(config)

	// input := &sns.PublishInput{
	// 	Message:  aws.String(msg),
	// 	TopicArn: aws.String(arn),
	// }

	// _, err = client.Publish(context.TODO(), input)

	return err
}
