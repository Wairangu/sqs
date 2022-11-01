package queues

import (
	//"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetQURL(sess *session.Session,queue *string) string {
	// queue := flag.String("q", "", "The name of the queue")
	// flag.Parse()

	// if *queue == "" {
	// 	fmt.Println("You must supply a queue name (-q QUEUE")
	// 	return ""
	// }

	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return *result.QueueUrl

}
