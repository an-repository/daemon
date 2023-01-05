/*
------------------------------------------------------------------------------------------------------------------------
####### Copyright (c) 2022-2023 Archivage Numérique.
####### All rights reserved.
####### Use of this source code is governed by a MIT style license that can be found in the LICENSE file.
------------------------------------------------------------------------------------------------------------------------
*/

package daemon

import (
	"context"
	"errors"

	"github.com/an-repository/zombie"
)

type (
	Logger interface {
		zombie.Logger
		Error(err error, msg string, kv ...any)
	}

	CheckFunc func() bool

	Daemon struct {
		fnCheck CheckFunc
		logger  Logger
		zombie  *zombie.Zombie
	}

	Option func(d *Daemon)
)

func WithCheckFunc(fn CheckFunc) Option {
	return func(d *Daemon) {
		d.fnCheck = fn
	}
}

func WithLogger(logger Logger) Option {
	return func(d *Daemon) {
		d.logger = logger
	}
}

func New(opts ...Option) *Daemon {
	d := &Daemon{}

	for _, option := range opts {
		option(d)
	}

	return d
}

func (d *Daemon) Start() (bool, error) {
	if d.zombie != nil {
		return false, errors.New("daemon already started") /////////////////////////////////////////////////////////////
	}

	interval, err := WatchdogInterval()
	if err != nil {
		return false, err
	}

	if interval == 0 {
		return false, nil
	}

	d.zombie = zombie.GoTicker(
		context.Background(),
		interval/2,
		func() error {
			if d.fnCheck != nil && !d.fnCheck() {
				return nil
			}

			_, err = Watchdog()
			if err != nil && d.logger != nil {
				d.logger.Error(err, "Daemon watchdog failure") //:::::::::::::::::::::::::::::::::::::::::::::::::::::::
			}

			return err
		},
		zombie.WithName("daemon"),
		zombie.WithLogger(d.logger),
	)

	return true, nil
}

func (d *Daemon) Stop() error {
	if d.zombie == nil {
		return errors.New("daemon not started") ////////////////////////////////////////////////////////////////////////
	}

	d.zombie.Stop()
	d.zombie.Wait()

	d.zombie = nil

	return nil
}

/*
######################################################################################################## @(^_^)@ #######
*/
