/*
 * Copyright (c) 2026 Mipsou <chpujol@gmail.com>
 *
 * Licensed under the EUPL, Version 1.2 or later.
 * You may obtain a copy at:
 * https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12
 */

package mcpserver

import (
	"github.com/mark3labs/mcp-go/server"
)

// Server wraps the MCP server with Lore-specific configuration.
type Server struct {
	mcp *server.MCPServer
}

// New creates a new Lore MCP server (stub — no tools registered yet).
func New() *Server {
	s := server.NewMCPServer(
		"lore",
		"0.1.0",
		server.WithToolCapabilities(false),
	)
	return &Server{mcp: s}
}

// MCPServer returns the underlying mcp-go server for wiring.
func (s *Server) MCPServer() *server.MCPServer {
	return s.mcp
}
