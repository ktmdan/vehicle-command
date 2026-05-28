package main

import (
	"errors"
	"strconv"
	"testing"

	"github.com/teslamotors/vehicle-command/pkg/vehicle"
)

func TestMinutesAfterMidnight(t *testing.T) {
	type params struct {
		str     string
		minutes int32
		err     error
	}
	testCases := []params{
		{str: "3:03", minutes: 183},
		{str: "0:00", minutes: 0},
		{str: "", err: ErrInvalidTime},
		{str: "3:", err: ErrInvalidTime},
		{str: ":40", err: ErrInvalidTime},
		{str: "3:40pm", err: ErrInvalidTime},
		{str: "25:40", err: ErrInvalidTime},
		{str: "23:40", minutes: 23*60 + 40},
		{str: "23:60", err: ErrInvalidTime},
		{str: "23:-01", err: ErrInvalidTime},
		{str: "24:00", err: ErrInvalidTime},
		{str: "-2:00", err: ErrInvalidTime},
	}
	for _, test := range testCases {
		minutes, err := MinutesAfterMidnight(test.str)
		if !errors.Is(err, test.err) {
			t.Errorf("expected '%s' to result in error %s, but got %s", test.str, test.err, err)
		} else if test.minutes != minutes {
			t.Errorf("expected MinutesAfterMidnight('%s') = %d, but got %d", test.str, test.minutes, minutes)
		}
	}
}

func requireCommand(t *testing.T, name string) {
	t.Helper()
	cmd, ok := commands[name]
	if !ok {
		t.Fatalf("command %q is not registered", name)
	}
	if cmd.help == "" {
		t.Errorf("command %q has empty help", name)
	}
	if cmd.handler == nil {
		t.Errorf("command %q has nil handler", name)
	}
}

func TestClimateKeeperAndSunroofCommands(t *testing.T) {
	for _, name := range []string{"climate-keeper", "sunroof"} {
		requireCommand(t, name)
	}
}

func TestSpeedLimitCommands(t *testing.T) {
	for _, name := range []string{
		"speed-limit-activate",
		"speed-limit-deactivate",
		"speed-limit-set-limit",
		"speed-limit-clear-pin",
		"speed-limit-clear-pin-admin",
	} {
		requireCommand(t, name)
	}
}

func TestComfortParityCommands(t *testing.T) {
	for _, name := range []string{
		"cabin-overheat-protection",
		"cop-temp",
		"bioweapon-mode",
		"precondition-max",
		"seat-cooler",
	} {
		requireCommand(t, name)
	}
}

func TestGetDays(t *testing.T) {
	type params struct {
		str   string
		mask  int32
		isErr bool
	}
	testCases := []params{
		{str: "SUN", mask: 1},
		{str: "SUN, WED", mask: 1 + 8},
		{str: "SUN, WEDnesday", mask: 1 + 8},
		{str: "sUN,wEd", mask: 1 + 8},
		{str: "all", mask: 127},
		{str: "sun,all", mask: 127},
		{str: "mon,tues,wed,thurs", mask: 2 + 4 + 8 + 16},
		{str: "marketday", isErr: true},
		{str: "sun mon", isErr: true},
	}
	for _, test := range testCases {
		mask, err := GetDays(test.str)
		if (err != nil) != test.isErr {
			t.Errorf("day string '%s' gave unexpected err = %s", test.str, err)
		} else if mask != test.mask {
			t.Errorf("day string '%s' gave mask %s instead of %s", test.str, strconv.FormatInt(int64(mask), 2), strconv.FormatInt(int64(test.mask), 2))
		}
	}
}

func TestChargeParityCommands(t *testing.T) {
	for _, name := range []string{
		"charge-max-range",
		"charge-standard",
		"schedule-departure",
		"clear-scheduled-departure",
	} {
		requireCommand(t, name)
	}
}

func TestPINCommands(t *testing.T) {
	for _, name := range []string{
		"pin-to-drive",
		"clear-pin-to-drive",
		"reset-pin-to-drive-pin",
		"reset-valet-pin",
	} {
		requireCommand(t, name)
	}
}

func TestParseChargingPolicy(t *testing.T) {
	cases := map[string]vehicle.ChargingPolicy{
		"":         vehicle.ChargingPolicyOff,
		"off":      vehicle.ChargingPolicyOff,
		"all":      vehicle.ChargingPolicyAllDays,
		"weekdays": vehicle.ChargingPolicyWeekdays,
	}
	for in, want := range cases {
		got, err := parseChargingPolicy(in)
		if err != nil || got != want {
			t.Errorf("parseChargingPolicy(%q) = %v, %v; want %v, nil", in, got, err, want)
		}
	}
	if _, err := parseChargingPolicy("bogus"); err == nil {
		t.Errorf("parseChargingPolicy(\"bogus\") should error")
	}
}

func TestMiscParityCommands(t *testing.T) {
	for _, name := range []string{"homelink", "set-vehicle-name"} {
		requireCommand(t, name)
	}
}
