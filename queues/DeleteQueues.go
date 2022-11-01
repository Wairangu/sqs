package queues

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteQueue(sess *session.Session, queueName *string) {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	queueURL := GetQURL(sess, queueName)

	// Create an SQS service client
	svc := sqs.New(sess)

	// snippet-start:[sqs.go.delete_queue.call]
	_, err := svc.DeleteQueue(&sqs.DeleteQueueInput{
		QueueUrl: &queueURL,
	})
	// snippet-end:[sqs.go.delete_queue.call]
	if err != nil {
		fmt.Println(err)
		return
	}

}
