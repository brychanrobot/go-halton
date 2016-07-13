package halton

type HaltonSampler struct {
	CurrentIndex uint64
	base         uint64
}

func NewHaltonSampler(base uint64) *HaltonSampler {
	return &HaltonSampler{CurrentIndex: base + 1, base: base}
}

func getValue(index, base uint64) float64 {
	result := 0.0
	f := 1.0
	i := index
	for i > 0 {
		f /= float64(base)
		result += f * float64(i%base)
		i = i / base
	}

	return result
}

func (h *HaltonSampler) Next() float64 {
	value := getValue(h.CurrentIndex, h.base)
	h.CurrentIndex++
	return value
}
