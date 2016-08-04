package subpackage

import "github.com/neocortical/glidetestc"

func GlideTestBSubpackage() string {
	return "GlideTestB subpackage called " + glidetestc.GlideTestC()
}
