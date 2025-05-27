package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/KassymbekovTimur/UTTMS/statistics/internal/model"
	"github.com/KassymbekovTimur/UTTMS/statistics/internal/repository"
)

type pgRepo struct {
	db *sql.DB
}

// NewPostgresRepo creates implementations of Repository on PostgreSQL
func NewPostgresRepo(db *sql.DB) repository.Repository {
	return &pgRepo{db: db}
}

func (r *pgRepo) SaveOrderEvent(evt model.OrderEvent) error {
	log.Printf("[SaveOrderEvent] %%+v", evt)
	_, err := r.db.Exec(
		`INSERT INTO order_events(user_id, order_id, ts) VALUES($1,$2,$3)`,
		evt.UserID, evt.OrderID, evt.Timestamp,
	)
	return err
}

func (r *pgRepo) SaveUserEvent(evt model.UserEvent) error {
	log.Printf("[SaveUserEvent] %%+v", evt)
	_, err := r.db.Exec(
		`INSERT INTO user_events(user_id, action, ts) VALUES($1,$2,$3)`,
		evt.UserID, evt.Action, evt.Timestamp,
	)
	return err
}

func (r *pgRepo) CountOrdersByUser(userID string) (int, error) {
	var cnt int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM order_events WHERE user_id=$1`, userID,
	).Scan(&cnt)
	log.Printf("[CountOrdersByUser] user=%s total=%d", userID, cnt)
	return cnt, err
}

func (r *pgRepo) CountOrdersByHour(userID string) (map[string]int, error) {
	rows, err := r.db.Query(
		`SELECT EXTRACT(HOUR FROM ts)::int AS hr, COUNT(*) FROM order_events
         WHERE user_id=$1 GROUP BY hr`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make(map[string]int)
	for rows.Next() {
		var hr, cnt int
		if err := rows.Scan(&hr, &cnt); err != nil {
			return nil, err
		}
		key := fmt.Sprintf("%02d", hr)
		res[key] = cnt
		log.Printf("[CountOrdersByHour] user=%s hour=%s cnt=%d", userID, key, cnt)
	}
	return res, nil
}

func (r *pgRepo) CountTotalUsers() (int, error) {
	var cnt int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM user_events`).Scan(&cnt)
	log.Printf("[CountTotalUsers] total=%d", cnt)
	return cnt, err
}

func (r *pgRepo) CountRegisteredUsers() (int, error) {
	var cnt int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM (SELECT DISTINCT user_id FROM user_events WHERE action='created') AS sub`,
	).Scan(&cnt)
	log.Printf("[CountRegisteredUsers] registered=%d", cnt)
	return cnt, err
}
