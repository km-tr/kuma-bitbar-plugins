package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	var icon string
	tellSpotify := "tell application \"Spotify\" "
	args := os.Args

	if len(args) > 1 && args[1] == "launch" {
		exec.Command("osascript", "-e", tellSpotify+"to activate").Run()
		os.Exit(0)
	}

	running, _ := exec.Command("osascript", "-e", "application \"Spotify\" is running").Output()
	isRunning := string(running) == "true\n"

	if !isRunning {
		fmt.Println("♫")
		fmt.Println("---")
		fmt.Println("Spotify is not running")
		fmt.Println("Launch Spotify | bash=" + args[0] + " param1=launch terminal=false")
		os.Exit(0)
	}

	track, _ := exec.Command("osascript", "-e", tellSpotify+"to name of current track as string").Output()
	artist, _ := exec.Command("osascript", "-e", tellSpotify+"to artist of current track as string").Output()
	state, _ := exec.Command("osascript", "-e", tellSpotify+"to player state as string").Output()

	if string(state) == "playing\n" {
		icon = "▶"
	} else {
		icon = "❚❚"
	}

	rep := regexp.MustCompile(`\n`)
	text := icon + " " + string(track) + " - " + string(artist)
	fmt.Println(rep.ReplaceAllString(text, ""))
}
