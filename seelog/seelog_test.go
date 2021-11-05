package seelog

import "testing"

func TestSeelog(t *testing.T) {
	t.Run("", func(t *testing.T) {
		Infof("%d", 1)
		Flush()
	})
}
