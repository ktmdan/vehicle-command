# Command coverage: Fleet API ↔ proxy ↔ tesla-control

Generated 2026-05-27. Cross-references every Tesla Fleet API *vehicle-commands*
endpoint against the `pkg/proxy` REST proxy, the `tesla-control` CLI, and the
backing `pkg/vehicle` method.

Legend: ✓ = supported. "REST fwd" = no signed-command protobuf exists; the proxy
returns `ErrCommandUseRESTAPI`, which forwards the request to Tesla's REST API.
"—" = intentionally absent (REST-only/data; the CLI speaks only the signed
protocol).

| Fleet endpoint | Proxy | CLI command | Library method | Notes |
|---|---|---|---|---|
| actuate_trunk | ✓ | trunk-move / trunk-open / trunk-close / frunk-open | ActuateTrunk / OpenTrunk / CloseTrunk / OpenFrunk | |
| add_charge_schedule | ✓ | charging-schedule-add | AddChargeSchedule | |
| add_precondition_schedule | ✓ | precondition-schedule-add | AddPreconditionSchedule | |
| adjust_volume | ✓ | media-set-volume | SetVolume | |
| auto_conditioning_start | ✓ | climate-on | ClimateOn | |
| auto_conditioning_stop | ✓ | climate-off | ClimateOff | |
| cancel_software_update | ✓ | software-update-cancel | CancelSoftwareUpdate | |
| charge_max_range | ✓ | charge-max-range | ChargeMaxRange | |
| charge_port_door_close | ✓ | charge-port-close | ChargePortClose | |
| charge_port_door_open | ✓ | charge-port-open | ChargePortOpen | |
| charge_standard | ✓ | charge-standard | ChargeStandardRange | |
| charge_start | ✓ | charging-start | ChargeStart | |
| charge_stop | ✓ | charging-stop | ChargeStop | |
| clear_pin_to_drive_admin | ✓ | clear-pin-to-drive | ClearPINToDrive | |
| door_lock | ✓ | lock | Lock | |
| door_unlock | ✓ | unlock | Unlock | |
| erase_user_data | ✓ | erase-guest-data | EraseGuestData | |
| flash_lights | ✓ | flash-lights | FlashLights | |
| guest_mode | ✓ | guest-mode-on / guest-mode-off | SetGuestMode | |
| honk_horn | ✓ | honk | HonkHorn | |
| media_next_fav | ✓ | media-next-favorite | MediaNextFavorite | |
| media_next_track | ✓ | media-next-track | MediaNextTrack | |
| media_prev_fav | ✓ | media-previous-favorite | MediaPreviousFavorite | |
| media_prev_track | ✓ | media-previous-track | MediaPreviousTrack | |
| media_toggle_playback | ✓ | media-toggle-playback | ToggleMediaPlayback | |
| media_volume_down | ✓ | media-volume-down | VolumeDown | |
| media_volume_up | ✓ | media-volume-up | VolumeUp | |
| navigation_gps_request | REST fwd | — | — | No protobuf action |
| navigation_request | REST fwd | — | — | No protobuf action |
| navigation_sc_request | REST fwd | — | — | No protobuf action |
| navigation_waypoints_request | REST fwd | — | — | No protobuf action |
| remote_auto_seat_climate_request | ✓ | auto-seat-and-climate | AutoSeatAndClimate | |
| remote_auto_steering_wheel_heat_climate_request | REST fwd | — | — | No protobuf action |
| remote_boombox | REST fwd | — | — | Works via REST per hass-teslemetry#31 |
| remote_seat_cooler_request | ✓ | seat-cooler | SetSeatCooler | |
| remote_seat_heater_request | ✓ | seat-heater | SetSeatHeater | |
| remote_start_drive | ✓ | drive | RemoteDrive | |
| remote_steering_wheel_heat_level_request | REST fwd | — | — | No protobuf action |
| remote_steering_wheel_heater_request | ✓ | steering-wheel-heater | SetSteeringWheelHeater | |
| remove_charge_schedule | ✓ | charging-schedule-remove | RemoveChargeSchedule | |
| remove_precondition_schedule | ✓ | precondition-schedule-remove | RemovePreconditionSchedule | |
| reset_pin_to_drive_pin | ✓ | reset-pin-to-drive-pin | ResetPIN | |
| reset_valet_pin | ✓ | reset-valet-pin | ResetValetPin | |
| schedule_software_update | ✓ | software-update-start | ScheduleSoftwareUpdate | |
| set_bioweapon_mode | ✓ | bioweapon-mode | SetBioweaponDefenseMode | |
| set_cabin_overheat_protection | ✓ | cabin-overheat-protection | SetCabinOverheatProtection | |
| set_charge_limit | ✓ | charging-set-limit | ChangeChargeLimit | |
| set_charging_amps | ✓ | charging-set-amps | SetChargingAmps | |
| set_climate_keeper_mode | ✓ | climate-keeper | SetClimateKeeperMode | Dog mode |
| set_cop_temp | ✓ | cop-temp | SetCabinOverheatProtectionTemperature | |
| set_pin_to_drive | ✓ | pin-to-drive | SetPINToDrive | |
| set_preconditioning_max | ✓ | precondition-max | SetPreconditioningMax | |
| set_scheduled_charging | ✓ | charging-schedule | ScheduleCharging | |
| set_scheduled_departure | ✓ | schedule-departure / clear-scheduled-departure | ScheduleDeparture / ClearScheduledDeparture | |
| set_sentry_mode | ✓ | sentry-mode | SetSentryMode | |
| set_temps | ✓ | climate-set-temp | ChangeClimateTemp | |
| set_valet_mode | ✓ | valet-mode-on / valet-mode-off | SetValetMode | |
| set_vehicle_name | ✓ | set-vehicle-name | SetVehicleName | |
| speed_limit_activate | ✓ | speed-limit-activate | ActivateSpeedLimit | |
| speed_limit_clear_pin | ✓ | speed-limit-clear-pin | ClearSpeedLimitPIN | |
| speed_limit_clear_pin_admin | ✓ | speed-limit-clear-pin-admin | ClearSpeedLimitPINAdminAction | |
| speed_limit_deactivate | ✓ | speed-limit-deactivate | DeactivateSpeedLimit | |
| speed_limit_set_limit | ✓ | speed-limit-set-limit | SpeedLimitSetLimitMPH | |
| sun_roof_control | ✓ | sunroof | VentSunroof / CloseSunroof / ChangeSunroofState | Signed command (this PR) |
| trigger_homelink | ✓ | homelink | TriggerHomelink | |
| upcoming_calendar_entries | REST fwd | — | — | REST/data; no protobuf action |
| window_control | ✓ | windows-vent / windows-close | VentWindows / CloseWindows | |

## Endpoints not on the Fleet *vehicle-commands* page

- `dashcam_save_clip` (legacy Owner API) and `take_drivenote` (Teslemetry) have
  no protobuf action; the proxy returns `ErrCommandUseRESTAPI` (REST fwd).
- `set_managed_charge_current_request`, `set_managed_charger_location`,
  `set_managed_scheduled_charging_time` are proxy extras (managed charging), not
  part of the vehicle-commands page.

## Signable vs REST-only

Teslemetry publishes a list of endpoints that "cannot be signed". It disagrees
with this repo, which *does* sign `media_*`, `window_control`, `erase_user_data`,
`clear_pin_to_drive_admin`, `speed_limit_clear_pin_admin`, and `sun_roof_control`.
The authoritative REST-only set for this protocol is the proxy's own
`ErrCommandUseRESTAPI` returns (the "REST fwd" rows above). Signing becomes the
default in 2025 per Teslemetry.
