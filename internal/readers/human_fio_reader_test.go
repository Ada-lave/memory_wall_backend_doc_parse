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

func TestHumanFIOReader_GetFIO(t *testing.T) {
	tests := []struct {
		name string
		HFR  *HumanFIOReader
		text string
		want []string
	}{
		{
			name: "Base good test case",
			text: "АНУФРИЕВ АЛЕКСАНДР ПЕТРОВИЧ<br>",
			HFR: initHumanFIOReader(),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
		{
			name: "Base medium test case",
			text: "АНУФРИЕВ АЛЕКСАНДР<br>ПЕТРОВИЧ<br>",
			HFR: initHumanFIOReader(),
			want: []string{"Ануфриев", "Александр", "Петрович"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.HFR.GetFIO(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HumanFIOReader.GetFIO() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
