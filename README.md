# gofixtures
Ruby on Rails style fixture for testing. No dependencies for other database driver and ORM.

#Usage
    type User struct{
      Name string  
      Age  int64  
    }

    import (
      github.com/kyawmyintthein/gofixtures
    )
    
    func main(){
      // setup config for fixtures
      config := gofixtures.Config{
        Path: "data",  //path to fixtures folder
      }
      _ := gofixtures.SetupConfig(config)
      fixture, _ := gofixtures.LoadFixture("example") // load example.yaml file 
      user := &User{}
      _ = fixture.Load("one",user)
    }
    
    
    // yaml file
    one:
      name: John
      age: 21
    two:
      Name: Kyaw
      age: 23
    
    
    


