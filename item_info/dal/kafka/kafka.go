package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"item_info/model"
)

var cli sarama.SyncProducer

const Topic = "moni"

func Init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//生产者
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}
	cli = client
}

func Close() {
	if cli != nil {
		cli.Close()
	}
}

func Produce(ctx context.Context, text string) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = Topic
	msg.Value = sarama.StringEncoder(text)

	pid, offset, err := cli.SendMessage(msg)
	if err != nil {
		return err
	}
	klog.CtxInfof(ctx, "msg produced, pid:%v offset:%v\n", pid, offset)
	return nil
}

func SendItemChange(ctx context.Context, ic *model.ItemChange) {
	bytes, err := json.Marshal(ic)
	if err != nil {
		klog.CtxErrorf(ctx, "marshal failed: %v", err)
		return
	}
	err = Produce(ctx, string(bytes))
	if err != nil {
		klog.CtxErrorf(ctx, "produce failed: %v", err)
		return
	}
}
