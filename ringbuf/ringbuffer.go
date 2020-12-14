package ringbuf

import "errors"

type Ringbuffer struct {
	buffer []int
	start  int
	length int
}

func New(size int) *Ringbuffer {
	return &Ringbuffer{
		buffer: make([]int, size),
	}
}

func (r *Ringbuffer) PushStart(v int) error {
	return r.Push(v, true)
}

func (r *Ringbuffer) PushEnd(v int) error {
	return r.Push(v, false)
}

func (r *Ringbuffer) Push(v int, head bool) error {
	if r.length == len(r.buffer) {
		return errors.New("buffer full")
	}
	if head {
		r.start--
		r.start = mod(r.start, len(r.buffer))
		r.buffer[r.start] = v
	} else {
		end := (r.start + r.length) % len(r.buffer)
		r.buffer[end] = v
	}
	r.length++
	return nil

}

func (r *Ringbuffer) PopStart() (int, error) {
	return r.Pop(true)
}

func (r *Ringbuffer) PopEnd() (int, error) {
	return r.Pop(false)
}

func (r *Ringbuffer) Pop(head bool) (int, error) {
	if r.length == 0 {
		return 0, errors.New("buffer empty")
	}
	r.length--
	var result int
	if head {
		result = r.buffer[r.start]
		r.start++
		r.start = mod(r.start, len(r.buffer))
	} else {
		end := (r.start + r.length) % len(r.buffer)
		result = r.buffer[end]
	}
	return result, nil
}

// mod handles taking the modulus of negative numbers
func mod(v, modulus int) int {
	return (v%modulus + modulus) % modulus
}
