package script

import goredis "github.com/adimax2953/go-redis"

// RedisResult -
type RedisResult struct {
	Value      string
	Value2     string
	CountDown  int64
	EndTime    int64
	ValueInt64 int64
	Key        string
}

type MyScriptor struct {
	Scriptor *goredis.Scriptor
}

var LuaScripts = map[string]string{
	UseUserPoolID: UseUserPoolTemplate,
}
