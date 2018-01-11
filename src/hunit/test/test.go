package test

import (
  "time"
)

import (
  "gopkg.in/yaml.v2"
)

// Options
type Options uint32

const (
  OptionNone                          = 0
  OptionDebug                         = 1 << 0
  OptionEntityTrimTrailingWhitespace  = 1 << 1
  OptionInterpolateVariables          = 1 << 2
  OptionDisplayRequests               = 1 << 3
  OptionDisplayResponses              = 1 << 4
  OptionDisplayRequestsOnFailure      = 1 << 5
  OptionDisplayResponsesOnFailure     = 1 << 6
)

// Basic credentials
type BasicCredentials struct {
  Username    string              `yaml:"username"`
  Password    string              `yaml:"password"`
}

// A test request
type Request struct {
  Method      string              `yaml:"method"`
  URL         string              `yaml:"url"`
  Headers     map[string]string   `yaml:"headers"`
  Params      map[string]string   `yaml:"params"`
  Entity      string              `yaml:"entity"`
  Format      string              `yaml:"format"`
  BasicAuth   *BasicCredentials   `yaml:"basic-auth"`
}

// A test response
type Response struct {
  Status      int                 `yaml:"status"`
  Headers     map[string]string   `yaml:"headers"`
  Entity      string              `yaml:"entity"`
  Comparison  Comparison          `yaml:"compare"`
  Format      string              `yaml:"format"`
}

// A test case
type Case struct {
  Id        string                `yaml:"id"`
  Wait      time.Duration         `yaml:"wait"`
  Repeat    int                   `yaml:"repeat"`
  Gendoc    bool                  `yaml:"gendoc"`
  Title     string                `yaml:"title"`
  Comments  string                `yaml:"doc"`
  Params    map[string]string     `yaml:"params"`
  Request   Request               `yaml:"request"`
  Response  Response              `yaml:"response"`
  Vars      yaml.MapSlice         `yaml:"vars"`
}

// Determine if this case is documented or not
func (c Case) Documented() bool {
  return c.Gendoc || c.Title != "" || c.Comments != ""
}
