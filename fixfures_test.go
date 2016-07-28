package gofixtures

import(
    "testing"
)


type User struct{
    Name string  `yaml:"name"`
    Age  int64  `yaml:"age"`
}

func TestInitFixture(t *testing.T) {
    config := Config{
        Path: "data",
    }
    c := SetupConfig(config)
    // fmt.Printf("%v",c)

    if config.Path != c.Path{
         t.Errorf("failed to setup config")
    }
}

func TestLoadFixtureFile(t *testing.T){
    b, err := loadFixtureFile("data/example.yaml")
    if err != nil{
        t.Errorf("failed to read yaml file: %v", err)
    }

    if len(b) <= 0 {
        t.Errorf("unable to read yaml file")
    }
}   

func TestLoadFixture(t *testing.T){
    fixture, err := LoadFixture("example")
    if err != nil{
        t.Errorf("failed to load fixture: %v", err)
    }

    if len(fixture.Content) <= 0 {
        t.Errorf("unable to load fixture")
    }
}

func TestLoadFixtureModel(t *testing.T){
    fixture, err := LoadFixture("example")
    if err != nil{
        t.Errorf("failed to load fixture: %v", err)
    }
    user := &User{}
    err = fixture.Load("one",user)
    if err != nil{
        t.Errorf("failed to load fixture model: %v", err)
    }

    if user.Name != "John"{
         t.Errorf("failed to load fixture model")
    }
}