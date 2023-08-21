package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) FormatJSON(text string) (string, error) {
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, []byte(text), "", "    "); err != nil {
		return "", fmt.Errorf("unable to format JSON: %w", err)
	}

	return prettyJson.String(), nil
}

func (a *App) FormatTOML(text string) (string, error) {
	var v interface{}
	if err := toml.Unmarshal([]byte(text), &v); err != nil {
		return "", fmt.Errorf("invalid TOML: %w", err)
	}
	out, err := toml.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("unable to marshall TOML: %w", err)
	}

	return string(out), nil
}

func (a *App) FormatYAML(text string) (string, error) {
	var v interface{}
	if err := yaml.Unmarshal([]byte(text), &v); err != nil {
		return "", fmt.Errorf("invalid YAML: %w", err)
	}
	out, err := yaml.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("unable to marshall YAML: %w", err)
	}

	return string(out), nil
}

func (a *App) MinifyJSON(text string) (string, error) {
	var minifyJson bytes.Buffer
	if err := json.Compact(&minifyJson, []byte(text)); err != nil {
		return "", fmt.Errorf("unable to pretty JSON: %w", err)
	}

	return minifyJson.String(), nil
}

func (a *App) Base2(text string) (string, error) {
	return fmt.Sprintf("2 called with %s", text), nil
}

func (a *App) Base8(text string) (string, error) {
	return fmt.Sprintf("8 called with %s", text), nil
}

func (a *App) Base10(text string) (string, error) {
	return fmt.Sprintf("10 called with %s", text), nil
}

func (a *App) Base16(text string) (string, error) {
	return fmt.Sprintf("16 called with %s", text), nil
}
