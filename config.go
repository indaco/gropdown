package gropdown

// Config represents a dropdown menu component.
type Config struct {
	Open                bool      // Open indicates whether the dropdown menu is currently open.
	Placement           Placement // Placement indicates the position of the dropdown content relative to the button.
	Animation           bool      // Animation indicates whether the dropdown button should use animations on open and close.
	CloseOnOutsideClick bool      // CloseOnOutsideClick indicates whether the dropdown should be closed when a click occurs outside of it.
}

// ConfigBuilder is used to construct Dropdown instances with options.
type ConfigBuilder struct {
	config Config
}

type ConfigMap struct {
	Data map[string]Config
}

// DefaultConfig returns the default configuration.
func DefaultConfig() Config {
	return Config{
		Open:                false,
		Placement:           Bottom,
		CloseOnOutsideClick: true,
		Animation:           true,
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

// WithPlacement sets the opening position for the dropdown menu.
func (b *ConfigBuilder) WithPlacement(position Placement) *ConfigBuilder {
	b.config.Placement = position
	return b
}

// WithCloseOnOutsideClick sets whether auto close when a click occurs outside of it..
func (b *ConfigBuilder) WithCloseOnOutsideClick(close bool) *ConfigBuilder {
	b.config.CloseOnOutsideClick = close
	return b
}

// WithAnimation sets the animations for the dropdown button icon when open/close.
func (b *ConfigBuilder) WithAnimation(animation bool) *ConfigBuilder {
	b.config.Animation = animation
	return b
}

func (b *ConfigBuilder) Build() Config {
	return b.config
}

// NewConfigMap initialize a new ConfigMap instance
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
