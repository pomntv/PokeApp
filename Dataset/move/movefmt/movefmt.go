package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := `

	Name
Type
Cat.
Power
Acc.
PP
Effect
Prob. (%)
10,000,000 Volt Thunderbolt	ELECTRIC	Special	195	—	1	Pikachu-exclusive Z-Move. High critical hit ratio.	—
Absorb	GRASS	Special	20	100	25	User recovers half the HP inflicted on opponent.	—
Accelerock	ROCK	Physical	40	100	20	User attacks first.	—
Acid	POISON	Special	40	100	30	May lower opponent's Special Defense.	10
Acid Armor	POISON	Status	—	—	20	Sharply raises user's Defense.	—
Acid Downpour	POISON	—	—	—	1	Poison type Z-Move.	—
Acid Spray	POISON	Special	40	100	20	Sharply lowers opponent's Special Defense.	100
Acrobatics	FLYING	Physical	55	100	15	Stronger when the user does not have a held item.	—
Acupressure	NORMAL	Status	—	—	30	Sharply raises a random stat.	—
Aerial Ace	FLYING	Physical	60	∞	20	Ignores Accuracy and Evasiveness.	—
Aeroblast	FLYING	Special	100	95	5	High critical hit ratio.	—
After You	NORMAL	Status	—	—	15	Gives target priority in the next turn.	—
Agility	PSYCHIC	Status	—	—	30	Sharply raises user's Speed.	—
Air Cutter	FLYING	Special	60	95	25	High critical hit ratio.	—
Air Slash	FLYING	Special	75	95	15	May cause flinching.	30
All-Out Pummeling	FIGHTING	—	—	—	1	Fighting type Z-Move.	—
Ally Switch	PSYCHIC	Status	—	—	15	User switches with opposite teammate.	—
Amnesia	PSYCHIC	Status	—	—	20	Sharply raises user's Special Defense.	—
Anchor Shot	STEEL	Physical	80	100	20	The user entangles the target with its anchor chain while attacking. The target becomes unable to flee.	—
Ancient Power	ROCK	Special	60	100	5	May raise all user's stats at once.	10
Apple Acid	GRASS	Special	80	100	10	Lowers target's Special Defense.	100
Aqua Cutter	WATER	Physical	70	100	20	High critical hit ratio.	—
Aqua Jet	WATER	Physical	40	100	20	User attacks first.	—
Aqua Ring	WATER	Status	—	—	20	Restores a little HP each turn.	—
Aqua Step	WATER	Physical	80	100	10	Raises user's Speed.	100
Aqua Tail	WATER	Physical	90	90	10		—
Arm Thrust	FIGHTING	Physical	15	100	20	Hits 2-5 times in one turn.	—
Armor Cannon	FIRE	Special	120	100	5	Lowers user's Defense and Special Defense.	—
Aromatherapy	GRASS	Status	—	—	5	Cures all status problems in your party.	—
Aromatic Mist	FAIRY	Status	—	—	20	Raises Special Defense of an ally.	—
Assist	NORMAL	Status	—	—	20	User performs a move known by its allies at random.	—
Assurance	DARK	Physical	60	100	10	Power doubles if opponent already took damage in the same turn.	—
Astonish	GHOST	Physical	30	100	15	May cause flinching.	30
Astral Barrage	GHOST	Special	120	100	5	The user attacks by sending a frightful amount of small ghosts at opposing Pokémon.	—
Attack Order	BUG	Physical	90	100	15	High critical hit ratio.	—
Attract	NORMAL	Status	—	100	15	If opponent is the opposite gender, it's less likely to attack.	—
Aura Sphere	FIGHTING	Special	80	∞	20	Ignores Accuracy and Evasiveness.	—
Aura Wheel	ELECTRIC	Physical	110	100	10	Changes type based on Morpeko's Mode.	100
Aurora Beam	ICE	Special	65	100	20	May lower opponent's Attack.	10
Aurora Veil	ICE	Status	—	—	20	Halves damage from Physical and Special attacks for five turns.	—
Autotomize	STEEL	Status	—	—	15	Reduces weight and sharply raises Speed.	—
Avalanche	ICE	Physical	60	100	10	Power doubles if user took damage first.	—
Axe Kick	FIGHTING	Physical	120	90	10	May confuse opponent. If it misses, the user loses HP.	—
Baby-Doll Eyes	FAIRY	Status	—	100	30	Always goes first. Lowers the target's attack.	—
Baddy Bad	DARK	Special	90	100	15	Reduces damage from Physical attacks.	—
Baneful Bunker	POISON	Status	—	—	10	Protects the user and poisons opponent on contact.	—
Barb Barrage	POISON	Physical	60	100	10	Inflicts double damage if the target has a status condition.	—
Barrage	NORMAL	Physical	15	85	20	Hits 2-5 times in one turn.	—
Barrier	PSYCHIC	Status	—	—	20	Sharply raises user's Defense.	—
Baton Pass	NORMAL	Status	—	—	40	User switches out and gives stat changes to the incoming Pokémon.	—
Beak Blast	FLYING	Physical	100	100	15	The user first heats up its beak, and then it attacks the target. Making direct contact with the Pokémon while it's heating up its beak results in a burn.	—
Beat Up	DARK	Physical	—	100	10	Each Pokémon in user's party attacks.	—
Behemoth Bash	STEEL	Physical	100	100	5	Damage doubles if target is Dynamaxed.	—
Behemoth Blade	STEEL	Physical	100	100	5	Damage doubles if target is Dynamaxed.	—
Belch	POISON	Special	120	90	10	User must have consumed a Berry.	—
Belly Drum	NORMAL	Status	—	—	10	User loses 50% of its max HP, but Attack raises to maximum.	—
Bestow	NORMAL	Status	—	—	15	Gives the user's held item to the target.	—
Bide	NORMAL	Physical	—	—	10	User takes damage for two turns then strikes back double.	—
Bind	NORMAL	Physical	15	85	20	Traps opponent, damaging them for 4-5 turns.	100
Bite	DARK	Physical	60	100	25	May cause flinching.	30
Bitter Blade	FIRE	Physical	90	100	10	User recovers half the HP inflicted on opponent.	—
Bitter Malice	GHOST	Special	75	100	10	Inflicts double damage if the target has a status condition.	—
Black Hole Eclipse	DARK	—	—	—	1	Dark type Z-Move.	—
Blast Burn	FIRE	Special	150	90	5	User must recharge next turn.	—
Blaze Kick	FIRE	Physical	85	90	10	High critical hit ratio. May burn opponent.	10
Blazing Torque	FIRE	Physical	80	100	10		—
Bleakwind Storm	FLYING	Special	100	80	10	May cause frostbite.	—
Blizzard	ICE	Special	110	70	5	May freeze opponent.	10
Block	NORMAL	Status	—	—	5	Opponent cannot flee or switch.	—
Bloom Doom	GRASS	—	—	—	1	Grass type Z-Move.	—
Blue Flare	FIRE	Special	130	85	5	May burn opponent.	20
Body Press	FIGHTING	Physical	80	100	10	The higher the user's Defense, the stronger the attack.	—
Body Slam	NORMAL	Physical	85	100	15	May paralyze opponent.	30
Bolt Beak	ELECTRIC	Physical	85	100	10	If the user attacks before the target, the power of this move is doubled.	—
Bolt Strike	ELECTRIC	Physical	130	85	5	May paralyze opponent.	20
Bone Club	GROUND	Physical	65	85	20	May cause flinching.	10
Bone Rush	GROUND	Physical	25	90	10	Hits 2-5 times in one turn.	—
Bonemerang	GROUND	Physical	50	90	10	Hits twice in one turn.	—
Boomburst	NORMAL	Special	140	100	10	Hits all adjacent Pokémon.	—
Bounce	FLYING	Physical	85	85	5	Springs up on first turn, attacks on second. May paralyze opponent.	30
Bouncy Bubble	WATER	Special	90	100	15	User recovers half the HP inflicted on opponent.	—
Branch Poke	GRASS	Physical	40	100	40		—
Brave Bird	FLYING	Physical	120	100	15	User receives recoil damage.	—
Breaking Swipe	DRAGON	Physical	60	100	15	Hits multiple opponents and lowers their attack.	100
Breakneck Blitz	NORMAL	—	—	—	1	Normal type Z-Move.	—
Brick Break	FIGHTING	Physical	75	100	15	Breaks through Reflect and Light Screen barriers.	—
Brine	WATER	Special	65	100	10	Power doubles if opponent's HP is less than 50%.	—
Brutal Swing	DARK	Physical	60	100	20	The user swings its body around violently to inflict damage on everything in its vicinity.	—
Bubble	WATER	Special	40	100	30	May lower opponent's Speed.	10
Bubble Beam	WATER	Special	65	100	20	May lower opponent's Speed.	10
Bug Bite	BUG	Physical	60	100	20	Receives the effect from the opponent's held berry.	—
Bug Buzz	BUG	Special	90	100	10	May lower opponent's Special Defense.	10
Bulk Up	FIGHTING	Status	—	—	20	Raises user's Attack and Defense.	—
Bulldoze	GROUND	Physical	60	100	20	Lowers opponent's Speed.	100
Bullet Punch	STEEL	Physical	40	100	30	User attacks first.	—
Bullet Seed	GRASS	Physical	25	100	30	Hits 2-5 times in one turn.	—
Burn Up	FIRE	Special	130	100	5	To inflict massive damage, the user burns itself out. After using this move, the user will no longer be Fire type.	—
Burning Jealousy	FIRE	Special	70	100	5	Hits all opponents, and burns any that have had their stats boosted.	—
Buzzy Buzz	ELECTRIC	Special	90	100	15	Paralyzes the opponent.	100
Calm Mind	PSYCHIC	Status	—	—	20	Raises user's Special Attack and Special Defense.	—
Camouflage	NORMAL	Status	—	—	20	Changes user's type according to the location.	—
Captivate	NORMAL	Status	—	100	20	Sharply lowers opponent's Special Attack if opposite gender.	—
Catastropika	ELECTRIC	Physical	210	—	1	Pikachu-exclusive Z-Move.	—
Ceaseless Edge	DARK	Physical	65	90	15	High critical hit ratio. Damages target with splinters each turn.	—
Celebrate	NORMAL	Status	—	—	40	The Pokémon congratulates you on your special day. No battle effect.	—
Charge	ELECTRIC	Status	—	—	20	Raises user's Special Defense and next Electric move's power increases.	—
Charge Beam	ELECTRIC	Special	50	90	10	May raise user's Special Attack.	70
Charm	FAIRY	Status	—	100	20	Sharply lowers opponent's Attack.	—
Chatter	FLYING	Special	65	100	20	Confuses opponent.	100
Chilling Water	WATER	Special	50	100	20	Lowers opponent's Attack.	—
Chilly Reception	ICE	Status	—	—	10	Switches out and summons a snowstorm lasting 5 turns.	—
Chip Away	NORMAL	Physical	70	100	20	Ignores opponent's stat changes.	—
Chloroblast	GRASS	Special	150	95	5	User receives recoil damage.	—
Circle Throw	FIGHTING	Physical	60	90	10	In battles, the opponent switches. In the wild, the Pokémon runs.	—
Clamp	WATER	Physical	35	85	15	Traps opponent, damaging them for 4-5 turns.	100
Clanging Scales	DRAGON	Special	110	100	5	Lowers user's Defense.	100
Clangorous Soul	DRAGON	Status	—	100	5	Raises all user's stats but loses HP.	100
Clangorous Soulblaze	DRAGON	Special	185	—	1	Kommo-o exclusive Z-Move.	—
Clear Smog	POISON	Special	50	—	15	Removes all of the target's stat changes.	—
Close Combat	FIGHTING	Physical	120	100	5	Lowers user's Defense and Special Defense.	100
Coaching	FIGHTING	Status	—	—	10	Boosts Attack and Defense of a teammate.	—
Coil	POISON	Status	—	—	20	Raises user's Attack, Defense and Accuracy.	—
Collision Course	FIGHTING	Physical	100	100	5	Boosted even more if it's super-effective.	—
Combat Torque	FIGHTING	Physical	100	100	10		—
Comet Punch	NORMAL	Physical	18	85	15	Hits 2-5 times in one turn.	—
Comeuppance	DARK	Physical	1	100	10	Deals more damage to the opponent that last inflicted damage on it.	—
Confide	NORMAL	Status	—	—	20	Lowers opponent's Special Attack.	100
Confuse Ray	GHOST	Status	—	100	10	Confuses opponent.	—
Confusion	PSYCHIC	Special	50	100	25	May confuse opponent.	10
Constrict	NORMAL	Physical	10	100	35	May lower opponent's Speed by one stage.	10
Continental Crush	ROCK	—	—	—	1	Rock type Z-Move.	—
Conversion	NORMAL	Status	—	—	30	Changes user's type to that of its first move.	—
Conversion 2	NORMAL	Status	—	—	30	User changes type to become resistant to opponent's last move.	—
Copycat	NORMAL	Status	—	—	20	Copies opponent's last move.	—
Core Enforcer	DRAGON	Special	100	100	10	Suppresses the target's ability if the target has already moved.	—
Corkscrew Crash	STEEL	—	—	—	1	Steel type Z-Move.	—
Corrosive Gas	POISON	Status	—	100	40	Removes opponent's items.	—
Cosmic Power	PSYCHIC	Status	—	—	20	Raises user's Defense and Special Defense.	—
Cotton Guard	GRASS	Status	—	—	10	Drastically raises user's Defense.	—
Cotton Spore	GRASS	Status	—	100	40	Sharply lowers opponent's Speed.	—
Counter	FIGHTING	Physical	—	100	20	When hit by a Physical Attack, user strikes back with 2x power.	—
Court Change	NORMAL	Status	—	100	10	Swaps the effects on either side of the field.	—
Covet	NORMAL	Physical	60	100	25	Opponent's item is stolen by the user.	—
Crabhammer	WATER	Physical	100	90	10	High critical hit ratio.	—
Crafty Shield	FAIRY	Status	—	—	10	Protects the Pokémon from status moves.	—
Cross Chop	FIGHTING	Physical	100	80	5	High critical hit ratio.	—
Cross Poison	POISON	Physical	70	100	20	High critical hit ratio. May poison opponent.	10
Crunch	DARK	Physical	80	100	15	May lower opponent's Defense.	20
Crush Claw	NORMAL	Physical	75	95	10	May lower opponent's Defense.	50
Crush Grip	NORMAL	Physical	—	100	5	More powerful when opponent has higher HP.	—
Curse	GHOST	Status	—	—	10	Ghosts lose 50% of max HP and curse the opponent; Non-Ghosts raise Attack, Defense and lower Speed.	—
Cut	NORMAL	Physical	50	95	30		—
Dark Pulse	DARK	Special	80	100	15	May cause flinching.	20
Dark Void	DARK	Status	—	50	10	Puts all adjacent opponents to sleep.	—
Darkest Lariat	DARK	Physical	85	100	10	Ignores opponent's stat changes.	—
Dazzling Gleam	FAIRY	Special	80	100	10	Hits all adjacent opponents.	—
Decorate	FAIRY	Status	—	—	15	Sharply raises target's Attack and Special Attack.	100
Defend Order	BUG	Status	—	—	10	Raises user's Defense and Special Defense.	—
Defense Curl	NORMAL	Status	—	—	40	Raises user's Defense.	—
Defog	FLYING	Status	—	—	15	Lowers opponent's Evasiveness and clears fog.	—
Destiny Bond	GHOST	Status	—	—	5	If the user faints, the opponent also faints.	—
Detect	FIGHTING	Status	—	—	5	Protects the user, but may fail if used consecutively.	—
Devastating Drake	DRAGON	—	—	—	1	Dragon type Z-Move.	—
Diamond Storm	ROCK	Physical	100	95	5	May sharply raise user's Defense.	50
Dig	GROUND	Physical	80	100	10	Digs underground on first turn, attacks on second. Can also escape from caves.	—
Dire Claw	POISON	Physical	80	100	15	High critical hit ratio. May poison, paralyze or make the opponent drowsy.	—
Disable	NORMAL	Status	—	100	20	Opponent can't use its last attack for a few turns.	—
Disarming Voice	FAIRY	Special	40	∞	15	Ignores Accuracy and Evasiveness.	—
Discharge	ELECTRIC	Special	80	100	15	May paralyze opponent.	30
Dive	WATER	Physical	80	100	10	Dives underwater on first turn, attacks on second turn.	—
Dizzy Punch	NORMAL	Physical	70	100	10	May confuse opponent.	20
Doodle	NORMAL	Status	—	100	10	Changes the abilities of the user and its teammates to that of the target.	—
Doom Desire	STEEL	Special	140	100	5	Damage occurs 2 turns later.	—
Double Hit	NORMAL	Physical	35	90	10	Hits twice in one turn.	—
Double Iron Bash	STEEL	Physical	60	100	5	Hits twice in one turn; may cause flinching.	30
Double Kick	FIGHTING	Physical	30	100	30	Hits twice in one turn.	—
Double Shock	ELECTRIC	Physical	120	100	5	After using this move, the user will no longer be Electric type.	—
Double Slap	NORMAL	Physical	15	85	10	Hits 2-5 times in one turn.	—
Double Team	NORMAL	Status	—	—	15	Raises user's Evasiveness.	—
Double-Edge	NORMAL	Physical	120	100	15	User receives recoil damage.	—
Draco Meteor	DRAGON	Special	130	90	5	Sharply lowers user's Special Attack.	100
Dragon Ascent	FLYING	Physical	120	100	5	Lowers user's Defense and Special Defense.	100
Dragon Breath	DRAGON	Special	60	100	20	May paralyze opponent.	30
Dragon Claw	DRAGON	Physical	80	100	15		—
Dragon Dance	DRAGON	Status	—	—	20	Raises user's Attack and Speed.	—
Dragon Darts	DRAGON	Physical	50	100	10	User attacks twice.	—
Dragon Energy	DRAGON	Special	150	100	5	The higher the user's HP, the higher the power.	—
Dragon Hammer	DRAGON	Physical	90	100	15	The user uses its body like a hammer to attack the target and inflict damage.	—
Dragon Pulse	DRAGON	Special	85	100	10		—
Dragon Rage	DRAGON	Special	—	100	10	Always inflicts 40 HP.	—
Dragon Rush	DRAGON	Physical	100	75	10	May cause flinching.	20
Dragon Tail	DRAGON	Physical	60	90	10	In battles, the opponent switches. In the wild, the Pokémon runs.	—
Drain Punch	FIGHTING	Physical	75	100	10	User recovers half the HP inflicted on opponent.	—
Draining Kiss	FAIRY	Special	50	100	10	User recovers most the HP inflicted on opponent.	—
Dream Eater	PSYCHIC	Special	100	100	15	User recovers half the HP inflicted on a sleeping opponent.	—
Drill Peck	FLYING	Physical	80	100	20		—
Drill Run	GROUND	Physical	80	95	10	High critical hit ratio.	—
Drum Beating	GRASS	Physical	80	100	10	Lowers opponent's Speed.	100
Dual Chop	DRAGON	Physical	40	90	15	Hits twice in one turn.	—
Dual Wingbeat	FLYING	Physical	40	90	10	The user slams the target with its wings. The target is hit twice in a row.	—
Dynamax Cannon	DRAGON	Special	100	100	5	Damage doubles if opponent is Dynamaxed.	—
Dynamic Punch	FIGHTING	Physical	100	50	5	Confuses opponent.	100
Earth Power	GROUND	Special	90	100	10	May lower opponent's Special Defense.	10
Earthquake	GROUND	Physical	100	100	10	Power is doubled if opponent is underground from using Dig.	—
Echoed Voice	NORMAL	Special	40	100	15	Power increases each turn.	—
Eerie Impulse	ELECTRIC	Status	—	100	15	Sharply lowers opponent's Special Attack.	—
Eerie Spell	PSYCHIC	Special	80	100	5	Deals damage and reduces opponent's PP.	—
Egg Bomb	NORMAL	Physical	100	75	10		—
Electric Terrain	ELECTRIC	Status	—	—	10	Prevents all Pokémon from falling asleep for 5 turns.	—
Electrify	ELECTRIC	Status	—	—	20	Changes the target's move to Electric type.	—
Electro Ball	ELECTRIC	Special	—	100	10	The faster the user, the stronger the attack.	—
Electro Drift	ELECTRIC	Special	100	100	5	Boosted even more if it's super-effective.	—
Electroweb	ELECTRIC	Special	55	95	15	Lowers opponent's Speed.	100
Embargo	DARK	Status	—	100	15	Opponent cannot use items.	—
Ember	FIRE	Special	40	100	25	May burn opponent.	10
Encore	NORMAL	Status	—	100	5	Forces opponent to keep using its last move for 3 turns.	—
Endeavor	NORMAL	Physical	—	100	5	Reduces opponent's HP to same as user's.	—
Endure	NORMAL	Status	—	—	10	Always left with at least 1 HP, but may fail if used consecutively.	—
Energy Ball	GRASS	Special	90	100	10	May lower opponent's Special Defense.	10
Entrainment	NORMAL	Status	—	100	15	Makes target's ability same as user's.	—
Eruption	FIRE	Special	150	100	5	Stronger when the user's HP is higher.	—
Esper Wing	PSYCHIC	Special	80	100	10	High critical hit ratio. Raises user's Speed.	—
Eternabeam	DRAGON	Special	160	90	5	User can't move on the next turn.	—
Expanding Force	PSYCHIC	Special	80	100	10	Increases power and hits all opponents on Psychic Terrain.	—
Explosion	NORMAL	Physical	250	100	5	User faints.	—
Extrasensory	PSYCHIC	Special	80	100	20	May cause flinching.	10
Extreme Evoboost	NORMAL	Status	—	—	1	Eevee-exclusive Z-Move. Sharply raises all stats.	100
Extreme Speed	NORMAL	Physical	80	100	5	User attacks first.	—
Facade	NORMAL	Physical	70	100	20	Power doubles if user is burned, poisoned, or paralyzed.	—
Fairy Lock	FAIRY	Status	—	—	10	Prevents fleeing in the next turn.	—
Fairy Wind	FAIRY	Special	40	100	30		—
Fake Out	NORMAL	Physical	40	100	10	User attacks first, foe flinches. Only usable on first turn.	100
Fake Tears	DARK	Status	—	100	20	Sharply lowers opponent's Special Defense.	—
False Surrender	DARK	Physical	80	∞	10	Ignores Accuracy and Evasiveness.	—
False Swipe	NORMAL	Physical	40	100	40	Always leaves opponent with at least 1 HP.	—
Feather Dance	FLYING	Status	—	100	15	Sharply lowers opponent's Attack.	—
Feint	NORMAL	Physical	30	100	10	Only hits if opponent uses Protect or Detect in the same turn.	—
Feint Attack	DARK	Physical	60	∞	20	Ignores Accuracy and Evasiveness.	—
Fell Stinger	BUG	Physical	50	100	25	Drastically raises user's Attack if target is KO'd.	—
Fiery Dance	FIRE	Special	80	100	10	May raise user's Special Attack.	50
Fiery Wrath	DARK	Special	90	100	10	May cause flinching.	20
Fillet Away	NORMAL	Status	—	—	10	Lowers HP but sharply boosts Attack, Special Attack, and Speed.	—
Final Gambit	FIGHTING	Special	—	100	5	Inflicts damage equal to the user's remaining HP. User faints.	—
Fire Blast	FIRE	Special	110	85	5	May burn opponent.	10
Fire Fang	FIRE	Physical	65	95	15	May cause flinching and/or burn opponent.	10
Fire Lash	FIRE	Physical	80	100	15	The user strikes the target with a burning lash. This also lowers the target's Defense stat.	100
Fire Pledge	FIRE	Special	80	100	10	Added effects appear if combined with Grass Pledge or Water Pledge.	—
Fire Punch	FIRE	Physical	75	100	15	May burn opponent.	10
Fire Spin	FIRE	Special	35	85	15	Traps opponent, damaging them for 4-5 turns.	100
First Impression	BUG	Physical	90	100	10	Although this move has great power, it only works the first turn the user is in battle.	—
Fishious Rend	WATER	Physical	85	100	10	If the user attacks before the target, the power of this move is doubled.	—
Fissure	GROUND	Physical	—	30	5	One-Hit-KO, if it hits.	—
Flail	NORMAL	Physical	—	100	15	The lower the user's HP, the higher the power.	—
Flame Burst	FIRE	Special	70	100	15	May also injure nearby Pokémon.	—
Flame Charge	FIRE	Physical	50	100	20	Raises user's Speed.	100
Flame Wheel	FIRE	Physical	60	100	25	May burn opponent.	10
Flamethrower	FIRE	Special	90	100	15	May burn opponent.	10
Flare Blitz	FIRE	Physical	120	100	15	User receives recoil damage. May burn opponent.	10
Flash	NORMAL	Status	—	100	20	Lowers opponent's Accuracy.	—
Flash Cannon	STEEL	Special	80	100	10	May lower opponent's Special Defense.	10
Flatter	DARK	Status	—	100	15	Confuses opponent, but raises its Special Attack.	—
Fleur Cannon	FAIRY	Special	130	90	5	Sharply lowers user's Special Attack.	100
Fling	DARK	Physical	—	100	10	Power depends on held item.	—
Flip Turn	WATER	Physical	60	100	20	After making its attack, the user rushes back to switch places with a party Pokémon in waiting.	—
Floaty Fall	FLYING	Physical	90	95	15	May cause flinching.	30
Floral Healing	FAIRY	Status	—	—	10	The user restores the target's HP by up to half of its max HP. It restores more HP when the terrain is grass.	—
Flower Shield	FAIRY	Status	—	—	10	Sharply raises Defense of all Grass-type Pokémon on the field.	100
Flower Trick	GRASS	Physical	70	∞	10	Never misses; always results in a critical hit.	—
Fly	FLYING	Physical	90	95	15	Flies up on first turn, attacks on second turn.	—
Flying Press	FIGHTING	Physical	100	95	10	Deals Fighting and Flying type damage.	—
Focus Blast	FIGHTING	Special	120	70	5	May lower opponent's Special Defense.	10
Focus Energy	NORMAL	Status	—	—	30	Increases critical hit ratio.	—
Focus Punch	FIGHTING	Physical	150	100	20	If the user is hit before attacking, it flinches instead.	—
Follow Me	NORMAL	Status	—	—	20	In Double Battle, the user takes all the attacks.	—
Force Palm	FIGHTING	Physical	60	100	10	May paralyze opponent.	30
Foresight	NORMAL	Status	—	—	40	Resets opponent's Evasiveness, and allows Normal- and Fighting-type attacks to hit Ghosts.	—
Forest's Curse	GRASS	Status	—	100	20	Adds Grass type to opponent.	—
Foul Play	DARK	Physical	95	100	15	Uses the opponent's Attack stat.	—
Freeze Shock	ICE	Physical	140	90	5	Charges on first turn, attacks on second. May paralyze opponent.	30
Freeze-Dry	ICE	Special	70	100	20	May freeze opponent. Super-effective against Water types.	10
Freezing Glare	PSYCHIC	Special	90	100	10	May freeze opponent.	10
Freezy Frost	ICE	Special	90	100	15	Resets all stat changes.	—
Frenzy Plant	GRASS	Special	150	90	5	User must recharge next turn.	—
Frost Breath	ICE	Special	60	90	10	Always results in a critical hit.	100
Frustration	NORMAL	Physical	—	100	20	Power decreases with higher Friendship.	—
Fury Attack	NORMAL	Physical	15	85	20	Hits 2-5 times in one turn.	—
Fury Cutter	BUG	Physical	40	95	20	Power increases each turn.	—
Fury Swipes	NORMAL	Physical	18	80	15	Hits 2-5 times in one turn.	—
Fusion Bolt	ELECTRIC	Physical	100	100	5	Power increases if Fusion Flare is used in the same turn.	—
Fusion Flare	FIRE	Special	100	100	5	Power increases if Fusion Bolt is used in the same turn.	—
Future Sight	PSYCHIC	Special	120	100	10	Damage occurs 2 turns later.	—
G-Max Befuddle	BUG	—	—	∞	5	Butterfree-exclusive G-Max Move. Poisons, paralyzes, or puts opponent to sleep.	100
G-Max Cannonade	WATER	—	—	∞	10	Blastoise-exclusive G-Max Move. Damages non-Water types for 4 turns.	—
G-Max Centiferno	FIRE	—	—	∞	5	Centiskorch-exclusive G-Max Move. Traps opponents for 4-5 turns.	100
G-Max Chi Strike	FIGHTING	—	—	∞	5	Machamp-exclusive G-Max Move. Increases critical hit ratio.	—
G-Max Cuddle	NORMAL	—	—	∞	5	Eevee-exclusive G-Max Move. Infatuates opponents.	100
G-Max Depletion	DRAGON	—	—	∞	5	Duraludon-exclusive G-Max Move. Reduces opponent's PP.	—
G-Max Drum Solo	GRASS	—	160	∞	5	Rillaboom-exclusive G-Max Move. Ignores target's ability.	—
G-Max Finale	FAIRY	—	—	∞	5	Alcremie-exclusive G-Max Move. Heals the user's team.	—
G-Max Fireball	FIRE	—	160	∞	5	Cinderace-exclusive G-Max Move. Ignores target's ability.	—
G-Max Foam Burst	WATER	—	—	∞	5	Kingler-exclusive G-Max Move. Harshly lowers opponents' Speed.	100
G-Max Gold Rush	NORMAL	—	—	∞	5	Meowth-exclusive G-Max Move. Confuses opponents and earns more money.	100
G-Max Gravitas	PSYCHIC	—	—	∞	5	Orbeetle-exclusive G-Max Move. Summons Gravity for 5 turns.	—
G-Max Hydrosnipe	WATER	—	160	∞	5	Inteleon-exclusive G-Max Move. Ignores target's ability.	—
G-Max Malodor	POISON	—	—	∞	5	Garbodor-exclusive G-Max Move. Poisons opponents.	100
G-Max Meltdown	STEEL	—	—	∞	5	Melmetal-exclusive G-Max Move. Prevents opponents using the same move twice in a row.	—
G-Max One Blow	DARK	—	—	∞	5	Urshifu Single-Strike Style-exclusive G-Max Move. Strikes through Max Guard and Protect.	—
G-Max Rapid Flow	WATER	—	—	∞	5	Urshifu Rapid-Strike Style-exclusive G-Max Move. Strikes through Max Guard and Protect.	—
G-Max Replenish	NORMAL	—	—	∞	5	Snorlax-exclusive G-Max Move. Recycles Berries.	—
G-Max Resonance	ICE	—	—	∞	5	Lapras-exclusive G-Max Move. Reduces damage for 5 turns.	—
G-Max Sandblast	GROUND	—	—	∞	5	Sandaconda-exclusive G-Max Move. Traps opponents for 4-5 turns.	100
G-Max Smite	FAIRY	—	—	∞	5	Hatterene-exclusive G-Max Move. Confuses opponents.	100
G-Max Snooze	DARK	—	—	∞	5	Grimmsnarl-exclusive G-Max Move. Makes opponents drowsy.	—
G-Max Steelsurge	STEEL	—	—	∞	5	Copperajah-exclusive G-Max Move. Sets up Spikes on the field.	—
G-Max Stonesurge	WATER	—	—	∞	5	Drednaw-exclusive G-Max Move. Sets up Stealth Rock.	—
G-Max Stun Shock	ELECTRIC	—	—	∞	10	Toxtricity-exclusive G-Max Move. Poisons or paralyzes opponents.	100
G-Max Sweetness	GRASS	—	—	∞	10	Appletun-exclusive G-Max Move. Heals status conditions of user's team.	—
G-Max Tartness	GRASS	—	—	∞	10	Flapple-exclusive G-Max Move. Reduces opponents' evasiveness.	100
G-Max Terror	GHOST	—	—	∞	10	Gengar-exclusive G-Max Move. Prevents opponent from switching out.	—
G-Max Vine Lash	GRASS	—	—	∞	10	Venusaur-exclusive G-Max Move. Damages non-Grass types for 4 turns.	—
G-Max Volcalith	ROCK	—	—	∞	10	Coalossal-exclusive G-Max Move. Deals damage for 4 turns.	100
G-Max Volt Crash	ELECTRIC	—	—	∞	10	Pikachu-exclusive G-Max Move. Paralyzes opponents.	100
G-Max Wildfire	FIRE	—	—	∞	10	Charizard-exclusive G-Max Move. Damages non-Fire types for 4 turns.	100
G-Max Wind Rage	FLYING	—	—	∞	10	Corviknight-exclusive G-Max Move. Removes battlefield hazards.	—
Gastro Acid	POISON	Status	—	100	10	Cancels out the effect of the opponent's Ability.	—
Gear Grind	STEEL	Physical	50	85	15	Hits twice in one turn.	—
Gear Up	STEEL	Status	—	—	20	The user engages its gears to raise the Attack and Sp. Atk stats of ally Pokémon with the Plus or Minus Ability.	—
Genesis Supernova	PSYCHIC	Special	185	—	1	Mew-exclusive Z-Move.	—
Geomancy	FAIRY	Status	—	—	10	Charges on first turn, sharply raises user's Sp. Attack, Sp. Defense and Speed on the second.	—
Giga Drain	GRASS	Special	75	100	10	User recovers half the HP inflicted on opponent.	—
Giga Impact	NORMAL	Physical	150	90	5	User must recharge next turn.	—
Gigaton Hammer	STEEL	Physical	160	100	5	Cannot be used twice in a row.	—
Gigavolt Havoc	ELECTRIC	—	—	—	1	Electric type Z-Move.	—
Glacial Lance	ICE	Physical	120	100	5	The user attacks by hurling a blizzard-cloaked icicle lance at opposing Pokémon.	—
Glaciate	ICE	Special	65	95	10	Lowers opponent's Speed.	100
Glaive Rush	DRAGON	Physical	120	100	5	Attacks from opposing Pokémon during the next turn cannot miss and will inflict double damage.	—
Glare	NORMAL	Status	—	100	30	Paralyzes opponent.	—
Glitzy Glow	PSYCHIC	Special	90	100	15	Reduces damage from Special attacks.	—
Grass Knot	GRASS	Special	—	100	20	The heavier the opponent, the stronger the attack.	—
Grass Pledge	GRASS	Special	80	100	10	Added effects appear if preceded by Water Pledge or succeeded by Fire Pledge.	—
Grass Whistle	GRASS	Status	—	55	15	Puts opponent to sleep.	—
Grassy Glide	GRASS	Physical	60	100	20	High priority during Grassy Terrain.	—
Grassy Terrain	GRASS	Status	—	—	10	Restores a little HP of all Pokémon for 5 turns.	—
Grav Apple	GRASS	Physical	80	100	10	Lowers the opponent's Defense stat.	100
Gravity	PSYCHIC	Status	—	—	5	Prevents moves like Fly and Bounce and the Ability Levitate for 5 turns.	—
Growl	NORMAL	Status	—	100	40	Lowers opponent's Attack.	—
Growth	NORMAL	Status	—	—	20	Raises user's Attack and Special Attack.	—
Grudge	GHOST	Status	—	—	5	If the users faints after using this move, the PP for the opponent's last move is depleted.	—
Guard Split	PSYCHIC	Status	—	—	10	Averages Defense and Special Defense with the target.	—
Guard Swap	PSYCHIC	Status	—	—	10	User and opponent swap Defense and Special Defense.	—
Guardian of Alola	FAIRY	Special	—	—	1	Tapu-exclusive Z-move. Cuts opponent's HP by 75%.	—
Guillotine	NORMAL	Physical	—	30	5	One-Hit-KO, if it hits.	—
Gunk Shot	POISON	Physical	120	80	5	May poison opponent.	30
Gust	FLYING	Special	40	100	35	Hits Pokémon using Fly/Bounce/Sky Drop with double power.	—
Gyro Ball	STEEL	Physical	—	100	5	The slower the user, the stronger the attack.	—
Hail	ICE	Status	—	—	10	Non-Ice types are damaged for 5 turns.	—
Hammer Arm	FIGHTING	Physical	100	90	10	Lowers user's Speed.	100
Happy Hour	NORMAL	Status	—	—	30	Doubles prize money from trainer battles.	—
Harden	NORMAL	Status	—	—	30	Raises user's Defense.	—
Haze	ICE	Status	—	—	30	Resets all stat changes.	—
Head Charge	NORMAL	Physical	120	100	15	User receives recoil damage.	—
Head Smash	ROCK	Physical	150	80	5	User receives recoil damage.	—
Headbutt	NORMAL	Physical	70	100	15	May cause flinching.	30
Headlong Rush	GROUND	Physical	120	100	5	Lowers user's Defense.	—
Heal Bell	NORMAL	Status	—	—	5	Heals the user's party's status conditions.	—
Heal Block	PSYCHIC	Status	—	100	15	Prevents the opponent from restoring HP for 5 turns.	—
Heal Order	BUG	Status	—	—	10	User recovers half its max HP.	—
Heal Pulse	PSYCHIC	Status	—	—	10	Restores half the target's max HP.	—
Healing Wish	PSYCHIC	Status	—	—	10	The user faints and the next Pokémon released is fully healed.	—
Heart Stamp	PSYCHIC	Physical	60	100	25	May cause flinching.	30
Heart Swap	PSYCHIC	Status	—	—	10	Stat changes are swapped with the opponent.	—
Heat Crash	FIRE	Physical	—	100	10	The heavier the user, the stronger the attack.	—
Heat Wave	FIRE	Special	95	90	10	May burn opponent.	10
Heavy Slam	STEEL	Physical	—	100	10	The heavier the user, the stronger the attack.	—
Helping Hand	NORMAL	Status	—	—	20	In Double Battles, boosts the power of the partner's move.	—
Hex	GHOST	Special	65	100	10	Inflicts more damage if the target has a status condition.	—
Hidden Power	NORMAL	Special	60	100	15	Type and power depends on user's IVs.	—
High Horsepower	GROUND	Physical	95	95	10	The user fiercely attacks the target using its entire body.	—
High Jump Kick	FIGHTING	Physical	130	90	10	If it misses, the user loses half their HP.	—
Hold Back	NORMAL	Physical	40	100	40	Always leaves opponent with at least 1 HP.	—
Hold Hands	NORMAL	Status	—	—	40	Makes the user and an ally very happy.	—
Hone Claws	DARK	Status	—	—	15	Raises user's Attack and Accuracy.	—
Horn Attack	NORMAL	Physical	65	100	25		—
Horn Drill	NORMAL	Physical	—	30	5	One-Hit-KO, if it hits.	—
Horn Leech	GRASS	Physical	75	100	10	User recovers half the HP inflicted on opponent.	—
Howl	NORMAL	Status	—	—	40	Raises Attack of allies.	—
Hurricane	FLYING	Special	110	70	10	May confuse opponent.	30
Hydro Cannon	WATER	Special	150	90	5	User must recharge next turn.	—
Hydro Pump	WATER	Special	110	80	5		—
Hydro Steam	WATER	Special	80	100	15	Power increases in harsh sunlight.	—
Hydro Vortex	WATER	—	—	—	1	Water type Z-Move.	—
Hyper Beam	NORMAL	Special	150	90	5	User must recharge next turn.	—
Hyper Drill	NORMAL	Physical	100	100	5	Can strike through Protect/Detect.	—
Hyper Fang	NORMAL	Physical	80	90	15	May cause flinching.	10
Hyper Voice	NORMAL	Special	90	100	10		—
Hyperspace Fury	DARK	Physical	100	∞	5	Lowers user's Defense. Can strike through Protect/Detect.	100
Hyperspace Hole	PSYCHIC	Special	80	∞	5	Can strike through Protect/Detect.	—
Hypnosis	PSYCHIC	Status	—	60	20	Puts opponent to sleep.	—
Ice Ball	ICE	Physical	30	90	20	Doubles in power each turn for 5 turns.	—
Ice Beam	ICE	Special	90	100	10	May freeze opponent.	10
Ice Burn	ICE	Special	140	90	5	Charges on first turn, attacks on second. May burn opponent.	30
Ice Fang	ICE	Physical	65	95	15	May cause flinching and/or freeze opponent.	10
Ice Hammer	ICE	Physical	100	90	10	The user swings and hits with its strong, heavy fist. It lowers the user's Speed, however.	100
Ice Punch	ICE	Physical	75	100	15	May freeze opponent.	10
Ice Shard	ICE	Physical	40	100	30	User attacks first.	—
Ice Spinner	ICE	Physical	80	100	15	Removes effects of Terrain.	—
Icicle Crash	ICE	Physical	85	90	10	May cause flinching.	30
Icicle Spear	ICE	Physical	25	100	30	Hits 2-5 times in one turn.	—
Icy Wind	ICE	Special	55	95	15	Lowers opponent's Speed.	100
Imprison	PSYCHIC	Status	—	—	10	Opponent is unable to use moves that the user also knows.	—
Incinerate	FIRE	Special	60	100	15	Destroys the target's held berry.	—
Infernal Parade	GHOST	Special	60	100	15	Inflicts double damage if the target has a status condition.	—
Inferno	FIRE	Special	100	50	5	Burns opponent.	100
Inferno Overdrive	FIRE	—	—	—	1	Fire type Z-Move.	—
Infestation	BUG	Special	20	100	20	Traps opponent, damaging them for 4-5 turns.	100
Ingrain	GRASS	Status	—	—	20	User restores HP each turn. User cannot escape/switch.	—
Instruct	PSYCHIC	Status	—	—	15	Allows an ally to use a move instead.	—
Ion Deluge	ELECTRIC	Status	—	—	25	Changes Normal-type moves to Electric-type.	—
Iron Defense	STEEL	Status	—	—	15	Sharply raises user's Defense.	—
Iron Head	STEEL	Physical	80	100	15	May cause flinching.	30
Iron Tail	STEEL	Physical	100	75	15	May lower opponent's Defense.	30
Jaw Lock	DARK	Physical	80	100	10	Prevents user and opponent from switching out.	—
Jet Punch	WATER	Physical	60	100	15	Always goes first.	—
Judgment	NORMAL	Special	100	100	10	Type depends on the Arceus Plate being held.	—
Jump Kick	FIGHTING	Physical	100	95	10	If it misses, the user loses half their HP.	—
Jungle Healing	GRASS	Status	—	—	10	Restores team's HP and cures status conditions.	—
Karate Chop	FIGHTING	Physical	50	100	25	High critical hit ratio.	—
Kinesis	PSYCHIC	Status	—	80	15	Lowers opponent's Accuracy.	—
King's Shield	STEEL	Status	—	—	10	Protects the user and lowers opponent's Attack on contact.	—
Knock Off	DARK	Physical	65	100	20	Removes opponent's held item for the rest of the battle.	—
Kowtow Cleave	DARK	Physical	85	∞	10	Always hits.	—
Land's Wrath	GROUND	Physical	90	100	10		—
Laser Focus	NORMAL	Status	—	—	30	User's next attack is guaranteed to result in a critical hit.	—
Lash Out	DARK	Physical	75	100	5	Double power if stats were lowered during the turn.	—
Last Resort	NORMAL	Physical	140	100	5	Can only be used after all other moves are used.	—
Last Respects	GHOST	Physical	50	100	10	Damages increases the more party Pokémon have been defeated.	—
Lava Plume	FIRE	Special	80	100	15	May burn opponent.	30
Leaf Blade	GRASS	Physical	90	100	15	High critical hit ratio.	—
Leaf Storm	GRASS	Special	130	90	5	Sharply lowers user's Special Attack.	100
Leaf Tornado	GRASS	Special	65	90	10	May lower opponent's Accuracy.	50
Leafage	GRASS	Physical	40	100	40	Strikes opponent with leaves.	—
Leech Life	BUG	Physical	80	100	10	User recovers half the HP inflicted on opponent.	—
Leech Seed	GRASS	Status	—	90	10	Drains HP from opponent each turn.	—
Leer	NORMAL	Status	—	100	30	Lowers opponent's Defense.	100
Let's Snuggle Forever	FAIRY	Physical	190	—	1	Mimikyu-exclusive Z-Move.	—
Lick	GHOST	Physical	30	100	30	May paralyze opponent.	30
Life Dew	WATER	Status	—	—	10	User and teammates recover HP.	—
Light of Ruin	FAIRY	Special	140	90	5	User receives recoil damage.	—
Light Screen	PSYCHIC	Status	—	—	30	Halves damage from Special attacks for 5 turns.	—
Light That Burns the Sky	PSYCHIC	Special	200	—	1	Ultra Necrozma-exclusive Z-Move. Ignores target's ability; uses highest Attack stat.	—
Liquidation	WATER	Physical	85	100	10	May lower opponent's Defense.	20
Lock-On	NORMAL	Status	—	—	5	User's next attack is guaranteed to hit.	—
Lovely Kiss	NORMAL	Status	—	75	10	Puts opponent to sleep.	—
Low Kick	FIGHTING	Physical	—	100	20	The heavier the opponent, the stronger the attack.	—
Low Sweep	FIGHTING	Physical	65	100	20	Lowers opponent's Speed.	100
Lucky Chant	NORMAL	Status	—	—	30	Opponent cannot land critical hits for 5 turns.	—
Lumina Crash	PSYCHIC	Special	80	100	10	Harshly lowers target’s Special Defense.	—
Lunar Blessing	PSYCHIC	Status	—	—	5	Heals user's status conditions and recovers HP.	—
Lunar Dance	PSYCHIC	Status	—	—	10	The user faints but the next Pokémon released is fully healed.	—
Lunge	BUG	Physical	80	100	15	The user makes a lunge at the target, attacking with full force. This also lowers the target's Attack stat.	100
Luster Purge	PSYCHIC	Special	70	100	5	May lower opponent's Special Defense.	50
Mach Punch	FIGHTING	Physical	40	100	30	User attacks first.	—
Magic Coat	PSYCHIC	Status	—	—	15	Reflects moves that cause status conditions back to the attacker.	—
Magic Powder	PSYCHIC	Status	—	100	20	Changes target's type to Psychic.	—
Magic Room	PSYCHIC	Status	—	—	10	Suppresses the effects of held items for five turns.	—
Magical Leaf	GRASS	Special	60	∞	20	Ignores Accuracy and Evasiveness.	—
Magical Torque	FAIRY	Physical	100	100	10		—
Magma Storm	FIRE	Special	100	75	5	Traps opponent, damaging them for 4-5 turns.	100
Magnet Bomb	STEEL	Physical	60	∞	20	Ignores Accuracy and Evasiveness.	—
Magnet Rise	ELECTRIC	Status	—	—	10	User becomes immune to Ground-type moves for 5 turns.	—
Magnetic Flux	ELECTRIC	Status	—	—	20	Raises Defense and Sp. Defense of Plus/Minus Pokémon.	—
Magnitude	GROUND	Physical	—	100	30	Hits with random power.	—
Make It Rain	STEEL	Special	120	100	5	Lowers user's Special Attack. Money is earned after the battle.	—
Malicious Moonsault	DARK	Physical	180	—	1	Incineroar-exclusive Z-Move.	—
Mat Block	FIGHTING	Status	—	—	10	Protects teammates from damaging moves.	—
Max Airstream	FLYING	—	—	—	—	Flying type Dynamax move. Raises the team's Speed.	100
Max Darkness	DARK	—	—	—	—	Dark type Dynamax move. Lowers the target's Special Defense.	100
Max Flare	FIRE	—	—	—	—	Fire type Dynamax move. Summons harsh sunlight.	—
Max Flutterby	BUG	—	—	—	—	Bug type Dynamax move. Lowers the target's Special Attack.	100
Max Geyser	WATER	—	—	—	—	Water type Dynamax move. Summons heavy rain.	—
Max Guard	NORMAL	—	—	—	—	Status category Dynamax move. Protects the user.	—
Max Hailstorm	ICE	—	—	—	—	Ice type Dynamax move. Summons hail.	—
Max Knuckle	FIGHTING	—	—	—	—	Fighting type Dynamax move. Increases the team's Attack.	100
Max Lightning	ELECTRIC	—	—	—	—	Electric type Dynamax move. Summons Electric Terrain.	—
Max Mindstorm	PSYCHIC	—	—	—	—	Psychic type Dynamax move. Summons Psychic Terrain.	—
Max Ooze	POISON	—	—	—	—	Poison type Dynamax move. Increases the team's Special Attack.	100
Max Overgrowth	GRASS	—	—	—	—	Grass type Dynamax move. Summons Grassy Terrain.	—
Max Phantasm	GHOST	—	—	—	—	Ghost type Dynamax move. Lowers the target's Defense.	100
Max Quake	GROUND	—	—	—	—	Ground type Dynamax move. Increases the team's Special Defense.	100
Max Rockfall	ROCK	—	—	—	—	Rock type Dynamax move. Summons a sandstorm.	—
Max Starfall	FAIRY	—	—	—	—	Fairy type Dynamax move. Summons Misty Terrain.	—
Max Steelspike	STEEL	—	—	—	—	Steel type Dynamax move. Raises the team's Defense.	100
Max Strike	NORMAL	—	—	—	—	Normal type Dynamax move. Lowers the target's Speed.	100
Max Wyrmwind	DRAGON	—	—	—	—	Dragon type Dynamax move. Lowers the target's Attack.	100
Me First	NORMAL	Status	—	—	20	User copies the opponent's attack with 1.5× power.	—
Mean Look	NORMAL	Status	—	—	5	Opponent cannot flee or switch.	—
Meditate	PSYCHIC	Status	—	—	40	Raises user's Attack.	—
Mega Drain	GRASS	Special	40	100	15	User recovers half the HP inflicted on opponent.	—
Mega Kick	NORMAL	Physical	120	75	5		—
Mega Punch	NORMAL	Physical	80	85	20		—
Megahorn	BUG	Physical	120	85	10		—
Memento	DARK	Status	—	100	10	User faints, sharply lowers opponent's Attack and Special Attack.	—
Menacing Moonraze Maelstrom	GHOST	Special	200	—	1	Lunala-exclusive Z-Move.	—
Metal Burst	STEEL	Physical	—	100	10	Deals damage equal to 1.5x opponent's attack.	—
Metal Claw	STEEL	Physical	50	95	35	May raise user's Attack.	10
Metal Sound	STEEL	Status	—	85	40	Sharply lowers opponent's Special Defense.	—
Meteor Assault	FIGHTING	Physical	150	100	5	User must recharge next turn.	—
Meteor Beam	ROCK	Special	120	90	10	User gathers space power and boosts its Sp. Atk stat, then attacks the target on the next turn.	—
Meteor Mash	STEEL	Physical	90	90	10	May raise user's Attack.	20
Metronome	NORMAL	Status	—	—	10	User performs almost any move in the game at random.	—
Milk Drink	NORMAL	Status	—	—	5	User recovers half its max HP.	—
Mimic	NORMAL	Status	—	—	10	Copies the opponent's last move.	—
Mind Blown	FIRE	Special	150	100	5	User receives recoil damage.	—
Mind Reader	NORMAL	Status	—	—	5	User's next attack is guaranteed to hit.	—
Minimize	NORMAL	Status	—	—	10	Sharply raises user's Evasiveness.	—
Miracle Eye	PSYCHIC	Status	—	—	40	Resets opponent's Evasiveness, removes Dark's Psychic immunity.	—
Mirror Coat	PSYCHIC	Special	—	100	20	When hit by a Special Attack, user strikes back with 2x power.	—
Mirror Move	FLYING	Status	—	—	20	User performs the opponent's last move.	—
Mirror Shot	STEEL	Special	65	85	10	May lower opponent's Accuracy.	30
Mist	ICE	Status	—	—	30	User's stats cannot be changed for a period of time.	—
Mist Ball	PSYCHIC	Special	70	100	5	May lower opponent's Special Attack.	50
Misty Explosion	FAIRY	Special	100	100	5	Power increases on Misty Terrain.	—
Misty Terrain	FAIRY	Status	—	—	10	Protects the field from status conditions for 5 turns.	—
Moonblast	FAIRY	Special	95	100	15	May lower opponent's Special Attack.	30
Moongeist Beam	GHOST	Special	100	100	5	Ignores the target's ability.	—
Moonlight	FAIRY	Status	—	—	5	User recovers HP. Amount varies with the weather.	—
Morning Sun	NORMAL	Status	—	—	5	User recovers HP. Amount varies with the weather.	—
Mortal Spin	POISON	Physical	30	100	15	Removes entry hazards and trap move effects, and poisons opposing Pokémon.	—
Mountain Gale	ICE	Physical	100	85	10	Lowers user's Speed.	—
Mud Bomb	GROUND	Special	65	85	10	May lower opponent's Accuracy.	30
Mud Shot	GROUND	Special	55	95	15	Lowers opponent's Speed.	100
Mud Sport	GROUND	Status	—	—	15	Weakens the power of Electric-type moves.	—
Mud-Slap	GROUND	Special	20	100	10	Lowers opponent's Accuracy.	100
Muddy Water	WATER	Special	90	85	10	May lower opponent's Accuracy.	30
Multi-Attack	NORMAL	Physical	120	100	10	Type matches Memory item held.	—
Mystical Fire	FIRE	Special	75	100	10	Lowers opponent's Special Attack.	100
Mystical Power	PSYCHIC	Special	70	90	10	Raises user's Attack or Defense.	—
Nasty Plot	DARK	Status	—	—	20	Sharply raises user's Special Attack.	—
Natural Gift	NORMAL	Physical	—	100	15	Power and type depend on the user's held berry.	—
Nature Power	NORMAL	Status	—	—	20	Uses a certain move based on the current terrain.	—
Nature's Madness	FAIRY	Special	—	90	10	Halves the foe's HP.	—
Needle Arm	GRASS	Physical	60	100	15	May cause flinching.	30
Never-Ending Nightmare	GHOST	—	—	—	1	Ghost type Z-Move.	—
Night Daze	DARK	Special	85	95	10	May lower opponent's Accuracy.	40
Night Shade	GHOST	Special	—	100	15	Inflicts damage equal to user's level.	—
Night Slash	DARK	Physical	70	100	15	High critical hit ratio.	—
Nightmare	GHOST	Status	—	100	15	The sleeping opponent loses 25% of its max HP each turn.	—
No Retreat	FIGHTING	Status	—	—	5	Raises all stats but user cannot switch out.	100
Noble Roar	NORMAL	Status	—	100	30	Lowers opponent's Attack and Special Attack.	100
Noxious Torque	POISON	Physical	100	100	10		—
Nuzzle	ELECTRIC	Physical	20	100	20	Paralyzes opponent.	100
Oblivion Wing	FLYING	Special	80	100	10	User recovers most of the HP inflicted on opponent.	—
Obstruct	DARK	Status	—	100	10	Protects the user and sharply lowers Defence on contact.	—
Oceanic Operetta	WATER	Special	195	—	1	Primarina-exclusive Z-Move.	—
Octazooka	WATER	Special	65	85	10	May lower opponent's Accuracy.	50
Octolock	FIGHTING	Status	—	100	15	Lowers opponent's Defense and Special Defense every turn, and they cannot flee or switch out.	—
Odor Sleuth	NORMAL	Status	—	—	40	Resets opponent's Evasiveness, and allows Normal- and Fighting-type attacks to hit Ghosts.	—
Ominous Wind	GHOST	Special	60	100	5	May raise all user's stats at once.	10
Order Up	DRAGON	Physical	80	100	10		—
Origin Pulse	WATER	Special	110	85	10	Hits all adjacent opponents.	—
Outrage	DRAGON	Physical	120	100	10	User attacks for 2-3 turns but then becomes confused.	—
Overdrive	ELECTRIC	Special	80	100	10	Hits all adjacent opponents.	—
Overheat	FIRE	Special	130	90	5	Sharply lowers user's Special Attack.	100
Pain Split	NORMAL	Status	—	—	20	The user's and opponent's HP becomes the average of both.	—
Parabolic Charge	ELECTRIC	Special	65	100	20	User recovers half the HP inflicted on opponent.	—
Parting Shot	DARK	Status	—	100	20	Lowers opponent's Attack and Special Attack then switches out.	100
Pay Day	NORMAL	Physical	40	100	20	Money is earned after the battle.	—
Payback	DARK	Physical	50	100	10	Power doubles if the user was attacked first.	—
Peck	FLYING	Physical	35	100	35		—
Perish Song	NORMAL	Status	—	—	5	Any Pokémon in play when this attack is used faints in 3 turns.	—
Petal Blizzard	GRASS	Physical	90	100	15	Hits all adjacent Pokémon.	—
Petal Dance	GRASS	Special	120	100	10	User attacks for 2-3 turns but then becomes confused.	—
Phantom Force	GHOST	Physical	90	100	10	Disappears on first turn, attacks on second. Can strike through Protect/Detect.	—
Photon Geyser	PSYCHIC	Special	100	100	5	Uses Attack or Special Attack stat, whichever is higher.	—
Pika Papow	ELECTRIC	Special	—	∞	20	Power increases when player's bond is stronger.	—
Pin Missile	BUG	Physical	25	95	20	Hits 2-5 times in one turn.	—
Plasma Fists	ELECTRIC	Physical	100	100	15	Changes Normal-type moves to Electric-type moves.	—
Play Nice	NORMAL	Status	—	—	20	Lowers opponent's Attack. Always hits.	100
Play Rough	FAIRY	Physical	90	90	10	May lower opponent's Attack.	10
Pluck	FLYING	Physical	60	100	20	If the opponent is holding a berry, its effect is stolen by user.	—
Poison Fang	POISON	Physical	50	100	15	May badly poison opponent.	50
Poison Gas	POISON	Status	—	90	40	Poisons opponent.	—
Poison Jab	POISON	Physical	80	100	20	May poison the opponent.	30
Poison Powder	POISON	Status	—	75	35	Poisons opponent.	—
Poison Sting	POISON	Physical	15	100	35	May poison the opponent.	30
Poison Tail	POISON	Physical	50	100	25	High critical hit ratio. May poison opponent.	10
Pollen Puff	BUG	Special	90	100	15	Deals damage to opponent or restores HP of teammate.	—
Poltergeist	GHOST	Physical	110	90	5	Fails if the target doesn’t have an item.	—
Population Bomb	NORMAL	Physical	20	90	10	Hits 1-10 times in a row.	—
Pounce	BUG	Physical	50	100	20	Lowers opponent's Speed.	100
Pound	NORMAL	Physical	40	100	35		—
Powder	BUG	Status	—	100	20	Damages Pokémon using Fire type moves.	—
Powder Snow	ICE	Special	40	100	25	May freeze opponent.	10
Power Gem	ROCK	Special	80	100	20		—
Power Shift	NORMAL	Status	—	—	10	Switches offensive and defensive stats.	—
Power Split	PSYCHIC	Status	—	—	10	Averages Attack and Special Attack with the target.	—
Power Swap	PSYCHIC	Status	—	—	10	User and opponent swap Attack and Special Attack.	—
Power Trick	PSYCHIC	Status	—	—	10	User's own Attack and Defense switch.	—
Power Trip	DARK	Physical	20	100	10	The user boasts its strength and attacks the target. The more the user's stats are raised, the greater the move's power.	—
Power Whip	GRASS	Physical	120	85	10		—
Power-Up Punch	FIGHTING	Physical	40	100	20	Raises Attack.	100
Precipice Blades	GROUND	Physical	120	85	10	Hits all adjacent opponents.	—
Present	NORMAL	Physical	—	90	15	Either deals damage or heals.	—
Prismatic Laser	PSYCHIC	Special	160	100	10	The user shoots powerful lasers using the power of a prism. The user can't move on the next turn.	—
Protect	NORMAL	Status	—	—	10	Protects the user, but may fail if used consecutively.	—
Psybeam	PSYCHIC	Special	65	100	20	May confuse opponent.	10
Psyblade	PSYCHIC	Physical	80	100	15	Power increases on Electric Terrain.	—
Psych Up	NORMAL	Status	—	—	10	Copies the opponent's stat changes.	—
Psychic	PSYCHIC	Special	90	100	10	May lower opponent's Special Defense.	10
Psychic Fangs	PSYCHIC	Physical	85	100	10	The user bites the target with its psychic capabilities. This can also destroy Light Screen and Reflect.	—
Psychic Terrain	PSYCHIC	Status	—	—	10	Prevents priority moves from being used for 5 turns.	—
Psycho Boost	PSYCHIC	Special	140	90	5	Sharply lowers user's Special Attack.	100
Psycho Cut	PSYCHIC	Physical	70	100	20	High critical hit ratio.	—
Psycho Shift	PSYCHIC	Status	—	100	10	Transfers user's status condition to the opponent.	—
Psyshield Bash	PSYCHIC	Physical	70	90	10	Raises user's Defense and Special Defense.	—
Psyshock	PSYCHIC	Special	80	100	10	Inflicts damage based on the target's Defense, not Special Defense.	—
Psystrike	PSYCHIC	Special	100	100	10	Inflicts damage based on the target's Defense, not Special Defense.	—
Psywave	PSYCHIC	Special	—	100	15	Inflicts damage 50-150% of user's level.	—
Pulverizing Pancake	NORMAL	Physical	210	—	1	Snorlax-exclusive Normal type Z-Move.	—
Punishment	DARK	Physical	—	100	5	Power increases when opponent's stats have been raised.	—
Purify	POISON	Status	—	—	20	The user heals the target's status condition. If the move succeeds, it also restores the user's own HP.	—
Pursuit	DARK	Physical	40	100	20	Double power if the opponent is switching out.	—
Pyro Ball	FIRE	Physical	120	90	5	May burn opponent.	10
Quash	DARK	Status	—	100	15	Makes the target act last this turn.	—
Quick Attack	NORMAL	Physical	40	100	30	User attacks first.	—
Quick Guard	FIGHTING	Status	—	—	15	Protects the user's team from high-priority moves.	—
Quiver Dance	BUG	Status	—	—	20	Raises user's Special Attack, Special Defense and Speed.	—
Rage	NORMAL	Physical	20	100	20	Raises user's Attack when hit.	—
Rage Fist	GHOST	Physical	50	100	10	The more times the user has been hit by attacks, the greater the move's power.	—
Rage Powder	BUG	Status	—	—	20	Forces attacks to hit user, not team-mates.	—
Raging Bull	NORMAL	Physical	90	100	10	Type depends on the user’s form. Breaks through Reflect and Light Screen barriers.	—
Raging Fury	FIRE	Physical	120	100	10	User keeps repeating the same move over and over.	—
Rain Dance	WATER	Status	—	—	5	Makes it rain for 5 turns.	—
Rapid Spin	NORMAL	Physical	50	100	40	Raises user's Speed and removes entry hazards and trap move effects.	100
Razor Leaf	GRASS	Physical	55	95	25	High critical hit ratio.	—
Razor Shell	WATER	Physical	75	95	10	May lower opponent's Defense.	50
Razor Wind	NORMAL	Special	80	100	10	Charges on first turn, attacks on second. High critical hit ratio.	—
Recover	NORMAL	Status	—	—	5	User recovers half its max HP.	—
Recycle	NORMAL	Status	—	—	10	User's used hold item is restored.	—
Reflect	PSYCHIC	Status	—	—	20	Halves damage from Physical attacks for 5 turns.	—
Reflect Type	NORMAL	Status	—	—	15	User becomes the target's type.	—
Refresh	NORMAL	Status	—	—	20	Cures paralysis, poison, and burns.	—
Relic Song	NORMAL	Special	75	100	10	May put the target to sleep.	10
Rest	PSYCHIC	Status	—	—	5	User sleeps for 2 turns, but user is fully healed.	—
Retaliate	NORMAL	Physical	70	100	5	Inflicts double damage if a teammate fainted on the last turn.	—
Return	NORMAL	Physical	—	100	20	Power increases with higher Friendship.	—
Revelation Dance	NORMAL	Special	90	100	15	Type changes based on Oricorio's form.	—
Revenge	FIGHTING	Physical	60	100	10	Power increases if user was hit first.	—
Reversal	FIGHTING	Physical	—	100	15	The lower the user's HP, the higher the power.	—
Revival Blessing	NORMAL	Status	—	—	1	Revives a fainted party Pokémon to half HP.	—
Rising Voltage	ELECTRIC	Special	70	100	20	Power doubles on Electric Terrain.	—
Roar	NORMAL	Status	—	—	20	In battles, the opponent switches. In the wild, the Pokémon runs.	—
Roar of Time	DRAGON	Special	150	90	5	User must recharge next turn.	—
Rock Blast	ROCK	Physical	25	90	10	Hits 2-5 times in one turn.	—
Rock Climb	NORMAL	Physical	90	85	20	May confuse opponent.	20
Rock Polish	ROCK	Status	—	—	20	Sharply raises user's Speed.	—
Rock Slide	ROCK	Physical	75	90	10	May cause flinching.	30
Rock Smash	FIGHTING	Physical	40	100	15	May lower opponent's Defense.	50
Rock Throw	ROCK	Physical	50	90	15		—
Rock Tomb	ROCK	Physical	60	95	15	Lowers opponent's Speed.	100
Rock Wrecker	ROCK	Physical	150	90	5	User must recharge next turn.	—
Role Play	PSYCHIC	Status	—	—	10	User copies the opponent's Ability.	—
Rolling Kick	FIGHTING	Physical	60	85	15	May cause flinching.	30
Rollout	ROCK	Physical	30	90	20	Doubles in power each turn for 5 turns.	—
Roost	FLYING	Status	—	—	5	User recovers half of its max HP and loses the Flying type temporarily.	—
Rototiller	GROUND	Status	—	—	10	Raises Attack and Special Attack of Grass-types.	100
Round	NORMAL	Special	60	100	15	Power increases if teammates use it in the same turn.	—
Ruination	DARK	Special	1	90	10	Halves the opponent's HP.	—
Sacred Fire	FIRE	Physical	100	95	5	May burn opponent.	50
Sacred Sword	FIGHTING	Physical	90	100	15	Ignores opponent's stat changes.	—
Safeguard	NORMAL	Status	—	—	25	The user's party is protected from status conditions.	—
Salt Cure	ROCK	Physical	40	100	15	Deals damage each turn; Steel and Water types are more affected.	—
Sand Attack	GROUND	Status	—	100	15	Lowers opponent's Accuracy.	—
Sand Tomb	GROUND	Physical	35	85	15	Traps opponent, damaging them for 4-5 turns.	100
Sandsear Storm	GROUND	Special	100	80	10	May burn target.	—
Sandstorm	ROCK	Status	—	—	10	Creates a sandstorm for 5 turns.	—
Sappy Seed	GRASS	Physical	90	100	15	Drains HP from opponent each turn.	100
Savage Spin-Out	BUG	—	—	—	1	Bug type Z-Move.	—
Scald	WATER	Special	80	100	15	May burn opponent.	30
Scale Shot	DRAGON	Physical	25	90	20	Hits 2-5 times in one turn. Boosts user's Speed but lowers its Defense.	—
Scary Face	NORMAL	Status	—	100	10	Sharply lowers opponent's Speed.	—
Scorching Sands	GROUND	Special	70	100	10	May burn the target.	—
Scratch	NORMAL	Physical	40	100	35		—
Screech	NORMAL	Status	—	85	40	Sharply lowers opponent's Defense.	—
Searing Shot	FIRE	Special	100	100	5	May burn opponent.	30
Searing Sunraze Smash	STEEL	Physical	200	—	1	Solgaleo-exclusive Z-Move.	—
Secret Power	NORMAL	Physical	70	100	20	Effects of the attack vary with the location.	30
Secret Sword	FIGHTING	Special	85	100	10	Inflicts damage based on the target's Defense, not Special Defense.	—
Seed Bomb	GRASS	Physical	80	100	15		—
Seed Flare	GRASS	Special	120	85	5	May lower opponent's Special Defense.	40
Seismic Toss	FIGHTING	Physical	—	100	20	Inflicts damage equal to user's level.	—
Self-Destruct	NORMAL	Physical	200	100	5	User faints.	—
Shadow Ball	GHOST	Special	80	100	15	May lower opponent's Special Defense.	20
Shadow Bone	GHOST	Physical	85	100	10	May lower opponent's Defense.	20
Shadow Claw	GHOST	Physical	70	100	15	High critical hit ratio.	—
Shadow Force	GHOST	Physical	120	100	5	Disappears on first turn, attacks on second. Can strike through Protect/Detect.	—
Shadow Punch	GHOST	Physical	60	∞	20	Ignores Accuracy and Evasiveness.	—
Shadow Sneak	GHOST	Physical	40	100	30	User attacks first.	—
Sharpen	NORMAL	Status	—	—	30	Raises user's Attack.	—
Shattered Psyche	PSYCHIC	—	—	—	1	Psychic type Z-Move.	—
Shed Tail	NORMAL	Status	—	—	10	Creates a substitute, then swaps places with a party Pokémon in waiting.	—
Sheer Cold	ICE	Special	—	30	5	One-Hit-KO, if it hits.	—
Shell Side Arm	POISON	Special	90	100	10	May poison opponent. Inflicts either Special or Physical damage, whichever is better.	—
Shell Smash	NORMAL	Status	—	—	15	Sharply raises user's Attack, Special Attack and Speed but lowers Defense and Special Defense.	—
Shell Trap	FIRE	Special	150	100	5	Deals more damage to opponent if hit by a Physical move.	—
Shelter	STEEL	Status	—	—	10	Raises Defense and Evasion.	—
Shift Gear	STEEL	Status	—	—	10	Raises user's Attack and sharply raises Speed.	—
Shock Wave	ELECTRIC	Special	60	∞	20	Ignores Accuracy and Evasiveness.	—
Shore Up	GROUND	Status	—	—	5	The user regains up to half of its max HP. It restores more HP in a sandstorm.	—
Signal Beam	BUG	Special	75	100	15	May confuse opponent.	10
Silk Trap	BUG	Status	—	—	10	Protects the user and lowers opponent's Speed on contact.	—
Silver Wind	BUG	Special	60	100	5	May raise all stats of user at once.	10
Simple Beam	NORMAL	Status	—	100	15	Changes target's ability to Simple.	—
Sing	NORMAL	Status	—	55	15	Puts opponent to sleep.	—
Sinister Arrow Raid	GHOST	Physical	180	—	1	Decidueye-exclusive Z-Move.	—
Sizzly Slide	FIRE	Physical	90	100	15	Burns the opponent.	100
Sketch	NORMAL	Status	—	—	1	Permanently copies the opponent's last move.	—
Skill Swap	PSYCHIC	Status	—	—	10	The user swaps Abilities with the opponent.	—
Skitter Smack	BUG	Physical	70	90	10	Lowers opponent's Sp. Attack.	—
Skull Bash	NORMAL	Physical	130	100	10	Raises Defense on first turn, attacks on second.	100
Sky Attack	FLYING	Physical	140	90	5	Charges on first turn, attacks on second. May cause flinching. High critical hit ratio.	30
Sky Drop	FLYING	Physical	60	100	10	Takes opponent into the air on first turn, drops them on second turn.	—
Sky Uppercut	FIGHTING	Physical	85	90	15	Hits the opponent, even during Fly.	—
Slack Off	NORMAL	Status	—	—	5	User recovers half its max HP.	—
Slam	NORMAL	Physical	80	75	20		—
Slash	NORMAL	Physical	70	100	20	High critical hit ratio.	—
Sleep Powder	GRASS	Status	—	75	15	Puts opponent to sleep.	—
Sleep Talk	NORMAL	Status	—	—	10	User performs one of its own moves while sleeping.	—
Sludge	POISON	Special	65	100	20	May poison opponent.	30
Sludge Bomb	POISON	Special	90	100	10	May poison opponent.	30
Sludge Wave	POISON	Special	95	100	10	May poison opponent.	10
Smack Down	ROCK	Physical	50	100	15	Makes Flying-type Pokémon vulnerable to Ground moves.	100
Smart Strike	STEEL	Physical	70	—	10	The user stabs the target with a sharp horn. This attack never misses.	—
Smelling Salts	NORMAL	Physical	70	100	10	Power doubles if opponent is paralyzed, but cures it.	—
Smog	POISON	Special	30	70	20	May poison opponent.	40
Smokescreen	NORMAL	Status	—	100	20	Lowers opponent's Accuracy.	—
Snap Trap	GRASS	Physical	35	100	15	Traps opponent, damaging them for 4-5 turns.	100
Snarl	DARK	Special	55	95	15	Lowers opponent's Special Attack.	100
Snatch	DARK	Status	—	—	10	Steals the effects of the opponent's next move.	—
Snipe Shot	WATER	Special	80	100	15	Ignores moves and abilities that draw in moves. High critical hit ratio.	—
Snore	NORMAL	Special	50	100	15	Can only be used if asleep. May cause flinching.	30
Snowscape	ICE	Status	—	—	10	Raises Defense of Ice types for 5 turns.	—
Soak	WATER	Status	—	100	20	Changes the target's type to water.	—
Soft-Boiled	NORMAL	Status	—	—	5	User recovers half its max HP.	—
Solar Beam	GRASS	Special	120	100	10	Charges on first turn, attacks on second.	—
Solar Blade	GRASS	Physical	125	100	10	Charges on first turn, attacks on second.	—
Sonic Boom	NORMAL	Special	—	90	20	Always inflicts 20 HP.	—
Soul-Stealing 7-Star Strike	GHOST	Physical	195	—	1	Marshadow-exclusive Z-Move.	—
Spacial Rend	DRAGON	Special	100	95	5	High critical hit ratio.	—
Spark	ELECTRIC	Physical	65	100	20	May paralyze opponent.	30
Sparkling Aria	WATER	Special	90	100	10	Heals the burns of its target.	—
Sparkly Swirl	FAIRY	Special	90	100	15	Cures all status problems in the party Pokémon.	—
Spectral Thief	GHOST	Physical	90	100	10	The user hides in the target's shadow, steals the target's stat boosts, and then attacks.	—
Speed Swap	PSYCHIC	Status	—	—	10	The user exchanges Speed stats with the target.	—
Spicy Extract	GRASS	Status	—	∞	15	Harshly lowers the opponent's Defense and sharply raises their Attack.	—
Spider Web	BUG	Status	—	—	10	Opponent cannot escape/switch.	—
Spike Cannon	NORMAL	Physical	20	100	15	Hits 2-5 times in one turn.	—
Spikes	GROUND	Status	—	—	20	Hurts opponents when they switch into battle.	—
Spiky Shield	GRASS	Status	—	—	10	Protects the user and inflicts damage on contact.	—
Spin Out	STEEL	Physical	100	100	5	Harshly lowers user’s Speed.	—
Spirit Break	FAIRY	Physical	75	100	15	Lowers opponent's Special Attack.	100
Spirit Shackle	GHOST	Physical	80	100	10	Prevents the opponent from switching out.	—
Spit Up	NORMAL	Special	—	100	10	Power depends on how many times the user performed Stockpile.	—
Spite	GHOST	Status	—	100	10	The opponent's last move loses 2-5 PP.	—
Splash	NORMAL	Status	—	—	40	Doesn't do ANYTHING.	—
Splintered Stormshards	ROCK	Physical	190	—	1	Lycanroc-exclusive Z-Move.	—
Splishy Splash	WATER	Special	90	100	15	May paralyze opponent.	30
Spore	GRASS	Status	—	100	15	Puts opponent to sleep.	—
Spotlight	NORMAL	Status	—	—	15	The user shines a spotlight on the target so that only the target will be attacked during the turn.	—
Springtide Storm	FAIRY	Special	100	80	5	Boosts user's stats in Incarnate Forme, or lowers opponent's stats in Therian Forme.	—
Stealth Rock	ROCK	Status	—	—	20	Damages opponent switching into battle.	—
Steam Eruption	WATER	Special	110	95	5	May burn opponent.	30
Steamroller	BUG	Physical	65	100	20	May cause flinching.	30
Steel Beam	STEEL	Special	140	95	5	User loses 50% of its HP.	—
Steel Roller	STEEL	Physical	130	100	5	Fails if no Terrain in effect.	—
Steel Wing	STEEL	Physical	70	90	25	May raise user's Defense.	10
Sticky Web	BUG	Status	—	—	20	Lowers opponent's Speed when switching into battle.	—
Stockpile	NORMAL	Status	—	—	20	Stores energy for use with Spit Up and Swallow.	—
Stoked Sparksurfer	ELECTRIC	Special	175	—	1	Alolan Raichu-exclusive Electric type Z-Move.	100
Stomp	NORMAL	Physical	65	100	20	May cause flinching.	30
Stomping Tantrum	GROUND	Physical	75	100	10	Driven by frustration, the user attacks the target. If the user's previous move has failed, the power of this move doubles.	—
Stone Axe	ROCK	Physical	65	90	15	High critical hit ratio. Damages target with splinters each turn.	—
Stone Edge	ROCK	Physical	100	80	5	High critical hit ratio.	—
Stored Power	PSYCHIC	Special	20	100	10	Power increases when user's stats have been raised.	—
Storm Throw	FIGHTING	Physical	60	100	10	Always results in a critical hit.	100
Strange Steam	FAIRY	Special	90	95	10	May confuse opponent.	20
Strength	NORMAL	Physical	80	100	15		—
Strength Sap	GRASS	Status	—	100	10	The user restores its HP by the same amount as the target's Attack stat. It also lowers the target's Attack stat.	100
String Shot	BUG	Status	—	95	40	Sharply lowers opponent's Speed.	—
Struggle	NORMAL	Physical	50	—	—	Only usable when all PP are gone. Hurts the user.	—
Struggle Bug	BUG	Special	50	100	20	Lowers opponent's Special Attack.	100
Stuff Cheeks	NORMAL	Status	—	—	10	The user eats its held Berry, then sharply raises its Defense stat.	100
Stun Spore	GRASS	Status	—	75	30	Paralyzes opponent.	—
Submission	FIGHTING	Physical	80	80	20	User receives recoil damage.	—
Substitute	NORMAL	Status	—	—	10	Uses HP to creates a decoy that takes hits.	—
Subzero Slammer	ICE	—	—	—	1	Ice type Z-Move.	—
Sucker Punch	DARK	Physical	70	100	5	User attacks first, but only works if opponent is readying an attack.	—
Sunny Day	FIRE	Status	—	—	5	Makes it sunny for 5 turns.	—
Sunsteel Strike	STEEL	Physical	100	100	5	Ignores the target's ability.	—
Super Fang	NORMAL	Physical	—	90	10	Always takes off half of the opponent's HP.	—
Superpower	FIGHTING	Physical	120	100	5	Lowers user's Attack and Defense.	100
Supersonic	NORMAL	Status	—	55	20	Confuses opponent.	—
Supersonic Skystrike	FLYING	—	—	—	1	Flying type Z-Move.	—
Surf	WATER	Special	90	100	15	Hits all adjacent Pokémon.	—
Surging Strikes	WATER	Physical	25	100	5	Always results in a critical hit and ignores stat changes.	—
Swagger	NORMAL	Status	—	85	15	Confuses opponent, but sharply raises its Attack.	—
Swallow	NORMAL	Status	—	—	10	The more times the user has performed Stockpile, the more HP is recovered.	—
Sweet Kiss	FAIRY	Status	—	75	10	Confuses opponent.	—
Sweet Scent	NORMAL	Status	—	100	20	Lowers opponent's Evasiveness.	—
Swift	NORMAL	Special	60	∞	20	Ignores Accuracy and Evasiveness.	—
Switcheroo	DARK	Status	—	100	10	Swaps held items with the opponent.	—
Swords Dance	NORMAL	Status	—	—	20	Sharply raises user's Attack.	—
Synchronoise	PSYCHIC	Special	120	100	10	Hits any Pokémon that shares a type with the user.	—
Synthesis	GRASS	Status	—	—	5	User recovers HP. Amount varies with the weather.	—
Tackle	NORMAL	Physical	40	100	35		—
Tail Glow	BUG	Status	—	—	20	Drastically raises user's Special Attack.	—
Tail Slap	NORMAL	Physical	25	85	10	Hits 2-5 times in one turn.	—
Tail Whip	NORMAL	Status	—	100	30	Lowers opponent's Defense.	—
Tailwind	FLYING	Status	—	—	15	Doubles Speed for 4 turns.	—
Take Down	NORMAL	Physical	90	85	20	User receives recoil damage.	—
Take Heart	PSYCHIC	Status	—	—	10	Heals user's status conditions and raises its stats.	—
Tar Shot	ROCK	Status	—	100	15	Lowers the opponent's Speed and makes them weaker to Fire-type moves.	100
Taunt	DARK	Status	—	100	20	Opponent can only use moves that attack.	—
Tearful Look	NORMAL	Status	—	—	20	The user gets teary eyed to make the target lose its combative spirit. This lowers the target's Attack and Sp. Atk stats.	100
Teatime	NORMAL	Status	—	—	10	Forces all Pokémon on the field to eat their berries.	—
Techno Blast	NORMAL	Special	120	100	5	Type depends on the Drive being held.	—
Tectonic Rage	GROUND	—	—	—	1	Ground type Z-Move.	—
Teeter Dance	NORMAL	Status	—	100	20	Confuses all Pokémon.	—
Telekinesis	PSYCHIC	Status	—	—	15	Ignores opponent's Evasiveness for three turns, add Ground immunity.	—
Teleport	PSYCHIC	Status	—	—	20	Allows user to flee wild battles; also warps player to last PokéCenter.	—
Tera Blast	NORMAL	Special	80	100	10	Changes type when the user has Terastallized.	—
Terrain Pulse	NORMAL	Special	50	100	10	Type and power change depending on the Terrain in effect.	—
Thief	DARK	Physical	60	100	25	Also steals opponent's held item.	—
Thousand Arrows	GROUND	Physical	90	100	10	Makes Flying-type Pokémon vulnerable to Ground moves.	—
Thousand Waves	GROUND	Physical	90	100	10	Opponent cannot flee or switch.	—
Thrash	NORMAL	Physical	120	100	10	User attacks for 2-3 turns but then becomes confused.	—
Throat Chop	DARK	Physical	80	100	15	Prevents use of sound moves for two turns.	100
Thunder	ELECTRIC	Special	110	70	10	May paralyze opponent.	30
Thunder Cage	ELECTRIC	Special	80	90	15	Deals damage and traps opponent, damaging them for 4-5 turns.	—
Thunder Fang	ELECTRIC	Physical	65	95	15	May cause flinching and/or paralyze opponent.	10
Thunder Punch	ELECTRIC	Physical	75	100	15	May paralyze opponent.	10
Thunder Shock	ELECTRIC	Special	40	100	30	May paralyze opponent.	10
Thunder Wave	ELECTRIC	Status	—	90	20	Paralyzes opponent.	—
Thunderbolt	ELECTRIC	Special	90	100	15	May paralyze opponent.	10
Thunderous Kick	FIGHTING	Physical	90	100	10	Lowers opponent's Defense.	100
Tickle	NORMAL	Status	—	100	20	Lowers opponent's Attack and Defense.	—
Tidy Up	NORMAL	Status	—	—	10	Removes the effects of entry hazards and Substitute, and boosts user’s Attack and Speed.	—
Topsy-Turvy	DARK	Status	—	—	20	Reverses stat changes of opponent.	—
Torch Song	FIRE	Special	80	100	10	Raises user's Special Attack.	—
Torment	DARK	Status	—	100	15	Opponent cannot use the same move in a row.	—
Toxic	POISON	Status	—	90	10	Badly poisons opponent.	—
Toxic Spikes	POISON	Status	—	—	20	Poisons opponents when they switch into battle.	—
Toxic Thread	POISON	Status	—	100	20	Poisons opponent and lowers its Speed.	100
Trailblaze	GRASS	Physical	50	100	20	Raises user's Speed.	—
Transform	NORMAL	Status	—	—	10	User takes on the form and attacks of the opponent.	—
Tri Attack	NORMAL	Special	80	100	10	May paralyze, burn or freeze opponent.	20
Trick	PSYCHIC	Status	—	100	10	Swaps held items with the opponent.	—
Trick Room	PSYCHIC	Status	—	—	5	Slower Pokémon move first in the turn for 5 turns.	—
Trick-or-Treat	GHOST	Status	—	100	20	Adds Ghost type to opponent.	—
Triple Arrows	FIGHTING	Physical	90	100	10	Raises critical hit ratio and lowers target's Defense.	—
Triple Axel	ICE	Physical	20	90	10	Attacks thrice with more power each time.	—
Triple Dive	WATER	Physical	30	95	10	Hits 3 times in a row.	—
Triple Kick	FIGHTING	Physical	10	90	10	Hits thrice in one turn at increasing power.	—
Trop Kick	GRASS	Physical	70	100	15	Lowers opponent's Attack.	100
Trump Card	NORMAL	Special	—	∞	5	The lower the PP, the higher the power.	—
Twin Beam	PSYCHIC	Special	40	100	10	Hits twice in one turn.	—
Twineedle	BUG	Physical	25	100	20	Hits twice in one turn. May poison opponent.	20
Twinkle Tackle	FAIRY	—	—	—	1	Fairy type Z-Move.	—
Twister	DRAGON	Special	40	100	20	May cause flinching. Hits Pokémon using Fly/Bounce with double power.	20
U-turn	BUG	Physical	70	100	20	User switches out immediately after attacking.	—
Uproar	NORMAL	Special	90	100	10	User attacks for 3 turns and prevents sleep.	—
V-create	FIRE	Physical	180	95	5	Lowers user's Defense, Special Defense and Speed.	100
Vacuum Wave	FIGHTING	Special	40	100	30	User attacks first.	—
Veevee Volley	NORMAL	Physical	—	∞	20	Power increases when player's bond is stronger.	—
Venom Drench	POISON	Status	—	100	20	Lowers poisoned opponent's Special Attack and Speed.	100
Venoshock	POISON	Special	65	100	10	Inflicts double damage if the target is poisoned.	—
Victory Dance	FIGHTING	Status	—	—	10	Raises Attack and Defense.	—
Vine Whip	GRASS	Physical	45	100	25		—
Vise Grip	NORMAL	Physical	55	100	30		—
Vital Throw	FIGHTING	Physical	70	∞	10	User attacks last, but ignores Accuracy and Evasiveness.	—
Volt Switch	ELECTRIC	Special	70	100	20	User must switch out after attacking.	—
Volt Tackle	ELECTRIC	Physical	120	100	15	User receives recoil damage. May paralyze opponent.	10
Wake-Up Slap	FIGHTING	Physical	70	100	10	Power doubles if opponent is asleep, but wakes it up.	—
Water Gun	WATER	Special	40	100	25		—
Water Pledge	WATER	Special	80	100	10	Added effects appear if preceded by Fire Pledge or succeeded by Grass Pledge.	—
Water Pulse	WATER	Special	60	100	20	May confuse opponent.	20
Water Shuriken	WATER	Special	15	100	20	Hits 2-5 times in one turn.	—
Water Sport	WATER	Status	—	—	15	Weakens the power of Fire-type moves.	—
Water Spout	WATER	Special	150	100	5	The higher the user's HP, the higher the damage caused.	—
Waterfall	WATER	Physical	80	100	15	May cause flinching.	20
Wave Crash	WATER	Physical	120	100	10	User receives recoil damage.	—
Weather Ball	NORMAL	Special	50	100	10	Move's power and type changes with the weather.	—
Whirlpool	WATER	Special	35	85	15	Traps opponent, damaging them for 4-5 turns.	100
Whirlwind	NORMAL	Status	—	—	20	In battles, the opponent switches. In the wild, the Pokémon runs.	—
Wicked Blow	DARK	Physical	75	100	5	Always results in a critical hit and ignores stat changes.	—
Wicked Torque	DARK	Physical	80	100	10		—
Wide Guard	ROCK	Status	—	—	10	Protects the user's team from multi-target attacks.	—
Wild Charge	ELECTRIC	Physical	90	100	15	User receives recoil damage.	—
Wildbolt Storm	ELECTRIC	Special	100	80	10	May paralyze target.	—
Will-O-Wisp	FIRE	Status	—	85	15	Burns opponent.	—
Wing Attack	FLYING	Physical	60	100	35		—
Wish	NORMAL	Status	—	—	10	The user recovers HP in the following turn.	—
Withdraw	WATER	Status	—	—	40	Raises user's Defense.	—
Wonder Room	PSYCHIC	Status	—	—	10	Swaps every Pokémon's Defense and Special Defense for 5 turns.	—
Wood Hammer	GRASS	Physical	120	100	15	User receives recoil damage.	—
Work Up	NORMAL	Status	—	—	30	Raises user's Attack and Special Attack.	—
Worry Seed	GRASS	Status	—	100	10	Changes the opponent's Ability to Insomnia.	—
Wrap	NORMAL	Physical	15	90	20	Traps opponent, damaging them for 4-5 turns.	100
Wring Out	NORMAL	Special	—	100	5	The higher the opponent's HP, the higher the damage.	—
X-Scissor	BUG	Physical	80	100	15		—
Yawn	NORMAL	Status	—	—	10	Puts opponent to sleep in the next turn.	—
Zap Cannon	ELECTRIC	Special	120	50	5	Paralyzes opponent.	100
Zen Headbutt	PSYCHIC	Physical	80	90	15	May cause flinching.	20
Zing Zap	ELECTRIC	Physical	80	100	10	May cause flinching.	30
Zippy Zap	ELECTRIC	Physical	50	100	15	Always results in a critical hit.	100

`

	rows := strings.Split(data, "\n")
	csvData := make([][]string, len(rows))

	for i, row := range rows {
		csvData[i] = strings.Fields(row)
		if i > 9 { // cut header\
			j := i - 9
			jj := strconv.Itoa(j)
			// csvData[i] = strings.Fields(row)
			csvData[i] = append([]string{jj}, strings.Fields(row)...)
		}
	}

	file, err := os.Create("move.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range csvData {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}

	fmt.Println("CSV data has been written to move.csv")
}
