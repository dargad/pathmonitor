package pathmonitor_test

import (
	"fmt"
	"github.com/dargad/pathmonitor/internal/app/pathmonitor"
	"io/ioutil"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

var configData = `log:
    # debug | info | error | critical
    level: debug
    # file://<path> | syslog
    output: file:///tmp/pathmonitor.log
paths:
    - path: /path1
      filter: .*.ac3|.*.flv|.*.mkv|.*.mp4|.*.avi
      command: qnapi -q {}
    - path: /path2
      filter: .*.ac3|.*.flv|.*.mkv|.*.mp4|.*.avi
      command: qnapi -q {}
`

func TestConfig(t *testing.T) {
	f, err := ioutil.TempFile("", "config.yaml")
	if err != nil {
		t.Fail()
	}
	f.Write([]byte(configData))
	f.Close()

	c, err := pathmonitor.ReadConfig(f.Name())
	AssertEqual(t, c.Log.Level, "debug", "")
	AssertEqual(t, c.Log.Output, "file:///tmp/pathmonitor.log", "")
	AssertEqual(t, len(c.Paths), 2, "")
	AssertEqual(t, c.Paths[0].Path, "/path1", "")
	AssertEqual(t, c.Paths[1].Path, "/path2", "")
}
