Problem: Plant Epidemic Simulation (Multiple Variants)
You are given a 2D grid representing a field of plants. Each cell contains one of the following characters:
•'X' — infected plant
•'.' — healthy, non-infected plant
•'I' — immune plant (cannot be infected)
•'#' — dead plant (cannot infect or be infected)
Cells are considered adjacent if they share a side (8-directional adjacency).


Part 1 — Basic Epidemic Spread
You are given:
•A 2D grid of 'X' and '.'
•An integer T
On each day, any healthy (.) plant that has at least T adjacent infected (X) plants becomes infected on the next day.
The epidemic ends when no additional plants become infected.
Return:

the number of days the epidemic lasts
(Count only days where new infections occur.)


Part 2 — Adding Immunity
Now the grid may also contain:
•'I' — immune plants, which never get infected.
Follow the same rules as Part 1, but ignore immune plants when considering infection.
Return:

the number of days the epidemic lasts


Part 3 — Infection → Immunity Timer
Now you are also given:
•An integer D: the number of days after infection for a plant to become immune.
Rules:
•When a plant becomes infected (. → X), it stays infected for D days, then becomes immune (I).
•Only currently-infected plants (X) can infect neighbors.
•The epidemic ends when no cell is infected (all cells are either I, . or #).
Return:

the number of days until no plants remain infected


Part 4 — Death Threshold and Final Output
Now the grid may contain:
•'X' — infected
•'.' — healthy
•'I' — immune
•'#' — dead
You are also given:
•An integer K — the number of infected neighbors required for a non-immune plant to die.
Rules are expanded:
Infection rule (same as earlier parts)
•A healthy plant (.) becomes infected (X) if it has at least T adjacent infected neighbors.
Death rule (new)
•Any non-immune, currently-alive plant (. or X) dies (#) if it has at least K adjacent infected (X) neighbors.
•A dead plant stays # forever and cannot infect others.
Immunity timer (same as Part 3)
•Infected plants become immune after D days.
The simulation ends when:
•No plant is infected (X) anymore.
Return:

(days the epidemic lasts, total number of plants that died)