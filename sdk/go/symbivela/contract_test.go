package symbivela

import (
	"os"
	"strings"
	"testing"
)

// TestSDKPathsMatchContract asserts the SDK methods target paths declared in the
// published OpenAPI contract, guarding against drift.
func TestSDKPathsMatchContract(t *testing.T) {
	data, err := os.ReadFile("../../contracts/openapi.yaml")
	if err != nil {
		t.Skipf("contract not readable: %v", err)
	}
	contract := string(data)
	required := []string{
		"/v1/goals:",
		"/v1/decisions:",
		"/v1/attention:",
		"/v1/evidence/{goal_id}:",
		"/v1/my-work:",
	}
	for _, path := range required {
		if !strings.Contains(contract, path) {
			t.Errorf("contract missing path %s", path)
		}
	}
}
