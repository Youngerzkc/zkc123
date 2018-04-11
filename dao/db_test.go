package dao

import(
	"testing"
)
// type RedisConf struct{
// 	Redis []Redisf
// }
// type Redisf struct{
// 	Port int
// 	Host string
// 	MaxActive int
//     MaxIdle int
// }

func TestGetRedisConf(t *testing.T )  {
	 rf:= &RedisConf{}
	 rf=GetRedisConf()
	if rf.Redis[0].Port==6379{
		t.Log("pass")
	}
}