package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os/exec"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	for {

		// specify terminal commands
		app := "adb"
		arg1 := "shell"
		arg2 := "screencap"
		arg3 := "-p"

		// run adb screencap command
		cmd := exec.Command(app, arg1, arg2, arg3)
		stdout, stdoutErr := cmd.Output()
		check(stdoutErr)

		// decode png
		img, imgErr := png.Decode(bytes.NewReader(stdout))
		check(imgErr)

		// var to store black point coordinates
		values := []int{}

		// iterate through png
		for y := 0; y < img.Bounds().Max.Y; y++ {
			for x := 0; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				// store coordinates if point color is black
				if r == 0 && g == 0 && b == 0 {
					values = append(values, x)
					values = append(values, y)
				}

			}
		}

		// println("Switch to Sketch")
		// println(len(values) - 1)
		// time.Sleep(5 * time.Second)

		// storing vars for computation

		x1 := values[0]
		y1 := values[1]

		x2 := values[len(values)-1]
		y2 := values[len(values)-1]

		xAdj := x2 - x1
		yAdj := y2 - y1
		print("\033[H\033[2J")
		fmt.Printf("xAdj\t%d\nyAdj\t%d\n", xAdj, yAdj)

		cmd2 := exec.Command("adb", "shell", "input", "touchscreen", "swipe", strconv.Itoa(values[len(values)-2]), strconv.Itoa(values[len(values)-1]), strconv.Itoa(values[0]), strconv.Itoa(values[1]), "500")
		_, stdoutErr2 := cmd2.Output()
		check(stdoutErr2)
		//print(stdout2)

	}

}
