package manager

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

type State struct {
	client client.Client
	store  string
}

func NewState(client client.Client, store string) *State {
	return &State{
		client: client,
		store:  store,
	}
}

func (state *State) Get(key string) (*client.StateItem, error) {
	ctx := context.Background()

	return state.client.GetState(ctx, state.store, key)
}

func (state *State) Set(key string, value []byte) error {
	ctx := context.Background()
	if err := state.client.SaveState(ctx, state.store, key, value); err != nil {
		return err
	}

	return nil
}

func (state *State) Delete(key string) error {
	ctx := context.Background()
	return state.client.DeleteState(ctx, state.store, key)
}

func (state *State) Update(oldKey string, newKey string) error {
	oldResult, err := state.Get(oldKey)
	if err != nil {
		return err
	}
	if err := state.Delete(oldKey); err != nil {
		return err
	}
	if err := state.Set(newKey, oldResult.Value); err != nil {
		return err
	}

	return nil
}
