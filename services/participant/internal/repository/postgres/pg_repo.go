package postgresrepo

import (
	"database/sql"
	"encoding/json"

	"github.com/KassymbekovTimur/UTTMS/participant/internal/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

//goland:noinspection ALL
func (r *PostgresRepository) Save(p *model.Participant) error {
	scheduleIDs, err := json.Marshal(p.ScheduleIDs)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(
		`INSERT INTO participants(id, name, email, schedule_ids, status) VALUES($1, $2, $3, $4, $5)`,
		p.ID, p.Name, p.Email, string(scheduleIDs), p.Status,
	)
	return err
}

func (r *PostgresRepository) FindByID(id string) (*model.Participant, error) {
	row := r.db.QueryRow(`SELECT id, name, email, schedule_ids, status FROM participants WHERE id = $1`, id)
	var p model.Participant
	var schedStr string
	if err := row.Scan(&p.ID, &p.Name, &p.Email, &schedStr, &p.Status); err != nil {
		return nil, err
	}
	_ = json.Unmarshal([]byte(schedStr), &p.ScheduleIDs)
	return &p, nil
}

func (r *PostgresRepository) FindByScheduleID(scheduleID string) ([]*model.Participant, error) {
	rows, err := r.db.Query(
		`SELECT id, name, email, schedule_ids, status FROM participants WHERE schedule_ids::text LIKE '%' || $1 || '%'`,
		scheduleID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.Participant
	for rows.Next() {
		var p model.Participant
		var schedStr string
		if err := rows.Scan(&p.ID, &p.Name, &p.Email, &schedStr, &p.Status); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(schedStr), &p.ScheduleIDs)
		result = append(result, &p)
	}
	return result, nil
}

func (r *PostgresRepository) Update(p *model.Participant) error {
	scheduleIDs, err := json.Marshal(p.ScheduleIDs)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(
		`UPDATE participants SET name = $1, email = $2, schedule_ids = $3, status = $4 WHERE id = $5`,
		p.Name, p.Email, string(scheduleIDs), p.Status, p.ID,
	)
	return err
}

func (r *PostgresRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM participants WHERE id = $1`, id)
	return err
}
