package global

// Gearman
const (
	ETCD_KEY_GEARMAN_ADDR = "/banerwai/gearman/conn"
)

// Gearman
const (
	ETCD_KEY_REDIS_ADDR = "/banerwai/redis/conn"
)

// micros/query
const (
	ETCD_KEY_MICROS_QUERY_CATEGORY = "banerwai/micros/query/category/addr"
	ETCD_KEY_MICROS_QUERY_PROFILE  = "banerwai/micros/query/profile/addr"
	ETCD_KEY_MICROS_QUERY_RENDER   = "banerwai/micros/query/render/addr"
)

// command/query
const (
	ETCD_KEY_MICROS_COMMAND_TOKEN   = "banerwai/micros/command/token/addr"
	ETCD_KEY_MICROS_COMMAND_AUTH    = "banerwai/micros/command/auth/addr"
	ETCD_KEY_MICROS_COMMAND_PROFILE = "banerwai/micros/command/profile/addr"
)
