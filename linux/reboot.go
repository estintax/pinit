package linux

/*
#include <unistd.h>
#include <linux/reboot.h>
#include <sys/reboot.h>
*/
import "C"

// From linux/reboot.h. Copyright (c) Linus Torvalds and Co.
const (
  LINUX_REBOOT_MAGIC1         int = 0xfee1dead
  LINUX_REBOOT_MAGIC2         int = 672274793
  LINUX_REBOOT_MAGIC2A        int = 85072278
  LINUX_REBOOT_MAGIC2B        int = 369367448
  LINUX_REBOOT_MAGIC2C        int = 537993216

  LINUX_REBOOT_CMD_RESTART    int = 0x01234567
  LINUX_REBOOT_CMD_HALT       int = 0xCDEF0123
  LINUX_REBOOT_CMD_CAD_ON     int = 0x89ABCDEF
  LINUX_REBOOT_CMD_CAD_OFF    int = 0x00000000
  LINUX_REBOOT_CMD_POWER_OFF  int = 0x4321FEDC
  LINUX_REBOOT_CMD_RESTART2   int = 0xA1B2C3D4
  LINUX_REBOOT_CMD_SW_SUSPEND int = 0xD000FCE2
  LINUX_REBOOT_CMD_KEXEC      int = 0x45584543
)

func Reboot(cmd int) int {
  result := C.reboot(C.int(cmd))
  return int(result)
}
