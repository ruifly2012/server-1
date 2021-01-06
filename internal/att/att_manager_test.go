package att

import (
	"testing"

	"github.com/east-eden/server/define"
	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
)

func TestAttManager(t *testing.T) {
	// reload to project root path
	if err := utils.RelocatePath(); err != nil {
		t.Fatalf("TestAttManager failed: %s", err.Error())
	}

	excel.ReadAllEntries("config/excel")

	attManager := NewAttManager(1)
	attManager.ModBaseAtt(define.Att_Str, 100)

	attManager2 := NewAttManager(2)
	attManager.ModAttManager(attManager2)
	attManager.CalcAtt()
	_, ok := attManager.GetAttValue(define.Att_Str)
	if !ok {
		t.Errorf("att manager calc failed")
	}
}
