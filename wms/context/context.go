package context

import (
	"github.com/seerwo/imigou/credential"
	"github.com/seerwo/imigou/wms/config"
)

//Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}