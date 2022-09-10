package ad

import (
  "fmt"
  "github.com/briancsparks/activedevelopment/ad/grumpy"
  "github.com/hypebeast/go-osc/osc"
  "sync"
)

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type Application struct {
  name        string

  // OSC
  port               int
  addr               string
  dispatcher        *osc.StandardDispatcher
  server            *osc.Server

  wantFeedback       bool
  //cAddr              string
  client            *osc.Client
}

func NewApplication(name string) *Application {
  app := Application{
    name:       name,
    dispatcher: osc.NewStandardDispatcher(),
  }

  return &app
}

func NewApplicationWithFeedback(name string) *Application {
  app := NewApplication(name)
  app.wantFeedback = true
  return app
}

func (app *Application) RegisterVariable(name string, handler osc.HandlerFunc) (string, error) {
  path := fmt.Sprintf("/%s/%s", app.name, name)
  return path, app.dispatcher.AddMsgHandler(path, handler)
}

func (app *Application) RegisterSubVariable(sub, name string, handler osc.HandlerFunc) (string, error) {
  path := fmt.Sprintf("/%s/%s/%s", app.name, sub, name)
  return path, app.dispatcher.AddMsgHandler(path, handler)
}

func (app *Application) RegisterDefaultHandler(handler osc.HandlerFunc) (string, error) {
  return "*", app.dispatcher.AddMsgHandler("*", handler)
}

func (app *Application) Start(ip string, port int) error {
  //app.port = port
  //addr := fmt.Sprintf("%s:%d", ip, app.port)
  //app.server = &osc.Server{
  //  Addr:        addr,
  //  Dispatcher:  app.dispatcher,
  //}

  // TODO: Proper error handling
  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    wg.Done()

    err := app.StartBlocking(ip, port)
    if err != nil {
      grumpy.Fatal(err)
    }
    // TODO: Signal on some channel that we are done listening
  }()

  wg.Wait()

  return nil
}

func (app *Application) StartBlocking(ip string, port int) error {
  app.port = port

  if app.wantFeedback {
    //app.cAddr = fmt.Sprintf("%s:%d", ip, app.port - 1)
    app.client = osc.NewClient(ip, app.port - 1)
  }

  addr := fmt.Sprintf("%s:%d", ip, app.port)
  app.server = &osc.Server{
    Addr:       addr,
    Dispatcher: app.dispatcher,
  }

  err := app.server.ListenAndServe()
  if err != nil {
    grumpy.Fatal(err)
  }
  return err
}

func (app *Application) Feedback(msg *osc.Message) error {
  return app.client.Send(msg)
}

func (app *Application) CloseConnection() error {
  return app.server.CloseConnection()
}


