HexTile
=======

Library for building and managing tiles in a hex map.

Like apparently every other hex map library out there, this is based on [the
amazing post on hex tiles by Amit
Patel](https://www.redblobgames.com/grids/hexagons).

One major difference is that this library doesn't manage any hex<->pixel or
display information. This library was designed as a way to manage data for a map
on a backend server for a game, not to handle any of the actual display.
