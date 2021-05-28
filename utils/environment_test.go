package utils

import (
	"github.com/spf13/viper"
	"testing"
)

func TestIsDev(t *testing.T) {
	t.Run("is development", func(t *testing.T) {
		viper.Set("mode", "development")
		if got := IsDev(); !got {
			t.Errorf("IsDev() = %v, want %v", got, !got)
		}
		viper.Reset()
	})
	t.Run("not development", func(t *testing.T) {
		if got := IsDev(); got {
			t.Errorf("IsDev() = %v, want %v", got, !got)
		}
	})
}

func TestIsLocal(t *testing.T) {
	t.Run("is local", func(t *testing.T) {
		viper.Set("mode", "local")
		if got := IsLocal(); !got {
			t.Errorf("IsLocal() = %v, want %v", got, !got)
		}
		viper.Reset()
	})
	t.Run("not local", func(t *testing.T) {
		if got := IsLocal(); got {
			t.Errorf("IsLocal() = %v, want %v", got, !got)
		}
	})
}

func TestIsProd(t *testing.T) {
	t.Run("is production", func(t *testing.T) {
		viper.Set("mode", "production")
		if got := IsProd(); !got {
			t.Errorf("IsProd() = %v, want %v", got, !got)
		}
		viper.Reset()
	})
	t.Run("not production", func(t *testing.T) {
		if got := IsProd(); got {
			t.Errorf("IsProd() = %v, want %v", got, !got)
		}
	})
}
