package ad

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

// -------------------------------------------------------------------------------------------------------------------

type Config struct {
	IsActiveDevelopment   bool
	IsShowDebug           bool
	IsLogApis             bool
	VerbosityLevel        int
	IsDoX                 bool

	FeatureFlags  map[string]struct{}
	FeatureTags   map[string]string
	FeatureTagsEx map[string]interface{}
}

// -------------------------------------------------------------------------------------------------------------------

func NewConfig() *Config {
  config := Config{}

  //config.IsActiveDevelopment = true
  //
  //// Release
  //config.IsShowDebug = false
  //config.VerbosityLevel = 0
  //config.IsDoX = false
  //config.IsLogApis = false
  //
  //if !config.IsActiveDevelopment {
  //  // Debug / ActiveDevelopment
  //  config.IsShowDebug = true
  //  config.VerbosityLevel = 1
  //  config.IsDoX = true
  //  config.IsLogApis = true
  //}

  config.FeatureFlags  = make(map[string]struct{}, 1)
  config.FeatureTags   = make(map[string]string, 1)
  config.FeatureTagsEx = make(map[string]interface{}, 1)

  config.SetReleaseMode()

  return &config
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) SetReleaseMode() {
  c.SetActiveD(false)
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) SetActiveD(isAD bool) {
  c.IsActiveDevelopment = isAD

  if !c.IsActiveDevelopment {
    // Release
    c.IsShowDebug = false
    c.VerbosityLevel = 0
    c.IsDoX = false
    c.IsLogApis = false

  } else {
    // Debug / ActiveDevelopment
    c.IsShowDebug = true
    c.VerbosityLevel = 1
    c.IsDoX = true
    c.IsLogApis = true
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) SetVerbosity(level int) {
  c.VerbosityLevel = level
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) SetDoX(isDOX bool) {
  c.IsDoX = isDOX
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) SetLogApis(isLogApis bool) {
  c.IsLogApis = isLogApis
}



// -------------------------------------------------------------------------------------------------------------------

func (c *Config) IsActiveD() bool {
  return TheConfig.IsActiveDevelopment
}

//// -------------------------------------------------------------------------------------------------------------------
//
//func (c *Config) ActiveD() bool {
//  return ConfigIsActiveD()
//}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) IsProd() bool {
  return !ActiveD()
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) Verbosity() int {
  return TheConfig.VerbosityLevel
}

//// -------------------------------------------------------------------------------------------------------------------
//
//func (c *Config) Verbosity() int {
//  return TheConfig.VerbosityLevel
//}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) DoX() bool {
  return TheConfig.IsDoX
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) LogApis() bool {
  return TheConfig.IsLogApis
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) If(n int) bool {
  if n == 0 {
    return false
  }
  return true
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) FeatureFlag(flag string) bool {
  _, present := TheConfig.FeatureTags[flag]
  return present
}

// -------------------------------------------------------------------------------------------------------------------

func (c *Config) OnFeatureFlag(flag string, fn func() error) error {
  if FeatureFlag(flag) {
    return fn()
  }
  return nil
}
















// -------------------------------------------------------------------------------------------------------------------

var TheConfig *Config

func init() {
  TheConfig = NewConfig()
	//TheConfig = &Config{}
  //
	//TheConfig.IsActiveDevelopment = true
  //
	//// Release
	//TheConfig.IsShowDebug = false
	//TheConfig.VerbosityLevel = 0
	//TheConfig.IsDoX = false
	//TheConfig.IsLogApis = false
  //
	//if !TheConfig.IsActiveDevelopment {
	//	// Debug / ActiveDevelopment
	//	TheConfig.IsShowDebug = true
	//	TheConfig.VerbosityLevel = 1
	//	TheConfig.IsDoX = true
	//	TheConfig.IsLogApis = true
	//}
}

// -------------------------------------------------------------------------------------------------------------------

func ConfigIsActiveD() bool {
	return TheConfig.IsActiveD()
}

// -------------------------------------------------------------------------------------------------------------------

func ActiveD() bool {
	return ConfigIsActiveD()
}

// -------------------------------------------------------------------------------------------------------------------

func IsProd() bool {
	return !ActiveD()
}

// -------------------------------------------------------------------------------------------------------------------

func ConfigVerbosity() int {
	return TheConfig.Verbosity()
}

// -------------------------------------------------------------------------------------------------------------------

func Verbosity() int {
	return TheConfig.Verbosity()
}

// -------------------------------------------------------------------------------------------------------------------

func DoX() bool {
	return TheConfig.DoX()
}

// -------------------------------------------------------------------------------------------------------------------

func LogApis() bool {
  return TheConfig.LogApis()
}

// -------------------------------------------------------------------------------------------------------------------

func ConfigIf(n int) bool {
	if n == 0 {
		return false
	}
	return true
}

// -------------------------------------------------------------------------------------------------------------------

func FeatureFlag(flag string) bool {
	_, present := TheConfig.FeatureTags[flag]
	return present
}

// -------------------------------------------------------------------------------------------------------------------

func OnFeatureFlag(flag string, fn func() error) error {
	if FeatureFlag(flag) {
		return fn()
	}
	return nil
}
