package main

import (
    "log"
    "os"
    "github.com/urfave/cli/v2"
    "os/exec"
    "fmt"
    "bufio"
    "time"
)

func convCommand(c *cli.Context) error {
    startTime := time.Now()  // Add this line to capture the start time

    if c.Args().Len() < 2 {
        return fmt.Errorf("insufficient arguments")
    }
    inputPath := c.Args().Get(0)
    outputPath := c.Args().Get(1)

    cmd := exec.Command("ffmpeg", "-i", inputPath, "-y", "-acodec", "libmp3lame", outputPath)
    if err := cmd.Start(); err != nil {
        return fmt.Errorf("error starting command: %w", err)
    }

    writer := bufio.NewWriter(os.Stdout)
    go func() {
        squareChars := []rune{'▖', '▘', '▝', '▗'}
        for {
            for _, r := range squareChars {
                fmt.Fprintf(writer, "\r%c", r)
                writer.Flush()
                time.Sleep(100 * time.Millisecond)
            }
        }
    }()

    if err := cmd.Wait(); err != nil {
        return fmt.Errorf("ffmpeg error: %w", err)
    }

    elapsed := time.Since(startTime)
    fmt.Printf("\nFinished! Took %dms.\n", elapsed.Milliseconds())
    return nil
}

func main() {
    app := &cli.App{
        Name:  "rtconv",
        Usage: "Convert anything to everything.",
        Commands: []*cli.Command{
            {
                Name:   "conv",
                Usage:  "Convert video to audio.",
                Action: convCommand,
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}