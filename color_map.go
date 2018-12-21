package pushshift

type ColorMap uint
const (
	// Perceptually Uniform Sequential Color Maps

	ColorMapViridis = ColorMap(iota)
	ColorMapPlasma
	ColorMapInferno
	ColorMapMagma

	// Sequential Color Maps

	ColorMapGreys
	ColorMapPurples
	ColorMapBlues
	ColorMapGreens
	ColorMapOranges
	ColorMapReds
	ColorMapYlOrBr
	ColorMapYlOrRd
	ColorMapOrRd
	ColorMapPuRd
	ColorMapRdPu
	ColorMapBuPu
	ColorMapGnBu
	ColorMapPuBu
	ColorMapYlGnBu
	ColorMapPuBuGn
	ColorMapBuGn
	ColorMapYlGn
	ColorMapBinary
	ColorMapGistYarg
	ColorMapGistGray
	ColorMapGray
	ColorMapBone
	ColorMapPink
	ColorMapSpring
	ColorMapSummer
	ColorMapAutumn
	ColorMapWinter
	ColorMapCool
	ColorMapWistia
	ColorMapHot
	ColorMapAfmHot
	ColorMapGistHeat
	ColorMapCopper

	// Diverging Color Maps

	ColorMapPiYG
	ColorMapPRGn
	ColorMapBrBG
	ColorMapPuOr
	ColorMapRdGy
	ColorMapRdBu
	ColorMapRdYlBu
	ColorMapRdYlGn
	ColorMapSpectral
	ColorMapCoolWarm
	ColorMapBwr
	ColorMapSeismic

	// Qualitative Color Maps

	ColorMapPastel1
	ColorMapPastel2
	ColorMapPaired
	ColorMapAccent
	ColorMapDark2
	ColorMapSet1
	ColorMapSet2
	ColorMapSet3
	ColorMapTab10
	ColorMapTab20
	ColorMapTab20b
	ColorMapTab20c

	// Miscellaneous Color Maps

	ColorMapFlag
	ColorMapPrism
	ColorMapOcean
	ColorMapGistEarth
	ColorMapTerrain
	ColorMapGistStern
	ColorMapGnuPlot
	ColorMapGnuPlot2
	ColorMapCMRmap
	ColorMapCubeHelix
	ColorMapBrg
	ColorMapHsv
	ColorMapGistRainbow
	ColorMapRainbow
	ColorMapJet
	ColorMapNiPySpectral
	ColorMapGistNcar
)

var colorMapCodes = map[ColorMap]string {
	ColorMapViridis: "viridis",
	ColorMapPlasma: "plasma",
	ColorMapInferno: "inferno",
	ColorMapMagma: "magma",
	ColorMapGreys: "Greys",
	ColorMapPurples: "Purples",
	ColorMapBlues: "Blues",
	ColorMapGreens: "Greens",
	ColorMapOranges: "Oranges",
	ColorMapReds: "Reds",
	ColorMapYlOrBr: "YlOrBr",
	ColorMapYlOrRd: "YlOrRd",
	ColorMapOrRd: "OrRd",
	ColorMapPuRd: "PuRd",
	ColorMapRdPu: "RdPu",
	ColorMapBuPu: "BuPu",
	ColorMapGnBu: "GnBu",
	ColorMapPuBu: "PuBu",
	ColorMapYlGnBu: "YlGnBu",
	ColorMapPuBuGn: "PuBuGn",
	ColorMapBuGn: "BuGn",
	ColorMapYlGn: "YlGn",
	ColorMapBinary: "binary",
	ColorMapGistYarg: "gist_yarg",
	ColorMapGistGray: "gist_gray",
	ColorMapGray: "gray",
	ColorMapBone: "bone",
	ColorMapPink: "pink",
	ColorMapSpring: "spring",
	ColorMapSummer: "summer",
	ColorMapAutumn: "autumn",
	ColorMapWinter: "winter",
	ColorMapCool: "cool",
	ColorMapWistia: "Wistia",
	ColorMapHot: "hot",
	ColorMapAfmHot: "afmhot",
	ColorMapGistHeat: "gist_heat",
	ColorMapCopper: "copper",
	ColorMapPiYG: "PiYG",
	ColorMapPRGn: "PRGn",
	ColorMapBrBG: "BrBG",
	ColorMapPuOr: "PuOr",
	ColorMapRdGy: "RdGy",
	ColorMapRdBu: "RdBu",
	ColorMapRdYlBu: "RdYlBu",
	ColorMapRdYlGn: "RdYlGn",
	ColorMapSpectral: "Spectral",
	ColorMapCoolWarm: "coolwarm",
	ColorMapBwr: "bwr",
	ColorMapSeismic: "seismic",
	ColorMapPastel1: "Pastel1",
	ColorMapPastel2: "Pastel2",
	ColorMapPaired: "Paired",
	ColorMapAccent: "Accent",
	ColorMapDark2: "Dark2",
	ColorMapSet1: "Set1",
	ColorMapSet2: "Set2",
	ColorMapSet3: "Set3",
	ColorMapTab10: "tab10",
	ColorMapTab20: "tab20",
	ColorMapTab20b: "tab20b",
	ColorMapTab20c: "tab20c",
	ColorMapFlag: "flag",
	ColorMapPrism: "prism",
	ColorMapOcean: "ocean",
	ColorMapGistEarth: "gist_earth",
	ColorMapTerrain: "terrain",
	ColorMapGistStern: "gist_stern",
	ColorMapGnuPlot: "gnuplot",
	ColorMapGnuPlot2: "gnuplot2",
	ColorMapCMRmap: "CMRmap",
	ColorMapCubeHelix: "cubehelix",
	ColorMapBrg: "brg",
	ColorMapHsv: "hsv",
	ColorMapGistRainbow: "gist_rainbow",
	ColorMapRainbow: "rainbow",
	ColorMapJet: "jet",
	ColorMapNiPySpectral: "nipy_spectral",
	ColorMapGistNcar: "gist_ncar",
}

func (c ColorMap) String() string {
	return colorMapCodes[c]
}
