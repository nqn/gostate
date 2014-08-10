package state

import (
  "github.com/nqn/gostate/storage"
)

type State struct {}

func New(storage.Storage) *State {
  return &State {}
}
