package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/andradesoares/mongo_crud/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb+srv://danielasoares23:vY03OoFsazfZdA63@cluster0.lyro71c.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	return s
}