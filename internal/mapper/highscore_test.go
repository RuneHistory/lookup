package mapper

import (
	"lookup/internal/domain/highscore"
	"testing"
	"time"
)

func TestAccountToHttpV1(t *testing.T) {
	t.Parallel()
	acc := highscore.NewHighScore("Jim", time.Now())
	mapped := HighScoreToHttpV1(acc)
	if acc.Nickname != mapped.Nickname {
		t.Errorf("Expecting %s, got %s", acc.Nickname, mapped.Nickname)
	}
}
