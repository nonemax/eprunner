package process

import (
	"time"

	"github.com/nonemax/eprunner/runner"
)

// Data represent external process data
type Data struct {
	Name    string
	ID      string
	Args    string
	Status  string
	LastOut string
	Exec    runner.External
	Channel chan int
}

// Run runs goroutin for process
func (d *Data) Run() {
	ch := make(chan int)
	go func() {
		s, err := d.Exec.Start(d.Name, d.Args)
		if err != nil {
			return
		}
		id, err := d.Exec.Check(d.Name)
		if err != nil {
			return
		}
		d.LastOut = s
		d.ID = id

		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			select {
			case <-ch:
				return
			default:
				s, err := d.Exec.Check(d.Name)
				if err != nil {
					continue
				}
				if len(s) == 0 {
					d.Exec.Restart(d.Name)
				}
			}
		}
	}()
	d.Channel = ch
}

// Stop kills external process and goroutine
func (d *Data) Stop() error {
	err := d.Exec.Stop(d.Name)
	if err != nil {
		return err
	}
	close(d.Channel)
	return nil
}
