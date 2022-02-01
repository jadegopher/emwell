package entities

import "testing"

func TestUpdateType_String_Unique(t *testing.T) {
	i := UpdateType(0)
	result := map[string]struct{}{}
	for {
		if i.String() == "" {
			break
		}

		if _, in := result[i.String()]; in {
			t.Error("String is not unique", i.String())
		} else {
			result[i.String()] = struct{}{}
		}
	}
}

func TestUpdateType_String(t *testing.T) {
	tests := []struct {
		name string
		u    UpdateType
		want string
	}{
		{
			name: "Too big number",
			u:    127,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
