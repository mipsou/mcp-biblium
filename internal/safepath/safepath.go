/*
 * Copyright (c) 2026 Mipsou <chpujol@gmail.com>
 *
 * Licensed under the EUPL, Version 1.2 or later.
 * You may obtain a copy at:
 * https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12
 */

package safepath

// Resolve validates and resolves a user-provided path segment
// against a trusted root directory.
// Returns the absolute resolved path or an error if the path
// escapes the root.
func Resolve(root, userPath string) (string, error) {
	return "", nil // TDD: stub
}
