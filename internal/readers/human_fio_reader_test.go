package readers

import (
	"memory_wall/lib/utils"
	"reflect"
	"testing"
)

func initHumanFIOReader() *HumanFIOReader {
	return &HumanFIOReader{
		HumanInfoReader: HumanInfoReader{
			textFormatter: &utils.TextFormatter{},
			dateTool: &utils.DateParseTool{},
		},
	}
}
// func Test_getFIO(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		text string
// 		want []string
// 	}{
// 		{
// 			name: "Base good case test",
// 			text: "АНУФРИЕВ АЛЕКСАНДР ПЕТРОВИЧ<br>",
// 			want: []string{"Ануфриев", "Александр", "Петрович"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			assert.Equal(t, tt.want, getFIO(tt.text))
// 		})
// 	}
// }

// TODO: Modify this test to new logic
func TestHumanFIOReader_GetFIO(t *testing.T) {
	tests := []struct {
		name string
		HFR  *HumanFIOReader
		text string
		want []string
	}{
		{
			name: "Base full name in one line test case",
			text: "АНУФРИЕВ АЛЕКСАНДР ПЕТРОВИЧ<br>",
			HFR: initHumanFIOReader(),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "Base test case with namw on different line",
			text: "АНУФРИЕВ АЛЕКСАНДР<br>ПЕТРОВИЧ<br>",
			HFR: initHumanFIOReader(),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "Base test case without middle name",
			text:"АНУФРИЕВ АЛЕКСАНДР<br>",
			HFR: initHumanFIOReader(),
			want: []string{"Ануфриев", "Александр"},
		},
		{
			name: "All on different line",
			text: "АНУФРИЕВ<br>АЛЕКСАНДР<br>ПЕТРОВИЧ<br>",
			HFR: initHumanFIOReader(),
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
