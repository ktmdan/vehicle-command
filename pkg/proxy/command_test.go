package proxy_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/teslamotors/vehicle-command/pkg/connector/inet"
	"github.com/teslamotors/vehicle-command/pkg/protocol"
	"github.com/teslamotors/vehicle-command/pkg/proxy"
	"github.com/teslamotors/vehicle-command/pkg/vehicle"
)

func TestExtractCommandAction(t *testing.T) {
	ctx := context.Background()
	params := proxy.RequestParameters{
		"volume":        5.0,
		"on":            true,
		"seat_position": 0,
		"level":         2.0,
		// Add more test cases for different commands and parameters
	}

	tests := []struct {
		command      string
		params       proxy.RequestParameters
		expectedFunc func(*vehicle.Vehicle) error
		expected     error
	}{
		{"adjust_volume", params, func(v *vehicle.Vehicle) error { return v.SetVolume(ctx, 0.0) }, nil},
		{"adjust_volume", nil, nil, &protocol.NominalError{Details: fmt.Errorf("missing volume param")}},
		{"remote_boombox", params, nil, proxy.ErrCommandUseRESTAPI},
		{"navigation_gps_request", params, nil, proxy.ErrCommandUseRESTAPI},
		{"navigation_sc_request", params, nil, proxy.ErrCommandUseRESTAPI},
		{"navigation_waypoints_request", params, nil, proxy.ErrCommandUseRESTAPI},
		{"dashcam_save_clip", params, nil, proxy.ErrCommandUseRESTAPI},
		{"take_drivenote", params, nil, proxy.ErrCommandUseRESTAPI},
		{"upcoming_calendar_entries", params, nil, proxy.ErrCommandUseRESTAPI},
		{"remote_auto_steering_wheel_heat_climate_request", params, nil, proxy.ErrCommandUseRESTAPI},
		{"remote_steering_wheel_heat_level_request", params, nil, proxy.ErrCommandUseRESTAPI},
		{"invalid_command", params, nil, &inet.HTTPError{Code: http.StatusBadRequest, Message: "{\"response\":null,\"error\":\"invalid_command\",\"error_description\":\"\"}"}},
	}

	for _, test := range tests {
		action, err := proxy.ExtractCommandAction(ctx, test.command, test.params)

		if errors.Is(err, test.expected) {
			if test.expected != nil && action != nil {

				t.Errorf("Expected error %#v but got action %p for command %#v", test.expected, action, test.command)
			}
		} else if err != nil && err.Error() != test.expected.Error() {
			t.Errorf("Unexpected error for command %s: %v", test.command, err)
		}
	}
}

func TestSunRoofControl(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		state    string
		expected error
	}{
		{"vent", nil},
		{"close", nil},
		{"bogus", errors.New("sun_roof_control state must be 'vent' or 'close'")},
	}
	for _, test := range tests {
		params := proxy.RequestParameters{"state": test.state}
		action, err := proxy.ExtractCommandAction(ctx, "sun_roof_control", params)
		if test.expected == nil {
			if err != nil || action == nil {
				t.Errorf("state %q: expected an action and no error, got action=%p err=%v", test.state, action, err)
			}
		} else if err == nil || err.Error() != test.expected.Error() {
			t.Errorf("state %q: expected error %v, got %v", test.state, test.expected, err)
		}
	}
}
