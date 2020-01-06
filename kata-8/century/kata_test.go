package kata

import "testing"

func TestIsDivisible(t *testing.T) {
	tests := []struct {
		title  string
		got    int
		expect int
	}{
		{
			title:  "century(int(1990))",
			got:    century(int(1990)),
			expect: 20,
		}, {
			title:  "century(int(1705))",
			got:    century(int(1705)),
			expect: 18,
		}, {
			title:  "century(int(1900))",
			got:    century(int(1900)),
			expect: 19,
		}, {
			title:  "century(int(1601))",
			got:    century(int(1601)),
			expect: 17,
		}, {
			title:  "century(int(2000))",
			got:    century(int(2000)),
			expect: 20,
		}, {
			title:  "century(int(89))",
			got:    century(int(89)),
			expect: 1,
		},
	}

	for _, test := range tests {
		if test.got != test.expect {
			t.Errorf("Invalid test '%v': got: %v / actual: %v", test.title, test.got, test.expect)
		}
	}
}

// var _ = Describe("Basic tests", func() {
// 	It("should return the correct values", func() {
// 		Expect(century(int(1990))).To(Equal(20))
// 		Expect(century(int(1705))).To(Equal(18))
// 		Expect(century(int(1900))).To(Equal(19))
// 		Expect(century(int(1601))).To(Equal(17))
// 		Expect(century(int(2000))).To(Equal(20))
// 		Expect(century(int(89))).To(Equal(1))
// 	})
// })
