package run

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/iosRealRun-cli/iOSRealRun-cli/internal/device"
)

func GeoDistance(p1 map[string]float64, p2 map[string]float64) (distance float64) {
	lat1, lng1 := p1["lat"], p1["lng"]
	lat2, lng2 := p2["lat"], p2["lng"]
	radius := 6371000.0 // unit: m
	rad := func(d float64) float64 {
		return d * math.Pi / 180.0
	}
	radLat1, lngLat1 := rad(lat1), rad(lng1)
	radLat2, lngLat2 := rad(lat2), rad(lng2)
	dradLat := radLat1 - radLat2
	dradLng := lngLat1 - lngLat2
	a := math.Pow(math.Sin(dradLat/2), 2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(dradLng/2), 2)
	distance = radius * 2 * math.Asin(math.Sqrt(a))
	return
}

// return an coefficient between 0 and 1 that makes the speed change smoothly
// then i == start or i == end, the coefficient is 0
func Smooth(start int, end int, i int) float64 {
	ii := float64(i-start) / float64(end-start)
	return math.Pow(math.Sin(ii*math.Pi), 2)
}

func randLoc(loc []map[string]float64, d float64, n int) []map[string]float64 {
	// deep copy
	result := make([]map[string]float64, 0)
	for _, i := range loc {
		result = append(result, MapCopy(i))
	}

	center := map[string]float64{"lat": 0, "lng": 0}
	for _, i := range result {
		center["lat"] += i["lat"]
		center["lng"] += i["lng"]
	}
	center["lat"] /= float64(len(result))
	center["lng"] /= float64(len(result))

	for i := 0; i < n; i++ {
		start := int(float64(len(result)) / float64(n) * float64(i))
		end := int(float64(len(result)) / float64(n) * float64(i+1))
		offset := (2*rand.Float64() - 1) * d
		for j := start; j < end; j++ {
			distance := math.Sqrt(math.Pow(result[j]["lat"]-center["lat"], 2) + math.Pow(result[j]["lng"]-center["lng"], 2))
			if distance == 0 {
				continue
			}
			result[j]["lat"] += (result[j]["lat"] - center["lat"]) / distance * offset * Smooth(start, end, j)
			result[j]["lng"] += (result[j]["lng"] - center["lng"]) / distance * offset * Smooth(start, end, j)
		}
	}
	return result
}

func fixLockT(loc []map[string]float64, v float64, dt float64) []map[string]float64 {
	fixedLoc := make([]map[string]float64, 0)
	t := float64(0)
	T := make([]float64, 0)
	T = append(T, GeoDistance(loc[1], loc[0])/v)
	a := MapCopy(loc[0])
	b := MapCopy(loc[1])
	j := 0
	for t < T[0] {
		xa := a["lat"] + float64(j)*(b["lat"]-a["lat"])/float64(IntMax(1, int(T[0]/dt)))
		xb := a["lng"] + float64(j)*(b["lng"]-a["lng"])/float64(IntMax(1, int(T[0]/dt)))
		fixedLoc = append(fixedLoc, map[string]float64{"lat": xa, "lng": xb})
		j += 1
		t += dt
	}
	for i := 1; i < len(loc); i++ {
		T = append(T, GeoDistance(loc[(i+1)%len(loc)], loc[i])/v+T[len(T)-1])
		a = MapCopy(loc[i])
		b = MapCopy(loc[(i+1)%len(loc)])
		j = 0
		for t < T[i] {
			xa := a["lat"] + float64(j)*(b["lat"]-a["lat"])/float64(IntMax(1, int((T[i]-T[i-1])/dt)))
			xb := a["lng"] + float64(j)*(b["lng"]-a["lng"])/float64(IntMax(1, int((T[i]-T[i-1])/dt)))
			fixedLoc = append(fixedLoc, map[string]float64{"lat": xa, "lng": xb})
			j += 1
			t += dt
		}
	}
	return fixedLoc
}

// run a circle
func run1(loc []map[string]float64, v float64, dt float64) {
	fixedLoc := fixLockT(loc, v, dt)
	nList := []int{5, 6, 7, 8, 9}
	n := nList[rand.Intn(len(nList))]
	fixedLoc = randLoc(fixedLoc, 0.000025, n)
	clock := time.Now()
	for _, i := range fixedLoc {
		device.SetLoc(i)
		for time.Since(clock).Seconds() < dt {
			// pass
		}
		clock = time.Now()
	}
}

// keep running;
// loc: the location list;
// v: the speed;
// d: the max deviation of pace, Unit: s;
func Run(loc []map[string]float64, v float64, d int) {
	for {
		vRand := 1000 / (1000/v - (2*rand.Float64()-1)*float64(d))
		run1(loc, vRand, 0.2)
		fmt.Println("跑完一圈了")
	}
}
