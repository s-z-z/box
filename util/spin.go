package util

import (
	"context"
	"time"

	"github.com/briandowns/spinner"
)

const (
	DefaultStyle = 9
	DefaultSpeed = 100
)

func Spin(ctx context.Context, style int, speed int, options ...spinner.Option) {
	s := spinner.New(spinner.CharSets[style], time.Duration(speed)*time.Millisecond, options...)
	s.Start()
	<-ctx.Done()
	s.Stop()
}

func SpinNDefault(ctx context.Context, style int, options ...spinner.Option) {
	Spin(ctx, style, DefaultSpeed, options...)
}

func SpinDefault(ctx context.Context, options ...spinner.Option) {
	Spin(ctx, DefaultStyle, DefaultSpeed, options...)
}

func Spin11Default(ctx context.Context, options ...spinner.Option) {
	Spin(ctx, 11, DefaultSpeed, options...)
}

func Spin13Default(ctx context.Context, options ...spinner.Option) {
	Spin(ctx, 13, DefaultSpeed, options...)
}

type Spinner struct {
	Style   int
	Speed   int
	Options []spinner.Option
	start   chan struct{}
	pause   chan struct{}
	done    chan struct{}
}

func NewSpinner(style int, speed int, options ...spinner.Option) *Spinner {
	if speed <= 0 {
		speed = DefaultSpeed
	}
	return &Spinner{
		Style:   style,
		Speed:   speed,
		Options: options,
		start:   make(chan struct{}),
		pause:   make(chan struct{}),
		done:    make(chan struct{}),
	}
}

func (s *Spinner) Run() {
	go func() {
		for {
			select {
			case <-s.start:
				ctx, cancel := context.WithCancel(context.Background())
				go func() {
					SpinNDefault(ctx, s.Style, s.Options...)
				}()
				select {
				case <-s.pause:
					cancel()
				case <-s.done:
					cancel()
					return
				}
			case <-s.done:
				return
			}
		}
	}()
}

func (s *Spinner) Start() {
	select {
	case s.start <- struct{}{}:
	default:
	}
}

func (s *Spinner) Pause() {
	select {
	case s.pause <- struct{}{}:
	default:
	}
}

func (s *Spinner) Done() {
	close(s.done)
}
