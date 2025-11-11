package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem occured while parsing league, %v", err)
	}
	return league, err
}
