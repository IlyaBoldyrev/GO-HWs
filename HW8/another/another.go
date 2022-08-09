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
