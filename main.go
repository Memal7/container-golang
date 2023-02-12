package main

// import necessary packages
import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// main function - check if the program is run or child
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

// run function - if run, then print the arguments and create a new process
func run() {
	fmt.Println("running \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// create an UTS (Unix Timesharing System) namespace for setting the hostname and PID namespace to isolate the process from the host
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	// execute the commands (above cmds)
	must(cmd.Run())
}

// child function - set hostname
func child() {
	fmt.Println("running \n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	// Chroot to the new root filesystem
	must(syscall.Chroot)("/home/memal7/ubuntu-fs"))
	// change the current working directory to the new root filesystem
	must(os.Chdir("/"))
	// mount the proc filesystem in the directory /proc
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	// mount the sysfs filesystem in the directory /sys
	must(syscall.Mount("temporary-fs", "mytmpfs", "tmpfs", 0, ""))
	
	must(cmd.Run())
	// unmount or clean up the proc filesystem
	must(syscall.Unmount("/proc", 0))
	// unmount or clean up the tmpfs filesystem
	must(syscall.Unmount("/mytmpfs", 0))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}