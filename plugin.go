package plugin

import (
	"log"

	"github.com/gdbu/jump/events"

	"github.com/vroomy/vroomy"
)

var p plugin

func init() {
	if err := vroomy.Register("jump-events", &p); err != nil {
		log.Fatalf("error loading jump plugin: %v", err)
	}
}

type plugin struct {
	vroomy.BasePlugin

	backend *events.Controller
}

// Load will be called by Vroomy on initialization
func (p *plugin) Load(env vroomy.Environment) (err error) {
	p.backend = events.New()
	return
}

// Backend will return the plugin's backend
func (p *plugin) Backend() interface{} {
	return p.backend
}
