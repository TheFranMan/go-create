package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Run("Can set env locale", func(t *testing.T) {
		for _, test := range []struct {
			env  string
			want Envs
		}{
			{
				env: "Prod",
				want: Envs{
					IsLocal: false,
					IsStage: false,
					IsProd:  true,
				},
			},
			{
				env: "prod",
				want: Envs{
					IsLocal: false,
					IsStage: false,
					IsProd:  true,
				},
			},
			{
				env: "Production",
				want: Envs{
					IsLocal: false,
					IsStage: false,
					IsProd:  true,
				},
			},
			{
				env: "Stage",
				want: Envs{
					IsLocal: false,
					IsStage: true,
					IsProd:  false,
				},
			},
			{
				env: "stage",
				want: Envs{
					IsLocal: false,
					IsStage: true,
					IsProd:  false,
				},
			},
			{
				env: "Staging",
				want: Envs{
					IsLocal: false,
					IsStage: true,
					IsProd:  false,
				},
			},
			{
				env: "Local",
				want: Envs{
					IsLocal: true,
					IsStage: false,
					IsProd:  false,
				},
			},
			{
				env: "Dev",
				want: Envs{
					IsLocal: true,
					IsStage: false,
					IsProd:  false,
				},
			},
			{
				env: "",
				want: Envs{
					IsLocal: true,
					IsStage: false,
					IsProd:  false,
				},
			},
		} {
			t.Run(test.env, func(t *testing.T) {
				os.Setenv("ENV", test.env)

				cfg, err := GetConfig()
				require.Nil(t, err)

				require.Equal(t, test.want.IsLocal, cfg.IsLocal)
				require.Equal(t, test.want.IsStage, cfg.IsStage)
				require.Equal(t, test.want.IsProd, cfg.IsProd)

				os.Unsetenv("ENV")
			})
		}

	})
}
