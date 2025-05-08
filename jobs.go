package context

type PoolKey string

const (
	// PoolName used in the Jobs drivers to indicate on which pool execute the job
	PoolName PoolKey = "pool" //nolint:gochecknoglobals
)
