package linux

// #include <sys/mount.h>
import "C"
import "unsafe"

func Mount(source string, target string, fstype string, mountflags uint32) (int, string) {
  csource := C.CString(source)
  ctarget := C.CString(target)
  cfstype := C.CString(fstype)
  cdata := C.CString("")

  result := C.mount(csource, ctarget, cfstype, C.ulong(mountflags), unsafe.Pointer(cdata))
  return int(result), C.GoString(cdata)
}

func Umount(target string) int {
  ctarget := C.CString(target)

  result := C.umount(ctarget)
  return int(result)
}
