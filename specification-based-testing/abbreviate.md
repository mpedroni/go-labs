# Specification-based testing applied to the Abbreviate func

## Analyzing each input individually and related with others

_Partitions marked with a "\*" won't be tested multiple times_

### Inputs individually

- str (string)
  - \*empty
  - \*len == 1
  - len > 1
- abbrevMarker (string)
  - empty
  - \*len == 1
  - len > 1
- offset (int)
  - \*0
  - higher than 0
- maxWidth (int)
  - \*0
  - higher than 0

### Inputs relation

- str & abbrevMarker
  - \*both empty
  - \*str empty & abbrevMarker not empty
  - both not empty
- str & offset
  - \*offset > len(str)
  - offset <= len(str)
- str & maxWidth
  - len(str) == maxWidth
  - len(str) < maxWidth
  - len(str) > maxWidth
- abbrevMarker & offset
  - can't abbrev start (offset <= len(abbrevMarker)+1)
  - can abbrev start (offset > len(abbrevMarker)+1)
- str & abbrevMarker & offset & maxWidth
  - offset+maxWidth-len(abbrevMarker) < len(str)
  - offset+maxWidth-len(abbrevMarker) >= len(str)
