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

// output to stderr
func Say2(a ...interface{}) {
  fmt.Fprintf(os.Stderr, arg(a...))
}

func Warn(a ...interface{}) {
  Say2(color.YellowString(arg(a...)))
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
    Say2(color.RedString(msg))
  }
}

// with exit 1
func ErrorExit(a ...interface{}) {
  Error(a...)
  os.Exit(1)
}

func Debug(a ...interface{}) {
  if len(os.Getenv("DEBUG")) > 0 {
    Say2(a...)
  }
}
