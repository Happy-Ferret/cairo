// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 06 Jan 2018 21:05:42 EST.
// By https://git.io/c-for-go. DO NOT EDIT.

package pugl

/*
#cgo pkg-config: cairo
#include "pugl.h"
#include "cairo_gl.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/golang-ui/cairo-go/cairo"
)

// Init function as declared in pugl/pugl.h:383
func Init(pargc *int32, argv **byte) *View {
	cpargc, _ := (*C.int)(unsafe.Pointer(pargc)), cgoAllocsUnknown
	cargv, _ := (**C.char)(unsafe.Pointer(argv)), cgoAllocsUnknown
	__ret := C.puglInit(cpargc, cargv)
	__v := *(**View)(unsafe.Pointer(&__ret))
	return __v
}

// InitWindowClass function as declared in pugl/pugl.h:390
func InitWindowClass(view *View, name string) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	name = safeString(name)
	cname, _ := unpackPCharString(name)
	C.puglInitWindowClass(cview, cname)
	runtime.KeepAlive(name)
}

// InitWindowParent function as declared in pugl/pugl.h:396
func InitWindowParent(view *View, parent NativeWindow) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cparent, _ := (C.PuglNativeWindow)(parent), cgoAllocsUnknown
	C.puglInitWindowParent(cview, cparent)
}

// InitWindowSize function as declared in pugl/pugl.h:402
func InitWindowSize(view *View, width int32, height int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cwidth, _ := (C.int)(width), cgoAllocsUnknown
	cheight, _ := (C.int)(height), cgoAllocsUnknown
	C.puglInitWindowSize(cview, cwidth, cheight)
}

// InitWindowMinSize function as declared in pugl/pugl.h:408
func InitWindowMinSize(view *View, width int32, height int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cwidth, _ := (C.int)(width), cgoAllocsUnknown
	cheight, _ := (C.int)(height), cgoAllocsUnknown
	C.puglInitWindowMinSize(cview, cwidth, cheight)
}

// InitWindowAspectRatio function as declared in pugl/pugl.h:417
func InitWindowAspectRatio(view *View, minX int32, minY int32, maxX int32, maxY int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cminX, _ := (C.int)(minX), cgoAllocsUnknown
	cminY, _ := (C.int)(minY), cgoAllocsUnknown
	cmaxX, _ := (C.int)(maxX), cgoAllocsUnknown
	cmaxY, _ := (C.int)(maxY), cgoAllocsUnknown
	C.puglInitWindowAspectRatio(cview, cminX, cminY, cmaxX, cmaxY)
}

// InitResizable function as declared in pugl/pugl.h:427
func InitResizable(view *View, resizable int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cresizable, _ := (C.bool)(resizable), cgoAllocsUnknown
	C.puglInitResizable(cview, cresizable)
}

// InitTransientFor function as declared in pugl/pugl.h:436
func InitTransientFor(view *View, parent uint) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cparent, _ := (C.uintptr_t)(parent), cgoAllocsUnknown
	C.puglInitTransientFor(cview, cparent)
}

// InitContextType function as declared in pugl/pugl.h:442
func InitContextType(view *View, kind ContextType) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	ckind, _ := (C.PuglContextType)(kind), cgoAllocsUnknown
	C.puglInitContextType(cview, ckind)
}

// CreateWindow function as declared in pugl/pugl.h:460
func CreateWindow(view *View, title string) int32 {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	title = safeString(title)
	ctitle, _ := unpackPCharString(title)
	__ret := C.puglCreateWindow(cview, ctitle)
	runtime.KeepAlive(title)
	__v := (int32)(__ret)
	return __v
}

// ShowWindow function as declared in pugl/pugl.h:466
func ShowWindow(view *View) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	C.puglShowWindow(cview)
}

// HideWindow function as declared in pugl/pugl.h:472
func HideWindow(view *View) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	C.puglHideWindow(cview)
}

// GetNativeWindow function as declared in pugl/pugl.h:478
func GetNativeWindow(view *View) NativeWindow {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	__ret := C.puglGetNativeWindow(cview)
	__v := (NativeWindow)(__ret)
	return __v
}

// SetHandle function as declared in pugl/pugl.h:491
func SetHandle(view *View, handle *Handle) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	chandle, _ := (C.PuglHandle)(unsafe.Pointer(handle)), cgoAllocsUnknown
	C.puglSetHandle(cview, chandle)
}

// GetHandle function as declared in pugl/pugl.h:497
func GetHandle(view *View) *Handle {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	__ret := C.puglGetHandle(cview)
	__v := *(**Handle)(unsafe.Pointer(&__ret))
	return __v
}

// GetVisible function as declared in pugl/pugl.h:503
func GetVisible(view *View) int32 {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	__ret := C.puglGetVisible(cview)
	__v := (int32)(__ret)
	return __v
}

// GetSize function as declared in pugl/pugl.h:509
func GetSize(view *View, width *int32, height *int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cwidth, _ := (*C.int)(unsafe.Pointer(width)), cgoAllocsUnknown
	cheight, _ := (*C.int)(unsafe.Pointer(height)), cgoAllocsUnknown
	C.puglGetSize(cview, cwidth, cheight)
}

// GetContext function as declared in pugl/pugl.h:523
func GetContext(view *View) unsafe.Pointer {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	__ret := C.puglGetContext(cview)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// EnterContext function as declared in pugl/pugl.h:535
func EnterContext(view *View) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	C.puglEnterContext(cview)
}

// LeaveContext function as declared in pugl/pugl.h:544
func LeaveContext(view *View, flush int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cflush, _ := (C.bool)(flush), cgoAllocsUnknown
	C.puglLeaveContext(cview, cflush)
}

// SetEventFunc function as declared in pugl/pugl.h:564
func SetEventFunc(view *View, eventFunc EventFunc) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	ceventFunc, _ := eventFunc.PassValue()
	C.puglSetEventFunc(cview, ceventFunc)
}

// IgnoreKeyRepeat function as declared in pugl/pugl.h:570
func IgnoreKeyRepeat(view *View, ignore int32) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	cignore, _ := (C.bool)(ignore), cgoAllocsUnknown
	C.puglIgnoreKeyRepeat(cview, cignore)
}

// GrabFocus function as declared in pugl/pugl.h:576
func GrabFocus(view *View) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	C.puglGrabFocus(cview)
}

// WaitForEvent function as declared in pugl/pugl.h:587
func WaitForEvent(view *View) Status {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	__ret := C.puglWaitForEvent(cview)
	__v := (Status)(__ret)
	return __v
}

// ProcessEvents function as declared in pugl/pugl.h:597
func ProcessEvents(view *View) Status {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	__ret := C.puglProcessEvents(cview)
	__v := (Status)(__ret)
	return __v
}

// PostRedisplay function as declared in pugl/pugl.h:607
func PostRedisplay(view *View) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	C.puglPostRedisplay(cview)
}

// Destroy function as declared in pugl/pugl.h:613
func Destroy(view *View) {
	cview, _ := (*C.PuglView)(unsafe.Pointer(view)), cgoAllocsUnknown
	C.puglDestroy(cview)
}

// CairoGLCreate function as declared in pugl/cairo_gl.h:32
func CairoGLCreate(ctx *CairoGL, width int32, height int32, bpp int32) *cairo.Surface {
	cctx, _ := ctx.PassRef()
	cwidth, _ := (C.int)(width), cgoAllocsUnknown
	cheight, _ := (C.int)(height), cgoAllocsUnknown
	cbpp, _ := (C.int)(bpp), cgoAllocsUnknown
	__ret := C.pugl_cairo_gl_create(cctx, cwidth, cheight, cbpp)
	__v := *(**cairo.Surface)(unsafe.Pointer(&__ret))
	return __v
}

// CairoGLFree function as declared in pugl/cairo_gl.h:36
func CairoGLFree(ctx *CairoGL) {
	cctx, _ := ctx.PassRef()
	C.pugl_cairo_gl_free(cctx)
}

// CairoGLConfigure function as declared in pugl/cairo_gl.h:39
func CairoGLConfigure(ctx *CairoGL, width int32, height int32) {
	cctx, _ := ctx.PassRef()
	cwidth, _ := (C.int)(width), cgoAllocsUnknown
	cheight, _ := (C.int)(height), cgoAllocsUnknown
	C.pugl_cairo_gl_configure(cctx, cwidth, cheight)
}

// CairoGLDraw function as declared in pugl/cairo_gl.h:42
func CairoGLDraw(ctx *CairoGL, width int32, height int32) {
	cctx, _ := ctx.PassRef()
	cwidth, _ := (C.int)(width), cgoAllocsUnknown
	cheight, _ := (C.int)(height), cgoAllocsUnknown
	C.pugl_cairo_gl_draw(cctx, cwidth, cheight)
}
