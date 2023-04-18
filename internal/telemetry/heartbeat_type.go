// Copyright (c) 2023 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package telemetry

const (
	dimensionON       = "on"
	dimensionOFF      = "off"
	dimensionDefault  = "default"
	dimensionCustom   = "custom"
	dimensionSSL      = "ssl"
	dimensionStartTLS = "starttls"
)

type HeartbeatValues struct {
	Rollout   int `json:"rollout"`
	NbAccount int `json:"nb_account"`
}

type HeartbeatDimensions struct {
	AutoUpdate         string `json:"auto_update"`
	AutoStart          string `json:"auto_start"`
	Beta               string `json:"beta"`
	Doh                string `json:"doh"`
	SplitMode          string `json:"split_mode"`
	ShowAllMail        string `json:"show_all_mail"`
	IMAPConnectionMode string `json:"imap_connection_mode"`
	SMTPConnectionMode string `json:"smtp_connection_mode"`
	IMAPPort           string `json:"imap_port"`
	SMTPPort           string `json:"smtp_port"`
	CacheLocation      string `json:"cache_location"`
	KeychainPref       string `json:"keychain_pref"`
	PrevVersion        string `json:"prev_version"`
}

type HeartbeatData struct {
	MeasurementGroup string
	Event            string
	Values           HeartbeatValues
	Dimensions       HeartbeatDimensions
}

type Heartbeat struct {
	Metrics HeartbeatData

	DefaultIMAPPort int
	DefaultSMTPPort int
	DefaultCache    string
	DefaultKeychain string
}
