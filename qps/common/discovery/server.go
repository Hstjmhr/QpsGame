package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Weight  int    `json:"weight"`
	Version string `json:"version"`
	Ttl     int64  `json:"ttl"`
}

func (s Server) BuildRegisterKey() string {
	if len(s.Version) == 0 {
		return fmt.Sprintf("/%s/%s", s.Name, s.Addr)
	}
	return fmt.Sprintf("/%s/%s/%s", s.Name, s.Version, s.Addr)
}

func ParseValue(v []byte) (Server, error) {
	var server Server
	if err := json.Unmarshal(v, &server); err != nil {
		return server, err
	}

	return server, nil
}

func ParseKey(key string) (Server, error) {
	strs := strings.Split(key, "/")
	if len(strs) == 2 {
		return Server{
			Name:    strs[0],
			Addr:    strs[2],
			Version: strs[1],
		}, nil
	}
	return Server{}, errors.New("invalid key")
}
