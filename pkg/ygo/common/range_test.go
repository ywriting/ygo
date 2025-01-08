package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"riguz.com/ygo/pkg/ygo/common"
)

func TestCheckRange_GivenNotIncludes(t *testing.T) {
	assert.Equal(t, false, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(0, 1)},
		&[]common.Range{common.NewRange(2, 3)}))
	assert.Equal(t, false, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(0, 1)},
		&[]common.Range{common.NewRange(1, 3)}))

	assert.Equal(t, false, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(1, 2), common.NewRange(2, 3), common.NewRange(3, 4)},
		&[]common.Range{common.NewRange(0, 3)}))
}

func TestCheckRange_GivenIncludes(t *testing.T) {
	assert.Equal(t, true, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(0, 1)},
		&[]common.Range{common.NewRange(0, 3)}))
	assert.Equal(t, true, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(1, 2)},
		&[]common.Range{common.NewRange(0, 3)}))
	assert.Equal(t, true, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(1, 2), common.NewRange(2, 3)},
		&[]common.Range{common.NewRange(0, 3)}))
	assert.Equal(t, true, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(0, 1), common.NewRange(2, 3)},
		&[]common.Range{common.NewRange(0, 2), common.NewRange(2, 4)}))
	assert.Equal(t, true, common.CheckRangeCovered(
		&[]common.Range{common.NewRange(1, 2), common.NewRange(2, 3), common.NewRange(3, 4)},
		&[]common.Range{common.NewRange(0, 2), common.NewRange(2, 4)}))
}

func TestRangeDiff_GivenNotIncludes(t *testing.T) {
	r1 := common.DiffRange(
		&[]common.Range{common.NewRange(0, 1)},
		&[]common.Range{common.NewRange(2, 3)})
	assert.Equal(t, 0, len(r1))
}

func TestRangeDiff_GivenHasSingleRangeDiff(t *testing.T) {
	r1 := common.DiffRange(
		&[]common.Range{common.NewRange(0, 10)},
		&[]common.Range{common.NewRange(0, 11)})
	assert.Equal(t, 1, len(r1))
	assert.Equal(t, uint64(10), r1[0].Start)
	assert.Equal(t, uint64(11), r1[0].End)
}

func TestRangeDiff_GivenHasSingleRangeDiff_MultipleRanges(t *testing.T) {
	r1 := common.DiffRange(
		&[]common.Range{common.NewRange(0, 10), common.NewRange(20, 30)},
		&[]common.Range{common.NewRange(0, 11), common.NewRange(20, 30)})
	assert.Equal(t, 1, len(r1))
	assert.Equal(t, uint64(10), r1[0].Start)
	assert.Equal(t, uint64(11), r1[0].End)
}

func TestRangeDiff_GivenHasMultipleRangeDiff_OldFragmented(t *testing.T) {
	r1 := common.DiffRange(
		&[]common.Range{common.NewRange(0, 3), common.NewRange(5, 7),
			common.NewRange(8, 10), common.NewRange(16, 18), common.NewRange(21, 23)},
		&[]common.Range{common.NewRange(0, 12), common.NewRange(15, 23)})
	assert.Equal(t, 5, len(r1))
	assertRange(t, []uint64{3, 5, 7, 8, 10, 12, 15, 16, 18, 21}, &r1)
}

func TestRangeDiff_GivenHasMultipleRangeDiff_NewFragmented(t *testing.T) {
	r1 := common.DiffRange(
		&[]common.Range{common.NewRange(1, 6), common.NewRange(8, 12)},
		&[]common.Range{common.NewRange(0, 12), common.NewRange(15, 23), common.NewRange(24, 28)})
	assert.Equal(t, 4, len(r1))
	assertRange(t, []uint64{0, 1, 6, 8, 15, 23, 24, 28}, &r1)
}

func assertRange(t *testing.T, expected []uint64, actual *[]common.Range) {
	i := 0
	for i < len(*actual) {
		start := expected[i*2]
		end := expected[i*2+1]
		r := (*actual)[i]
		assert.Equal(t, start, r.Start)
		assert.Equal(t, end, r.End)
		i += 1
	}
}
