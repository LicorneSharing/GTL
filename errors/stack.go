package errors

import (
	"errors"
	"fmt"
	"runtime"
)

type Stack struct {
	err   error
	stack []stack
}

func (s *Stack) Error() string {
	msg := s.err.Error()
	for _, stack := range s.stack {
		msg += "\n|> from " + stack.String()
	}

	return msg
}

type stack struct {
	pc             uintptr
	file, funcName string
	line           int
}

func (s stack) String() string {
	return fmt.Sprintf("%s:%d %s", s.file, s.line, s.funcName)
}

func WithStack(err error) error {
	if err == nil {
		return nil
	}

	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		return err
	}

	funcName := runtime.FuncForPC(pc).Name()

	var stackError *Stack

	if !errors.As(err, &stackError) || stackError == nil {
		return &Stack{
			err:   err,
			stack: []stack{{pc, file, funcName, line}},
		}
	}

	stackError.stack = append(
		stackError.stack, stack{pc, file, funcName, line},
	)

	return stackError
}
