#!/usr/bin/env python3

import xml.etree.ElementTree
import re
import math

### Data to be gathered for the variant. ###

# The name of the variant
VARIANT = 'Hundred'
# The starting units
START_UNITS = {'Burgundy': {'Army': ['dij', 'lux', 'fla'], 'Fleet': ['hol']},
               'England': {'Army': ['cal', 'guy', 'nmd'], 'Fleet': ['lon', 'dev']},
               'France': {'Army': ['dau', 'orl', 'par', 'tou', 'pro'], 'Fleet': []}}
# The nations in the variant
NATIONS = START_UNITS.keys()
# The first year of the game
START_YEAR = 1425

### Constants ###

INK = '{http://www.inkscape.org/namespaces/inkscape}'
SVG = '{http://www.w3.org/2000/svg}'
MAP = 'input.svg'
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

root = xml.etree.ElementTree.parse(MAP).getroot()

def getLayer(root, label):
    return root.find('{}g[@{}label="{}"]'.format(SVG, INK, label))

def removeAllLayers(root):
    for layer in root.findall('{}g'.format(SVG)):
        root.remove(layer)

def addLayer(root, name, visible):
    display = 'inline' if visible else 'none'
    'svg:style="display:none"'
    attrs = {'id': name, '{}groupmode'.format(INK): 'layer', '{}label'.format(INK): name, 'style': 'display:' + display}
    xml.etree.ElementTree.SubElement(root, '{}g'.format(SVG), attrs)
    return getLayer(root, name)

def locFrom(locString):
    loc = locString.split(',')
    return (float(loc[0]), float(loc[1]))

def addLocs(locA, locB):
    return (locA[0] + locB[0], locA[1] + locB[1])

def subLocs(locA, locB):
    return (locA[0] - locB[0], locA[1] - locB[1])

def strFrom(loc):
    return ','.join(map(lambda x: str(x), loc))

def findDist(locA, locB):
    dx = locA[0] - locB[0]
    dy = locA[1] - locB[1]
    return math.sqrt(dx*dx + dy*dy)

def getCorners(root):
    viewBox = map(float, root.get('viewBox').split(' '))
    return ((viewBox[0], viewBox[1]), (viewBox[2], viewBox[1]), (viewBox[2], viewBox[3]), (viewBox[0], viewBox[3]))

def getCentersWithin(root, layerLabel):
    centers = {}
    layer = getLayer(root, layerLabel)
    if layer != None:
        for path in layer.findall('{}path'.format(SVG)):
            name = path.get('id').split('Center')[0]
            loc = locFrom(path.get('d').split(' ')[1])
            centers[name] = loc
    return centers

def getJunctions(root, corners):
    junctions = []
    layer = getLayer(root, 'points')
    for circle in layer.findall('{}circle'.format(SVG)):
        x = float(circle.get('cx'))
        y = float(circle.get('cy'))
        if abs(x - corners[0][0]) < GUTTER:
            x = corners[0][0]
        elif abs(x - corners[2][0]) < GUTTER:
            x = corners[2][0]
        if abs(y - corners[0][1]) < GUTTER:
            y = corners[0][1]
        elif abs(y - corners[2][1]) < GUTTER:
            y = corners[2][1]
        junctions.append((x, y))
    return junctions

def getToolParts(d):
    return re.findall(r'[LM] [0-9\.\-e,]+', d)

def makeToolPart(tool, loc):
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
            if bit in ['M', 'm', 'c', 'v', 'H', 'h', 'l', 'L']:
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

def reverseD(d):
    toolParts = re.findall(r'[a-zA-Z][ 0-9\.,\-]*', d)
    ret = ''
    for toolPart in reversed(toolParts):
        tool, instruction = toolPart.split(' ', 1)
        if tool == 'c':
            ret += tool
            for triple in reversed(re.findall(r'[0-9\.,\-]* [0-9\.,\-]* [0-9\.,\-]* *', instruction)):
                locA, locB, locC = map(locFrom, triple.strip().split(' '))
                ret += ' ' + strFrom(subLocs(locB, locC))
                ret += ' ' + strFrom(subLocs(locA, locC))
                ret += ' ' + strFrom(subLocs((0,0), locC))
        elif tool == 'v':
            ret += tool
            ret += ' ' + str(-float(instruction))
        else:
            raise Exception('Unsuported tool: ' + tool)
    return ret

def findClosestJunction(junctions, point):
    closestJunction = None
    closestDist = 1000000
    for junction in junctions:
        dist = findDist(junction, point)
        if dist < closestDist:
            closestJunction = junction
            closestDist = dist
    return closestJunction

def findDesiredEdges(junctions, edges):
    desiredEdges = {}
    for edge, d in edges.items():
        start = findClosestJunction(junctions, edge[0])
        end = findClosestJunction(junctions, edge[1])
        d0 = [makeToolPart('M', start)] + getToolParts(d[0])[1:-1] + [makeToolPart('L', end)]
        d1 = [makeToolPart('M', end)] + getToolParts(d[1])[1:-1] + [makeToolPart('L', start)]
        desiredEdges[(start, end)] = (' '.join(d0), ' '.join(d1))
    return desiredEdges

def updateEdges(root, oldEdges, edges):
    # This doesn't really work like I hoped.
    layer = getLayer(root, 'edges')
    i = 0
    for edge in layer.findall('{}path'.format(SVG)):
        d = edge.get('d')
        # Assume that the edge starts with an M or m and doesn't ends with a pair of coordinates.
        if not d.lower().startswith('m ') and ',' not in d.split(' ')[-1]:
            raise Exception('Unsupported update to edge: ' + d)
        bits = d.split(' ')
        s = subLocs(addLocs(locFrom(bits[1]), edges[i][0]), oldEdges[i][0])
        e = edges[i][1]
        newBits = bits[:1] + [strFrom(s)] + bits[2:] + ['M', strFrom(e)]
        edge.set('d', ' '.join(newBits))
        i += 1

# See http://bryceboe.com/2006/10/23/line-segment-intersection-algorithm/
def ccw(A,B,C):
    return (C[1]-A[1]) * (B[0]-A[0]) > (B[1]-A[1]) * (C[0]-A[0])
def intersect(first, second):
    A,B = first
    C,D = second
    return ccw(A,C,D) != ccw(B,C,D) and ccw(A,B,C) != ccw(A,B,D)

def singleIntersection(segment, edges):
    count = 0
    for edge in edges:
        if intersect(segment, edge):
            count += 1
            if count > 1:
                return False
    return count == 1

def findAdjacencyGraph(allCenters, edges):
    centerList = allCenters.items()
    adjacentPairs = []
    for i, a in enumerate(centerList):
        for b in centerList[i+1:]:
            adjEdge = (a[1], b[1])
            if singleIntersection(adjEdge, edges):
                adjacentPairs.append((a[0], b[0]))
    return adjacentPairs

def vectorAngle(start, end):
    return math.atan2(end[1] - start[1], end[0] - start[0])

def findLinks(junction, edges):
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
    # Estimate the middle of the region by the middle of its bounding box
    xs = map(lambda r: r[0], region)
    ys = map(lambda r: r[1], region)
    return ((max(xs) + min(xs)) / 2.0, (max(ys) + min(ys)) / 2.0)

def findClosestPair(pointsA, pointsB):
    bestDist = 1000000
    bestPair = None
    for a in pointsA:
        for b in pointsB:
            d = findDist(a, b)
            if d < bestDist:
                bestDist = d
                bestPair = (a, b)
    return bestPair

def guessRegionNames(regions, allCenters):
    # Use a greedy algorithm to name regions based on how close their middle is to a center.
    middleToRegion = {}
    for region in regions:
        middleToRegion[middleOfRegion(region)] = region
    centerToName = {}
    for name, center in allCenters.items():
        centerToName[center] = name
    middlePoints = list(middleToRegion.keys())
    centerPoints = list(allCenters.values())
    regionNames = {}
    while len(middlePoints) > 0:
        m, c = findClosestPair(middlePoints, centerPoints)
        regionNames[centerToName[c]] = middleToRegion[m]
        middlePoints.remove(m)
        centerPoints.remove(c)
    return regionNames
    

def addLayerWithEdges(root, edges):
    layer = getLayer(root, 'names')
    for edge in edges:
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'d': 'M {} c 0,0 0,0 {}'.format(strFrom(edge[0]), strFrom(subLocs(edge[1], edge[0]))), 'style': 'opacity:1;vector-effect:none;fill:none;fill-opacity:1;fill-rule:evenodd;stroke:#000100;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1'})

def calculateCurvePoints(lastLoc, loc, nextLoc):
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

def addRectToLayer(layer, corners, fill):
    fillStyle = 'fill:{};fill-opacity:1;'.format(LAND_COLOR) if fill else 'fill:none;'
    style= fillStyle + 'display:inline;stroke:#000000;stroke-width:{};stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1'.format(THICK)
    width = '{}'.format(corners[2][0] - corners[0][0])
    height = '{}'.format(corners[2][1] - corners[0][1])
    xml.etree.ElementTree.SubElement(layer, '{}rect'.format(SVG), {'id': 'bg_rect', 'style': style, 'width': width, 'height': height, 'x': '0', 'y': '0'})

def addLayerWithRegions(root, regionNames, edgeToDMap, layerName, color, visible, corners = None):
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
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': name, 'd': d, 'style': 'fill:{0};fill-opacity:1;vector-effect:none;fill-rule:evenodd'.format(color)})

def getEdgeThickness(edges, regionNames, sea, impassable):
    edgeThickness = {}
    edgeToNames = {}
    for edge in edges:
        touches = set()
        names = []
        for name, region in regionNames.items():
            if edge[0] in region and edge[1] in region and abs(region.index(edge[0]) - region.index(edge[1])) in [1, len(region)-1]:
                if name in sea:
                    touches.add('sea')
                else:
                    touches.add('land')
                if name in impassable:
                    touches.add('impassible')
                names.append(name)
        edgeToNames[edge] = names
        edgeThickness[edge] = THICK if len(touches) > 1 else THIN
    return edgeThickness, edgeToNames

def addForeground(root, edgeToDMap, edgeThickness, edgeToNames, corners):
    layer = addLayer(root, 'foreground', True)
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
        xml.etree.ElementTree.SubElement(layer, '{}path'.format(SVG), {'id': 'e_'+'_'.join(edgeToNames[edge]), 'd': d, 'style': 'fill:none;vector-effect:none;fill-rule:evenodd;stroke:#000100;stroke-width:{};stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1'.format(thickness)})
    addRectToLayer(layer, corners, False)

def getNeighbours(center, regionNames):
    region = regionNames[center]
    neighbours = []
    for edge in zip(region, region[1:] + [region[0]]):
        reverse = (edge[1], edge[0])
        selected = []
        if edge in edgeToNames.keys() and center in edgeToNames[edge]:
            selected = list(edgeToNames[edge])
        elif reverse in edgeToNames.keys() and center in edgeToNames[reverse]:
            selected = list(edgeToNames[reverse])
        if len(selected) == 2:
            selected.remove(center)
            neighbours.append(selected[0])
    return neighbours

def createGraphFile(fileName, passableCenters, supplyCenters, seaCenters, regionNames, edgeToNames):
    f = open(fileName, 'w')
    f.write('package {}\n'.format(VARIANT.lower()))
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
    f.write('\nvar {}Variant = common.Variant{{\n'.format(VARIANT))
    f.write('\tName:        "{}",\n'.format(VARIANT))
    f.write('\tGraph:       func() dip.Graph {{ return {}Graph() }},\n'.format(VARIANT))
    f.write('\tStart:       {}Start,\n'.format(VARIANT))
    f.write('\tBlank:       {}Blank,\n'.format(VARIANT))
    f.write("""	Phase:       classical.Phase,
	ParseOrders: orders.ParseAll,
	ParseOrder:  orders.Parse,
	OrderTypes:  orders.OrderTypes(),
	Nations:     Nations,
	PhaseTypes:  cla.PhaseTypes,
	Seasons:     cla.Seasons,
	UnitTypes:   cla.UnitTypes,""")
    f.write('\n\tSoloSupplyCenters: {},\n'.format(int(round(len(supplyCenters) / 2.0))))
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
""".format(VARIANT.lower()))
    f.write("""
func {0}Blank(phase dip.Phase) *state.State {{
	return state.New({0}Graph(), phase, classical.BackupRule)
}}

func {0}Start() (result *state.State, err error) {{
	startPhase := classical.Phase({1}, cla.Spring, cla.Movement)
	result = state.New({0}Graph(), startPhase, classical.BackupRule)
	if err = result.SetUnits(map[dip.Province]dip.Unit{{
""".format(VARIANT, START_YEAR))
    for nation, units in START_UNITS.items():
        for unitType in units.keys():
            for region in units[unitType]:
                f.write('\t\t"{}": dip.Unit{{cla.{}, {}}},\n'.format(region, unitType, nation))
    f.write("""	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[dip.Province]dip.Nation{
""")
    for nation, units in START_UNITS.items():
        for unitType in units.keys():
            for region in units[unitType]:
                if region in supplyCenters:
                    f.write('\t\t"{}": {},\n'.format(region, nation))
    f.write("""	}})
	return
}}

func {0}Graph() *graph.Graph {{
	return graph.New().
""".format(VARIANT, START_YEAR))
    flags = {}
    for center in passableCenters:
        if center in seaCenters:
            flags[center] = 'Sea'
        else:
            flag = 'Land'
            for neighbour in getNeighbours(center, regionNames):
                if neighbour in seaCenters:
                    flag = 'Coast...'
            flags[center] = flag
    for center in passableCenters:
        f.write('\t\t// {}\n'.format(center))
        f.write('\t\tProv("{}").'.format(center))
        region = regionNames[center]
        for edge in zip(region, region[1:] + [region[0]]):
            reverse = (edge[1], edge[0])
            if edge in edgeToNames.keys() and center in edgeToNames[edge]:
                selected = edgeToNames[edge]
            elif reverse in edgeToNames.keys() and center in edgeToNames[reverse]:
                selected = edgeToNames[reverse]
            for other in selected:
                if other != center and other in passableCenters:
                    if center in seaCenters and other in seaCenters:
                        borderType = 'Sea'
                    elif center not in seaCenters and other not in seaCenters:
                        borderType = 'Land'
                    else:
                        borderType = 'Coast...'
                    f.write('Conn("{}", cla.{}).'.format(other, borderType))
        f.write('Flag(cla.{}).'.format(flags[center]))
        if center in supplyCenters:
            owner = 'cla.Neutral'
            for nation, units in START_UNITS.items():
                for regions in units.values():
                    if center in regions:
                        owner = nation
            f.write('SC({}).'.format(owner))
        f.write('\n')
    f.write("""		Done()
}
""")
    f.close()

# Load data from the svg file.
corners = getCorners(root)
supplyCenters = getCentersWithin(root, 'supply-centers')
regionCenters = getCentersWithin(root, 'province-centers')
seaCenters = getCentersWithin(root, 'sea')
impassableCenters = getCentersWithin(root, 'impassable')
allCenters = dict(supplyCenters)
allCenters.update(regionCenters)
passableCenters = dict(allCenters)
allCenters.update(impassableCenters)
junctions = getJunctions(root, corners)
oldEdgeToDMap = getEdges(root)

edgeToDMap = findDesiredEdges(junctions, oldEdgeToDMap)

#updateEdges(root, oldEdges, edges)

regions = makeRegions(junctions, edgeToDMap.keys(), corners)
regionNames = guessRegionNames(regions, allCenters)

#adjacencyGraph = findAdjacencyGraph(allCenters, edges)

scLayer = getLayer(root, 'supply-centers')
pcLayer = getLayer(root, 'province-centers')
removeAllLayers(root)
#addLayerWithEdges(root, edges)
backgroundRegionNames = {}
for regionName in regionNames.keys():
    if regionName in seaCenters.keys():
        backgroundRegionNames[regionName + '_background'] = regionNames[regionName]
addLayerWithRegions(root, backgroundRegionNames, edgeToDMap, 'background', SEA_COLOR, True, corners)
edgeThickness, edgeToNames = getEdgeThickness(edgeToDMap.keys(), regionNames, seaCenters.keys(), impassableCenters.keys())
passableNames = {}
for name, region in regionNames.items():
    if name in passableCenters:
        passableNames[name] = region
addLayerWithRegions(root, passableNames, edgeToDMap, 'provinces', '#000000', False)
root.append(scLayer)
root.append(pcLayer)
addLayer(root, 'highlights', True)
addForeground(root, edgeToDMap, edgeThickness, edgeToNames, corners)
addLayer(root, 'names', True)
addLayer(root, 'units', True)
addLayer(root, 'orders', True)

# Create an output svg file.
xml.etree.ElementTree.ElementTree(root).write(VARIANT.lower() + 'map.svg')

createGraphFile(VARIANT.lower() + '.go', passableCenters, supplyCenters, seaCenters, regionNames, edgeToNames)
