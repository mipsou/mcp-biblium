/*
 * Copyright (c) 2026 Mipsou <chpujol@gmail.com>
 *
 * Licensed under the EUPL, Version 1.2 or later.
 * You may obtain a copy at:
 * https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12
 */

package config

// Config holds all runtime configuration.
type Config struct {
	DataDir       string
	SearchBackend string
	LogLevel      string
	OllamaURL     string
	OllamaModel   string
}

// Load reads configuration from environment variables with defaults.
func Load() (*Config, error) {
	return nil, nil // TDD: stub
}
