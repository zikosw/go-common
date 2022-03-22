package envconfig

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name string
		want map[string]interface{}
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type config struct {
				Port   int
				User   string
				APIKEY string `envconfig:"api_key"`
			}
			cfg, err := Read[config]("TEST")
			if err != nil {
				panic(err)
			}
			fmt.Printf("%#+v\n", cfg)
			// TODO: compare `want`

		})
	}
}
