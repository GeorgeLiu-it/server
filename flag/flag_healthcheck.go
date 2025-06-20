package flag

import (
	"fmt"
	"net/http"
	"time"
)

func runHealthCheck() error {
	client := &http.Client{
		Timeout: 3 * time.Second, // fast fail
	}

	resp, err := client.Get("http://localhost:8080/api/base/health")
	if err != nil {
		return fmt.Errorf("failed to request health endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unhealthy status code: %d", resp.StatusCode)
	}

	return nil
}
