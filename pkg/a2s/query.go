package a2s

import "github.com/rumblefrog/go-a2s"

type QueryResponse struct {
	Players *a2s.PlayerInfo
	Rules   *a2s.RulesInfo
	Info    *a2s.ServerInfo
}
