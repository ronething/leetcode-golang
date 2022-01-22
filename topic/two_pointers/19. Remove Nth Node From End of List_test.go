package two_pointers

import (
	"reflect"
	"testing"

	"github.com/InterviewTips/algorithm-coding/guns"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one",
			args: args{
				head: guns.GenLinkList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: guns.GenLinkList([]int{1, 2, 3, 5}),
		},
		{
			name: "two",
			args: args{
				head: guns.GenLinkList([]int{1}),
				n:    1,
			},
			want: guns.GenLinkList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}