// Package react implements a basic reactive system
package react

const testVersion = 4

type sheet struct {
	cells []*cell
}

type cell struct {
	sheet     *sheet
	val       int
	ref1      *Cell
	ref2      *Cell
	onUpdate  func()
	callbacks []callback
}

type callback *func(int)

// New creates a new empty sheet
func New() Reactor {
	return &sheet{}
}

// CreateInput creates a new InputCell with value val
func (s *sheet) CreateInput(val int) InputCell {
	newCell := cell{sheet: s}
	newCell.onUpdate = func() {}
	s.cells = append(s.cells, &newCell)
	newCell.SetValue(val)
	return &newCell
}

// CreateCompute1 creates a new ComputeCell that depends on an InputCell
func (s *sheet) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	newCell := cell{sheet: s, ref1: &c}
	newCell.onUpdate = func() {
		r1 := *newCell.ref1
		newCell.val = f(r1.Value())
	}
	s.cells = append(s.cells, &newCell)
	newCell.onUpdate()
	return &newCell
}

// CreateCompute2 creates a new ComputeCell that depends on two InputCells
func (s *sheet) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	newCell := cell{sheet: s, ref1: &c1, ref2: &c2}
	newCell.onUpdate = func() {
		r1 := *newCell.ref1
		r2 := *newCell.ref2
		newCell.val = f(r1.Value(), r2.Value())
	}
	s.cells = append(s.cells, &newCell)
	newCell.onUpdate()
	return &newCell
}

// Value returns the value of a cell
func (c *cell) Value() int {
	return c.val
}

// SetValue sets the value of a cell
func (c *cell) SetValue(val int) {
	oldValue := c.val
	if val == oldValue {
		return
	}
	c.val = val
	c.sheet.updateAll()
}

// updateAll updates the whole sheet according to the references between cells and calls callbacks
func (s *sheet) updateAll() {
	for _, c := range s.cells {
		oldValue := c.val
		c.onUpdate()
		if oldValue != c.val {
			for _, cb := range c.callbacks {
				callback := *cb
				callback(c.val)
			}
		}
	}
}

// AddCallback registers a callback function for a cell
func (c *cell) AddCallback(f func(int)) CallbackHandle {
	c.callbacks = append(c.callbacks, &f)
	return &f
}

// RemoveCallback removes a callback from a cell
func (c *cell) RemoveCallback(ch CallbackHandle) {
	chToRemove := callback(ch.(*func(int)))
	for i, cb := range c.callbacks {
		if cb == chToRemove {
			l := len(c.callbacks)
			c.callbacks[i] = c.callbacks[l-1]
			c.callbacks = c.callbacks[:l-1]
			return
		}
	}
}
