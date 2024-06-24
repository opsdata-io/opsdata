package health

import (
	"encoding/json"
	"net/http"

	"github.com/opsdata-io/opsdata/pkg/logging"
	"github.com/opsdata-io/opsdata/pkg/version"
)

var logger = logging.SetupLogging()

// VersionInfo represents the structure of version information.
type VersionInfo struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	BuildTime string `json:"buildTime"`
}

// VersionHandler returns version information as JSON.
func VersionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		versionInfo := VersionInfo{
			Version:   version.Version,
			GitCommit: version.GitCommit,
			BuildTime: version.BuildTime,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(versionInfo); err != nil {
			logger.Error("Failed to encode version info to JSON", err)
			http.Error(w, "Failed to encode version info", http.StatusInternalServerError)
		}
		logger.Debug("Version info is successfully returned")
	}
}
