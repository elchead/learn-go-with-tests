package main

import "sort"

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (l League) Sort() League {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Score > l[j].Score
	})
	return l
}
