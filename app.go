package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (a *App) Beautify(text string) (string, error) {
	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, []byte(text), "", "    "); err != nil {
		return "", fmt.Errorf("unable to pretty JSON: %w", err)
	}

	return prettyJson.String(), nil
}

func (a *App) Minify(text string) (string, error) {
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
