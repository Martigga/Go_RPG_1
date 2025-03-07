package enemy

import (
	"math/rand"
	"start/game"
)

// Erstes Wesen
type Wesen struct {
	level int
	hp    int
	att   int
	def   int
	name  string
}

///////////////////////////Create Enemy////////////////////////////////

// Gets Game_level
// Creates Enemy depending on Game_level
// Returns Enemy Name and Enemy Stats
func CreateEnemy(world string, world_barrier int) (string, [4]int) {
	var enemy_level = game.SetEnemyLevel(world, world_barrier)

	// Decides on Enemy Typ + creates empty Variables for it
	var typ int
	var enemyStats [4]int
	var enemyName string

	// Creates Stats and Name of Enemy depending on Enemy Typ + Enemy level
	var gegner = NewEnemy()
	gegner.SetEnemyTyp(typ)
	enemyStats = gegner.GetStatsEnemy(enemy_level)
	enemyName = gegner.name

	switch enemyName {
	case "Ork":
		typ = 1
	case "Wolf":
		typ = 2
	case "Baer":
		typ = 3
	case "Raeuber":
		typ = 4
	case "Burgritter":
		typ = 5
	case "betrunkener Dorfbewohner":
		typ = 6
	case "Mutantenratte":
		typ = 7
	case "dreikoepfige Schlange":
		typ = 8
	case "laufende Makrowelle":
		typ = 9
	case "Pluenderer":
		typ = 10
	case "Triceratops":
		typ = 11
	case "Pterodactylus":
		typ = 12
	case "prehistorische Bergsteigerziege":
		typ = 13
	case "Johannes Bleileber":
		typ = 14
	case "Juan Pared Caballero":
		typ = 15
	case "Sheriff Joe Swanson":
		typ = 16
	case "Koyot":
		typ = 17
	default:
		typ = rand.Intn(17) + 1
	}

	return enemyName, enemyStats
}

// Gets Enemy_level
// Creates stats based on typ + level
// Returns Stats
func (w *Wesen) GetStatsEnemy(enemy_level int) [4]int {

	w.level = enemy_level

	var stats [4]int

	switch w.name {
	case "Ork":
		w.hp = 13 + (7 * w.level)
		w.att = 4 + (3 * w.level)
		w.def = 4 + w.level
	case "Wolf":
		w.hp = 11 + (3 * w.level)
		w.att = 6 + (4 * w.level)
		w.def = 3 + (w.level / 2)
	case "Baer":
		w.hp = 13 + (7 * w.level)
		w.att = 6 + (4 * w.level)
		w.def = 3 + w.level
	case "Raeuber":
		w.hp = 12 + (4 * w.level)
		w.att = 7 + (4 * w.level)
		w.def = 3 + w.level
	case "Burgritter":
		w.hp = 8 + (2 * w.level)
		w.att = 8 + (4 * w.level)
		w.def = 7 + (5 * w.level)
	case "betrunkener Dorfbewohner":
		w.hp = 5 + (2 * w.level)
		w.att = 6 + (5 * w.level)
		w.def = 6 + (7 * w.level)
	case "Mutantenratte":
		w.hp = 15 + (6 * w.level)
		w.att = 4 + (2 * w.level)
		w.def = 3 + w.level
	case "dreikoepfige Schlange":
		w.hp = 15 + (3 * w.level)
		w.att = 7 + (3 * w.level)
		w.def = 5 + (3 * w.level)
	case "laufende Makrowelle":
		w.hp = 4 + w.level
		w.att = 4 + (5 * w.level)
		w.def = 3 + (3 * w.level)
	case "Pluenderer":
		w.hp = 10 + (3 * w.level)
		w.att = 5 + (5 * w.level)
		w.def = 3 + w.level
	case "Triceratops":
		w.hp = 16 + (7 * w.level)
		w.att = 4 + (4 * w.level)
		w.def = 3 + (2 * w.level)
	case "Pterodactylus":
		w.hp = 7 + (7 * w.level)
		w.att = 6 + (5 * w.level)
		w.def = 3 + w.level
	case "prehistorische Bergsteiger Ziege":
		w.hp = 5 + (5 * w.level)
		w.att = 6 + (5 * w.level)
		w.def = 6 + (5 * w.level)
	case "Johannes Bleileber":
		w.hp = 10 + (6 * w.level)
		w.att = 6 + (6 * w.level)
		w.def = 5 + w.level
	case "Juan Pared Caballero":
		w.hp = 14 + (6 * w.level)
		w.att = 5 + (4 * w.level)
		w.def = 5 + w.level
	case "Sheriff Joe Swanson":
		w.hp = 12 + (6 * w.level)
		w.att = 5 + (4 * w.level)
		w.def = 5 + w.level
	case "Koyot":
		w.hp = 7 + (3 * w.level)
		w.att = 8 + (3 * w.level)
		w.def = 2 + w.level
	}

	stats[0] = w.level
	stats[1] = w.hp
	stats[2] = w.att
	stats[3] = w.def

	return stats
}

// Gets Enemy Typ
// Sets Name of enemy based on given typ
// Returns Nothing
func (w *Wesen) SetEnemyTyp(typ int) {
	switch typ {
	case 1:
		w.name = "Ork"
	case 2:
		w.name = "Wolf"
	case 3:
		w.name = "Baer"
	case 4:
		w.name = "Raeuber"
	case 5:
		w.name = "Burgritter"
	case 6:
		w.name = "betrunkener Dorfbewohner"
	case 7:
		w.name = "Mutantenratte"
	case 8:
		w.name = "dreiköpfige Schlange"
	case 9:
		w.name = "laufende Makrowelle"
	case 10:
		w.name = "Pluenderer"
	case 11:
		w.name = "Triceratops"
	case 12:
		w.name = "Pterodactylus"
	case 13:
		w.name = "prehistorische Bergsteiger Ziege"
	case 14:
		w.name = "Johannes Bleileber"
	case 15:
		w.name = "Juan Pared Caballero"
	case 16:
		w.name = "Sheriff Joe Swanson"
	case 17:
		w.name = "Koyot"
	default:
		w.name = "Ork"
	}
}

/////////////////// Creating New Enemy //////////////////////

func NewEnemy() *Wesen {
	var enemy *Wesen = new(Wesen)
	return enemy
}
