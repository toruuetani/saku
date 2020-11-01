package saku

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

func emojiEnabled() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}

// Returns the string prepended by the given emoji when the terminal is tty, otherwise drops emoji and returns the string.
func prependEmoji(e string, str string, useEmoji bool) string {
	if useEmoji {
		str = e + "  " + str
	}

	return str
}

// Returns true if the process is invoked in saku.
func invokedInSaku() bool {
	return os.Getenv("IN_SAKU") == "true"
}

func forceLfReadFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	content = []byte(strings.NewReplacer(
		"\r\n", "\n",
		"\r", "\n",
	).Replace(string(content)))
	return content, err
}
