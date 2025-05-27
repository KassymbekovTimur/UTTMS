package redisrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/KassymbekovTimur/UTTMS/participant/internal/model"
	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisRepository(rdb *redis.Client) *RedisRepository {
	return &RedisRepository{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (r *RedisRepository) keyByID(id string) string {
	return fmt.Sprintf("participants:item:%s", id)
}

func (r *RedisRepository) keyBySchedule(scheduleID string) string {
	return fmt.Sprintf("participants:list:%s", scheduleID)
}

func (r *RedisRepository) Save(p *model.Participant) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return r.rdb.Set(r.ctx, r.keyByID(p.ID), data, 10*time.Minute).Err()
}

func (r *RedisRepository) FindByID(id string) (*model.Participant, error) {
	val, err := r.rdb.Get(r.ctx, r.keyByID(id)).Result()
	if err != nil {
		return nil, err
	}
	var p model.Participant
	if err := json.Unmarshal([]byte(val), &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *RedisRepository) FindByScheduleID(scheduleID string) ([]*model.Participant, error) {
	val, err := r.rdb.Get(r.ctx, r.keyBySchedule(scheduleID)).Result()
	if err != nil {
		return nil, err
	}
	var list []*model.Participant
	if err := json.Unmarshal([]byte(val), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *RedisRepository) Update(p *model.Participant) error {
	return r.Save(p)
}

func (r *RedisRepository) Delete(id string) error {
	return r.rdb.Del(r.ctx, r.keyByID(id)).Err()
}
