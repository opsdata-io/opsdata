package version

import (
	"encoding/json"
	"log"
	"net/http"
)

// Version is the version of the application
var Version = "v0.0.0"

// GitCommit is the git commit hash of the application
var GitCommit = "MISSING GIT COMMIT"

// BuildTime is the time the application was built
var BuildTime = "MISSING BUILD TIME"

// VersionInfo holds the version information
type info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildTime string `json:"build_time"`
}

// VersionResponse holds the version information
var VersionResponse = info{
	Version:   Version,
	GitCommit: GitCommit,
	BuildTime: BuildTime,
}

// GetVersion returns the application version information
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(VersionResponse); err != nil {
		http.Error(w, "Failed to encode version response", http.StatusInternalServerError)
		log.Printf("Failed to encode version response: %v", err)
	}
}
