package imigou

import (
	"github.com/seerwo/imigou/cache"
	"github.com/seerwo/imigou/wms"
	wmsconfig "github.com/seerwo/imigou/wms/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	//Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	//Output to stdout instead of the default stderr
	//Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	//Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

//Imigou struct
type Imigou struct {
	cache cache.Cache
}

//NewImigou init
func NewImigou() *Imigou {
	return &Imigou{}
}

//SetCache set cache
func (c *Imigou) SetCache(cache cache.Cache){
	c.cache = cache
}

//GetWms get Wms instance
func (c *Imigou) GetWms(cfg *wmsconfig.Config) *wms.Wms {
	if cfg.Cache == nil {
		cfg.Cache = c.cache
	}
	return wms.NewWms(cfg)
}
