package redis

import "testing"

func TestNewMutexRedis(t *testing.T) {
	_ = NewMutexRedis("localhost:5432", 0, "password")
}
