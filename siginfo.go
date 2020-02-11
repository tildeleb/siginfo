package siginfo

import (
	"os"
	"os/signal"
	"syscall"
)

var signals = map[string]int{
	"SIGHUP":    1,
	"SIGINT":    2,
	"SIGQUIT":   3,
	"SIGILL":    4,
	"SIGTRAP":   5,
	"SIGABRT":   6,
	"SIGIOT":    7,
	"SIGPOLL":   7,
	"SIGFPE":    8,
	"SIGKILL":   9,
	"SIGBUS":    10,
	"SIGSEGV":   11,
	"SIGSYS":    12,
	"SIGPIPE":   13,
	"SIGALRM":   14,
	"SIGTERM":   15,
	"SIGURG":    16,
	"SIGSTOP":   17,
	"SIGTSTP":   18,
	"SIGCONT":   19,
	"SIGCHLD":   20,
	"SIGTTIN":   21,
	"SIGTTOU":   22,
	"SIGIO":     23,
	"SIGXCPU":   24,
	"SIGXFSZ":   25,
	"SIGVTALRM": 26,
	"SIGPROF":   27,
	"SIGWINCH":  28,
	"SIGINFO":   29, // SIGINFO isn't part of the stdlib, but it's 29 on most systems
	"SIGUSR1":   30,
	"SIGUSR2":   31,
}

// SetHandler arranges for function f to be called when signal sig is signalled
func SetHandler(sig string, f func(os.Signal)) bool {

	sigstr, ok := signals[sig]
	if !ok {
		return true
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.Signal(sigstr))

	go func() {
		for s := range ch {
			f(s)
		}
	}()
	return false
}
