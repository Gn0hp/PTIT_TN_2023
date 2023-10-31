package voter

import "PTIT_TN/internal/entities"

func extractId(e []*entities.ElectionRole) []uint64 {
	var ids []uint64
	for _, v := range e {
		ids = append(ids, v.ID)
	}
	return ids
}
