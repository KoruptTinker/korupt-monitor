package korupt_monitor_server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KoruptTinker/korupt-monitor/config"
	httpClient "github.com/KoruptTinker/korupt-monitor/internal/core/http_client"
)

func New(config *config.Config) *Service {
	return &Service{
		BaseExternal: httpClient.BaseExternal{
			Hostname: config.External.KoruptMonitorServer.Hostname,
		},
	}
}

func (s *Service) RecordClickData(ctx context.Context, leftClickCount int, rightClickCount int) error {
	reqPayload := RecordClicksRequest{
		LeftClickCount:  leftClickCount,
		RightClickCount: rightClickCount,
	}

	_, err := s.Request(&ctx, fmt.Sprintf("%s/api/v1/clicks", s.Hostname), http.MethodPut, reqPayload, nil)
	if err != nil {
		fmt.Printf("Error transmitting click data: %v", err.Error())
		return err
	}

	return nil
}

func (s *Service) RecordKeypressData(ctx context.Context, keyPressCount int) error {
	reqPayload := RecordKeypressRequest{
		KeypressCount: keyPressCount,
	}

	_, err := s.Request(&ctx, fmt.Sprintf("%s/api/v1/keypresses", s.Hostname), http.MethodPut, reqPayload, nil)
	if err != nil {
		fmt.Printf("Error transmitting keypress data: %v", err.Error())
		return err
	}

	return nil
}
