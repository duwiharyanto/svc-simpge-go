package repo

import (
	"fmt"
	"testing"

	guuid "github.com/google/uuid"
)

func TestGenerateSKID(t *testing.T) {
	uuid := guuid.New()
	id := fmt.Sprintf("%d", uuid.ID())
	t.Logf("id: %s, length: %d\n", id, len(id))
}
