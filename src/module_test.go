package src

import (
	"testing"
)

func TestModuleName(t *testing.T) {
	if ProjectName() != "redhat-openai-bot" {
		t.Errorf("Project name `%s` incorrect", ProjectName())
	}
}
