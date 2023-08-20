package env

type Env struct {
  Mode string
}

func (self Env) IsProd() bool {
  var res bool
  res = true // assume the worst
  switch self.Mode {
  case "localhost":
    fallthrough
  case "remotedev":
    res = false
  }

  return res
}

func (self Env) IsLocalhost() bool {
  var res bool
  res = true // assume the worst
  switch self.Mode {
  case "prod":
    fallthrough
  case "remotedev":
    res = false
  }

  return res
}
