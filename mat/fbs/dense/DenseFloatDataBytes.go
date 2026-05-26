package dense

const (
	DTypeFloat32 int32 = 0
	DTypeFloat64 int32 = 1
)

func (rcv *DenseFloat32) DataBytes() []byte { _ = "STUB: not implemented"; return nil }

func (rcv *DenseFloat64) DataBytes() []byte { _ = "STUB: not implemented"; return nil }
