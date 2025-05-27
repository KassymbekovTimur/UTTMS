package nats

import (
	"encoding/json"
	"log"

	"github.com/KassymbekovTimur/UTTMS/statistics/internal/model"
	"github.com/KassymbekovTimur/UTTMS/statistics/internal/usecase"
	"github.com/nats-io/nats.go"
)

// NewHandler subscribes to NATS events and passes them to usecase
func NewHandler(nc *nats.Conn, uc *usecase.StatsUsecase) {
	nc.Subscribe("orders.*", func(m *nats.Msg) {
		log.Printf("[NATS] Received %s: %s", m.Subject, string(m.Data))
		var evt model.OrderEvent
		if err := json.Unmarshal(m.Data, &evt); err != nil {
			log.Println("[NATS] order unmarshal error:", err)
			return
		}
		if err := uc.HandleOrderEvent(evt); err != nil {
			log.Println("[Usecase] HandleOrderEvent error:", err)
		}
	})
	nc.Subscribe("users.*", func(m *nats.Msg) {
		log.Printf("[NATS] Received %s: %s", m.Subject, string(m.Data))
		var evt model.UserEvent
		if err := json.Unmarshal(m.Data, &evt); err != nil {
			log.Println("[NATS] user unmarshal error:", err)
			return
		}
		if err := uc.HandleUserEvent(evt); err != nil {
			log.Println("[Usecase] HandleUserEvent error:", err)
		}
	})
}
