package main

import (
	"fmt"
	"strings"
	"testing"

	 "github.com/gooolib/testing/table"
)

func TestErrors(t *testing.T) {
	err := New("msg")

	test := table.NewTable([]table.Record[string, []string]{
		{
			Name: "Stacktrace",
			Subject: func(_t *testing.T) (string, error) {
				return err.StackTrace(), nil
			},
			Expect: func(t *testing.T) ([]string, error) {
				return []string{
					"errors/errors_test.go. method: TestErrors. line: 12",
					"src/testing/testing.go. method: tRunner. line: ",
					"src/runtime/asm_amd64.s. method: goexit. line: ",
					"",
				}, nil
			},
			Assert: func(t *testing.T, r *table.Record[string, []string]) bool {
				e, _ := r.Expect(t)
				s, _ := r.Subject(t)
				lines := strings.Split(s, "\n")
				for i, line := range lines {
					if !strings.Contains(line, e[i]) {
						t.Errorf("Expected(line %d) %s to contain %s", i, line, e[i])
						return false
					}
				}
				return true
			},
		},
		{
			Name: "Print Error with +v",
			Subject: func(_t *testing.T) (string, error) {
				return fmt.Sprintf("%+v", err), nil
			},
			Expect: func(t *testing.T) ([]string, error) {
				return []string{
					"pkg/errors : msg",
					"",
					"errors/errors_test.go. method: TestErrors. line: 12",
					"src/testing/testing.go. method: tRunner. line: ",
					"src/runtime/asm_amd64.s. method: goexit. line: ",
					"",
					"",
				}, nil
			},
			Assert: func(t *testing.T, r *table.Record[string, []string]) bool {
				e, _ := r.Expect(t)
				s, _ := r.Subject(t)
				lines := strings.Split(s, "\n")
				for i, line := range lines {
					if !strings.Contains(line, e[i]) {
						t.Errorf("Expected(line %d) %s to contain %s", i, line, e[i])
						return false
					}
				}
				return true
			},
		},
		{
			Name: "Print Error with v",
			Subject: func(_t *testing.T) (string, error) {
				return fmt.Sprintf("%v", err), nil
			},
			Expect: func(t *testing.T) ([]string, error) {
				return []string{"pkg/errors : msg"}, nil
			},
			Assert: func(t *testing.T, r *table.Record[string, []string]) bool {
				e, _ := r.Expect(t)
				s, _ := r.Subject(t)
				return s == e[0]
			},
		},
		{
			Name: "Print Error with s",
			Subject: func(_t *testing.T) (string, error) {
				return fmt.Sprintf("%s", err), nil
			},
			Expect: func(t *testing.T) ([]string, error) {
				return []string{"pkg/errors : msg"}, nil
			},
			Assert: func(t *testing.T, r *table.Record[string, []string]) bool {
				e, _ := r.Expect(t)
				s, _ := r.Subject(t)
				return s == e[0]
			},
		},
	})

	test.Run(t)
}

