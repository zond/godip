#!/usr/bin/env python2

import xml.etree.ElementTree
import re
import math
import itertools
import collections
import random

### Data to be gathered for the variant. ###

# The name of the variant
VARIANT = 'Youngstown Redux'
# The starting units
#START_UNITS = {'NATO': {'Army': ['New York', 'Los Angeles', 'Paris'], 'Fleet': ['London', 'Istanbul', 'Australia']},
#               # Fleet should be Leningrad South Coast
#               'USSR': {'Army': ['Moscow', 'Shanghai', 'Vladivostok'], 'Fleet': ['Leningrad', 'Albania', 'Havana']}}
START_UNITS = {'Russia': {'Army': ['Moscow', 'Omsk', 'Warsaw'], 'Fleet': ['Sevastopol', 'St. Petersburg', 'Vladivostok']},
               'China': {'Army': ['Peking', 'Guangzhou', 'Wuhan'], 'Fleet': ['Shanghai']},
               'Japan': {'Army': ['Kyoto'], 'Fleet': ['Tokyo', 'Osaka', 'Sapporo']},
               'India': {'Army': ['Delhi', 'Calcutta'], 'Fleet': ['Bombay', 'Madras']},
               'Turkey': {'Army': ['Constantinople', 'Baghdad', 'Mecca'], 'Fleet': ['Ankara']},
               'France': {'Army': ['Paris', 'Marseilles'], 'Fleet': ['Brest', 'Saigon']},
               'Britain': {'Fleet': ['London', 'Liverpool', 'Edinburgh', 'Aden', 'Singapore']},
               'Italy': {'Army': ['Rome', 'Milan'], 'Fleet': ['Naples', 'Mogadishu']},
               'Germany': {'Army': ['Berlin', 'Munich', 'Cologne'], 'Fleet': ['Tsingtao', 'Kiel']},
               'Austria': {'Army': ['Vienna', 'Budapest', 'Trieste'], 'Fleet': ['Sarajevo']}}
# The nations in the variant
NATIONS = START_UNITS.keys()
# The first year of the game
START_YEAR = 1901
# Abbreviations that should be used (rather than letting the script try to guess an abbreviation).
#ABBREVIATIONS = {'Iran': 'irn', 'Iraq': 'irq', 'Japan': 'jap', 'Arabia': 'ara', 'India': 'ind', 'Sea of Japan': 'soj'}
#ABBREVIATIONS = {'North Atlantic': 'nat', 'Norwegian Sea': 'nrg', 'St Petersburg': 'stp', 'North Africa': 'naf', 'Liverpool': 'lvp', 'North Sea': 'nth', 'Norway': 'nwy', 'Livonia': 'lvn', 'Gulf of Bothnia': 'bot', 'Gulf of Lyon': 'gol', 'Tyrolia': 'tyr', 'Tyrrhenian Sea': 'tys'}
ABBREVIATIONS = {'Box A bcd': 'bxa', 'Box B ace': 'bxb', 'Box C abfgh': 'bxc', 'Box D aef': 'bxd', 'Box E bdf': 'bxe', 'Box F cdegh': 'bxf', 'Box G cfh': 'bxg', 'Box H cfg': 'bxh', 'Java Sea': 'jvs', 'Arabian Sea': 'ars', 'Persian Gulf': 'psg'}
# Overrides to swap centers. This only needs to contain something if the greedy algorithm fails.
#CENTER_OVERRIDES = [('Caribbean Sea', 'Havana'), ('West Atlantic', 'Brazil'), ('Black Sea', 'Istanbul'), ('Indian Ocean', 'Arabian Sea'), ('Caribbean Sea', 'Colombia'), ('Caribbean Sea', 'Venezuala'), ('Finland', 'Leningrad')]
#CENTER_OVERRIDES = [('Sweden', 'Gulf of Bothnia'), ('Mid Atlantic', 'Portugal')]
CENTER_OVERRIDES = [('Kamchatka', 'North Pacific Ocean'), ('Awdal', 'Gulf of Aden'), ('Hebei', 'Tsingtao'), ('Red Sea', 'Mecca'), ('Galicia', 'Vienna'), ('Awdal', 'Mogadishu')]
# Overrides to swap region names. This only needs to contain something if the greedy algorithm fails.
#REGION_OVERRIDES = [('West Atlantic', 'Brazil'), ('South China Sea', 'Saigon'), ('Black Sea', 'Istanbul')]
#REGION_OVERRIDES = [('Finland', 'Gulf of Bothnia'), ('Mid Atlantic', 'Portugal')]
REGION_OVERRIDES = [('Red Sea', 'Mecca')]#, ('Galicia', 'Vienna'), ('Awdal', 'Mogadishu')]
# Whether to highlight the region abbreviation in bold or not.
BOLD_ABBREVIATIONS = True

### Constants ###

INK = '{http://www.inkscape.org/namespaces/inkscape}'
SVG = '{http://www.w3.org/2000/svg}'
MAP = 'youngstownredux_input.svg'
# Any junctions within GUTTER pixels from the edge of the page will be moved to the edge.
GUTTER = 5
# How curvy the edges should be made
CURVE_WEIGHT = 0.5
# The background colour of sea regions
SEA_COLOR = '#c6efed'
# The background colour of the land
LAND_COLOR = '#ffffff'
# The thickness of thick lines
THICK = 2.225
# The thickness of thin lines
THIN = 1
# A path for the supply center symbol. This should be formatted to include the absolute start location.
CENTER_PATH = 'm {0} c 4.9401,2.5533 2.6167,9.4913 -3.1784,9.4913 -1.1467,0 -2.1967,-0.5071 -3.2688,-1.5786 -4.3606,-4.3582 0.88,-10.79 6.4472,-7.9127 z m -6.6684,-0.6788 c -5.0525,3.972 -2.4142,11.2502 4.078,11.2502 4.0998,0 6.1809,-2.0911 6.1809,-6.2105 0,-5.4556 -5.9882,-8.3973 -10.2589,-5.0397 z m 8.3437,-3.3432 c 6.6105,3.3706 6.5264,13.4768 -0.1388,16.7017 -8.795,4.2552 -17.5176,-5.4815 -12.1036,-13.511 2.8103,-4.1678 7.7839,-5.4641 12.2426,-3.1907 z m -8.8915,-1.3825 c -8.2929,3.6996 -8.0443,15.936 0.3953,19.4604 10.0229,4.1855 19.0983,-7.6593 12.5008,-16.3155 -2.6922,-3.5324 -8.7281,-5.0043 -12.8961,-3.1449 z'
#CENTER_PATH = 'm {0} c 4.88873,-2.52807 2.58951,-9.39762 -3.14536,-9.39762 -1.13481,0 -2.17382,0.50204 -3.23479,1.56302 -4.31527,4.31523 0.87083,10.68356 6.38015,7.8346 z m -6.59908,0.67219 c -4.99986,-3.93285 -2.38906,-11.13924 4.0356,-11.13924 4.05721,0 6.11664,2.07045 6.11664,6.14924 0,5.40176 -5.92591,8.31446 -10.15224,4.99 z m 8.257,3.31016 c 6.5417,-3.33731 6.45845,-13.34383 -0.13743,-16.5369 -8.7036,-4.21333 -17.33545,5.4274 -11.97776,13.37771 2.78105,4.12675 7.70298,5.4102 12.11528,3.15919 z m -8.79902,1.36888 c -8.20665,-3.66309 -7.96067,-15.77877 0.39112,-19.2684 9.91863,-4.14426 18.8997,7.5837 12.37085,16.15454 -2.66427,3.49755 -8.63734,4.95493 -12.76197,3.11386 z'
#CENTER_PATH = 'm {0} c 1.30948,0.67717 0.69362,2.51722 -0.84251,2.51722 -0.30396,0 -0.58227,-0.13447 -0.86646,-0.41866 -1.15587,-1.15587 0.23326,-2.86167 1.70897,-2.09856 z m -1.76761,-0.18005 c -1.33925,1.05344 -0.63993,2.98373 1.08096,2.98373 1.08676,0 1.63839,-0.55459 1.63839,-1.64712 0,-1.4469 -1.5873,-2.22709 -2.71935,-1.33661 z m 2.2117,-0.88665 c 1.75224,0.89393 1.72994,3.57424 -0.0368,4.42953 -2.33133,1.12857 -4.64343,-1.45377 -3.20833,-3.58332 0.74492,-1.10537 2.06329,-1.44916 3.24516,-0.84621 z m -2.35689,-0.36666 c -2.1982,0.98118 -2.13232,4.22645 0.10477,5.161181 2.65678,1.110069 5.06242,-2.031351 3.31362,-4.327111 -0.71364,-0.93684 -2.31357,-1.32722 -3.41839,-0.83407 z'
#CENTER_PATH = 'm {0} c 1.30948,0.67717 0.69362,2.51722 -0.84251,2.51722 -0.30396,0 -0.58227,-0.13447 -0.86646,-0.41866 -1.15587,-1.15587 0.23326,-2.86167 1.70897,-2.09856 z m -1.76761,-0.18005 c -1.33925,1.05344 -0.63993,2.98373 1.08096,2.98373 1.08676,0 1.63839,-0.55459 1.63839,-1.64712 0,-1.4469 -1.5873,-2.22709 -2.71935,-1.33661 z m 2.2117,-0.88665 c 1.75224,0.89393 1.72994,3.57424 -0.0368,4.42953 -2.33133,1.12857 -4.64343,-1.45377 -3.20833,-3.58332 0.74492,-1.10537 2.06329,-1.44916 3.24516,-0.84621 z m -2.35689,-0.36666 c -2.1982,0.98118 -2.13232,4.22645 0.10477,5.16118 2.65678,1.11007 5.06242,-2.03135 3.31362,-4.32711 -0.71364,-0.93684 -2.31357,-1.32722 -3.41839,-0.83407 z'
#d="m {0} c 4.94922,2.55935 2.62155,9.5139 -3.18428,9.5139 -1.14885,0 -2.20071,-0.50825 -3.27481,-1.58236 -4.36866,-4.36863 0.88161,-10.81575 6.45909,-7.93154 z m -6.68073,-0.68051 c -5.06172,3.98151 -2.41862,11.27707 4.08553,11.27707 4.10742,0 6.19233,-2.09607 6.19233,-6.22533 0,-5.4686 -5.99924,-8.41734 -10.27786,-5.05174 z m 8.35917,-3.35112 c 6.62264,3.37861 6.53836,13.50894 -0.13913,16.74152 -8.8113,4.26546 -17.54995,-5.49456 -12.12597,-13.54324 2.81546,-4.17781 7.79829,-5.47714 12.26519,-3.19828 z m -8.9079,-1.38582 c -8.30819,3.70842 -8.05917,15.97401 0.39596,19.50682 10.04136,4.19554 19.13356,-7.67754 12.52392,-16.35443 -2.69723,-3.54082 -8.74421,-5.01624 -12.91988,-3.15239 z"

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

def getLayer(root, label):
    """Get the layer from root with the given Inkscape label."""
    return root.find('{}g[@{}label="{}"]'.format(SVG, INK, label))

def removeAllLayers(root):
    """Remove all layers from root."""
    for layer in root.findall('{}g'.format(SVG)):
        root.remove(layer)

def addLayer(root, name, visible):
    """Add a new layer to root with the given name and visibility."""
    display = 'inline' if visible else 'none'
    'svg:style="display:none"'
    attrs = {'id': name, '{}groupmode'.format(INK): 'layer', '{}label'.format(INK): name, 'style': 'display:' + display}
    xml.etree.ElementTree.SubElement(root, '{}g'.format(SVG), attrs)
    return getLayer(root, name)

def locFrom(locString):
    """Convert a coordinate string (e.g. from a path) into a coordinate pair."""
    loc = locString.split(',')
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
        for bit in d.split(' '):
            # Check if this is a supported tool
            if bit in ['M', 'm', 'c', 'V', 'v', 'H', 'h', 'l', 'L']:
                tool = bit
                toolStart = loc
                if tool == 'c':
                    # Coordinates for the c tool come in threes
                    cIgnoreCount = 0
            elif re.match(r'[0-9\.,\-]', bit):
                # Apply the tool to the numbers
                if tool == 'M' or tool == 'L':
                    loc = locFrom(bit)
                elif tool == 'm' or tool == 'l':
                    loc = addLocs(loc, locFrom(bit))
                elif tool == 'c':
                    cIgnoreCount = (cIgnoreCount + 1) % 3
                    if cIgnoreCount == 0:
                        loc = addLocs(loc, locFrom(bit))
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
                raise Exception('Unsupported tool: ' + bit)
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
            directedEdges.remove((previousJunction, currentJunction))
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
        output[originalIdToAbbr[originalId]] = centers[originalId]
    return output

def makeLocsToNames(namesLayer):
    """Create a map from coordinates to names of provinces."""
    locsToNames = {}
    for text in namesLayer.findall('{}text'.format(SVG)):
        transform = text.get('transform')
        x, y = float(text.get('x')), float(text.get('y'))
        if transform != None:
            if re.match(r'^rotate\([^\(\)\,]*\)$', transform):
                angle = math.radians(float(transform.split('(')[1].split(')')[0]))
                x, y = x * math.cos(angle) - y * math.sin(angle), x * math.sin(angle) + y * math.cos(angle)
            else:
                print('Unsupported text transformation: ' + transform)
        loc = (x, y)
        name = []
        for tspan in text.findall('.//{}tspan'.format(SVG)):
            if tspan.text != None:
                name += re.split(r' +', tspan.text)
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
        abbr += name.replace('.','')[i]
    return abbr

def findTupleFromName(name, fullNamesTuples):
    """Find the name tuple that has the same letters (in the same order) as the given name."""
    name = name.replace(' ', '').replace('.', '').lower()
    for fullNameTuple in fullNamesTuples:
        if ''.join(fullNameTuple).lower() == name:
            return fullNameTuple
    raise Exception('Couldn\'t find tuple matching {0}'.format(name))

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

def firstLetterAbbr(nTuple):
    """Create an abbreviation from the first letter of the first three words."""
    return (nTuple[0][0]+nTuple[1][0]+nTuple[2][0]).lower()

def abbreviationsForNames_firstLetter(fullNamesTuples, abbrCount):
    """Try to create unique abbreviations for the given name tuples by considering the first letters of the first three words."""
    abbrMap = {}
    for nTuple in fullNamesTuples:
        if len(nTuple) >= 3:
            abbrCount[firstLetterAbbr(nTuple)] += 1
    for fullNameTuple, abbr in abbrMap.items():
        if abbrCount[abbr] > 1:
            del abbrMap[fullNameTuple]
    for nTuple in fullNamesTuples:
        if len(nTuple) >= 3:
            if nTuple not in abbrMap.keys() and abbrCount[firstLetterAbbr(nTuple)] == 1:
                abbrMap[nTuple] = firstLetterAbbr(nTuple)
    return abbrMap

def inventAbbreviations(fullNamesTuples):
    """Determine a suitable set of unique abbreviations for the given name tuples. Use any values from
    the user ABBREVIATIONS override list."""
    fixedAbbrs = {}
    abbrCount = collections.Counter()
    for name, abbr in map(lambda na: (na[0].replace(' ', '').replace('.', '').lower(), na[1]), ABBREVIATIONS.items()):
        fixedAbbrs[findTupleFromName(name, fullNamesTuples)] = abbr
        abbrCount[abbr] += 1
    # Start by taking any unique first three letters.
    remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
    fixedAbbrs.update(abbreviationsForNames(remainingNames, [(0, 1, 2)], abbrCount))
    # Add any abbreviations from the first three words (if the name has three words).
    remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
    fixedAbbrs.update(abbreviationsForNames_firstLetter(remainingNames, abbrCount))
    # combinations returns the indexes in lexigographical order, which is basically what we want.
    remainingNames = set(fullNamesTuples).difference(set(fixedAbbrs.keys()))
    maxLength = max(map(lambda nameTuple: len(''.join(nameTuple)), remainingNames))
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
        name = []
        for tspan in text.findall('.//{}tspan'.format(SVG)):
            if tspan.text != None:
                name += re.split(r' +', tspan.text)
        abbr = fullNameToAbbr[tuple(name)]
        if abbr not in passableCenterAbbrs:
            namesLayer.remove(text)
            continue
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
        if BOLD_ABBREVIATIONS and boldAbbr.replace(' ', '') != abbr:
            print 'Failed to automatically bold the abbreviation for {0} (got "{1}" rather than "{2}")'.format(name, boldAbbr, abbr)
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

def addPattern(root):
    """Add the pattern that's used for indicating provinces that can be clicked on by the player."""
    '''<pattern xmlns="http://www.w3.org/2000/svg" id="stripes" patternUnits="userSpaceOnUse" width="6" height="6" patternTransform="rotate(45 2 2)">
    <path xmlns="http://www.w3.org/2000/svg" d="M -1,2 l 6,0" stroke="#ff0000" stroke-width="1" id="path4424"></path>
    </pattern>'''
    xml.etree.ElementTree.SubElement(root, '{}pattern'.format(SVG), {'id': 'stripes', 'patternUnits': 'userSpaceOnUse', 'width': '6', 'height': '6', 'patternTransform': 'rotate(45 2 2)'})
    stripes = root.find('{0}pattern[@id="stripes"]'.format(SVG))
    xml.etree.ElementTree.SubElement(stripes, '{}path'.format(SVG), {'id': 'stripePath', 'd': 'M -1,2 l 6,0', 'stroke': '#ff0000', 'stroke-width': '1'})

def addRectToLayer(layer, corners, fill):
    """Draw a rectangle using the given corners and fill in the specified color."""
    fillStyle = 'fill:{};fill-opacity:1;'.format(LAND_COLOR) if fill else 'fill:none;'
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
        regionColor = color if color is not None else '#' + ''.join(random.sample('0123456789abcdef', 6))
        style = 'fill:{0};fill-opacity:1;vector-effect:none;fill-rule:evenodd'.format(regionColor)
        if edgesOnly:
            style += ';stroke:#000000;stroke-width:2px'
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': name, 'd': d, 'style': style})

def getEdgeThickness(edges, provinces):
    """Create a map giving the desired thickness of each edge.  Coastal edges are thick,
    as are edges involving impassable areas. All other edges are thin."""
    edgeThickness = {}
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
        edgeThickness[edge] = THICK if len(touches) > 1 else THIN
    return edgeThickness, edgeToNames

def addForeground(root, edgeToDMap, edgeThickness, edgeToNames, corners):
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
        thickness = edgeThickness[edge]
        edgeId = 'e_'+'_'.join(edgeToNames[edge])
        n = 2
        while edgeId in edgeIds:
            edgeId = 'e_'+'_'.join(edgeToNames[edge]) + '_{0}'.format(n)
            n += 1
        edgeIds.add(edgeId)
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': edgeId, 'd': d, 'style': 'fill:none;vector-effect:none;fill-rule:evenodd;stroke:#000100;stroke-width:{};stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1'.format(thickness)})
    addRectToLayer(layer, corners, False)

def addCenterPath(layer, province):
    """Add a supply center marker to a given layer."""
    centerId = '{0}Center'.format(province.abbreviation)
    d = CENTER_PATH.format(strFrom(province.center))
    xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': centerId, 'd': d, 'style': 'display:inline;fill:none;stroke:#9f9f9f;stroke-width:2.25273085;stroke-opacity:1;enable-background:new'})

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
    f = open(fileName, 'w')
    f.write('package {}\n'.format(VARIANT.lower().replace(' ', '')))
    f.write("""
import (
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/classical/orders"
	"github.com/zond/godip/variants/common"

	dip "github.com/zond/godip/common"
	cla "github.com/zond/godip/variants/classical/common"
)

const (
""")
    nationLength = max(map(len, START_UNITS.keys()))
    for nation in START_UNITS.keys():
        f.write('\t{{0:<{}}} dip.Nation = "{{0}}"\n'.format(nationLength).format(nation))
    f.write(')\n\nvar Nations = []dip.Nation{{{}}}\n'.format(', '.join(START_UNITS.keys())))
    f.write('\nvar {}Variant = common.Variant{{\n'.format(VARIANT.replace(' ', '')))
    f.write('\tName:        "{}",\n'.format(VARIANT))
    f.write('\tGraph:       func() dip.Graph {{ return {}Graph() }},\n'.format(VARIANT.replace(' ', '')))
    f.write('\tStart:       {}Start,\n'.format(VARIANT.replace(' ', '')))
    f.write('\tBlank:       {}Blank,\n'.format(VARIANT.replace(' ', '')))
    f.write("""	Phase:       classical.Phase,
	ParseOrders: orders.ParseAll,
	ParseOrder:  orders.Parse,
	OrderTypes:  orders.OrderTypes(),
	Nations:     Nations,
	PhaseTypes:  cla.PhaseTypes,
	Seasons:     cla.Seasons,
	UnitTypes:   cla.UnitTypes,""")
    scCount = int(round(len([province for province in provinces if province.flags.supplyCenter]) / 2.0))
    f.write('\n\tSoloSupplyCenters: {},\n'.format(scCount))
    f.write("""	SVGMap: func() ([]byte, error) {{
		return Asset("svg/{}map.svg")
	}},
	SVGVersion: "1",
	SVGUnits: map[dip.UnitType]func() ([]byte, error){{
		cla.Army: func() ([]byte, error) {{
			return classical.Asset("svg/army.svg")
		}},
		cla.Fleet: func() ([]byte, error) {{
			return classical.Asset("svg/fleet.svg")
		}},
	}},
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules: "",
}}
""".format(VARIANT.lower().replace(' ', '')))
    f.write("""
func {0}Blank(phase dip.Phase) *state.State {{
	return state.New({0}Graph(), phase, classical.BackupRule)
}}

func {0}Start() (result *state.State, err error) {{
	startPhase := classical.Phase({1}, cla.Spring, cla.Movement)
	result = state.New({0}Graph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[dip.Province]dip.Unit{{
""".format(VARIANT.replace(' ', ''), START_YEAR))
    for nation, units in START_UNITS.items():
        for unitType in units.keys():
            for region in units[unitType]:
                if len([province.abbreviation for province in provinces if province.name == tuple(region.split(' '))]) == 0:
                    raise Exception('Could not find region {} when setting starting units.'.format(region))
                abbr = [province.abbreviation for province in provinces if province.name == tuple(region.split(' '))][0]
                f.write('\t\t"{}": dip.Unit{{cla.{}, {}}},\n'.format(abbr, unitType, nation))
    f.write("""	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
""")
    for nation, units in START_UNITS.items():
        for unitType in units.keys():
            for region in units[unitType]:
                province = [province for province in provinces if province.name == tuple(region.split(' '))][0]
                if province.flags.supplyCenter:
                    f.write('\t\t"{}": {},\n'.format(province.abbreviation, nation))
    f.write("""	}})
	return
}}

func {0}Graph() *graph.Graph {{
	return graph.New().
""".format(VARIANT.replace(' ', ''), START_YEAR))
    flags = {}
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
        if province.flags.impassable:
            continue
        f.write('\t\t// {}\n'.format(' '.join(province.name)))
        f.write('\t\tProv("{}").'.format(province.abbreviation))
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
                f.write('Conn("{}", cla.{}).'.format(neighbour.abbreviation, borderType))
        f.write('Flag(cla.{}).'.format(flags[province.abbreviation]))
        if province.flags.supplyCenter:
            owner = 'cla.Neutral'
            for nation, units in START_UNITS.items():
                for regions in units.values():
                    if province.name in map(lambda name: tuple(name.split(' ')), regions):
                        owner = nation
            f.write('SC({}).'.format(owner))
        f.write('\n')
    f.write("""		Done()
}
""")
    f.close()

def createDebuggingMap(root, regions, edgeToDMap, corners):
    """Create a map for use when looking for problems with the input map. Draw all the regions in
    random colors so that it's easy to see any adjacent regions which have been merged by the script."""
    debugNames = {}
    for i, region in enumerate(regions):
        debugNames['region{0}'.format(i)] = region
    addLayerWithRegions(root, debugNames, edgeToDMap, 'background', None, True, corners, True)
    xml.etree.ElementTree.ElementTree(root).write(VARIANT.lower().replace(' ', '') + 'debug.svg')

# Load data from the svg file.
root = xml.etree.ElementTree.parse(MAP).getroot()
corners = getCorners(root)
junctions = getJunctions(root, corners)
oldEdgeToDMap = getEdges(root)
edgeToDMap = findDesiredEdges(junctions, oldEdgeToDMap)

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
# Add the pattern layer for when provinces are selectable.
addPattern(root)
# Add all the layers to the output.
backgroundRegionNames = {}
for province in provinces:
    if province.flags.sea:
        backgroundRegionNames[province.abbreviation + '_background'] = province.edges
addLayerWithRegions(root, backgroundRegionNames, edgeToDMap, 'background', SEA_COLOR, True, corners)
edgeThickness, edgeToNames = getEdgeThickness(edgeToDMap.keys(), provinces)
passableNames = {}
for province in provinces:
    if not province.flags.impassable:
        passableNames[province.abbreviation] = province.edges
addLayerWithRegions(root, passableNames, edgeToDMap, 'provinces', '#000000', False)
addCenterLayer(root, 'supply-centers', True, [province for province in provinces if province.flags.supplyCenter])
addCenterLayer(root, 'province-centers', False, [province for province in provinces if province.flags.province or province.flags.sea])
addLayer(root, 'highlights', True)
addForeground(root, edgeToDMap, edgeThickness, edgeToNames, corners)
addNamesLayer(root, namesLayer, abbreviations, passableNames.keys())
addLayer(root, 'units', True)
addLayer(root, 'orders', True)
# Create the output svg file.
xml.etree.ElementTree.ElementTree(root).write(VARIANT.lower().replace(' ', '') + 'map.svg')

# Create the output go file.
createGraphFile(VARIANT.lower().replace(' ', '') + '.go', provinces)
