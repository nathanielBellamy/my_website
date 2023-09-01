package env

type Env int

func IsProd(mode string) bool {
  var res bool
  res = true // assume the worst
  switch mode {
  case "localhost":
    fallthrough
  case "remotedev":
    res = false
  }

  return res
}

func IsLocalhost(mode string) bool {
  var res bool
  res = true // assume the worst
  switch mode {
  case "prod":
    fallthrough
  case "remotedev":
    res = false
  }

  return res
}


func IsRemotedev(mode string) bool {
  var res bool
  res = true // assume the worst
  switch mode {
  case "prod":
    fallthrough
  case "localhost":
    res = false
  }

  return res
}
