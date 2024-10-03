package handlers

// ValidSubscriptionStatus defines the list of valid subscription status values
var ValidSubscriptionStatus = []string{"Active", "Inactive", "Pending"}

// isValidSubscriptionStatus checks if the provided subscription status is valid
func isValidSubscriptionStatus(status string) bool {
	for _, validStatus := range ValidSubscriptionStatus {
		if status == validStatus {
			return true
		}
	}
	return false
}
