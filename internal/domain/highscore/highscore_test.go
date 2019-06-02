package highscore

import (
	"testing"
	"time"
)

func TestNewHighScore(t *testing.T) {
	t.Parallel()
	nickname := "Jim"
	acc := NewHighScore(nickname, time.Now())
	if nickname != acc.Nickname {
		t.Errorf("expected lookup nickname to equal %s, got %s", nickname, acc.Nickname)
	}
}
