package main

import (
	"bytes"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
)

func main() {
	authenticator := &core.IamAuthenticator{
		ApiKey: "Z_K7U-xwdpgrQuiTBXj4LPDyBGcaznGXAtkbI3rh-BH3",
	}

	options := &texttospeechv1.TextToSpeechV1Options{
		Authenticator: authenticator,
	}

	textToSpeech, textToSpeechErr := texttospeechv1.NewTextToSpeechV1(options)

	if textToSpeechErr != nil {
		panic(textToSpeechErr)
	}

	textToSpeech.SetServiceURL("https://api.eu-gb.text-to-speech.watson.cloud.ibm.com/instances/53d93b89-c456-43fe-ad88-7b80206c5c30")

	result, _, responseErr := textToSpeech.Synthesize(
		&texttospeechv1.SynthesizeOptions{
			Text:   core.StringPtr("Hello World, I'm a bot"),
			Accept: core.StringPtr("audio/wav"),
			Voice:  core.StringPtr(texttospeechv1.SynthesizeOptions_Voice_EnUsKevinv3voice),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}
	if result != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(result)
		file, _ := os.Create("hello_world.wav")
		file.Write(buff.Bytes())
		file.Close()
	}
}
