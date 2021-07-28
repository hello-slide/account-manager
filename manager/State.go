package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

type State struct {
	client *client.Client
	ctx    *context.Context
	store  string
}

func NewState(client *client.Client, ctx *context.Context, store string) *State {
	return &State{
		client: client,
		ctx:    ctx,
		store:  store,
	}
}

func (s *State) Get(key string) (*client.StateItem, error) {
	return (*s.client).GetState(*s.ctx, s.store, key)
}

func (s *State) Set(key string, value []byte) error {
	if err := (*s.client).SaveState(*s.ctx, s.store, key, value); err != nil {
		return err
	}

	return nil
}

func (s *State) Delete(key string) error {
	return (*s.client).DeleteState(*s.ctx, s.store, key)
}

func (s *State) Update(oldKey string, newKey string) error {
	oldResult, err := s.Get(oldKey)
	if err != nil {
		return err
	}
	if err := s.Delete(oldKey); err != nil {
		return err
	}
	if err := s.Set(newKey, oldResult.Value); err != nil {
		return err
	}

	return nil
}
