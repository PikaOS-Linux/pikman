package loader

import "testing"

func Test_getCommand(t *testing.T) {
	type args struct {
		command       string
		osType        OSType
		containerName string
		packageName   []string
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
				command:       "install",
				osType:        Ubuntu,
				containerName: "",
				packageName:   []string{"testPackage"},
			},
			want:    "sudo -S nala install testPackage",
			wantErr: false,
		},
		{
			name: "Arch single package",
			args: args{
				command:       "install",
				osType:        Arch,
				containerName: "",
				packageName:   []string{"testPackage"},
			},
			want:    "apx --aur install  testPackage",
			wantErr: false,
		},
		{
			name: "Arch single package with container name",
			args: args{
				command:       "install",
				osType:        Arch,
				containerName: "--name testName",
				packageName:   []string{"testPackage"},
			},
			want:    "apx --aur install --name testName testPackage",
			wantErr: false,
		},
		{
			name: "Ubuntu single package, container name not used",
			args: args{
				command:       "install",
				osType:        Ubuntu,
				containerName: "--name testName",
				packageName:   []string{"testPackage"},
			},
			want:    "sudo -S nala install testPackage",
			wantErr: false,
		},
		{
			name: "Ubuntu invalid command should return nothing and error",
			args: args{
				command:       "init",
				osType:        Ubuntu,
				containerName: "",
				packageName:   []string{"testPackage"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCommand(tt.args.command, tt.args.osType, tt.args.containerName, tt.args.packageName)
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
