The godip variant generator
===========================

A generator to help creating godip compatible Diplomacy variants.

Creating a map-only variant involves a lot of manual effort, and this generator aims to remove most of the labour.  Note that it is expected that the output from the generator will still need to be tweaked by hand.

In particular the current generator does not cope with:

* Provinces with multiple coasts (e.g. Spain in the Classical game).
* Build anywhere.
* Victory conditions other than 'more than half'.
* Non-planar maps (the extra edges must be added by hand afterwards).

The generator is written in [Python 2](https://www.python.org/downloads/release/python-2714/), and it is must be installed to run the generator.

### General advice

Raise an issue on the godip GitHub page to show that you are working on a variant. This will provide a convenient place to ask questions you might have, or give updates about your progress. It also
helps to avoid (the highly unlikely) situation where two people are working on the same variant at once.

The generator script has only been used to make a handful of variants and so it almost certainly has bugs in it. Please feel free to ask questions about any issues you encounter, and hopefully we can
iron out any issues.

### Workflow to create a variant with the generator

1. Create an image of the map you are aiming to end up with
2. Create a svg by tracing over the map
3. Run the generator and fix any mistakes it makes
4. Manually finish off the variant
5. Add variant tests

#### Create an image of the map you are aiming to end up with

Either by using image editing software such as [Microsoft Paint](https://en.wikipedia.org/wiki/Microsoft_Paint) or [Gimp](https://www.gimp.org/), or by scanning a hand-drawn image, create an image showing the regions of the map.

#### Create a svg by tracing over the map

Using [Inkscape](https://inkscape.org/en/) create a new document with your sketch as the bottom layer.  Save it in the `generator` directory with a name ending in "_input.svg".

Create a new layer above called "points". Pick a junction between three or more regions of the map and use the circles tool (`F5`) to draw a small circle (holding `ctrl` while dragging may help with this).
The circle should be centred on the junction. Copy this circle and paste a new circle for each junction on your map. Also paste a circle at any point two regions touch the edge of the map.  If any
regions touch fewer than three circles then add additional circles spread around their perimeter to make it up to three (e.g. this was done for Portugal in the Classical map).

Create a new layer above called "edges". Use the Bezier curve tool (`shift+F6`) to draw a straight edged approximation of the borders between each junction. For example, on the Classical map, the border
between Kiel and Helgoland Bight was drawn as an edge with three straight line sections.  Don't worry about the ends of the lines being exactly at the junctions, the generator will move them to the
centre of the nearest circles for you.  When you have finished you should have a very jagged outline of a map.

Create a new layer above called "supply-centers". Copy the grey circle supply center symbol from any existing svg map and paste copies onto your new map for each supply center (including home centers).

Create a new layer above this called "province-centers". Paste more supply center symbols onto this layer whereever an army should be placed in a non-supply center province.

Create a new layer above this called "sea". Paste more supply center symbols onto this layer whereever a fleet should be placed in a sea region. If a province is coastal then it should only have a supply
center on the "province-centers" layer.

Create a new layer above this called "impassable". Paste more supply center symbols onto this layer in the middle of any region of the map that no units may enter (e.g. Switzerland in the Classical map).
Once this is done then there should be one supply center symbol for each region of the map.

Add a new layer above this called "names". Using the Text tool (`F8`) add a name for each region. The name must be within a single text box, but can contain new lines or multiple spaces.  The name will
be used to generate suitable abbreviations.

#### Run the generator and fix any mistakes it makes

Open the generator script and set the values of the variables in the section entitled "Data to be gathered for the variant". The abbreviations, center_overrides and region_overrides should be left blank
for the first run.

Run the generator script with Python 2:

```python generate.py```

It is very likely that an exception will be thrown, and the svg will need fixing.  Some example issues:

* `Warning: Duplicate edges: ...`: A region has only two junctions around it (and also two edges). Replace an edge with a junction and two edges. (This error can also be thrown if an edge has been copied
  and pasted, but this is less common).
* `XXX appears twice as a region name`: Two regions have the same name. They must be distinguished, or the generator cannot create ids for them.
* `Could not determine abbreviation for these names: ... Please add a suitable abbreviation to the ABBREVIATIONS config option.`: The generator couldn't come up with a suitable abbreviation, often because
  two names are similar or overlap (e.g. "Java" and "Java Sea"). Add an entry to the `ABBREVIATIONS` input variable to pick a suitable abbreviation for one or both of the regions, and then re-run.

Once the generator gets past a certain stage, it will create a multi-coloured debug map. This can be very useful for spotting regions that have 'bled' into each other (i.e. a single joined region which
should be two separate regions). The colours on this map are random, but by running the script a few times you can check if adjacent regions are different or the same.

Once the generator creates an output map then the ids of the supply centers, province centers and provinces can be checked by opening the map in Inkscape and using the "Object Properties" view. It's likely
that some will be swapped, and the `CENTER_OVERRIDES` and `REGION_OVERRIDES` variables should be used to swap these back. (It's possible to do this by hand too, but using the generator will also fix all
the connections in the map for you).

#### Manually finish off the variant

Once everything is roughly right then the generated map and go file should be moved to a variant directory.  The svg map file goes in a "svg" subdirectory, and go-bindata should be used to generate the
bindata.go file (see the main README for more details).  Update the variants.go file with a new import statement, and add the variant to the list in the file.

The map may need some manual tweaks (e.g. to add canals, coasts or other details), and once this is done then you can generate the bindata.go file.

You can generate some other maps to check that everything generally looks right by running `env DRAW_MAPS=true go test -v ./...` from the variants directory, and looking in the `test_output_maps` directory afterwards. One
thing to look out for is erroneous sea-connections between coastal regions that fleets shouldn't be able to travel between.

#### Add variant tests

It's also good to add some tests for your new variant, such as sample games or tests for special rules.
