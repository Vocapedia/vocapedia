package speechtotext

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	whisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/wav"
	"github.com/go-chi/render"
)

func SpeechToText(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "Ses dosyası gelmedi", http.StatusBadRequest)
		return
	}
	defer file.Close()

	audioData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Dosya okunamadı", http.StatusInternalServerError)
		return
	}

	samples, err := wavToFloat32Samples(audioData)
	if err != nil {
		http.Error(w, "Ses verisi işlenemedi", http.StatusInternalServerError)
		return
	}
	modelpath := "models/ggml-small.bin"

	model, err := whisper.New(modelpath)
	if err != nil {
		http.Error(w, "Model açılamadı", http.StatusInternalServerError)
		return
	}
	defer model.Close()

	context, err := model.NewContext()
	if err != nil {
		http.Error(w, "Model context oluşturulamadı", http.StatusInternalServerError)
		return
	}
	if err := context.Process(samples, nil, nil, nil); err != nil {
		http.Error(w, "Ses işlenemedi", http.StatusInternalServerError)
		return
	}
	var fullText string

	for {
		segment, err := context.NextSegment()
		if err != nil {
			break
		}
		fullText += segment.Text + " "
	}
	render.JSON(w, r, map[string]string{"text": fullText})
}

func wavToFloat32Samples(data []byte) ([]float32, error) {

	reader := bytes.NewReader(data)
	decoder := wav.NewDecoder(reader)
	if !decoder.IsValidFile() {
		return nil, errors.New("geçersiz wav dosyası")
	}

	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	samples := make([]float32, len(buf.Data))
	for i, v := range buf.Data {
		samples[i] = float32(v) / float32(1<<15)
	}
	return samples, nil
}
