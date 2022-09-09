package ad

import (
  "fmt"
  "github.com/hypebeast/go-osc/osc"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Application struct {
  name        string

  addr        string
  dispatcher *osc.StandardDispatcher
  server     *osc.Server
}

func NewApplication(name string) *Application {
  app := Application{
    name:       name,
    dispatcher: osc.NewStandardDispatcher(),
  }

  return &app
}

func (a *Application) RegisterVariable(name string, handler osc.HandlerFunc) (string, error) {
  path := fmt.Sprintf("%s/%s", a.name, name)
  return path, a.dispatcher.AddMsgHandler(path, handler)
}

func (a *Application) RegisterSubVariable(sub, name string, handler osc.HandlerFunc) (string, error) {
  path := fmt.Sprintf("%s/%s/%s", a.name, sub, name)
  return path, a.dispatcher.AddMsgHandler(path, handler)
}

func (a *Application) Start(addr string) error {
  a.server = &osc.Server{
    Addr:        addr,
    Dispatcher:  a.dispatcher,
  }
  return a.server.ListenAndServe()
}
