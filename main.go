package main

import (
	"fmt"
	"github.com/Wairangu/sqs/queues"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	q := "testQ2"
	timeout := aws.Int64(60)
	//queues.CreateQueues(&q)
	//queues.ListQueues()
	//queues.DeleteQueue(&q)
	queueURL := queues.GetQURL(sess, &q)
	//queues.SendMsg(&queueURL)

	msg, err := queues.GetMessages(sess, &queueURL, timeout)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(*msg.Messages[0].MessageAttributes["Author"].StringValue)
	// fmt.Println(*msg.Messages[0].MessageAttributes["Title"].StringValue)
	// fmt.Println(*msg.Messages[0].MessageAttributes["WeeksOn"].StringValue)
	// fmt.Println(*msg.Messages[0].Body)

	fmt.Println(msg)

}
