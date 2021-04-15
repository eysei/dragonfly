package block

import (
	"github.com/df-mc/dragonfly/dragonfly/block/cube"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
)

// LitPumpkin is a decorative light emitting block crafted with a Carved Pumpkin & Torch
type LitPumpkin struct {
	noNBT
	solid

	// Facing is the direction the pumpkin is facing.
	Facing cube.Direction
}

// LightEmissionLevel ...
func (l LitPumpkin) LightEmissionLevel() uint8 {
	return 15
}

// UseOnBlock ...
func (l LitPumpkin) UseOnBlock(pos cube.Pos, face cube.Face, _ mgl64.Vec3, w *world.World, user item.User, ctx *item.UseContext) (used bool) {
	pos, _, used = firstReplaceable(w, pos, face, l)
	if !used {
		return
	}
	l.Facing = user.Facing().Opposite()

	place(w, pos, l, user, ctx)
	return placed(ctx)
}

// BreakInfo ...
func (l LitPumpkin) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    1,
		Harvestable: alwaysHarvestable,
		Effective:   axeEffective,
		Drops:       simpleDrops(item.NewStack(l, 1)),
	}
}

// EncodeItem ...
func (l LitPumpkin) EncodeItem() (id int32, meta int16) {
	return 91, 0
}

// EncodeBlock ...
func (l LitPumpkin) EncodeBlock() (name string, properties map[string]interface{}) {
	direction := 2
	switch l.Facing {
	case cube.South:
		direction = 0
	case cube.West:
		direction = 1
	case cube.East:
		direction = 3
	}

	return "minecraft:lit_pumpkin", map[string]interface{}{"direction": int32(direction)}
}

// Hash ...
func (l LitPumpkin) Hash() uint64 {
	return hashLitPumpkin | (uint64(l.Facing) << 32)
}

func allLitPumpkins() (pumpkins []canEncode) {
	for i := cube.Direction(0); i <= 3; i++ {
		pumpkins = append(pumpkins, LitPumpkin{Facing: i})
	}
	return
}
