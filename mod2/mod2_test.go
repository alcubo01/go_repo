// Este modulo traduce cadenas de texto al complicado lenguaje Kirgon
package mod2

import "testing"

func TestTranslate(t *testing.T) {
	type args struct {
		cadena string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Translate(tt.args.cadena); got != tt.want {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
