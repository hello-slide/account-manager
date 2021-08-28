package state

import (
	"context"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/dapr/go-sdk/client"
)

type state struct {
	client *client.Client
	ctx    *context.Context
	store  string
}

func NewState(ctx context.Context, client *client.Client, store string) *state {
	return &state{
		client: client,
		ctx:    &ctx,
		store:  store,
	}
}

func (s *state) Get(key string) (*client.StateItem, error) {
	return (*s.client).GetState(*s.ctx, s.store, key)
}

func (s *state) Set(key string, value []byte) error {
	if err := (*s.client).SaveState(*s.ctx, s.store, key, value); err != nil {
		return err
	}

	return nil
}

func (s *state) SetTTL(key string, value []byte, ttl string) error {
	item := &client.SetStateItem{
		Key: key,
		Metadata: map[string]string{
			"created-on":   time.Now().UTC().String(),
			"ttlInSeconds": ttl,
		},
		Value: value,
	}
	if err := (*s.client).SaveBulkState(*s.ctx, s.store, item); err != nil {
		return err
	}
	return nil
}

func (s *state) Delete(key string) error {
	return (*s.client).DeleteState(*s.ctx, s.store, key)
}

func (s *state) Update(oldKey string, newKey string, ttl string) (string, error) {
	oldResult, err := s.Get(oldKey)
	if err != nil {
		return "", err
	}
	if utf8.RuneCount(oldResult.Value) == 0 {
		return "", fmt.Errorf("token does not exist")
	}

	if err := s.Delete(oldKey); err != nil {
		return "", err
	}
	if err := s.SetTTL(newKey, oldResult.Value, ttl); err != nil {
		return "", err
	}

	return string(oldResult.Value), nil
}
