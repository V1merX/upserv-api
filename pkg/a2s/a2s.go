package a2s

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/rumblefrog/go-a2s"
	"github.com/sirupsen/logrus"
)

func Run(ip string, port int) error {
	client, err := a2s.NewClient(
		fmt.Sprintf("%s:%d", ip, port),
		a2s.SetMaxPacketSize(14000), // Some engine does not follow the protocol spec, and may require bigger packet buffer
		a2s.TimeoutOption(3*time.Second),
	)

	defer func() {
		if err = client.Close(); err != nil {
			slog.Error("failed to close client", slog.Any("err", err))
		}
	}()

	var resp QueryResponse

	resp.Info, err = client.QueryInfo()
	if err != nil {
		return err
	}

	resp.Rules, err = client.QueryRules()
	if err != nil {
		return err
	}

	resp.Players, err = client.QueryPlayer()
	if err != nil {
		return err
	}

	logrus.Info(resp)

	return nil
}
