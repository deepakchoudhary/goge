package extras

/*
// Works also for 64 bits
#ifdef _WIN32

// Lib for console management in windows
#include "conio.h"

#else

// Libs terminal management in Unix, Linux...
#include <stdio.h>
#include <unistd.h>
#include <termios.h>

// Implement reading a key pressed in terminal
char getch(){
    char ch = 0;
    struct termios old = {0};
    fflush(stdout);
    if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
    old.c_lflag &= ~ICANON;
    old.c_lflag &= ~ECHO;
    old.c_cc[VMIN] = 1;
    old.c_cc[VTIME] = 0;
    if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
    if( read(0, &ch,1) < 0 ) perror("read()");
    old.c_lflag |= ICANON;
    old.c_lflag |= ECHO;
    if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
    return ch;
}
#endif
*/
//import "C"
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func Times(str string, length int) string {
	stop := true
	counter := 0
	returnString := ""

	for stop {
		counter += 1
		if counter == length {
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
	cmd := exec.Command("cmd", "/c", "cls")
	if IsWindows() {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Delay(millis int) {
	time.Sleep(time.Millisecond * time.Duration(millis))
}

func GetInput() string {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	//return string(byte(C.getch()))
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	return string(b)
}

func Flip(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
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
