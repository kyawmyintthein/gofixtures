package gofixtures

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"path/filepath"
	"os"
    "reflect"
    "strings"
)

type (
	Config struct {
		Path string
	}

	Fixture struct {
		Name    string
		Content []byte
	}

     Content map[string]map[string]interface{}
)

var (
	config *Config
)

func (fixture Fixture) Load(name string, result interface{}) error{
	var content Content
	err := yaml.Unmarshal(fixture.Content, &content)
	t := reflect.ValueOf(result).Elem()
	for k, v := range content[name] {
		field := strings.ToUpper(k[:1]) + k[1:]
		val := t.FieldByName(field)
		val.Set(reflect.ValueOf(v))
	}
	return err
}



// func FieldByNameOrTag(v reflect.Value, name string) (ret reflect.Value) {
// 	defer func() { fmt.Println(v, name, ret) }()
// 	field := strings.ToUpper(name[:1]) + name[1:]
// 	ret = v.FieldByName(field)
// 	if ret.IsValid() {
// 		if ret.Kind() == reflect.Ptr {
// 			if ret.Elem().IsNil() {
// 				ret.Set(reflect.New(ret.Type().Elem()))
// 			}
// 			ret = ret.Elem()
// 			return
// 		}
// 		return
// 	}

// 	t := v.Type()
// 	for i, n := 0, t.NumField(); i < n; i++ {
// 		f := t.Field(i)
// 		if name == f.Tag.Get("yaml") {
// 			ret = v.Field(i)
// 			if ret.Kind() == reflect.Ptr {
// 				if ret.Elem().IsNil() {
// 					ret.Set(reflect.New(ret.Type().Elem()))
// 				}
// 				ret = ret.Elem()
// 				return
// 			}
// 			return
// 		}
// 	}
// 	return
// }

//load fixture with name
func LoadFixture(name string) (Fixture, error) {
	var (
		fixture Fixture
		err     error
		content []byte
	)

	filePath, _ := filepath.Abs(config.Path + "/" + name + ".yaml")
	content, err = loadFixtureFile(filePath)
	fixture = Fixture{
		Name:    name,
		Content: content,
	}
	return fixture, err
}

// read yaml file as byte array
func loadFixtureFile(path string) ([]byte, error) {
	var (
		content []byte
		err     error
	)

	if _, err = os.Stat(path);  os.IsNotExist(err){
		return content, err
	}

	content, err = ioutil.ReadFile(path)
	return content, err
}

// Init fixtures
func SetupConfig(c Config) *Config{
	// fmt.Printf("%+v",c)
	config.Path = c.Path
	return config
}

func init(){
	if config == nil {
		config = &Config{}
	}
}