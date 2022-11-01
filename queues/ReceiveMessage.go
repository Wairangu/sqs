package queues

// snippet-start:[sqs.go.receive_messages.imports]
import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetMessages(sess *session.Session,queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {

	// Create an SQS service client
	svc := sqs.New(sess)

	// snippet-start:[sqs.go.receive_messages.call]
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   timeout,
	})
	// snippet-end:[sqs.go.receive_messages.call]
	if err != nil {
		return nil, err
	}

	if len(msgResult.Messages) ==0 {
		return nil, errors.New("queue is empty")
	}
	receiptHandle := msgResult.Messages[0].ReceiptHandle
	DeleteMessage(sess,queueURL, receiptHandle)

	return msgResult, nil
}
