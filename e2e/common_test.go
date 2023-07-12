package e2e_test

import (
	"github.com/open-feature/go-sdk/pkg/openfeature"
	mp "github.com/open-feature/go-sdk/pkg/openfeature/testing"
)

// ctxFunction is a context based evaluation callback
var ctxFunction = func(this mp.InMemoryFlag, evalCtx openfeature.FlattenedContext) (
	interface{}, openfeature.ProviderResolutionDetail) {

	defaultValue := this.Variants[this.DefaultVariant]
	defaultResolution := openfeature.ProviderResolutionDetail{
		Reason:  openfeature.DefaultReason,
		Variant: this.DefaultVariant,
	}

	expects := openfeature.FlattenedContext{
		"fn":       "Sulisław",
		"ln":       "Świętopełk",
		"age":      int64(29),
		"customer": false,
	}

	for k, v := range expects {
		if v != evalCtx[k] {
			return defaultValue, defaultResolution
		}
	}

	return this.Variants["internal"], openfeature.ProviderResolutionDetail{
		Reason:  openfeature.TargetingMatchReason,
		Variant: "internal",
	}
}

var memoryFlags = map[string]mp.InMemoryFlag{
	"boolean-flag": {
		Key:            "boolean-flag",
		State:          mp.Enabled,
		DefaultVariant: "on",
		Variants: map[string]interface{}{
			"on":  true,
			"off": false,
		},
		ContextEvaluator: nil,
	},
	"string-flag": {
		Key:            "string-flag",
		State:          mp.Enabled,
		DefaultVariant: "greeting",
		Variants: map[string]interface{}{
			"greeting": "hi",
			"parting":  "bye",
		},
		ContextEvaluator: nil,
	},
	"integer-flag": {
		Key:            "integer-flag",
		State:          mp.Enabled,
		DefaultVariant: "ten",
		Variants: map[string]interface{}{
			"one": 1,
			"ten": 10,
		},
		ContextEvaluator: nil,
	},
	"float-flag": {
		Key:            "float-flag",
		State:          mp.Enabled,
		DefaultVariant: "half",
		Variants: map[string]interface{}{
			"tenth": 0.1,
			"half":  0.5,
		},
		ContextEvaluator: nil,
	},
	"object-flag": {
		Key:            "object-flag",
		State:          mp.Enabled,
		DefaultVariant: "template",
		Variants: map[string]interface{}{
			"empty": map[string]interface{}{},
			"template": map[string]interface{}{
				"showImages":    true,
				"title":         "Check out these pics!",
				"imagesPerPage": 100,
			},
		},
		ContextEvaluator: nil,
	},
	"wrong-flag": {
		Key:            "wrong-flag",
		State:          mp.Enabled,
		DefaultVariant: "one",
		Variants: map[string]interface{}{
			"one": "uno",
			"two": "dos",
		},
		ContextEvaluator: nil,
	},
	"context-aware": {
		Key:            "context-aware",
		State:          mp.Enabled,
		DefaultVariant: "external",
		Variants: map[string]interface{}{
			"internal": "INTERNAL",
			"external": "EXTERNAL",
		},
		ContextEvaluator: &ctxFunction,
	},
}
