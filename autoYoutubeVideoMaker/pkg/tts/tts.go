package main

import (
    "context"
    "io/ioutil"
    texttospeech "cloud.google.com/go/texttospeech/apiv1"
    texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func main() {
    ctx := context.Background()

    // Creates a client.
    client, err := texttospeech.NewClient(ctx)
    if err != nil {
        // Handle error
    }

    // Perform the text-to-speech request on the text input with the selected voice parameters and audio file type.
    response, err := client.SynthesizeSpeech(ctx, &texttospeechpb.SynthesizeSpeechRequest{
        Input: &texttospeechpb.SynthesisInput{
            InputSource: &texttospeechpb.SynthesisInput_Text{Text: "Hello, World!"},
        },
        // Build the voice request, select the language code and the SSML voice gender.
        Voice: &texttospeechpb.VoiceSelectionParams{
            LanguageCode: "en-US",
            SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
        },
        // Select the type of audio file you want returned.
        AudioConfig: &texttospeechpb.AudioConfig{
            AudioEncoding: texttospeechpb.AudioEncoding_MP3,
        },
    })
    if err != nil {
        // Handle error
    }

    // The response's AudioContent is binary.
    if err := ioutil.WriteFile("output.mp3", response.AudioContent, 0644); err != nil {
        // Handle error
    }
}
