package helper

import "fmt"

// ObfuscateErrorWithContext returns a safe error message for the client (skeleton: same as err).
func ObfuscateErrorWithContext(err error, operation string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", operation, err)
}

// ObfuscateError returns a safe error for the client (skeleton: same as err).
func ObfuscateError(err error) error {
	return err
}
