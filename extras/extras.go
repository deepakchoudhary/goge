package extras

import (
  "os"
  "os/exec"
  "runtime"
  "time"
  "fmt"
  "strings"
)

func Times(str string, length int) string {
  stop := true
  counter := 0
  returnString := ""
  
  for (stop) {
    counter += 1
    if (counter == length) {
      stop = false
    }
    returnString += str
  }
  
  return returnString
}

func IsWindows() bool {
  if runtime.GOOS == "windows" {
    return true
  } else {
    return false
  }
}

func ClearScreen() {
  cmd := "cls"
  if (IsWindows()) {
    cmd = "cls"
  } else {
    cmd = "clear"
  }
  
  com := exec.Command(cmd)
  stdout, _ := com.Output()
  os.Stdout.Write(stdout)
}

func Delay(millis int) {
  time.Sleep(time.Millisecond * time.Duration(millis))
}

func InitInput() {
  exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
  exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
}
func StopInput() {
  exec.Command("stty", "-f", "/dev/tty", "echo").Run()
}
func GetInput() string {
  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    return string(b)
  }
}

func Flip(str string) string {
  runes := []rune(str)
  for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  
  returnString := string(runes)
  returnString = __FlipChars__(returnString)
  
  return returnString
}
func __FlipChars__(str string) string {
  replacer := strings.NewReplacer("(", ")", ")", "(", "{", "}", "}", "{", "[", "]", "]", "[", "\\", "/", "/", "\\", "<", ">", ">", "<")
  return replacer.Replace(str)
}

func Exit() {
  os.Exit(0)
}

func Error(text string) {
  fmt.Println("\nerror: " + text)
  os.Exit(1)
}