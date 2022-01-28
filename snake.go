package main

type Snake struct {
	positions []Pos
	direction Dir
}

func (s *Snake) checkIsDead() bool {
	length := len(s.positions)
	head := s.positions[length-1]
	for _, p := range s.positions[:length-1] {
		if p.x == head.x && p.y == head.y {
			return true
		}
	}
	return false
}

func (s *Snake) checkIsInSnake(pos Pos) bool {
	for _, p := range s.positions {
		if p.x == pos.x && p.y == pos.y {
			return true
		}
	}
	return false
}
