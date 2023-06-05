package command

import (
	"testing"

	"pikman/types"
)

func Test_getCommand(t *testing.T) {
	type args struct {
		command     string
		osType      types.OSType
		packageName []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Ubuntu single package",
			args: args{
				command:     "install",
				osType:      types.Ubuntu,
				packageName: []string{"testPackage"},
			},
			want:    "sudo -S apt install testPackage",
			wantErr: false,
		},
		{
			name: "Arch single package",
			args: args{
				command:     "install",
				osType:      types.Arch,
				packageName: []string{"testPackage"},
			},
			want:    "apx --aur install testPackage",
			wantErr: false,
		},
		{
			name: "Arch single package with container name",
			args: args{
				command:     "install",
				osType:      types.Arch,
				packageName: []string{"--name testName", "testPackage"},
			},
			want:    "apx --aur install --name testName testPackage",
			wantErr: false,
		},
		{
			name: "Ubuntu invalid command should return nothing and error",
			args: args{
				command:     "init",
				osType:      types.Ubuntu,
				packageName: []string{"testPackage"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command := &Command{
				Command:     tt.args.command,
				OsType:      tt.args.osType,
				PackageName: tt.args.packageName,
			}
			got, err := command.getCommand()
			if (err != nil) != tt.wantErr {
				t.Errorf("getCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCommand() got = %v, want %v", got, tt.want)
			}
		})
	}
}
