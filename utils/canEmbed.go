package utils

import (
	"net/http"
	"strings"
)

func CanEmbed(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// --- X-Frame-Options ---
	xfo := resp.Header.Get("X-Frame-Options")
	if xfo != "" {
		xfo = strings.ToUpper(xfo)
		if xfo == "DENY" {
			return false
		}
		if xfo == "SAMEORIGIN" {
			return false
		}
		if strings.HasPrefix(xfo, "ALLOW-FROM") {
			return false
		}
	}

	// --- Content-Security-Policy ---
	csp := resp.Header.Get("Content-Security-Policy")
	if csp != "" {
		// frame-ancestors を探す
		parts := strings.Split(csp, ";")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if strings.HasPrefix(p, "frame-ancestors") {
				if strings.Contains(p, "'none'") {
					return false
				}
				if !strings.Contains(p, "*") {
					return false
				}
			}
		}
	}

	return true
}
