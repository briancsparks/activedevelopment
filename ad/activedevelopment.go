package ad

import "fmt"

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

type ActiveDevelopment struct {
  App *Application
  Cfg *Config
}

func NewActiveDevelopment(appName string) *ActiveDevelopment {
  ad := ActiveDevelopment{
    App:  NewApplication(appName),
    Cfg:  NewConfig(),
  }
  TheConfig = ad.Cfg

  return &ad
}

func NewActiveDevelopmentWithFeedback(appName string) *ActiveDevelopment {
  ad := ActiveDevelopment{
    App:  NewApplicationWithFeedback(appName),
    Cfg:  NewConfig(),
  }
  TheConfig = ad.Cfg

  return &ad
}

func (ad *ActiveDevelopment) String() string {
  return fmt.Sprintf("App: %+v Cfg: %+v", ad.App, ad.Cfg)
}



// -------------------------------------------------------------------------------------------------------------------

func (ad *ActiveDevelopment) SetVerbosity(level int) {
  ad.Cfg.VerbosityLevel = level
  TheConfig.VerbosityLevel = level
}

// -------------------------------------------------------------------------------------------------------------------

func (ad *ActiveDevelopment) SetActiveD(isAD bool) {
  ad.Cfg.SetActiveD(isAD)
  TheConfig.SetActiveD(isAD)
}

// -------------------------------------------------------------------------------------------------------------------

func (ad *ActiveDevelopment) SetDoX(isDOX bool) {
  ad.Cfg.SetDoX(isDOX)
  TheConfig.SetDoX(isDOX)
}

// -------------------------------------------------------------------------------------------------------------------

func (ad *ActiveDevelopment) SetLogApis(isLogApis bool) {
  ad.Cfg.SetLogApis(isLogApis)
  TheConfig.SetLogApis(isLogApis)
}
