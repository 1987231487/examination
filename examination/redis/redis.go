package redis
import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client
func init(){
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func Set(key string,value string){
	err:=client.Set(key,value,5*time.Minute).Err()
	if err != nil {
		panic(err)
	}
}
func Get(Key string)(string,error){
	val2, err := client.Get(Key).Result()
	if err == redis.Nil {
		log.Println("key does not exists")
	} else if err != nil {
		panic(err)
	} else {
		return val2,err
	}
	return val2,err
}
