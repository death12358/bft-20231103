package script

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UseUserPool function - keys []string,  args ...interface{} - return *[]RedisResult , error
func (s *MyScriptor) UseUserPool(keys []string, args []string) (int64, error) {

	res, err := s.Scriptor.ExecSha(UseUserPoolID, keys, args)
	if err != nil {
		logtool.LogError("UseUserPool ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("UseUserPool Value Error", err)
		return 0, err
	}

	return result.ValueInt64, nil
}

// UseUserPool - 批量增加數字
const (
	UseUserPoolID       = "UseUserPool"
	UseUserPoolTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UseUserPool
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = KEYS[4]
		local v1                                            = ARGV
		local sender                                        = "UseUserPool.lua"
		
		---@return number 
		local function getTime()
			return redis.call("TIME")[1]
		end

		if DBKey and ProjectKey and TagKey and k1 and v1 then

			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
			redis.call("select",DBKey)

			local count = #v1
			for i = 1,count,2 do
				redis.call('hincrby',MAIN_KEY , v1[i] , - v1[i+1])
			end

			redis.call('hset',MAIN_KEY, 'lastUpdateTime', getTime())
			
			local r1 = ""
			local Tmp = redis.call('hgetall',MAIN_KEY)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return v1
		end
    `
)
