Part 1
Two teams of monsters fight based on a set of rules until one team is fully defeated. On a high level this can be represented as a function taking two teams as parameters and returning the outcome containing battle log. At a high level, the API looks like

outcome = battle (team1, team2)
A team has a name and be or more monsters. Each monster has:

Name
Number of life points
Attack strength
The monsters fight one-on-one and attack each other in turns:

The team that is the first argument to battle goes first
The monsters fight in the same order they are listed in their team
On each turn a monster attacks by subtracting its "attack" from the opponents life points
Then its time for the opponent to attack and do the same until one has life points reduced to 0
When a monster from a team is defeated the next one from the same team takes its place and continues
The outcome contains:

Name of the winning team
Log of messages on the events e.g. "Wolf attacks Bear for 5 damage", "Bear is defeated",...
Please design the types for representing monsters, teams and outcomes, and implement the battle function. Example:

Team Blue
Dog: 3 life points, 2 attack 
Wolf: 4 life points, 1 attack

Team Red
Cat: 3 life points, 3 attack 
Tiger: 4 life points, 5 attack

Battle log:
Team Blue turn
Dog attacks Cat for 2 damage
Team Red turn
Cat attacks Dog for 3 damage 
Dog is defeated
Team Blue turn
Wolf attacks Cat for 1 damage
...
Part 2
Each monster has a weakness attribute representing a specific type. Add a type attribute to all monsters. During an attack:

If the attacker's type matches the defender's weakness, the damage is doubled.
If the defender's type matches the attacker's weakness, the damage is halved (rounded down).
Update the event log to include information about type advantages or disadvantages when they affect the damage calculation.

Part 3
Change the rules so the attacking team acts smarter. Instead of always using the first monster, they should pick the monster that deals the most damage.