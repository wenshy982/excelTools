package validator

import (
	"testing"
)

func Test_ValidatePrompt(t *testing.T) {
	type args struct {
		prompt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				prompt: "test1 hello world tiger bird",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				prompt: "abc hello barely dressed test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PromptValidator.FilterPrompt(tt.args.prompt); (err != nil) != tt.wantErr {
				t.Errorf("FilterPrompt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
