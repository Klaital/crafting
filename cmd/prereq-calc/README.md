# Crafting Prerequisite Calculator

Given a desired item to be produced, answer questions:

1. What inputs will be needed?
2. What byproducts will be produced?
3. How many production buildings will I need producing each thing,
   including intermediate stages?

For example:

Given a set of recipes for producing Mineral Sludge from Blue Geodes,
what inputs are needed in order to produce 35 sludge per second?

    prereq-calc -library items.json -recipes sludge_from_blue_geodes.json \
                -desire-item mineral-sludge -desire-rate 35


I would expect to see something like

|        | Item | Rate |
| ---    | ---  | ---  |
| INPUT  | Blue Geode | 13.3 |
| INPUT  | Sulfuric Acid | 40 |
| OUTPUT | Mineral Sludge | 35 |
| OUTPUT | Sulfuric Waste Water | 56 |
| OUTPUT | Oxygen | 102 |
| OUTPUT | Hydrogen | 163 |