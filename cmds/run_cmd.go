package cmds

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const self = "/proc/self/exe"
const alpine = "/root/alpine"

var runCMD = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		runCmd := exec.Command(self, "exec", "/bin/busybox", "/bin/sh")

		runCmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
			UidMappings: []syscall.SysProcIDMap{
				{
					ContainerID: 0,
					HostID:      os.Getuid(),
					Size:        1,
				},
			},
			GidMappings: []syscall.SysProcIDMap{
				{
					ContainerID: 0,
					HostID:      os.Getgid(),
					Size:        1,
				},
			},
		}

		runCmd.Stdin = os.Stdin
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr

		if err := runCmd.Start(); err != nil {
			log.Fatal(err)
		}

		err := runCmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	},
}
