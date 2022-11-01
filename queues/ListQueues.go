package queues

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ListQueues() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}
