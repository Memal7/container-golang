package main

// import the necessary packages
import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// main function
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

// run function - create a new process
func run() {
	fmt.Println("running \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
// create a new UTS namespace for the process (set hostname, domainname, etc.)
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS,
	}

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
	// Change the current working directory to the new root filesystem
	must(os.Chdir("/"))
	
	must(cmd.Run())

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}