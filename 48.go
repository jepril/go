package main
	
import (
"context"
"fmt"
"os"

"github.com/apache/rocketmq-client-go"
//"github.com/apache/rocketmq-client-go/primitive"
//"github.com/apache/rocketmq-client-go/producer"
)

// Package main implements a simple producer to send message.
func main() {
p, _ := rocketmq.NewProducer(
producer.WithNameServer([]string{"127.0.0.1:9876"}),
producer.WithRetry(2),
)
err := p.Start()
if err != nil {
fmt.Printf("start producer error: %s", err.Error())
os.Exit(1)
}
for i := 0; i < 1000; i++ {
res, err := p.SendSync(context.Background(), primitive.NewMessage("test",
[]byte("Hello RocketMQ Go Client!")))

if err != nil {
fmt.Printf("send message error: %s\n", err)
} else {
fmt.Printf("send message success: result=%s\n", res.String())
}
}
err = p.Shutdown()
if err != nil {
fmt.Printf("shutdown producer error: %s", err.Error())
}
}