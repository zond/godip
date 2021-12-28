#!/usr/bin/env python2
# -*- coding: utf-8 -*-

import xml.etree.ElementTree
import re
import math
import itertools
import collections
import random
import os
import yaml
from string import Template

# The name of the variant
VARIANT = 'Unconstitutional'

# Set to true to create an output map where it's easier to check the regions and centers have the right ids.
OVERRIDE_CHECK_MODE = False
# Whether to highlight the region abbreviation in bold or not.
BOLD_ABBREVIATIONS = True

### Constants ###

INK = '{http://www.inkscape.org/namespaces/inkscape}'
SVG = '{http://www.w3.org/2000/svg}'
# Any junctions within GUTTER pixels from the edge of the page will be moved to the edge.
GUTTER = 6
# How curvy the edges should be made
CURVE_WEIGHT = 0.5
# The background colour of sea regions
SEA_COLOR = '#d4d0ad'
# The background colour of the land
LAND_COLOR = '#f4d7b5'
# The thickness of thick lines
THICK = 2.225
#THICK = 1.1125
# The thickness of thin lines
THIN = 1
#THIN = 0.5

IMPASSABLE_PATTERN = '<pattern id="impassable" patternUnits="userSpaceOnUse" width="16" height="16" patternTransform="rotate(35)"><line x1="0" y="0" x2="0" y2="16" stroke="#000000" stroke-opacity="0.1" stroke-width="18" id="impassableLine" /></pattern>'
IMPASSABLE_STYLE = 'fill:url(#impassable);fill-rule:evenodd;stroke:#000000;stroke-width:1'

# Globals
reportedErrors = set()

class Flags:
    """A class to hold boolean attributes of a province."""
    def __init__(self, supplyCenter, province, sea, impassable):
        self.supplyCenter = supplyCenter
        self.province = province
        self.sea = sea
        self.impassable = impassable
    def __repr__(self):
        return 'sc:{0},p:{1},s:{2},i:{3}'.format(self.supplyCenter, self.province, self.sea, self.impassable)
class Province:
    """A DTO for province details constructed from the inputs."""
    def __init__(self, name, abbreviation, center, flags, edges):
        self.name = name
        self.abbreviation = abbreviation
        self.center = center
        self.flags = flags
        self.edges = edges
    def __repr__(self):
        return '{0}: {1}'.format(self.abbreviation, self.name)

def toCamelCase(string):
    """Remove all spaces without changing the case."""
    return string.replace(' ', '')

def toLowerAlphaNumeric(string):
    """Convert to lower case and only keep the letters and numbers."""
    return re.sub(r'[^a-z0-9]', '', string.lower())

MAP = toLowerAlphaNumeric(VARIANT) + '_input.svg'
configFile = toLowerAlphaNumeric(VARIANT) + '.yml'

if not os.path.isfile(configFile):
    print('Missing config file: ', configFile)
with open(configFile, 'r') as y:
    config = yaml.safe_load(y)
    # The first year of the game
    START_YEAR = config['START_YEAR']
    # The starting units
    START_UNITS = {} if config['START_UNITS'] == None else config['START_UNITS']
    # The nations in the variant
    NATIONS = START_UNITS.keys()
    # Abbreviations that should be used (rather than letting the script try to guess an abbreviation).
    ABBREVIATIONS = config['ABBREVIATIONS']
    # Overrides to swap centers. This only needs to contain something if the greedy algorithm fails.
    CENTER_OVERRIDES = config['CENTER_OVERRIDES']
    # Overrides to swap region names. This only needs to contain something if the greedy algorithm fails.
    REGION_OVERRIDES = config['REGION_OVERRIDES']

def getLayer(root, label):
    """Get the layer from root with the given Inkscape label."""
    layer = root.find('{}g[@{}label="{}"]'.format(SVG, INK, label))
    if layer == None:
        print('Input svg has no layer called: {}'.format(label))
    return layer

def removeAllLayers(root):
    """Remove all layers from root."""
    for layer in root.findall('{}g'.format(SVG)):
        root.remove(layer)

def addLayer(root, name, visible):
    """Add a new layer to root with the given name and visibility."""
    display = 'inline' if visible else 'none'
    if OVERRIDE_CHECK_MODE:
        display = 'inline; opacity: 1' if visible else 'inline; opacity: 0.5'
    'svg:style="display:none"'
    attrs = {'id': name, '{}groupmode'.format(INK): 'layer', '{}label'.format(INK): name, 'style': 'display:' + display}
    xml.etree.ElementTree.SubElement(root, '{}g'.format(SVG), attrs)
    return getLayer(root, name)

def locFrom(locString):
    """Convert a coordinate string (e.g. from a path) into a coordinate pair."""
    loc = locString.split(',')
    if len(loc) != 2:
        print('Could not get parts from loc: {}'.format(loc))
    return (float(loc[0]), float(loc[1]))

def addLocs(locA, locB):
    """Add two coordinate pairs."""
    return (locA[0] + locB[0], locA[1] + locB[1])

def subLocs(locA, locB):
    """Subtract locB from locA."""
    return (locA[0] - locB[0], locA[1] - locB[1])

def strFrom(loc):
    """Convert a coordinate pair into a string suitable for use in a path."""
    return ','.join(map(lambda x: str(x), loc))

def getCentersPath(root):
    """Use the supply-centers layer to get an SC path with the right radius."""
    layer = getLayer(root, 'supply-centers')
    for path in layer.findall('{}path'.format(SVG)):
        scPath = path.get('d').split(' ')
        bits = scPath[:1] + ['{0}'] + scPath[2:]
        return ' '.join(bits)
    raise Exception('Couldn\'t find a SC in the supply-centers layer.')

def reverseMapLookup(inputMap, value):
    """Determine the key which has the given value in the map. This returns the first such
    key found if there are many, and throws an exception if there are none."""
    for k, v in inputMap.items():
        if v == value:
            return k
    raise Exception('Could not find value {} in map {}'.format(value, inputMap))

def findDist(locA, locB):
    """Find the euclidean distance between two coordinates."""
    dx = locA[0] - locB[0]
    dy = locA[1] - locB[1]
    return math.sqrt(dx*dx + dy*dy)

def getCorners(root):
    """Determine the coordinates of the corners of an svg view box."""
    viewBox = map(float, root.get('viewBox').split(' '))
    return ((viewBox[0], viewBox[1]), (viewBox[2], viewBox[1]), (viewBox[2], viewBox[3]), (viewBox[0], viewBox[3]))

def getCentersWithin(root, layerLabel):
    """For the given layer, create a map from the id of the paths in it to the location of their starts."""
    centers = {}
    layer = getLayer(root, layerLabel)
    if layer != None:
        for path in layer.findall('{}path'.format(SVG)):
            name = path.get('id')
            loc = locFrom(path.get('d').split(' ')[1])
            centers[name] = loc
    return centers

def findMiddleOfEllipse(ellipse):
    """For the given svg ellipse (or circle) find the coordinates of the center."""
    x = float(ellipse.get('cx'))
    y = float(ellipse.get('cy'))
    if abs(x - corners[0][0]) < GUTTER:
        x = corners[0][0]
    elif abs(x - corners[2][0]) < GUTTER:
        x = corners[2][0]
    if abs(y - corners[0][1]) < GUTTER:
        y = corners[0][1]
    elif abs(y - corners[2][1]) < GUTTER:
        y = corners[2][1]
    return (x, y)

def getJunctions(root, corners):
    """Find the coordinates of the junction points in the 'points' layer."""
    junctions = []
    layer = getLayer(root, 'points')
    for circle in layer.findall('{}circle'.format(SVG)):
        junctions.append(findMiddleOfEllipse(circle))
    for ellipse in layer.findall('{}ellipse'.format(SVG)):
        junctions.append(findMiddleOfEllipse(ellipse))
    return junctions

def getToolParts(d):
    """Get all the instructions in the given 'd' path, grouped by the tool used."""
    return re.findall(r'[LM] [0-9\.\-e,]+', d)

def makeToolPart(tool, loc):
    """Create a 'd' path string using the given tool and location."""
    return '{0} {1}'.format(tool, strFrom(loc))

def getEdges(root):
    """Create a map from (start, end) to the d attribute containing all the edges"""
    edges = {}
    layer = getLayer(root, 'edges')
    for edge in layer.findall('{}path'.format(SVG)):
        d = edge.get('d')
        tool = None
        loc = (0,0)
        start = None
        toolStart = None
        for bit in d.strip().split(' '):
            # Check if this is a supported tool
            if bit in ['M', 'm', 'c', 'C', 'V', 'v', 'H', 'h', 'l', 'L', 'z', 'Z']:
                tool = bit
                toolStart = loc
                # Needed in case the x and y coordinates are space separated.
                xCoord = None
                if tool == 'c':
                    # Coordinates for the c tool come in threes
                    cIgnoreCount = 0
                elif tool in ['z', 'Z']:
                    # No need to close the path, so ignore.
                    continue
            elif re.match(r'[0-9\.,\-]+', bit):
                # Apply the tool to the numbers
                if tool == 'M' or tool == 'L':
                    if ',' in bit:
                        loc = locFrom(bit)
                    elif xCoord == None:
                        xCoord = float(bit)
                        # Assume the y coordinate is in the next bit.
                        continue
                    else:
                        yCoord = float(bit)
                        loc = (xCoord, yCoord)
                        xCoord = None
                elif tool == 'm' or tool == 'l':
                    loc = addLocs(loc, locFrom(bit))
                elif tool == 'c':
                    cIgnoreCount = (cIgnoreCount + 1) % 3
                    if cIgnoreCount != 0:
                        continue
                    loc = addLocs(loc, locFrom(bit))
                elif tool == 'C':
                    cIgnoreCount = (cIgnoreCount + 1) % 3
                    if cIgnoreCount != 0:
                        continue
                    if ',' in bit:
                        loc = locFrom(bit)
                    elif xCoord == None:
                        xCoord = float(bit)
                        # Assume the y coordinate is in the next bit.
                        continue
                    else:
                        yCoord = float(bit)
                        loc = (xCoord, yCoord)
                elif tool == 'V':
                    loc = (loc[0], float(bit))
                elif tool == 'v':
                    loc = addLocs(toolStart, (0, float(bit)))
                elif tool == 'H':
                    loc = (float(bit), loc[1])
                elif tool == 'h':
                    loc = addLocs(toolStart, (float(bit), 0))
                else:
                    raise Exception('Unrecognised tool')
                if start == None:
                    d = ''
                d += 'L {0} '.format(strFrom(loc))
                start = (loc if start == None else start)
            else:
                raise Exception('Unsupported tool: ' + bit + ' in ' + edge.get('d'))
        end = loc
        forwardD = 'M'.join(d.split('L', 1))
        reversedD = 'M'.join(d.rsplit('L', 1))
        reversedD = ' '.join(reversed(getToolParts(reversedD)))
        edges[(start, end)] = (forwardD, reversedD)
    return edges

def findClosestJunction(junctions, point):
    """Find the closest junction point to the given location."""
    closestJunction = None
    closestDist = 1000000
    for junction in junctions:
        dist = findDist(junction, point)
        if dist < closestDist:
            closestJunction = junction
            closestDist = dist
    return closestJunction

def findDesiredEdges(junctions, edges):
    """Create a set of edges by snapping the ends of the given edges to the nearest junction points."""
    desiredEdges = {}
    for edge, d in edges.items():
        start = findClosestJunction(junctions, edge[0])
        end = findClosestJunction(junctions, edge[1])
        d0 = [makeToolPart('M', start)] + getToolParts(d[0])[1:-1] + [makeToolPart('L', end)]
        d1 = [makeToolPart('M', end)] + getToolParts(d[1])[1:-1] + [makeToolPart('L', start)]
        desiredEdges[(start, end)] = (' '.join(d0), ' '.join(d1))
    return desiredEdges

def vectorAngle(start, end):
    """Find the angle of the vector from start to end."""
    return math.atan2(end[1] - start[1], end[0] - start[0])

def findLinks(junction, edges):
    """Get all the junctions that neighbour the given junction, ordered by the angle of the link."""
    links = []
    linkAngles = {}
    for edge in edges:
        if edge[0] == junction:
            links.append(edge[1])
            linkAngles[edge[1]] = vectorAngle(junction, edge[1])
        elif edge[1] == junction:
            links.append(edge[0])
            linkAngles[edge[0]] = vectorAngle(junction, edge[0])
    # Sort the linked junctions to be in order of angle
    for i in range(len(links)):
        for j in range(i + 1, len(links)):
            if linkAngles[links[i]] > linkAngles[links[j]]:
                links[i], links[j] = links[j], links[i]
    return links

def ccwBorderDist(borderJunction, corners):
    """Find the distance to the given junction (that is on the border of the map) by following the edges of the map."""
    width = abs(corners[2][0] - corners[0][0])
    height = abs(corners[2][1] - corners[0][1])
    left = corners[0][0]
    top = corners[0][1]
    if borderJunction[0] == corners[0][0]:
        return borderJunction[1] - top
    elif borderJunction[1] == corners[2][1]:
        return height + borderJunction[0] - left
    elif borderJunction[0] == corners[2][0]:
        return 2 * height + width - (borderJunction[1] - top)
    elif borderJunction[1] == corners[0][1]:
        return 2 * height + 2 * width - (borderJunction[0] - left)
    else:
        raise Exception('Border junction is not on border: ' + borderJunction)

def makeRegions(junctions, edges, corners):
    """Take the given edges and the four sides of the map, and create the set of cycles that
    represents regions on the map."""
    global reportedErrors
    
    for corner in corners:
        if corner not in junctions:
            junctions.append(corner)
            
    # Add in edges from junctions on border rectangle (each once only)
    borderJunctions = list(corners)
    for junction in junctions:
        if junction[0] in (corners[0][0], corners[2][0]) or junction[1] in (corners[0][1], corners[2][1]):
            if junction not in borderJunctions:
                borderJunctions.append(junction)
    # Sort the border junctions starting from corner[0]
    for i in range(len(borderJunctions)):
        for j in range(i + 1, len(borderJunctions)):
            if ccwBorderDist(borderJunctions[i], corners) > ccwBorderDist(borderJunctions[j], corners):
                borderJunctions[i], borderJunctions[j] = borderJunctions[j], borderJunctions[i]
    for i, borderJunction in enumerate(borderJunctions):
        nextJunction = borderJunctions[(i + 1) % len(borderJunctions)]
        edges.append((nextJunction, borderJunction))
            
    # A region is a list of junctions in order.
    junctionLinks = {}
    for junction in junctions:
        junctionLinks[junction] = findLinks(junction, edges)
    directedEdges = []
    for edge in edges:
        directedEdges.append(edge)
        directedEdges.append((edge[1], edge[0]))
    if len(set(edges)) < len(edges):
        print('Warning: Duplicate edges: {0}'.format([edge for edge in edges if edges.count(edge) > 1]))

    regions = []
    while len(directedEdges) > 0:
        firstJunction, currentJunction = directedEdges.pop()
        region = [firstJunction]
        previousJunction = firstJunction
        while currentJunction != firstJunction:
            region.append(currentJunction)
            links = junctionLinks[currentJunction]
            nextJunction = links[(links.index(previousJunction) + 1) % len(links)]
            previousJunction, currentJunction = currentJunction, nextJunction
            try:
                directedEdges.remove((previousJunction, currentJunction))
            except:
                if (previousJunction, currentJunction) not in reportedErrors:
                    print('Issue removing directed edge:', (previousJunction, currentJunction), str(region)[:100] + ('...' if len(str(region)) > 100 else ''))
                    reportedErrors.add((previousJunction, currentJunction))
                #raise
        # Don't include the region that contains all four corners
        allFourCorners = True
        for corner in corners:
            if corner not in region:
                allFourCorners = False
                break
        if not allFourCorners:
            regions.append(region)
    return regions

def middleOfRegion(region):
    """Estimate the middle of the region by the middle of its bounding box."""
    xs = map(lambda r: r[0], region)
    ys = map(lambda r: r[1], region)
    return ((max(xs) + min(xs)) / 2.0, (max(ys) + min(ys)) / 2.0)

def findClosestPair(pointsA, pointsB):
    """Find the shortest line segment from a point in the first set to a point in the second set."""
    bestDist = 1000000
    bestPair = None
    for a in pointsA:
        for b in pointsB:
            d = findDist(a, b)
            if d < bestDist:
                bestDist = d
                bestPair = (a, b)
    return bestPair

def matchRegionMarkerToRegion(regions, allMarkers):
    """Use a greedy algorithm to match regions based on how close their middle is to a marker."""
    middleToRegion = {}
    for region in regions:
        middleToRegion[middleOfRegion(region)] = region
    centerToName = {}
    for name, center in allMarkers.items():
        centerToName[center] = name
    middlePoints = list(middleToRegion.keys())
    centerPoints = list(allMarkers.values())
    if len(middlePoints) != len(centerPoints):
        raise Exception('There are {0} areas on the map, and {1} markers.'.format(len(middlePoints), len(centerPoints)))
    regionNames = {}
    while len(middlePoints) > 0:
        m, c = findClosestPair(middlePoints, centerPoints)
        regionNames[centerToName[c]] = middleToRegion[m]
        middlePoints.remove(m)
        centerPoints.remove(c)
    return regionNames

def makeIdToAbbrMap(originalIdToRegion, regionFullNames, abbreviations):
    """Create a map from the original id of a center in the input to a region abbreviation
    that will be in the output."""
    idToAbbrMap = {}
    for originalId, region in originalIdToRegion.items():
        for name, anotherRegion in regionFullNames.items():
            if region == anotherRegion:
                idToAbbrMap[originalId] = abbreviations[name]
                break
    return idToAbbrMap

def replaceOriginalIds(centers, originalIdToAbbr, abbreviations):
    """Take a map where the keys are ids from the input svg, and return a map where the
    keys are abbreviations used in the output."""
    output = {}
    for originalId in centers.keys():
        try:
            output[originalIdToAbbr[originalId]] = centers[originalId]
        except:
            print originalId
            if originalId in originalIdToAbbr:
                print originalIdToAbbr[originalId]
            else:
                print originalIdToAbbr.values()
            raise
    return output

def nameFromTextElement(text):
    """Get a name tuple from a text element."""
    name = []
    for tspan in text.findall('.//{}tspan'.format(SVG)):
        if tspan.text != None:
            name += re.split(r' +', tspan.text)
    name = map(lambda bit: bit.strip(), name)
    name = filter(lambda bit: bit != '', name)
    return name

def makeLocsToNames(namesLayer):
    """Create a map from coordinates to names of provinces."""
    locsToNames = {}
    for text in namesLayer.findall('{}text'.format(SVG)):
        name = nameFromTextElement(text)
        transform = text.get('transform')
        x, y = float(text.get('x')), float(text.get('y'))
        if transform != None:
            if re.match(r'^rotate\([^\(\)\,]*\)$', transform):
                angle = math.radians(float(transform.split('(')[1].split(')')[0]))
                x, y = x * math.cos(angle) - y * math.sin(angle), x * math.sin(angle) + y * math.cos(angle)
            elif re.match(r'^matrix\([^\(\)]*\)$', transform):
                args = map(float, transform.split('(')[1].split(')')[0].split(','))
                x, y = args[0] * x + args[2] * y + args[4], args[1] * x + args[3] * y + args[5]
            elif re.match(r'^scale\([^\(\)]*\)$', transform):
                args = map(float, transform.split('(')[1].split(')')[0].split(','))
                x, y = args[0] * x, args[1] * y
            else:
                print('Unsupported text transformation: ' + transform + ' used for ' + ' '.join(name))
        loc = (x, y)
        if ''.join(name).strip() != '':
            locsToNames[loc] = tuple(name)
    return locsToNames

def guessRegionFullNames(regions, namesLayer):
    """Use a greedy algorithm to name regions based on how close their middle is to a center."""
    locsToNames = makeLocsToNames(namesLayer)
    middleToRegion = {}
    for region in regions:
        middleToRegion[middleOfRegion(region)] = region
    middlePoints = list(middleToRegion.keys())
    centerPoints = list(locsToNames.keys())
    print 'Number of middle points: {0}, Name points: {1}, Regions: {2}'.format(len(middlePoints), len(centerPoints), len(regions))
    regionNames = {}
    while len(centerPoints) > 0:
        m, c = findClosestPair(middlePoints, centerPoints)
        if locsToNames[c] in regionNames:
            raise Exception('{0} appears twice as a region name'.format(locsToNames[c]))
        regionNames[locsToNames[c]] = middleToRegion[m]
        middlePoints.remove(m)
        centerPoints.remove(c)
    return regionNames

def regionsDifference(regionsA, regionsB):
    """Find the regions in the first iterable, but not the second."""
    difference = []
    for region in regionsA:
        if region not in regionsB:
            difference.append(region)
    return difference

def abbrFromName(name, indexes):
    """Create a potential abbreviation from a name by picking out the letters at the given indexes."""
    abbr = ''
    for i in indexes:
        abbr += toLowerAlphaNumeric(name)[i]
    return abbr

def findTupleFromName(name, fullNamesTuples):
    """Find the name tuple that has the same letters (in the same order) as the given name."""
    name = toLowerAlphaNumeric(name)
    for fullNameTuple in fullNamesTuples:
        if toLowerAlphaNumeric(''.join(fullNameTuple)) == name:
            return fullNameTuple
    raise Exception('Couldn\'t find tuple matching {0}'.format(name))

nameToColorMap = {}
def colorForName(name):
    """Come up with a random color for the given name."""
    global nameToColorMap
    if name not in nameToColorMap.keys():
        nameToColorMap[name] = '#' + ''.join(random.sample('0123456789abcdef', 6))
    return nameToColorMap[name]

def abbreviationsForNames(fullNamesTuples, indexSets, abbrCount):
    """Try to create unique abbreviations for the given name tuples by considering the letters at the given indexes."""
    abbrMap = {}
    for indexes in indexSets:
        for nTuple in fullNamesTuples:
            n = ''.join(nTuple).lower()
            if indexes[-1] < len(n):
                abbrCount[abbrFromName(n, indexes)] += 1
        for fullNameTuple, abbr in abbrMap.items():
            if abbrCount[abbr] > 1:
                del abbrMap[fullNameTuple]
        for nTuple in fullNamesTuples:
            n = ''.join(nTuple).lower()
            if indexes[-1] < len(n):
                if nTuple not in abbrMap.keys() and abbrCount[abbrFromName(n, indexes)] == 1:
                    # Find the first suitable abbreviation (as we may have skipped a sensible one).
                    for nIndexes in indexSets:
                        if nIndexes[-1] < len(n):
                            if abbrCount[abbrFromName(n, nIndexes)] == 1:
                                abbrMap[nTuple] = abbrFromName(n, indexes)
                                break
        if len(abbrMap) == len(fullNamesTuples):
            break
    return abbrMap

def threeWordAbbrABCCheck(nTuple):
    """Check if the name has three or more words."""
    return len(nTuple) >= 3

def threeWordAbbrABC(nTuple):
    """Create an abbreviation from the first letter of the first three words."""
    abbr = ''
    if threeWordAbbrABCCheck(nTuple):
        for word in nTuple[:3]:
            if len(word) > 0:
                abbr += word[0].lower()
    return abbr

def twoWordAbbrAABCheck(nTuple):
    """Check if an abbreviation of the form aab can be formed."""
    return len(nTuple) >= 2 and len(nTuple[0]) >= 2

def twoWordAbbrAAB(nTuple):
    """Create an abbreviation from the first two letters of the first word and the first letter of the second."""
    abbr = ''
    if twoWordAbbrAABCheck(nTuple):
        abbr = nTuple[0].lower()[:2] + nTuple[1].lower()[0]
    return abbr

def twoWordAbbrABBCheck(nTuple):
    """Check if an abbreviation of the form abb can be formed."""
    return len(nTuple) >= 2 and len(nTuple[1]) >= 2

def twoWordAbbrABB(nTuple):
    """Create an abbreviation from the first letter of the first word and the first two letters of the second."""
    abbr = ''
    if twoWordAbbrABBCheck(nTuple):
        abbr = nTuple[0].lower()[0] + nTuple[1].lower()[:2]
    return abbr

def abbreviationsForNamesUsingFunctions(fullNamesTuples, abbrCount, checkFunction, abbrFunction):
    """Try to create unique abbreviations for the given name tuples by considering the first letters of the first three words."""
    abbrMap = {}
    for nTuple in fullNamesTuples:
        if checkFunction(nTuple):
            abbrCount[abbrFunction(nTuple)] += 1
    for fullNameTuple, abbr in abbrMap.items():
        if abbrCount[abbr] > 1:
            del abbrMap[fullNameTuple]
    for nTuple in fullNamesTuples:
        if checkFunction(nTuple):
            if nTuple not in abbrMap.keys() and abbrCount[abbrFunction(nTuple)] == 1:
                abbrMap[nTuple] = abbrFunction(nTuple)
    return abbrMap

def inventAbbreviations(fullNamesTuples):
    """Determine a suitable set of unique abbreviations for the given name tuples. Use any values from
    the user ABBREVIATIONS override list."""
    fixedAbbrs = {}
    abbrCount = collections.Counter()
    for name, abbr in map(lambda na: (toLowerAlphaNumeric(na[0]), na[1]), ABBREVIATIONS.items()):
        fixedAbbrs[findTupleFromName(name, fullNamesTuples)] = abbr
        abbrCount[abbr] += 1
    # Repeat a few times in case it helps.
    for i in range(4):
        # Start by taking any unique first three letters.
        remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
        fixedAbbrs.update(abbreviationsForNames(remainingNames, [(0, 1, 2)], abbrCount))
        # Add any abbreviations from the first three words (if the name has three words).
        remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
        fixedAbbrs.update(abbreviationsForNamesUsingFunctions(remainingNames, abbrCount, threeWordAbbrABCCheck, threeWordAbbrABC))
        # Add any abbreviations from the first two words (if the name has two words).
        remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
        fixedAbbrs.update(abbreviationsForNamesUsingFunctions(remainingNames, abbrCount, twoWordAbbrAABCheck, twoWordAbbrAAB))
        remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
        fixedAbbrs.update(abbreviationsForNamesUsingFunctions(remainingNames, abbrCount, twoWordAbbrABBCheck, twoWordAbbrABB))
    # combinations returns the indexes in lexigographical order, which is basically what we want.
    remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
    if len(remainingNames) > 0:
        try:
            maxLength = max(map(lambda nameTuple: len(''.join(nameTuple)), remainingNames))
        except:
            print 'Failed to abbreviate these: {0}. Had {1} and made {2}'.format(remainingNames, fullNamesTuples, fixedAbbrs)
            raise
        fixedAbbrs.update(abbreviationsForNames(remainingNames, list(itertools.combinations(range(maxLength), 3)), abbrCount))
        if len(fixedAbbrs) != len(fullNamesTuples):
            print fixedAbbrs
            print 'Managed to abbreviate {0} regions.'.format(len(fixedAbbrs))
            raise Exception('Could not determine abbreviation for these names: {0}. Please add a suitable abbreviation to the ABBREVIATIONS config option.'.format(set(fullNamesTuples).difference(set(fixedAbbrs.keys()))))
    return fixedAbbrs

def makeFullNameToAbbr(regionFullNames, regionNames):
    """Create a map from region abbreviation to region name. Note that currently this takes the region abbreviations
    as an input, but at some point it might be useful to rewrite it so that it determines a set of sensible abbreviations."""
    fullNameToAbbr = {}
    for fullName, region in regionFullNames.items():
        for abbr, abbrRegion in regionNames.items():
            if region == abbrRegion:
                fullNameToAbbr[fullName] = abbr
    return fullNameToAbbr

def performOverrides(provinces):
    """Swap province centers, flags and edge sets according to the configured manual override values."""
    for nameA, nameB in CENTER_OVERRIDES:
        nameA, nameB = tuple(nameA.split(' ')), tuple(nameB.split(' '))
        provinceA = [province for province in provinces if province.name == nameA][0]
        provinceB = [province for province in provinces if province.name == nameB][0]
        provinceA.center, provinceB.center = provinceB.center, provinceA.center
        provinceA.flags, provinceB.flags = provinceB.flags, provinceA.flags
    for nameA, nameB in REGION_OVERRIDES:
        nameA, nameB = tuple(nameA.split(' ')), tuple(nameB.split(' '))
        if nameA not in [p.name for p in provinces]:
            raise Exception('{0} not found on map.'.format(nameA))
        if nameB not in [p.name for p in provinces]:
            raise Exception('{0} not found on map.'.format(nameB))
        provinceA = [province for province in provinces if province.name == nameA][0]
        provinceB = [province for province in provinces if province.name == nameB][0]
        provinceA.edges, provinceB.edges = provinceB.edges, provinceA.edges

def addNamesLayer(root, namesLayer, fullNameToAbbr, passableCenterAbbrs):
    """Add the names layer to root and try to highlight the abbreviation in bold (if BOLD_ABBREVIATIONS is set)."""
    for text in namesLayer.findall('{}text'.format(SVG)):
        name = nameFromTextElement(text)
        if ''.join(name) == '':
            continue
        abbr = fullNameToAbbr[tuple(name)]
        if abbr not in passableCenterAbbrs:
            namesLayer.remove(text)
            continue
        elif OVERRIDE_CHECK_MODE:
            style = text.get('style')
            style = re.sub(r'fill:[^;]*;', 'fill:{};'.format(colorForName(abbr)), style)
            text.set('style', style)
        boldAbbr = ''
        i = 0
        for tspan in text.findall('.//{}tspan'.format(SVG)):
            if tspan.text != None:
                oldText = tspan.text
                if BOLD_ABBREVIATIONS:
                    newParts = []
                    j = len(abbr)
                    while j > i:
                        if abbr[i:j] in oldText.lower():
                            start = oldText.lower().index(abbr[i:j])
                            # Fix to try to ensure a part isn't just whitespace (which gets stripped by some renderers).
                            if oldText[:start].endswith(' '):
                                newParts += [oldText[:start-1], oldText[start-1:start+j-i]]
                            else:
                                newParts += [oldText[:start], oldText[start:start+j-i]]
                            oldText = oldText[start+j-i:]
                            i = j
                            j = len(abbr)
                            if i >= len(abbr):
                                break
                        else:
                            j -= 1
                    if len(newParts) > 0:
                        boldAbbr += ''.join(newParts[1::2]).lower()
                        newParts.append(re.sub('.*' + newParts[-1], '', oldText))
                        tspan.text = None
                        for index, part in enumerate(newParts):
                            if len(part) != 0:
                                fontWeight = 'normal' if index % 2 == 0 else 'bold'
                                attributes = {'style': 'font-weight:' + fontWeight}
                                e = xml.etree.ElementTree.Element('{}tspan'.format(SVG), attributes)
                                e.text = part
                                tspan.append(e)
        if BOLD_ABBREVIATIONS and toCamelCase(boldAbbr) != abbr:
            print 'Failed to automatically bold the abbreviation for {0} (got "{1}" rather than "{2}")'.format(name, boldAbbr, abbr)
            # TODO Fall back to a greedy algorithm here.
    root.append(namesLayer)

def calculateCurvePoints(lastLoc, loc, nextLoc):
    """Calculate a pair of bezier curve handle points. These points will be on the straight line
    parallel to the line joining lastLoc and nextLoc, which passes through loc."""
    # Check if gradient is infinity
    if nextLoc[0] == lastLoc[0]:
        yA = (loc[1] + CURVE_WEIGHT * lastLoc[1]) / (CURVE_WEIGHT + 1.0)
        locA = (loc[0], yA)
        yB = (loc[1] + CURVE_WEIGHT * nextLoc[1]) / (CURVE_WEIGHT + 1.0)
        locB = (loc[0], yB)
    else:
        gradient = (nextLoc[1] - lastLoc[1]) / (nextLoc[0] - lastLoc[0])
        offset = loc[1] - gradient * loc[0]
        if abs(gradient) < 1:
            xA = (loc[0] + CURVE_WEIGHT * lastLoc[0]) / (CURVE_WEIGHT + 1.0)
            locA = (xA, gradient * xA + offset)
            xB = (loc[0] + CURVE_WEIGHT * nextLoc[0]) / (CURVE_WEIGHT + 1.0)
            locB = (xB, gradient * xB + offset)
        else:
            yA = (loc[1] + CURVE_WEIGHT * lastLoc[1]) / (CURVE_WEIGHT + 1.0)
            locA = ((yA - offset) / gradient, yA)
            yB = (loc[1] + CURVE_WEIGHT * nextLoc[1]) / (CURVE_WEIGHT + 1.0)
            locB = ((yB - offset) / gradient, yB)
    return locA, locB

def addStripesPattern(root):
    """Add the pattern that's used for indicating provinces that can be clicked on by the player."""
    # Nb. It's called "stripes" but it's actually squares.
    '''<pattern id="stripes" patternUnits="userSpaceOnUse" x="0" y="0" width="24" height="24">
        <rect id="Rectangle" fill="#000000" fill-opacity="0.13" x="0" y="0" width="6" height="6"></rect>
        <rect id="Rectangle" fill="#000000" fill-opacity="0.13" x="12" y="12" width="6" height="6"></rect>
        <rect id="Rectangle" fill="#ffffff" fill-opacity="0.13" x="12" y="0" width="6" height="6"></rect>
        <rect id="Rectangle" fill="#ffffff" fill-opacity="0.13" x="0" y="12" width="6" height="6"></rect>
    </pattern>'''
    xml.etree.ElementTree.SubElement(root, '{}pattern'.format(SVG), {'id': 'stripes', 'patternUnits': 'userSpaceOnUse', 'x': '0', 'y': '0', 'width': '24', 'height': '24'})
    stripes = root.find('{0}pattern[@id="stripes"]'.format(SVG))
    xml.etree.ElementTree.SubElement(stripes, '{}rect'.format(SVG), {'id': 'square0', 'fill': '#000000', 'fill-opacity': '0.13', 'x': '0', 'y': '0', 'width': '6', 'height': '6'})
    xml.etree.ElementTree.SubElement(stripes, '{}rect'.format(SVG), {'id': 'square1', 'fill': '#000000', 'fill-opacity': '0.13', 'x': '12', 'y': '12', 'width': '6', 'height': '6'})
    xml.etree.ElementTree.SubElement(stripes, '{}rect'.format(SVG), {'id': 'square2', 'fill': '#ffffff', 'fill-opacity': '0.13', 'x': '12', 'y': '0', 'width': '6', 'height': '6'})
    xml.etree.ElementTree.SubElement(stripes, '{}rect'.format(SVG), {'id': 'square3', 'fill': '#ffffff', 'fill-opacity': '0.13', 'x': '0', 'y': '12', 'width': '6', 'height': '6'})

def addImpassablePattern(root):
    """Add the pattern that's used for indicating regions are impassable."""
    '''<pattern id="impassableStripes" patternUnits="userSpaceOnUse" width="16" height="16" patternTransform="rotate(35)">
        <line id="impassableStripesLine" x1="0" y="0" x2="0" y2="16" stroke="#000000" stroke-opacity="0.1" stroke-width="18" />
    </pattern>'''
    xml.etree.ElementTree.SubElement(root, '{}pattern'.format(SVG), {'id': 'impassableStripes', 'patternUnits': 'userSpaceOnUse', 'width': '16', 'height': '16', 'patternTransform': 'rotate(35)'})
    impassableStripes = root.find('{0}pattern[@id="impassableStripes"]'.format(SVG))
    xml.etree.ElementTree.SubElement(impassableStripes, '{}line'.format(SVG), {'id': 'impassableStripesLine', 'x1': '0', 'y':'0', 'x2': '0', 'y2': '16', 'stroke': '#000000', 'stroke-opacity': '0.1', 'stroke-width': '18'})

def addBlurFilter(root):
    """Add the Gaussian blur filter that's used for the shadow effect along the coast."""
    xml.etree.ElementTree.SubElement(root, '{}defs'.format(SVG), {'id': 'defs'})
    defs = root.find('{0}defs[@id="defs"]'.format(SVG))
    xml.etree.ElementTree.SubElement(defs, '{}filter'.format(SVG), {'id': 'blur', 'style': 'color-interpolation-filters:sRGB', 'x':'-10', 'y':'-10', 'width': '20', 'height': '20'})
    blur = defs.find('{0}filter[@id="blur"]'.format(SVG))
    xml.etree.ElementTree.SubElement(blur, '{}feGaussianBlur'.format(SVG), {'id': 'feBlur', 'stdDeviation': '2.5'})

def addPatterns(root):
    """Add the common patterns that will be used by map features."""
    addStripesPattern(root)
    addImpassablePattern(root)
    addBlurFilter(root)

def addRectToLayer(layer, corners, fill):
    """Draw a rectangle using the given corners and fill in the specified color."""
    fillStyle = 'fill:{};fill-opacity:1;'.format(SEA_COLOR) if fill else 'fill:none;'
    style= fillStyle + 'display:inline;stroke:#000000;stroke-width:{};stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1'.format(THICK)
    width = '{}'.format(corners[2][0] - corners[0][0])
    height = '{}'.format(corners[2][1] - corners[0][1])
    xml.etree.ElementTree.SubElement(layer, '{}rect'.format(SVG), {'id': 'bg_rect', 'style': style, 'width': width, 'height': height, 'x': '0', 'y': '0'})

def addLayerWithRegions(root, regionNames, edgeToDMap, layerName, color, visible, corners = None, edgesOnly = False):
    """Create a layer containing the provinces coloured in."""
    layer = addLayer(root, layerName, visible)
    if corners != None:
        addRectToLayer(layer, corners, True)
    for name, region in regionNames.items():
        d = 'M {} C '.format(strFrom(region[0]))
        lastR = region[0]
        for r in region[1:] + [region[0]]:
            if (r, lastR) in edgeToDMap.keys():
                instruction = edgeToDMap[(r, lastR)][1] + ' '
            elif (lastR, r) in edgeToDMap.keys():
                instruction = edgeToDMap[(lastR, r)][0] + ' '
            else:
                instruction = 'M {0} L {1} '.format(strFrom(lastR), strFrom(r))
            toolParts = getToolParts(instruction)
            for i, toolPart in enumerate(toolParts):
                tool, inst = toolPart.split(' ', 1)
                loc = locFrom(inst)
                if i == 0:
                    d += '{0} '.format(strFrom(loc))
                elif i == len(toolParts) - 1:
                    d += '{0} {1} '.format(strFrom(loc), strFrom(loc))
                else:
                    lastLoc = locFrom(toolParts[i-1].split(' ', 1)[1])
                    nextLoc = locFrom(toolParts[i+1].split(' ', 1)[1])
                    locA, locB = calculateCurvePoints(lastLoc, loc, nextLoc)
                    d += '{0} {1} {2} '.format(strFrom(locA), strFrom(loc), strFrom(locB))
            lastR = r
        d += ' z'
        regionColor = color if color is not None else colorForName(name)
        style = 'fill:{0};fill-opacity:1;vector-effect:none;fill-rule:evenodd'.format(regionColor)
        if edgesOnly:
            style += ';stroke:#000000;stroke-width:2px'
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': name, 'd': d, 'style': style})

def addShadowsToBackground(root, edgeStyle, edgeToNames):
    layer = getLayer(root, 'background')
    edgeIds = set()
    for edge, biedge in edgeToDMap.items():
        thickness = float(re.search(r'stroke-width:([^;]*)', edgeStyle[edge]).group(1))
        if thickness != THICK:
            continue
        edgePath = biedge[0]
        toolParts = getToolParts(edgePath)
        tool, inst = toolParts[0].split(' ', 1)
        loc = locFrom(inst)
        d = 'M {} C '.format(strFrom(loc))
        for i, toolPart in enumerate(toolParts):
            tool, inst = toolPart.split(' ', 1)
            loc = locFrom(inst)
            if i == 0:
                d += '{0} '.format(strFrom(loc))
            elif i == len(toolParts) - 1:
                d += '{0} {1} '.format(strFrom(loc), strFrom(loc))
            else:
                lastLoc = locFrom(toolParts[i-1].split(' ', 1)[1])
                nextLoc = locFrom(toolParts[i+1].split(' ', 1)[1])
                locA, locB = calculateCurvePoints(lastLoc, loc, nextLoc)
                d += '{0} {1} {2} '.format(strFrom(locA), strFrom(loc), strFrom(locB))
        edgeId = 's_'+'_'.join(edgeToNames[edge])
        n = 2
        while edgeId in edgeIds:
            edgeId = 's_'+'_'.join(edgeToNames[edge]) + '_{0}'.format(n)
            n += 1
        edgeIds.add(edgeId)
        shadow = xml.etree.ElementTree.Element('{}path'.format(SVG), {'id': edgeId, 'd': d, 'style': 'fill:none;vector-effect:none;fill-rule:evenodd;stroke:#000100;stroke-width:{};stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1;filter:url(#blur)'.format(thickness)})
        layer.insert(1, shadow)

def getEdgeStyle(edges, provinces):
    """Create a map giving the desired style of each edge.  Coastal edges are thick, as are
    edges involving impassable areas. All other edges are thin. Sea-sea edges are dashed."""
    edgeStyle = {}
    edgeToNames = {}
    for edge in edges:
        touches = set()
        abbreviations = []
        for province in provinces:
            if edge[0] in province.edges and edge[1] in province.edges and abs(province.edges.index(edge[0]) - province.edges.index(edge[1])) in [1, len(province.edges)-1]:
                if province.flags.sea:
                    touches.add('sea')
                else:
                    touches.add('land')
                if province.flags.impassable:
                    touches.add('impassible')
                abbreviations.append(province.abbreviation)
        edgeToNames[edge] = abbreviations
        if touches == set(['sea']):
            edgeStyle[edge] = 'fill:none;fill-rule:evenodd;stroke:#000000;stroke-width:{};stroke-dasharray:4'.format(THIN)
        else:
            thickness = THICK if len(touches) > 1 else THIN
            edgeStyle[edge] = 'fill:none;fill-rule:evenodd;stroke:#000100;stroke-width:{};stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none'.format(thickness)
    return edgeStyle, edgeToNames

def addForeground(root, edgeToDMap, edgeStyle, edgeToNames, corners):
    """Create the foreground layer, consisting of all the edges and a map border."""
    layer = addLayer(root, 'foreground', True)
    edgeIds = set()
    for edge, biedge in edgeToDMap.items():
        edgePath = biedge[0]
        toolParts = getToolParts(edgePath)
        tool, inst = toolParts[0].split(' ', 1)
        loc = locFrom(inst)
        d = 'M {} C '.format(strFrom(loc))
        for i, toolPart in enumerate(toolParts):
            tool, inst = toolPart.split(' ', 1)
            loc = locFrom(inst)
            if i == 0:
                d += '{0} '.format(strFrom(loc))
            elif i == len(toolParts) - 1:
                d += '{0} {1} '.format(strFrom(loc), strFrom(loc))
            else:
                lastLoc = locFrom(toolParts[i-1].split(' ', 1)[1])
                nextLoc = locFrom(toolParts[i+1].split(' ', 1)[1])
                locA, locB = calculateCurvePoints(lastLoc, loc, nextLoc)
                d += '{0} {1} {2} '.format(strFrom(locA), strFrom(loc), strFrom(locB))
        edgeId = 'e_'+'_'.join(edgeToNames[edge])
        n = 2
        while edgeId in edgeIds:
            edgeId = 'e_'+'_'.join(edgeToNames[edge]) + '_{0}'.format(n)
            n += 1
        edgeIds.add(edgeId)
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': edgeId, 'd': d, 'style': edgeStyle[edge]})
    addRectToLayer(layer, corners, False)

def addCenterPath(layer, province):
    """Add a supply center marker to a given layer."""
    centerId = '{0}Center'.format(province.abbreviation)
    d = CENTER_PATH.format(strFrom(province.center))
    centerColor = (colorForName(province.abbreviation) if OVERRIDE_CHECK_MODE else '#000000')
    opacity = (1 if OVERRIDE_CHECK_MODE else 0.5)
    xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': centerId, 'd': d, 'style': 'display:inline;fill:none;stroke:{0};stroke-width:{1};stroke-opacity:{2};enable-background:new'.format(centerColor, THIN, opacity)})

def addCenterLayer(root, layerName, visible, provinces):
    """Add a layer of supply centers for all the given provinces."""
    newLayer = addLayer(root, layerName, visible)
    for province in provinces:
        addCenterPath(newLayer, province)

def getNeighbours(province, provinces):
    """Get the neighbours of the specified province. Note that this assumes provinces don't contain holes - e.g. provinces completely enclosed by other provinces."""
    neighbours = []
    region = province.edges
    for edge in zip(region, region[1:] + [region[0]]):
        reverse = (edge[1], edge[0])
        for other in provinces:
            if other.abbreviation == province.abbreviation:
                continue
            otherRegion = other.edges
            otherEdges = zip(otherRegion, otherRegion[1:] + [otherRegion[0]])
            if edge in otherEdges or reverse in otherEdges:
                neighbours.append(other)
    return neighbours

def createGraphFile(fileName, provinces):
    """Create a *.go file for the variant."""
    templateFile = open('variant.go.template')
    template = Template(templateFile.read())
    
    nationLength = max(map(len, START_UNITS.keys()))
    nation_declarations = []
    for nation in START_UNITS.keys():
        nation_declarations.append(u'\t{{0:<{}}} godip.Nation = "{{1}}"'.format(nationLength).format(toCamelCase(nation), nation))
    nation_list = u'var Nations = []godip.Nation{{{}}}'.format(u', '.join(map(toCamelCase, START_UNITS.keys())))
    
    scCount = int(round((len([province for province in provinces if province.flags.supplyCenter]) + 1) / 2.0))
    
    unitsStrs = []
    for nation, units in START_UNITS.items():
        for unitType in units.keys():
            for region in units[unitType]:
                if len([province.abbreviation for province in provinces if province.name == tuple(region.split(' '))]) == 0:
                    raise Exception('Could not find region {} when setting starting units.'.format(region))
                abbr = [province.abbreviation for province in provinces if province.name == tuple(region.split(' '))][0]
                unitsStrs.append(u'\t\t"{}": godip.Unit{{godip.{}, {}}},'.format(abbr, unitType, toCamelCase(nation)))
                
    supplyCenterStrs = []
    for nation, units in START_UNITS.items():
        for unitType in units.keys():
            for region in units[unitType]:
                province = [province for province in provinces if province.name == tuple(region.split(' '))][0]
                if province.flags.supplyCenter:
                    supplyCenterStrs.append(u'\t\t"{}": {},'.format(province.abbreviation, toCamelCase(nation)))
    
    graphStrs = []
    flags = {}
    longNameStrs = []
    for province in provinces:
        if province.flags.impassable:
            continue
        if province.flags.sea:
            flags[province.abbreviation] = 'Sea'
        else:
            flag = 'Land'
            for neighbour in getNeighbours(province, provinces):
                if neighbour.flags.sea:
                    flag = 'Coast...'
            flags[province.abbreviation] = flag
    for province in provinces:
        graphStr = ''
        if province.flags.impassable:
            continue
        graphStr += '\t\t// {}\n'.format(' '.join(province.name))
        graphStr += '\t\tProv("{}").'.format(province.abbreviation)
        for neighbour in getNeighbours(province, provinces):
            if not neighbour.flags.impassable:
                if province.flags.sea or neighbour.flags.sea:
                    borderType = 'Sea'
                else:
                    # Assume coastal border if regions share a common sea neighbour
                    abbrsA = set([n.abbreviation for n in getNeighbours(province, provinces) if n.flags.sea])
                    abbrsB = set([n.abbreviation for n in getNeighbours(neighbour, provinces) if n.flags.sea])
                    if len(abbrsA.intersection(abbrsB)) > 0:
                        borderType = 'Coast...'
                    else:
                        borderType = 'Land'
                graphStr += 'Conn("{}", godip.{}).'.format(neighbour.abbreviation, borderType)
        graphStr += 'Flag(godip.{}).'.format(flags[province.abbreviation])
        if province.flags.supplyCenter:
            owner = 'godip.Neutral'
            for nation, units in START_UNITS.items():
                for regions in units.values():
                    if province.name in map(lambda name: tuple(name.split(' ')), regions):
                        owner = toCamelCase(nation)
            graphStr += u'SC({}).'.format(owner)
        graphStrs.append(graphStr)
        longNameStrs.append('\t"{}": "{}",'.format(province.abbreviation, ' '.join(province.name)))
    
    parameters = {
            'variant': VARIANT,
            'variant_lower': toLowerAlphaNumeric(VARIANT),
            'variant_camel': toCamelCase(VARIANT),
            'nation_declarations': '\n'.join(nation_declarations),
            'nation_list': nation_list,
            'sc_count': scCount,
            'start_year': str(START_YEAR),
            'units': '\n'.join(unitsStrs),
            'supply_centers': '\n'.join(supplyCenterStrs),
            'graph': '\n'.join(graphStrs),
            'long_names': '\n'.join(longNameStrs)
            }
    
    output = template.substitute(parameters)
    templateFile.close()
    f = open(fileName, 'w')
    f.write(output.encode('utf8'))
    f.close()

def createDebuggingMap(root, regions, edgeToDMap, corners):
    """Create a map for use when looking for problems with the input map. Draw all the regions in
    random colors so that it's easy to see any adjacent regions which have been merged by the script."""
    debugNames = {}
    for i, region in enumerate(regions):
        debugNames['region{0}'.format(i)] = region
    addLayerWithRegions(root, debugNames, edgeToDMap, 'background', None, True, corners, True)
    xml.etree.ElementTree.ElementTree(root).write(toLowerAlphaNumeric(VARIANT) + 'debug.svg')

# Load data from the svg file.
root = xml.etree.ElementTree.parse(MAP).getroot()
corners = getCorners(root)
junctions = getJunctions(root, corners)
oldEdgeToDMap = getEdges(root)
edgeToDMap = findDesiredEdges(junctions, oldEdgeToDMap)

# A path for the supply center symbol. This should be formatted to include the absolute start location.
CENTER_PATH = getCentersPath(root)

namesLayer = getLayer(root, 'names')
regions = makeRegions(junctions, edgeToDMap.keys(), corners)

# At this point we have enough information to create a map that's useful for investigating errors in the input map.
createDebuggingMap(root, regions, edgeToDMap, corners)

regionFullNames = guessRegionFullNames(regions, namesLayer)
abbreviations = inventAbbreviations(regionFullNames.keys())

# Here we assume that each region has a marker on exactly one of the four layers.
supplyCenters = getCentersWithin(root, 'supply-centers')
regionCenters = getCentersWithin(root, 'province-centers')
seaCenters = getCentersWithin(root, 'sea')
impassableCenters = getCentersWithin(root, 'impassable')
allMarkers = dict(supplyCenters)
allMarkers.update(regionCenters)
allMarkers.update(seaCenters)
passableCenters = dict(allMarkers)
allMarkers.update(impassableCenters)

originalIdToRegion = matchRegionMarkerToRegion(regions, allMarkers)
originalIdToAbbr = makeIdToAbbrMap(originalIdToRegion, regionFullNames, abbreviations)

supplyCenters = replaceOriginalIds(supplyCenters, originalIdToAbbr, abbreviations)
regionCenters = replaceOriginalIds(regionCenters, originalIdToAbbr, abbreviations)
seaCenters = replaceOriginalIds(seaCenters, originalIdToAbbr, abbreviations)
impassableCenters = replaceOriginalIds(impassableCenters, originalIdToAbbr, abbreviations)

print('Centers by type: SCs {}, Other Land {}, Sea {}, Impassable {}'.format(len(supplyCenters), len(regionCenters), len(seaCenters), len(impassableCenters)))

# Put all the data into the DTO.
provinces = []
for name in regionFullNames.keys():
    abbr = abbreviations[name]
    oldId = reverseMapLookup(originalIdToAbbr, abbr)
    center = allMarkers[oldId]
    flags = Flags(abbr in supplyCenters, abbr in regionCenters, abbr in seaCenters, abbr in impassableCenters)
    provinces.append(Province(name, abbr, center, flags, regionFullNames[name]))
# Swap any details that the user has manually overridden.
performOverrides(provinces)

scLayer = getLayer(root, 'supply-centers')
pcLayer = getLayer(root, 'province-centers')
seaLayer = getLayer(root, 'sea')
# Remove all layers from the root, ready to construct the output svg.
removeAllLayers(root)
addPatterns(root)
# Add all the layers to the output.
backgroundRegionNames = {}
for province in provinces:
    if not province.flags.sea:
        backgroundRegionNames[province.abbreviation + '_background'] = province.edges
addLayerWithRegions(root, backgroundRegionNames, edgeToDMap, 'background', LAND_COLOR, True, corners)
edgeStyle, edgeToNames = getEdgeStyle(edgeToDMap.keys(), provinces)
#addShadowsToBackground(root, edgeStyle, edgeToNames)
passableNames = {}
for province in provinces:
    if not province.flags.impassable:
        passableNames[province.abbreviation] = province.edges
addLayerWithRegions(root, passableNames, edgeToDMap, 'provinces', (None if OVERRIDE_CHECK_MODE else '#000000'), False)
addCenterLayer(root, 'supply-centers', True, [province for province in provinces if province.flags.supplyCenter])
addCenterLayer(root, 'province-centers', False, [province for province in provinces if province.flags.province or province.flags.sea])
addLayer(root, 'highlights', True)
addForeground(root, edgeToDMap, edgeStyle, edgeToNames, corners)
addNamesLayer(root, namesLayer, abbreviations, passableNames.keys())
addLayer(root, 'units', True)
addLayer(root, 'orders', True)
# Create the output svg file.
xml.etree.ElementTree.ElementTree(root).write(toLowerAlphaNumeric(VARIANT) + 'map.svg')

# Create the output go file.
createGraphFile(toLowerAlphaNumeric(VARIANT) + '.go', provinces)
