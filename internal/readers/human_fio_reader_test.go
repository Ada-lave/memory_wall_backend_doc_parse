package readers

import (
	"memory_wall/lib/utils"
	"reflect"
	"testing"
)

func initHumanFIOReader(text string) *HumanFIOReader {
	return &HumanFIOReader{
		textFormatter: &utils.TextFormatter{},
		text:          text,
	}
}

// TODO: Сделать рабочаю логику

func TestHumanFIOReader_GetFIO(t *testing.T) {
	tests := []struct {
		name string
		HFR  *HumanFIOReader
		want []string
	}{
		{
			name: "Base full name in one line test case",
			HFR:  initHumanFIOReader("АНУФРИЕВ АЛЕКСАНДР ПЕТРОВИЧ<br>"),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "Base test case with name on different line",
			HFR:  initHumanFIOReader("АНУФРИЕВ АЛЕКСАНДР<br>ПЕТРОВИЧ<br>"),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "bad test",
			HFR:  initHumanFIOReader("АНУФРИЕВ<br>АЛЕКСАНДР ПЕТРОВИЧ<br>"),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "bad test with whitespaces",
			HFR:  initHumanFIOReader("АНУФРИЕВ <br>АЛЕКСАНДР ПЕТРОВИЧ<br>"),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "Base test case without middle name",
			HFR:  initHumanFIOReader("АНУФРИЕВ АЛЕКСАНДР<br>"),
			want: []string{"Ануфриев", "Александр"},
		},
		{
			name: "All on different line",
			HFR:  initHumanFIOReader("АНУФРИЕВ<br>АЛЕКСАНДР<br>ПЕТРОВИЧ<br>"),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.HFR.GetFIO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HumanFIOReader.GetFIO() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
