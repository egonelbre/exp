package wave

import "github.com/egonelbre/exp/audio"

type Clip struct {
	head   int
	buffer audio.BufferInt16
}

func (clip *Clip) Process32(output *audio.Buffer32) error {
	tmp := clip.buffer
	// TODO: fix bounds
	tmp.Data = tmp.Data[clip.head : clip.head+len(output.Data)]
	clip.head += len(output.Data)
	return tmp.Process32(output)
}

func (clip *Clip) Process64(output *audio.Buffer64) error {
	tmp := clip.buffer
	// TODO: fix bounds
	tmp.Data = tmp.Data[clip.head : clip.head+len(output.Data)]
	clip.head += len(output.Data)
	return tmp.Process64(output)
}
