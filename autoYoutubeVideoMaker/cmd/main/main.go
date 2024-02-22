package main

import (
	"fmt"
	"log"
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/TardisDent/youtuberBot/pkg/wikipedia"
	"github.com/TardisDent/youtuberBot/pkg/pdfreader"
	"github.com/TardisDent/youtuberBot/pkg/textprocessor"
	"github.com/TardisDent/youtuberBot/pkg/tts"
	"github.com/TardisDent/youtuberBot/pkg/videoassembly"
	"github.com/TardisDent/youtuberBot/pkg/youtubeuploader"
	
)

func main() {
	mt.Print("Enter the title of the Wikipedia page you'd like to fetch: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input:", err)
	}
	title = strings.TrimSpace(title) // Trim newline and spaces

	article, err := wikipedia.GetArticle(title)
	if err != nil {
		log.Fatalf("Failed to fetch article for title %s: %v", title, err)

	// Initialize the program's state.
	log.Println("Initializing the program's state...")

	// Call other functions to perform the desired tasks.
	article, err := wikipedia.GetArticle("The Meaning of Life")
	if err != nil {
		log.Fatal(err)
	}

	pdf, err := pdfreader.ReadPDF("The Meaning of Life.pdf")
	if err != nil {
		log.Fatal(err)
	}

	text, err := textprocessor.ProcessText(article, pdf)
	if err != nil {
		log.Fatal(err)
	}

	audio, err := tts.SynthesizeSpeech(text)
	if err != nil {
		log.Fatal(err)
	}

	video, err := videoassembly.AssembleVideo(audio, text)
	if err != nil {
		log.Fatal(err)
	}

	err = youtubeuploader.UploadVideo(video, "The Meaning of Life")
	if err != nil {
		log.Fatal(err)
	}

	// Print the results of the program's execution.
	log.Println("The program has completed successfully.")
}
