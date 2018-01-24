package market_ws

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/bitly/go-simplejson"
)

func TestMarket(t *testing.T) {
	m, err := NewMarket()
	assert.NoError(t, err)
	err = m.Subscribe("market.eosusdt.kline.1min", func(topic string, json *simplejson.Json, raw []byte) {
		fmt.Println(topic, json, raw)
	})
	assert.NoError(t, err)
	err = m.Subscribe("market.eosusdt.trade.detail", func(topic string, json *simplejson.Json, raw []byte) {
		fmt.Println(topic, json, raw)
	})
	assert.NoError(t, err)
	fmt.Println(m)

	go func() {
		time.Sleep(time.Second * 20)
		m.Close()
	}()
	m.Loop()
}
