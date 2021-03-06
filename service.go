package service

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Service interface is use to handle all the possible dependencies of the
// service routes this will allow the user to define any kind of struct for
// the service without compromising the integrity of the router.
type Service interface {
	GetRoutes() Routes
	InitRouter(*httprouter.Router) *httprouter.Router
}

// Start this will bind the service routes to the adddress passed
// auth and other middleware definition shouldn't be define as part
// of the service biding they should be define before the handler are
// set to the Route to keep service consisten.
func Start(addrs string, s Service) {
	router := s.InitRouter(httprouter.New())

	router.Handle(http.MethodGet, "/_ls", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		b, _ := json.Marshal(s.GetRoutes())
		w.Write(b)
	})
	for _, r := range s.GetRoutes() {
		if log != nil {
			log.Infof("Register endpoint %s with the method %s and handler %T \n", r.Path, r.Method, r.Handler)
		}
		router.Handle(r.Method, r.Path, r.Handler)
	}
	if log != nil {
		log.Infof("Starting server at port: %s and service %T ", addrs, s)
	}
	err := http.ListenAndServe(addrs, &logHandler{router})
	if err != nil {
		panic(err)
	}
}

// Routes a map of each Route on a Service.
type Routes = map[string]Route

// Route on a service.
type Route struct {
	Path    string
	Method  string
	Handler httprouter.Handle
}
