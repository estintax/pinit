package linux

// #include <unistd.h>
import "C"

func SetHostname(name string) int {
  hostname := C.CString(name)
  size := C.size_t(len(name))

  result := C.sethostname(hostname, size)
  return int(result)
}
