package glidetestb

import "github.com/neocortical/glidetestb/subpackage"

func GlideTestB() string {
	return "GlideTestB() fired. Result: " + subpackage.GlideTestBSubpackage()
}
