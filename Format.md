# Field
- `spell` for spell text
- `battlecry` for Battlecry
- `deathrattle` for Deathrattle
- `inspire` for Inspire
- `triggers` for other triggers

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
- `DRAW: [target, cards]`
  - `target`: Target, either friendly hero or enemy hero, to draw cards for
  - `cards`: int, number of cards to draw
- `DESTROY: [target]`
  - `target`: Target, the target(s) to destroy
- `RETURN_TO_OWNER_HAND: [target]`
  - `target`: Target, the target(s) to return
- `RESTORE: [target, amount]`
  - `target`: Target, the target(s) to restore health to
  - `amount`: int, the amount of health to restore to the target(s)
- `MANA_THIS_TURN: [target, amount]`
  - `target`: Target, either friendly hero or enemy hero, to grant bonus mana crystal(s) for a single turn
  - `amount`: int, number of mana crystals to add
- `ENCHANTMENT: [target, enchantment]`
  - `target`: Target, the target to apply the enchantment to
  - `enchantment`: string, the ID of the enchantment to apply

## Targets
```
{
  "BASE": [groups, ...]
  "EXCEPT": [groups, ...]
  "RAND": [num, spell_dmg]
}
```

- `BASE`: Groups to include
- `EXCEPT`: Groups to exclude from pool of included targets
- `RAND`: Include if targets should be randomly selected
  - `num`: int, number of targets to randomly choose
  - `spell_dmg`: bool, whether or not the number of cards is affected by spell damage

**Targets may also be simply replaced with a single string of the base group to include**

### Groups
- `TARGET`: The user-selected target of this spell or battlecry, restrictions defined in `playReqirements` field
- `ADJACENT`: The minions adjacent to the target, not applicable if the target is not a minion in play
- `WEAPON`: The weapon currently equipped to the friendly hero
- `ENEMY_WEAPON`: The weapon currently equipped to the enemy hero
- `HERO`: The friendly hero
- `ENEMY_HERO`: The enemy hero
- `BOTH_HEROS`: Both heros
- `ALL`: All characters, including all minions and both heros
- `ALL_MINIONS`: All minions
- `ALL_FRIENDLY_MINIONS`: All friendly minions
- `ALL_ENEMY_MINIONS`: All enemy minions
- `ALL_ENEMIES`: All enemy minions and the enemy hero
- `ALL_FRIENDLIES`: All friendly minions

## Triggers
- `ON_USE`: Only for weapons?, triggered when a hero attacks while the weapon is equipped
