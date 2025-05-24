package ai

import (
	"log"
	"net/http"
	"strings"
)

type IntentPredictor interface {
	PredictIntent(*http.Request) (IntentInfo, error)
}

type IntentInfo struct {
	Intent    string  // e.g., "normal", "bot", "attacker", "scraper"
	RiskScore float64 // 0.0 (harmless) to 1.0 (super sus)
}

type RuleBasedPredictor struct{}

func NewRuleBasedPredictor() *RuleBasedPredictor {
	return &RuleBasedPredictor{}
}

func (r *RuleBasedPredictor) PredictIntent(req *http.Request) (IntentInfo, error) {
	ua := strings.ToLower(req.UserAgent())
	path := req.URL.Path

	log.Println("User Agent details: ", ua)

	switch {
	case strings.Contains(ua, "curl"),
		strings.Contains(ua, "httpie"),
		strings.Contains(ua, "wget"):
		return IntentInfo{Intent: "scripted", RiskScore: 0.8}, nil

	case strings.Contains(ua, "bot"),
		strings.Contains(ua, "crawler"),
		strings.Contains(ua, "spider"),
		strings.Contains(ua, "scan"):
		return IntentInfo{Intent: "bot", RiskScore: 0.9}, nil

	case strings.Contains(ua, "python-requests"),
		strings.Contains(ua, "java"),
		strings.Contains(ua, "go-http-client"):
		return IntentInfo{Intent: "programmatic", RiskScore: 0.7}, nil

	case strings.Contains(path, "/admin"),
		strings.Contains(path, "/wp-"),
		strings.Contains(path, "/login"),
		strings.Contains(path, "/.env"):
		return IntentInfo{Intent: "prober", RiskScore: 0.85}, nil

	default:
		return IntentInfo{Intent: "normal", RiskScore: 0.1}, nil
	}

}
