package utils

// Function to get standardised error response
func GetErrorResponse(reason string, err error) map[string]any {
	if err == nil {
		return map[string]interface{}{
			"message": reason,
			"error":   "An error occured",
		}
	}

	return map[string]interface{}{
		"reason": reason,
		"error":  err,
	}
}

// Function to get standardised success response
func GetSuccessResponse(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}
