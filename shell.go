package shell

import (
  "os"
  "fmt"
  "github.com/fatih/color"
)

func arg(a ...interface{}) string {
  if len(a) == 1 {
    return fmt.Sprintln(a[0])
  } else {
    return fmt.Sprintf(a[0].(string), a[1:]...)
  }
}

func Say(a ...interface{}) {
  fmt.Print(arg(a...))
}

func Warn(a ...interface{}) {
  color.Yellow(arg(a...))
}

// to stderr
func Error(a ...interface{}) {
  var msg string
  if len(a) == 1 {
    switch v := a[0].(type) {
    case string:
      msg = v
    case error:
      msg = v.Error()
    default:
      panic("unkown data type")
    }
  } else {
    msg = arg(a...)
  }
  if os.Getenv("DEBUG") != "" {
    panic(msg)
  } else {
    fmt.Fprintln(os.Stderr, color.RedString(msg))
  }
}

func ErrorExit(a ...interface{}) {
  Error(a...)
  os.Exit(1)
}

func Debug(a ...interface{}) {
  Say(a...)
}
