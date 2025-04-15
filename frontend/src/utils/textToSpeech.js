let voices = []

speechSynthesis.onvoiceschanged = () => {
  voices = speechSynthesis.getVoices()
}

export const speak = (text, lang = 'en-US') => {
  const utterance = new SpeechSynthesisUtterance(text)

  if (voices.length === 0) {
    voices = speechSynthesis.getVoices()
  }

  const voice = voices.find(v => v.lang.startsWith(lang))
  if (voice) utterance.voice = voice

  utterance.lang = lang
  speechSynthesis.speak(utterance)
}
