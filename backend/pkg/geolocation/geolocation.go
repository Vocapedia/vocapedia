package geolocation

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// GeoLocationService handles IP geolocation
type GeoLocationService struct {
	client *http.Client
}

// GeoLocationResponse represents the response from geolocation service
type GeoLocationResponse struct {
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	IP          string `json:"ip"`
	Error       string `json:"error,omitempty"`
}

// NewGeoLocationService creates a new geolocation service
func NewGeoLocationService() *GeoLocationService {
	return &GeoLocationService{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// GetCountryFromIP returns the country code from IP address with fallback strategies
func (g *GeoLocationService) GetCountryFromIP(ctx context.Context, ip string) (string, error) {
	// Skip private IPs
	if isPrivateIP(ip) {
		return "TR", nil // Default to Turkey for local development
	}

	// Try multiple geolocation services for reliability
	services := []string{
		"https://ipapi.co/%s/json/",
		"https://api.ipgeolocation.io/ipgeo?apiKey=free&ip=%s",
		"https://ipinfo.io/%s/json",
	}

	for _, serviceURL := range services {
		url := fmt.Sprintf(serviceURL, ip)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			continue
		}

		resp, err := g.client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			continue
		}

		var geoResp GeoLocationResponse
		if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
			continue
		}

		if geoResp.Error != "" || geoResp.CountryCode == "" {
			continue
		}

		// Validate country code format
		if len(geoResp.CountryCode) == 2 {
			return strings.ToUpper(geoResp.CountryCode), nil
		}
	}

	// All services failed, return default
	return "TR", fmt.Errorf("all geolocation services failed for IP: %s", ip)
}

// GetRealIP extracts the real IP address from request
func GetRealIP(r *http.Request) string {
	// Check X-Forwarded-For header
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Check CF-Connecting-IP header (Cloudflare)
	if cfip := r.Header.Get("CF-Connecting-IP"); cfip != "" {
		return cfip
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// isPrivateIP checks if an IP address is private
func isPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// Check for private IP ranges
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
	}

	for _, cidr := range privateRanges {
		_, network, _ := net.ParseCIDR(cidr)
		if network.Contains(parsedIP) {
			return true
		}
	}

	return false
}

// PaymentProviderConfig holds provider configuration
type PaymentProviderConfig struct {
	Provider            string
	Currency            string
	Confidence          int    // 0-100, how confident we are about this choice
	Reason              string // Why this provider was chosen
	AlternativeProvider string // Alternative provider
	AlternativeCurrency string
}

// GetPaymentProviderForCountry returns the appropriate payment provider with smart fallbacks
func GetPaymentProviderForCountry(countryCode string, acceptLanguage string, timezone string) *PaymentProviderConfig {
	countryCode = strings.ToUpper(countryCode)

	// High confidence Turkish detection
	if countryCode == "TR" {
		return &PaymentProviderConfig{
			Provider:            "iyzico",
			Currency:            "TRY",
			Confidence:          90,
			Reason:              "Turkey IP detected",
			AlternativeProvider: "paypal",
			AlternativeCurrency: "USD",
		}
	}

	// Check browser language for additional context
	if strings.Contains(strings.ToLower(acceptLanguage), "tr") &&
		(countryCode == "" || countryCode == "TR") {
		return &PaymentProviderConfig{
			Provider:            "iyzico",
			Currency:            "TRY",
			Confidence:          70,
			Reason:              "Turkish language preference detected",
			AlternativeProvider: "paypal",
			AlternativeCurrency: "USD",
		}
	}

	// European countries - might prefer EUR but we use USD for simplicity
	europeanCountries := map[string]bool{
		"DE": true, "FR": true, "IT": true, "ES": true, "NL": true,
		"BE": true, "AT": true, "PT": true, "GR": true, "IE": true,
		"FI": true, "LU": true, "MT": true, "CY": true, "EE": true,
		"LV": true, "LT": true, "SI": true, "SK": true, "HR": true,
		"BG": true, "RO": true, "HU": true, "PL": true, "CZ": true,
		"DK": true, "SE": true, "NO": true, "IS": true, "CH": true,
	}

	if europeanCountries[countryCode] {
		return &PaymentProviderConfig{
			Provider:            "paypal",
			Currency:            "USD",
			Confidence:          85,
			Reason:              fmt.Sprintf("European country (%s) detected", countryCode),
			AlternativeProvider: "iyzico",
			AlternativeCurrency: "TRY",
		}
	}

	// Default fallback - always provide both options
	return &PaymentProviderConfig{
		Provider:            "paypal",
		Currency:            "USD",
		Confidence:          50,
		Reason:              fmt.Sprintf("Default fallback for country: %s", countryCode),
		AlternativeProvider: "iyzico",
		AlternativeCurrency: "TRY",
	}
}

// GetPaymentProviderForCountry maintains backward compatibility
func GetPaymentProviderForCountryOld(countryCode string) (string, string) {
	config := GetPaymentProviderForCountry(countryCode, "", "")
	return config.Provider, config.Currency
}
