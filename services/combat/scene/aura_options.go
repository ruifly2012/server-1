package scene

import (
	"github.com/yokaiio/yokai_server/define"
)

type AuraOption func(*AuraOptions)
type AuraOptions struct {
	Owner        SceneUnit
	Caster       SceneUnit
	Amount       int32
	Level        int32
	RagePctMod   float32
	CurWrapTimes uint8
	SpellType    define.ESpellType
	Entry        *define.AuraEntry
}

func DefaultAuraOptions() *AuraOptions {
	return &AuraOptions{
		Owner:        nil,
		Caster:       nil,
		Amount:       0,
		Level:        1,
		RagePctMod:   0.0,
		CurWrapTimes: 0,
		SpellType:    define.SpellType_Melee,
		Entry:        nil,
	}
}

func WithAuraCaster(caster SceneUnit) AuraOption {
	return func(o *AuraOptions) {
		o.Caster = caster
	}
}

func WithAuraTarget(owner SceneUnit) AuraOption {
	return func(o *AuraOptions) {
		o.Owner = owner
	}
}

func WithAuraAmount(amount int32) AuraOption {
	return func(o *AuraOptions) {
		o.Amount = amount
	}
}

func WithAuraLevel(level int32) AuraOption {
	return func(o *AuraOptions) {
		o.Level = level
	}
}

func WithAuraRagePctMod(ragePctMod float32) AuraOption {
	return func(o *AuraOptions) {
		o.RagePctMod = ragePctMod
	}
}

func WithAuraCurWrapTimes(curWrapTimes uint8) AuraOption {
	return func(o *AuraOptions) {
		o.CurWrapTimes = curWrapTimes
	}
}

func WithAuraType(tp define.ESpellType) AuraOption {
	return func(o *AuraOptions) {
		o.SpellType = tp
	}
}

func WithAuraEntry(entry *define.AuraEntry) AuraOption {
	return func(o *AuraOptions) {
		o.Entry = entry
	}
}
