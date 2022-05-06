package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/neovim/go-client/nvim"
	"github.com/spf13/pflag"
)

var logger Logger

func main() {
	var err error
	dbg := pflag.BoolP("debug", "d", false, "Log everything")
	pflag.Parse()
	fnames := pflag.Args()
	sanitizeFileNames(fnames)
	if *dbg {
		logger = NewLogger("/dev/stdout", "debug")
	} else {
		logger = NewLogger("/dev/stdout", "info")
	}
	nvla := nvimListenAddress()
	err = openRemote(fnames, nvla)
	if err != nil {
		logger.Debug(err)
		err := startFresh(fnames, nvla)
		if err != nil {
			os.Exit(1)
		}
	}
}

func nvimListenAddress() string {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Error(err)
	}
	u, err := user.Current()
	if err != nil {
		logger.Error(err)
	}
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%s%s", cwd, u.Username))
	nvla := fmt.Sprintf("/tmp/%x_nvr.sock", h.Sum(nil))
	logger.Debugf("%s%s -> %s", cwd, u.Username, nvla)
	return nvla
}

func sanitizeFileNames(fnames []string) {
	for i := range fnames {
		fnames[i] = strings.ReplaceAll(fnames[i], "|", "")
		fnames[i] = strings.ReplaceAll(fnames[i], "!", "")
	}
}

// startFresh takes fnames and returns error
func startFresh(fnames []string, nvla string) error {
	logger.Debugf(`start new session with "%s" NVIM_LISTEN_ADDRESS=%s`, fnames, nvla)
	nvim := "/usr/bin/nvim"
	cmd := exec.Command(nvim, fnames...)
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("NVIM_LISTEN_ADDRESS=%s", nvla),
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func openRemote(fnames []string, nvla string) error {
	v, err := nvim.Dial(nvla)
	if err != nil {
		return err
	}
	defer v.Close()
	for _, fname := range fnames {
		vimCmd := fmt.Sprintf("edit %s", fname)
		logger.Debugf(`sending "%s" to %s`, vimCmd, nvla)
		v.Command(vimCmd)
	}
	return nil
}
