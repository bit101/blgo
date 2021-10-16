package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func MakeGIF(tool, folder, outFileName string, fps float64) {
	if tool == "convert" {
		ConvertToGIF(folder, outFileName, fps)
	} else if tool == "ffmpeg" {
		FfmpegToGIF(folder, outFileName, fps)
	}

}

// ConvertToGIF converts a folder of pngs into an animated gif using imagemagick convert.
func ConvertToGIF(folder, outFileName string, fps float64) {
	delay := fmt.Sprintf("%f", 1000.0/fps/10.0)
	path := folder + "/*.png"
	cmd := exec.Command("convert", "-delay", delay, "-layers", "Optimize", path, outFileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// FfmpegToGIF converts a folder of pngs into an animated gif using ffmpeg.
func FfmpegToGIF(folder, outFileName string, fps float64) {
	path := folder + "/frame_%04d.png"
	fpsArg := fmt.Sprintf("%d", int(fps))

	paletteCmd := exec.Command("ffmpeg", "-y", "-i", path, "-vf", "palettegen", "palette.png")
	err := paletteCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	outCmd := exec.Command("ffmpeg", "-y", "-framerate", fpsArg, "-i", path, "-i", "palette.png", "-filter_complex", "paletteuse", outFileName)
	err = outCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ConvertToYoutube converts a folder of pngs into a Youtube compatible mp4 video file. Requires ffmpeg.
func ConvertToYoutube(folder, outFileName string, fps int) {
	path := folder + "/frame_%04d.png"
	fpsArg := fmt.Sprintf("%d", fps)

	cmd := exec.Command("ffmpeg", "-framerate", fpsArg, "-i", path, "-s:v", "1280x720",
		"-c:v", "libx264", "-profile:v", "high", "-crf", "20",
		"-pix_fmt", "yuv420p", outFileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ViewImage displays an image using installed image viewer.
func ViewImage(imagePath string) {
	cmd := exec.Command("eog", imagePath)
	if runtime.GOOS == "darwin" {
		cmd = exec.Command("qlmanage", "-p", imagePath)
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// VLC launches vlc to play a video
func VLC(fileName string) {
	cmd := exec.Command("vlc", fileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ParentDir returns the immediated directory name of the current working directory.
func ParentDir() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get directory.")
		os.Exit(1)
	}
	return filepath.Base(wd)
}
