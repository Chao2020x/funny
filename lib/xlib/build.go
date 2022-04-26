package xlib

// windows 注意:X11环境可能无法初始化
// #cgo CFLAGS: -I${SRCDIR}/X11/include
// #cgo LDFLAGS: -llibX11
import "C"
