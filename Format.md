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

### Actions
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
