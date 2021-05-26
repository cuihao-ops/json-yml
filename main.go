package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Groups []Groups //`yaml:"groups"`
}

type Groups struct {
	Name  string  //`yaml:"name"`
	Rules []Rules //`yaml:"rules"`
}

type Rules struct {
	Alert       string      //`yaml:"alert"`
	Expr        string      //`yaml:"expr"`
	For         string      // `yaml:"for"`
	Labels      Labels      //`yaml:"labels"`
	Annotations Annotations ///`yaml:"annotations"`
}

type Labels struct {
	Severity string //`yaml:"severity"`
}

type Annotations struct {
	Summary     string //`yaml:"summary"`
	Description string //`yaml:"description"`
}

//read yaml config
//注：path为yaml或yml文件的路径
func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}

//test yaml
func main() {
	conf, err := ReadYamlConfig("ssl_expire.yml")
	if err != nil {
		log.Fatal(err)
	}

	byts, err := json.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byts))

	var s Config
	str := []byte(`{"Groups":[{"Name":"ssl_expiry","Rules":[{"Alert":"HTTPS SSL证书到期!","Expr":"(probe_ssl_earliest_cert_expiry - time() ) / 86400 \u003c 30","For":"24h","Labels":{"Severity":"warning"},"Annotations":{"Summary":"SSL证书将要过期 (instance {{ $labels.instance }})","Description":"HTTPS SSL证书将要在{{ $value }}天过期！"}}]}]}`)
	err = json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

}
