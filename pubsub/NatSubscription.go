package pubsub

import (
	"encoding/json"
	"github.com/devtron-labs/external-app-crawler/client"
	"github.com/devtron-labs/external-app-crawler/common"
	"github.com/nats-io/stan"
	"go.uber.org/zap"
	"time"
)

type NatSubscription interface {
	Subscribe() error
}

type NatSubscriptionImpl struct {
	pubSubClient *client.PubSubClient
	logger       *zap.SugaredLogger
}

func NewNatSubscription(pubSubClient *client.PubSubClient,
	logger *zap.SugaredLogger) (*NatSubscriptionImpl, error) {
	ns := &NatSubscriptionImpl{
		pubSubClient: pubSubClient,
		logger:       logger,
	}
	return ns, ns.Subscribe()
}

func (impl *NatSubscriptionImpl) Subscribe() error {
	_, err := impl.pubSubClient.Conn.QueueSubscribe(client.TOPIC_CI_SCAN, client.TOPIC_CI_SCAN_GRP, func(msg *stan.Msg) {
		impl.logger.Debugw("received msg", "msg", msg)
		defer msg.Ack()
		config := &common.TelemetryUserAnalyticsDto{}
		err := json.Unmarshal(msg.Data, config)
		if err != nil {
			impl.logger.Errorw("err in reading msg", "err", err, "msg", string(msg.Data))
			return
		}
		impl.logger.Infow("scanConfig unmarshal data", "config", config)

	}, stan.DurableName(client.TOPIC_CI_SCAN_DURABLE), stan.StartWithLastReceived(), stan.AckWait(time.Duration(impl.pubSubClient.AckDuration)*time.Second), stan.SetManualAckMode(), stan.MaxInflight(1))
	//s.Close()
	return err
}
