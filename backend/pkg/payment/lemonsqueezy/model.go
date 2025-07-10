package lemonsqueezy

import "net/http"

type LemonSqueezyProvider struct {
	apiKey     string
	httpClient *http.Client
}
