package cmds

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const ENV = "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

var execCMD = &cobra.Command{
	Use: "exec",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("error args")
		}

		runArgs := []string{}
		if len(args) > 1 {
			runArgs = args[1:]
		}

		err := syscall.Chroot(alpine)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Chdir("/")
		if err != nil {
			log.Fatal(err)
		}

		err = syscall.Mount("proc", "/proc", "proc", 0, "")
		if err != nil {
			log.Fatal(err)
		}

		runCmd := exec.Command(args[0], runArgs...)
		runCmd.Stdin = os.Stdin
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		runCmd.Env = []string{ENV}

		if err = runCmd.Start(); err != nil {
			log.Fatal(err)
		}

		runCmd.Wait()
	},
}
