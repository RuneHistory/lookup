package service

import (
	"bytes"
	"fmt"
	"lookup/internal/domain"
	"lookup/internal/errs"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HighScore interface {
	GetByNickname(nickname string) (*domain.HighScore, error)
}

func NewHighScoreService() HighScore {
	return &HighScoreService{}
}

type HighScoreService struct {
}

func (s *HighScoreService) GetByNickname(nickname string) (*domain.HighScore, error) {
	h := domain.NewHighScore(nickname, time.Now())
	data, err := getHighScoreData(nickname)
	if err != nil {
		return nil, err
	}
	rows := strings.Split(data, "\n")

	orderedSkills := domain.OrderedSkills()

	for i, name := range orderedSkills {
		cols := strings.Split(rows[i], ",")

		rank, err := strconv.Atoi(cols[0])
		if err != nil {
			return nil, err
		}
		level, err := strconv.Atoi(cols[1])
		if err != nil {
			return nil, err
		}
		experience, err := strconv.Atoi(cols[2])
		if err != nil {
			return nil, err
		}

		h.Skills[i] = domain.NewSkill(name, rank, level, experience)
	}
	return h, nil
}

func getHighScoreData(nickname string) (string, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "http://services.runescape.com/m=hiscore_oldschool/index_lite.ws", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("player", nickname)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyBuffer := new(bytes.Buffer)
	_, err = bodyBuffer.ReadFrom(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode == http.StatusNotFound {
		return "", errs.NotFound(fmt.Sprintf("nickname %s not found on lookup", nickname))
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d error on lookup for %s", res.StatusCode, nickname)
	}

	return bodyBuffer.String(), nil
}
