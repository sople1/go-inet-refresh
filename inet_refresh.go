/**
 * go wininet.dll refresh Internet Setting for proxy
 * using wininet.dll
 *
 * you may can use InternetSetOptionA rather then InternetSetOptionW
 *
 * @author SnooeyNET <sople1@snooey.net>
 */
package main

import (
	"fmt"
	"syscall"
)

const INTERNET_OPTION_REFRESH = 37
const INTERNET_OPTION_SETTINGS_CHANGED = 39

func main() {
	var inet     = syscall.NewLazyDLL("wininet.dll")
	var inetSetOption = inet.NewProc("InternetSetOptionW")

	inetSetOption.Call(0, INTERNET_OPTION_REFRESH, uintptr(0), 0)            // InternetSetOption(0, INTERNET_OPTION_REFRESH, IntPtr.Zero, 0);
	inetSetOption.Call(0, INTERNET_OPTION_SETTINGS_CHANGED, uintptr(0), 0)   // InternetSetOption(0, INTERNET_OPTION_SETTINGS_CHANGED, IntPtr.Zero, 0);

	fmt.Printf("InternetSetOption refresh launched.")

}
