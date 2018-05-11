package graph

type Queue []Vertex

func (h Queue) Len() int           { return len(h) }
func (h Queue) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h Queue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Queue) IsEmpty() bool      { return len(h) == 0 }

func (h *Queue) Push(v Vertex) {
	var changed bool
	old := *h
	updated := *h
	// insert Vertex at the correct position (keyed by distance)
	for i, w := range old {
		if v.id == w.id {
			if changed {
				if i+1 < len(updated) {
					updated = append(updated[:i], updated[i+1:]...)
				} else {
					updated = updated[:i]
				}
			} else if v.dist < w.dist {
				updated[i] = v
			}
			changed = true
		} else if v.dist < w.dist {
			changed = true
			updated = append(old[:i], v)
			updated = append(updated, w)
			updated = append(updated, old[i+1:]...)
		}
	}
	if !changed {
		updated = append(old, v)
	}
	*h = updated
}

func (h *Queue) Pop() (v Vertex) {
	old := *h
	v = old[0]
	*h = old[1:]
	return
}
