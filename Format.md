# Field
- `"spell"` for spell text
- `"battlecry"` for Battlecry
- `"deathrattle"` for Deathrattle

# Basic Format
```
"spell": [{
  "<ACTION>": [args]
  "<NEXT_ACTION_AS_SAME_STEP>":
},
{
  "<NEXT_ACTION": [args]
}]
```

## Actions
- `"DAMAGE": [target, amount, spell_dmg]`:
  - `target`: Target, the target(s) to damage
  - `amount`: int, the amount of damage dealt to the target(s)
  - `spell_dmg`: bool, whether or not the damage is buffed by spell damage
- `DRAW: [cards]`
  - `cards`: int, number of cards to draw
- `DESTROY: [target]`
  - `target`: Target, the target(s) to destroy
- `RETURN_TO_OWNER_HAND: [target]`
  - `target`: Target, the target(s) to return
- `RESTORE: [target, amount]`
  - `target`: Target, the target(s) to restore health to
  - `amount`: int, the amount of health to restore to the target(s)

## Targets
```
{
  "BASE": [<BASE_GROUP1>, <BASE_GROUP2>]
  
}
```
### Base Groups
- `TARGET`: The user-selected target of this spell or battlecry, restrictions defined in `playReqirements` field
- `ADJACENT`: The minions adjacent to the target, not applicable if the target is not a minion in play
- `HERO`: The friendly hero
- `ENEMY_HERO`: The enemy hero
- `BOTH_HEROS`: Both heros
- `ALL`: All characters, including all minions and both heros
- `ALL_MINIONS`: All minions
- `ALL_FRIENDLY_MINIONS`: All friendly minions
- `ALL_ENEMY_MINIONS`: All enemy minions
- `ALL_ENEMIES`: All enemy minions and the enemy hero
- `ALL_FRIENDLIES`: All friendly minions
