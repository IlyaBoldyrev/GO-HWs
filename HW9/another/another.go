package another

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/url"
	"os"
)

type myURL url.URL

func (mu *myURL) Decode(value string) error {
	temp, err := url.Parse(value)
	if err != nil {
		return err
	}
	*mu = myURL(*temp)
	return err
}

func (mu *myURL) UnmarshalJSON(value []byte) error {
	temp, err := url.Parse(string(value[1 : len(value)-1]))
	if err != nil {
		return err
	}
	*mu = myURL(*temp)
	return err
}

type Config struct {
	Port         int
	Db_url       myURL
	Jaeger_url   myURL
	Sentry_url   myURL
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func GetConfig() Config {
	fmt.Println(os.Getenv("GO111MODULE"))
	var s Config
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s
}

/*
Переменные окружения для ввода в терминал:
export MYAPP_PORT=8080
export MYAPP_DB_URL=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
export MYAPP_JAEGER_URL=http://jaeger:16686
export MYAPP_SENTRY_URL=http://sentry:9000
export MYAPP_KAFKA_BROKER=kafka:9092
export MYAPP_SOME_APP_ID=testid
export MYAPP_SOME_APP_KEY=testkey
*/
