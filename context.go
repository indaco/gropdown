package gropdown

import (
	"context"
)

// Define the context key type.
type contextKey string

// GetConfigMapFromContext retrieves the tab configuration from the context.
func GetConfigMapFromContext(ctx context.Context) *ConfigMap {
	if opts, ok := ctx.Value(ConfigContextKey).(*ConfigMap); ok {
		return opts
	}
	return &ConfigMap{
		Data: make(map[string]Config),
	}
}

// GetConfigFromContextById retrieves the tab configuration from the context.
func GetConfigFromContextById(ctx context.Context, id string) *Config {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return &config
	}
	return &Config{}
}

// getOpenFromContextById retrieves the configured tab position from the context.
func getOpenFromContextById(ctx context.Context, id string) Position {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return config.Position
	}
	return DefaultConfig().Position
}

// getOpenAsStringFromContextById retrieves the configured tab position from the context.
func getOpenAsStringFromContextById(ctx context.Context, id string) string {
	return getOpenFromContextById(ctx, id).String()
}

// getAnimationFromContextById retrieves the configured tab position from the context.
func getAnimationFromContextById(ctx context.Context, id string) Position {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return config.Position
	}
	return DefaultConfig().Position
}

// getAnimationAsStringFromContextById retrieves the configured tab position from the context.
func getAnimationAsStringFromContextById(ctx context.Context, id string) string {
	return getAnimationFromContextById(ctx, id).String()
}

// getPositionFromContextById retrieves the configured tab position from the context.
func getPositionFromContextById(ctx context.Context, id string) Position {
	if config, ok := GetConfigMapFromContext(ctx).Get(id); ok {
		return config.Position
	}
	return DefaultConfig().Position
}

// getPositionAsStringFromContextById retrieves the configured tab position from the context.
func getPositionAsStringFromContextById(ctx context.Context, id string) string {
	return getPositionFromContextById(ctx, id).String()
}
