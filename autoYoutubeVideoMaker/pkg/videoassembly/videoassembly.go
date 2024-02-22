package videoassembly

import (
    "github.com/u2takey/ffmpeg-go"
    "log"
    "strconv"
)

// CreateVideo combines audio and images into a video file.
func CreateVideo(audioPath string, images []string, outputPath string, imageDuration int) error {
    // Build the complex filter string to concatenate images and set durations
    var filterComplex string
    for i, image := range images {
        filterComplex += ffmpeg-go.FilterInput(image).Filter("scale", ffmpeg-go.Args{"640:-1"}).Filter("setsar", ffmpeg-go.Args{"1"}).Get("v")
        if i != len(images)-1 { // No need to add duration for the last image
            filterComplex += ";"
        }
    }

    // Concatenate all inputs and add audio
    filterComplex += "concat=n=" + strconv.Itoa(len(images)) + ":v=1:a=0 [v]"

    // Execute FFmpeg command with ffmpeg-go
    err := ffmpeg-go.Input("").
        Input(audioPath).
        FilterComplex(filterComplex).
        Output(outputPath, ffmpeg-go.KwArgs{
            "map":  "[v]",
            "map":  "1:a",
            "c:v":  "libx264",
            "c:a":  "aac",
            "vsync": "vfr",
            "pix_fmt": "yuv420p",
        }).Run()

    if err != nil {
        log.Printf("Failed to create video: %v", err)
        return err
    }

    return nil
}
