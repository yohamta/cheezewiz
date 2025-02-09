package scene

import (
	"cheezewiz/internal/attacks"
	"cheezewiz/internal/entity"
	"cheezewiz/internal/input"
	"cheezewiz/internal/mediator"
	"cheezewiz/internal/system"
	"cheezewiz/pkg/taskrunner"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}
type Scene struct {
	world     donburi.World
	systems   []System
	drawables []Drawable
}

const level1 string = "level1.json"

func Init() *Scene {
	// World
	world := donburi.NewWorld()

	level := loadWorld(level1)

	// Mediators
	attackMediator := &mediator.Attack{}

	// System
	renderer := system.NewRender()
	collision := system.NewCollision(attackMediator)
	timer := system.NewTimer()
	exp := system.NewExpbar()
	damageGroup := system.NewDamagebufferGroup()
	aicontroller := system.NewEnemyControl()

	attackMediator.D = &damageGroup
	attackMediator.C = collision

	taskrunner.Add(time.Millisecond*800, attacks.CheeseMissile(world, attackMediator))

	s := &Scene{
		world: world,
		systems: []System{
			renderer,
			system.NewPlayerControl(),
			timer,
			system.NewRegisterPlayer(),
			&damageGroup,
			aicontroller,
			collision,
			system.NewScheduler(level.Events, world),
			system.NewWorldViewPortLocation(),
			system.NewProjectileContol(),
			exp,
		},
		drawables: []Drawable{
			collision,
			renderer,
			timer,
			// exp,
		},
	}

	addEntities(world)

	return s
}

func addEntities(world donburi.World) {
	entity.MakeExpBar(world)
	entity.MakeWorld(world)
	entity.MakeBackground(world)
	entity.MakeTimer(world)
	entity.MakePlayer(world, input.Keyboard{})
	entity.MakeSlot(world)
}

func (s *Scene) Update() {
	for _, sys := range s.systems {
		sys.Update(s.world)
	}
	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}
func (s *Scene) Draw(screen *ebiten.Image) {
	for _, sys := range s.drawables {
		sys.Draw(s.world, screen)
	}
}

func (s *Scene) Exit() {
	os.Exit(0)
}
