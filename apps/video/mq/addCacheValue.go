package mq

import (
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

const TypeAddCacheValue = "cache#add"

type AddCacheValuePayload struct {
	Key   string
	Field string
	Add   int64
}

// NewAddCacheValueTask creates a new asynq task for adding a value to a cache.
func NewAddCacheValueTask(key string, field string, add int64) (*asynq.Task, error) {
	payload, err := json.Marshal(AddCacheValuePayload{Key: key, Field: field, Add: add})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal payload for task %q", TypeAddCacheValue)
	}
	return asynq.NewTask(TypeAddCacheValue, payload), nil
}
