package hunit

import (
  "fmt"
)

/**
 * Comparison types
 */
type Comparison uint32

const (
  CompareLiteral  Comparison  = iota
  CompareSemantic Comparison  = iota
)

var comparisonNames = []string{
  "literal",
  "values",
}

/**
 * Stringer
 */
func (c Comparison) String() string {
  if c < 0 || c > CompareSemantic {
    return "<invalid>"
  }else{
    return comparisonNames[int(c)]
  }
}

/**
 * Marshal
 */
func (c Comparison) MarshalYAML() (interface{}, error) {
  return c.String(), nil
}

/**
 * Unmarshal
 */
func (c *Comparison) UnmarshalYAML(unmarshal func(interface{}) error) error {
  var s string
  err := unmarshal(&s)
  if err != nil {
    return err
  }
  switch s {
    case "literal", "":
      *c = CompareLiteral
    case "smart":
      *c = CompareSemantic
    default:
      return fmt.Errorf("Unsupported comparison type: %v", s)
  }
  return nil
}
