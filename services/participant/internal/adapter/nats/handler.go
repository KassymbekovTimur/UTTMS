package nats

import (
	"encoding/json"
	"log"

	"github.com/KassymbekovTimur/UTTMS/participant/internal/usecase"
	"github.com/nats-io/nats.go"
)

type NATSHandler struct {
	nc *nats.Conn
	uc *usecase.ParticipantUsecase
}

func NewNATSHandler(nc *nats.Conn, uc *usecase.ParticipantUsecase) *NATSHandler {
	return &NATSHandler{nc: nc, uc: uc}
}

func (h *NATSHandler) Start() error {
	_, err := h.nc.Subscribe("schedule.deleted", func(m *nats.Msg) {
		var event struct {
			ScheduleID string `json:"schedule_id"`
		}
		if err := json.Unmarshal(m.Data, &event); err != nil {
			log.Printf("[NATS] Failed to parse schedule.deleted: %v", err)
			return
		}
		log.Printf("[NATS] Received schedule.deleted: %s", event.ScheduleID)
		// Здесь можно добавить вызов usecase: RemoveScheduleFromAll(event.ScheduleID)
	})
	if err != nil {
		return err
	}

	log.Println("[NATS] Subscribed to schedule.deleted")
	return nil
}
