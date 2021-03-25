package tests

import (
	"adventure-with-friends/src/models"
	"fmt"
	"testing"
)

func TestPlayer_RetrieveInfo(t *testing.T) {
	type args struct {
		name string
		level int
	}

	tests := []struct{
		testName string
		args args
		want string
	}{
		{"Create models with name Tommy and level 1", args{"Tommy", 1}, "Name: Tommy, Level: 1"},
	}

	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			player := models.Player{
				Name: testCase.args.name,
				Level: testCase.args.level,
			}

			if got := player.RetrieveInfo(); got != testCase.want {
				t.Errorf("RetrieveInfo() = %v, want %v", got, testCase.want)
			}
		})
	}
}


func BenchmarkPlayer_RetrieveInfo(b *testing.B) {
	player := models.Player{
		Name: "Jane Doe",
		Level: 1,
	}
	for i := 0; i < b.N; i++ {
		player.RetrieveInfo()
	}
}

func ExampleRetrieveInfo() {
	player := models.Player{
		Name: "Jane Doe",
		Level: 1,
	}
	fmt.Println(player.RetrieveInfo())
	// Output:
	// Name: Jane Doe, Level: 1
}

