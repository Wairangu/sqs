package queues

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)



func SetMsgVisibility(sess *session.Session, handle *string, queueURL *string, visibility *int64) error {
    // Create an SQS service client
    svc := sqs.New(sess)

    // snippet-start:[sqs.go.change_message_visibility.op]
    _, err := svc.ChangeMessageVisibility(&sqs.ChangeMessageVisibilityInput{
        ReceiptHandle:     handle,
        QueueUrl:          queueURL,
        VisibilityTimeout: visibility,
    })
    // snippet-end:[sqs.go.change_message_visibility.op]
    if err != nil {
        return err
    }

    return nil
}