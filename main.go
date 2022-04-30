package main

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/micmonay/keybd_event"
	"noidle/win32"
)

const (
	RDP_PROCESS_NAME    = ""
	RDP_WINDOW_NAME     = ""
	RDP_CLASS_NAME      = ""
	CITRIX_PROCESS_NAME = "CDViewer"
	CITRIX_WINDOW_NAME  = "IL-VDI-PO - Desktop Viewer"
	CITRIX_CLASS_NAME   = "WindowsForms10.Window.8.app.0.2b89eaa_r7_ad1"
)

const (
	INTERVAL = 60 // In seconds
)

var (
	countSimulatedEvents = 0
)

func printWelcome() {
	fmt.Println()
	fmt.Println(color.InGray("\t███╗   ██╗ ██████╗       ██╗██████╗ ██╗     ███████╗"))
	fmt.Println(color.InGray("\t████╗  ██║██╔═══██╗      ██║██╔══██╗██║     ██╔════╝"))
	fmt.Println(color.InGray("\t██╔██╗ ██║██║   ██║█████╗██║██║  ██║██║     █████╗ "))
	fmt.Println(color.InGray("\t██║╚██╗██║██║   ██║╚════╝██║██║  ██║██║     ██╔══╝"))
	fmt.Println(color.InGray("\t██║ ╚████║╚██████╔╝      ██║██████╔╝███████╗███████╗"))
	fmt.Println(color.InGray("\t╚═╝  ╚═══╝ ╚═════╝       ╚═╝╚═════╝ ╚══════╝╚══════╝"))
	fmt.Println(color.InBold("\twindows event generator v0.1"))
	fmt.Println()
}

func simulateKeyEvent() {
	keyboard, _ := keybd_event.NewKeyBonding()
	keyboard.SetKeys(keybd_event.VK_SCROLLLOCK)

	// Toggle Scroll Lock
	keyboard.Press()
	time.Sleep(5 * time.Millisecond)
	keyboard.Release()

	keyboard.Press()
	time.Sleep(5 * time.Millisecond)
	keyboard.Release()
}

func main() {
	printWelcome()
	startTime := time.Now()
	fmt.Println(color.InWhite("==> Timestamp " + startTime.Format(time.RFC1123)))
	fmt.Printf(color.InYellow("==> Time interval is set to %d seconds\n"), INTERVAL)

	for {
		originalWindowHandle := win32.GetForegroundWindow()

		/**
		If citrix is running then simulate targeted event
		Otherwise simulate global event
		*/
		if windowHandle := win32.FindWindow(CITRIX_CLASS_NAME, CITRIX_WINDOW_NAME); windowHandle != 0 {
			fmt.Println(color.InGreen("==> found Citrix process"))
			fmt.Printf("\tWindow handle: 0x%X\n", windowHandle)

			win32.SetForegroundWindow(windowHandle)
			simulateKeyEvent()
			win32.SetForegroundWindow(originalWindowHandle)
			countSimulatedEvents++

		} else if windowHandle := win32.FindWindow(RDP_CLASS_NAME, RDP_WINDOW_NAME); windowHandle != 0 {
			fmt.Println(color.InGreen("==> found RDP process"))
			fmt.Printf("\tWindow handle: 0x%X\n", windowHandle)

			// Simulate global keyboard event
			simulateKeyEvent()
			countSimulatedEvents++
		} else {
			fmt.Println(color.InRed("==> Unable to locate Citrix/RDP"))

			// Simulate global keyboard event
			simulateKeyEvent()
			countSimulatedEvents++
		}
		fmt.Printf("\tGenerated %d events\n", countSimulatedEvents)
		fmt.Printf("\tRunning for %d minutes\n", int(time.Since(startTime).Minutes()))

		time.Sleep(INTERVAL * time.Second)
	}
}
