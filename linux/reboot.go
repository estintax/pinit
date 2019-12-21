package main

/*
#include <unistd.h>
#include <linux/reboot.h>
#include <sys/reboot.h>

void halt() {
  reboot(LINUX_REBOOT_CMD_HALT);
}

void restart() {
  reboot(LINUX_REBOOT_CMD_RESTART);
}

void shutdown() {
  reboot(LINUX_REBOOT_CMD_POWER_OFF);
}
*/
import "C"

func Halt() {
  C.halt()
}

func Reboot() {
  C.restart()
}

func Poweroff() {
  C.shutdown()
}
