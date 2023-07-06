package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewClickHouseClient,
	NewRedisClient,

	NewRealtimeWarehousingRepo,
	NewAcceptStatusRepo,
)
