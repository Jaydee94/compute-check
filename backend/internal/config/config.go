package config

import(
 "ioutil"
 "yaml"
)

type Config struct {
  // Config data goes here, along with yaml bindings
  // We could keep a sample config in config/ and always
  // autogenerate this config struct from it using
  // https://yaml.to-go.online/
}

// LoadConfig takes a yaml file path and load it into
// a Config object
func LoadConfig(path string) (*Config, error) {
  cfg := &Config{}
  data := ioutil.ReadFile(path)
  err := yaml.Unmarshal([]byte(data), cfg)
  if err != nil {
	  return nil, err
  }
  return cfg, nil
}
