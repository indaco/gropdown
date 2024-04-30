package gropdown

// Config represents a dropdown menu component.
type Config struct {
	Open      bool     // Open indicates whether the dropdown menu is currently open.
	Animation bool     // Animation indicates whether the dropdown button should use animations on open and close.
	Position  Position // Position indicates the position of the dropdown content relative to the button.
}

// DropdownBuilder is used to construct Dropdown instances with options.
type ConfigBuilder struct {
	config Config
}

type ConfigMap struct {
	Data map[string]Config
}

// DefaultConfig returns the default configuration.
func DefaultConfig() Config {
	return Config{
		Open:      false,
		Animation: true,
		Position:  Bottom,
	}
}

// NewConfigBuilder creates a new ConfigBuilder instance with default settings.
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: DefaultConfig(),
	}
}

// WithOpen sets the Open field of the dropdown.
func (b *ConfigBuilder) WithOpen(open bool) *ConfigBuilder {
	b.config.Open = open
	return b
}

// WithAnimation sets the animations for the dropdown button icon when open/close.
func (b *ConfigBuilder) WithAnimation(animation bool) *ConfigBuilder {
	b.config.Animation = animation
	return b
}

func (b *ConfigBuilder) WithPosition(position Position) *ConfigBuilder {
	b.config.Position = position
	return b
}

func (b *ConfigBuilder) Build() Config {
	return b.config
}

// Initialize a new ConfigMap instance
func NewConfigMap() *ConfigMap {
	return &ConfigMap{
		Data: make(map[string]Config),
	}
}

// Add adds a configuration to the map.
func (m *ConfigMap) Add(key string, config Config) {
	m.Data[key] = config
}

// Get retrieves a configuration from the map.
func (m *ConfigMap) Get(key string) (Config, bool) {
	config, ok := m.Data[key]
	return config, ok
}
