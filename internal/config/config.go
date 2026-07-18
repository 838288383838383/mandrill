package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Setup    SetupConfig    `toml:"setup"`
	GitHub   GitHubConfig   `toml:"github"`
	Registry RegistryConfig `toml:"registry"`
	Install  InstallConfig  `toml:"install"`
	Theme    ThemeConfig    `toml:"theme"`
	TUI      TUIConfig      `toml:"tui"`
}

type SetupConfig struct {
	SetupComplete bool   `toml:"setup_complete"`
	FirstRunAt    string `toml:"first_run_at"`
}

type GitHubConfig struct {
	Enabled      bool   `toml:"enabled"`
	PairWithGh   bool   `toml:"pair_with_gh"`
	AuthMethod   string `toml:"auth_method"`
	TokenSource  string `toml:"token_source"`
	Username     string `toml:"username"`
}

type RegistryConfig struct {
	Default   string   `toml:"default"`
	CustomTaps []string `toml:"custom_taps"`
}

type InstallConfig struct {
	Prefix      string `toml:"prefix"`
	AutoAddPath bool   `toml:"auto_add_path"`
}

type ThemeConfig struct {
	Name   string            `toml:"name"`
	Accent string            `toml:"accent"`
	Custom map[string]string `toml:"custom"`
}

type TUIConfig struct {
	ScreensaverEnabled bool `toml:"screensaver_enabled"`
	ScreensaverTimeout int  `toml:"screensaver_timeout_secs"`
}

func MandrillDir() string {
	home, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "darwin":
		appSupport := filepath.Join(home, "Library", "Application Support", "mandrill")
		if _, err := os.Stat(appSupport); err == nil {
			return appSupport
		}
	case "windows":
		appData := os.Getenv("LOCALAPPDATA")
		if appData != "" {
			return filepath.Join(appData, "mandrill")
		}
		return filepath.Join(home, "AppData", "Local", "mandrill")
	}

	return filepath.Join(home, ".mandrill")
}

func ConfigPath() string {
	return filepath.Join(MandrillDir(), "config.toml")
}

func EnsureMandrillDir() error {
	dir := MandrillDir()
	subdirs := []string{"bin", "packages", "registry"}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	for _, sub := range subdirs {
		if err := os.MkdirAll(filepath.Join(dir, sub), 0755); err != nil {
			return err
		}
	}

	if runtime.GOOS == "windows" {
		binDir := filepath.Join(dir, "bin")
		_ = os.MkdirAll(binDir, 0755)
	}

	return nil
}

func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	prefix := filepath.Join(home, ".mandrill", "bin")

	switch runtime.GOOS {
	case "darwin":
		prefix = filepath.Join(home, ".mandrill", "bin")
	case "windows":
		appData := os.Getenv("LOCALAPPDATA")
		if appData == "" {
			appData = filepath.Join(home, "AppData", "Local")
		}
		prefix = filepath.Join(appData, "mandrill", "bin")
	}

	return &Config{
		Setup: SetupConfig{
			SetupComplete: false,
		},
		GitHub: GitHubConfig{
			Enabled:     true,
			PairWithGh:  true,
			AuthMethod:  "gh-cli",
			TokenSource: "gh",
		},
		Registry: RegistryConfig{
			Default: "mandrill/mandrill-registry",
		},
		Install: InstallConfig{
			Prefix:      prefix,
			AutoAddPath: true,
		},
		Theme: ThemeConfig{
			Name:   "catppuccin-mocha",
			Accent: "mauve",
		},
		TUI: TUIConfig{
			ScreensaverEnabled: true,
			ScreensaverTimeout: 30,
		},
	}
}

func Load() (*Config, error) {
	path := ConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func Save(cfg *Config) error {
	path := ConfigPath()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := toml.NewEncoder(f)
	return encoder.Encode(cfg)
}
