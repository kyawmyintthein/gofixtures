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
		// fmt.Println(val.Kind().String())
		switch val.Kind().String(){
		case "string":
			val.Set(reflect.ValueOf(v))
		case "int64":
			i := v.(int)
			val.Set(reflect.ValueOf(int64(i)))
		case "int32":
			i := v.(int)
			val.Set(reflect.ValueOf(int32(i)))
		default:
			val.Set(reflect.ValueOf(v))
		}
		
	}
	return err
}


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
func SetupConfig(path string){
	config.Path = path
}

func init(){
	if config == nil {
		config = &Config{}
	}
}