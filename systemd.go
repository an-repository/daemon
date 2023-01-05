/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package daemon

import (
	"time"

	"github.com/coreos/go-systemd/v22/daemon"
)

func Ready() (bool, error) {
	return daemon.SdNotify(false, daemon.SdNotifyReady)
}

func Stopping() (bool, error) {
	return daemon.SdNotify(false, daemon.SdNotifyStopping)
}

func Watchdog() (bool, error) {
	return daemon.SdNotify(false, daemon.SdNotifyWatchdog)
}

func WatchdogInterval() (time.Duration, error) {
	return daemon.SdWatchdogEnabled(false)
}

/*
######################################################################################################## @(^_^)@ #######
*/
