package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

func TestPlayerFlashed_FlashDuration(t *testing.T) {
	p := common.NewPlayer(demoInfoProviderMock{})
	e := PlayerFlashed{Player: p}

	assert.Equal(t, time.Duration(0), e.FlashDuration())

	p.FlashDuration = 2.3

	assert.Equal(t, 2300*time.Millisecond, e.FlashDuration())
}

func TestGrenadeEvent_Base(t *testing.T) {
	base := GrenadeEvent{GrenadeEntityID: 1}
	flashEvent := FlashExplode{base}

	assert.Equal(t, base, flashEvent.Base())
}

func TestBombEvents(t *testing.T) {
	events := []BombEventIf{
		BombDefuseStart{},
		BombDefuseAborted{},
		BombDefused{},
		BombExplode{},
		BombPlantBegin{},
		BombPlanted{},
	}

	for _, e := range events {
		e.implementsBombEventIf()
	}
}

func TestItemPickup_WeaponTraceable_PlayerNil(t *testing.T) {
	e := ItemPickup{
		Weapon: common.Equipment{Weapon: common.EqAK47},
		Player: nil,
	}

	assert.Equal(t, e.Weapon, *e.WeaponTraceable())
}

func TestItemPickup_WeaponTraceable_WeaponFound(t *testing.T) {
	wep := &common.Equipment{
		EntityID: 1,
		Weapon:   common.EqAK47,
	}
	e := ItemPickup{
		Weapon: common.Equipment{Weapon: common.EqAK47},
		Player: &common.Player{RawWeapons: map[int]*common.Equipment{
			1: wep,
		}},
	}

	assert.Equal(t, wep, e.WeaponTraceable())
}

func TestItemPickup_WeaponTraceable_WeaponNotFound(t *testing.T) {
	wep := &common.Equipment{
		EntityID: 1,
		Weapon:   common.EqAK47,
	}
	e := ItemPickup{
		Weapon: common.Equipment{Weapon: common.EqKnife},
		Player: &common.Player{RawWeapons: map[int]*common.Equipment{
			1: wep,
		}},
	}

	assert.Equal(t, e.Weapon, *e.WeaponTraceable())
}

type demoInfoProviderMock struct {
}

func (p demoInfoProviderMock) IngameTick() int {
	return 0
}

func (p demoInfoProviderMock) TickRate() float64 {
	return 128
}

func (p demoInfoProviderMock) FindPlayerByHandle(handle int) *common.Player {
	return nil
}
