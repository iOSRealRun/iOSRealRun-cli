def geodistance(p1, p2):
    lat1, lng1 = p1["lat"], p1["lng"]
    lat2, lng2 = p2["lat"], p2["lng"]
    from math import radians, cos, sin, asin, sqrt
    #lng1,lat1,lng2,lat2 = (120.12802999999997,30.28708,115.86572000000001,28.7427)
    lng1, lat1, lng2, lat2 = map(radians, [float(lng1), float(lat1), float(lng2), float(lat2)]) # 经纬度转换成弧度
    dlon=lng2-lng1
    dlat=lat2-lat1
    a=sin(dlat/2)**2 + cos(lat1) * cos(lat2) * sin(dlon/2)**2
    distance=2*asin(sqrt(a))*6371*1000 # 地球平均半径，6371km
    distance=round(distance,3)
    return distance

def randLoc(loc: list, d=0.0000005):
    import random
    random.seed()
    result = loc.copy()
    for i in result:
        i["lat"] += (2*random.random()-1) * d
        i["lng"] += (2*random.random()-1) * d
    return result

def fixLockT(loc, v, dt):
    fixedLoc = []
    T = []
    T.append(geodistance(loc[(1)],loc[0])/v)
    for i in range(1, len(loc)):
        T.append(geodistance(loc[(i+1)%len(loc)],loc[i])/v + T[-1])
    T.append(0)
    return fixedLoc, T

def run1(loc, v, dt=0.2):
    import time
    import tools.utils as utils
    fixedLoc, T = fixLockT(loc, v, dt)
    t = 0
    for i in range(len(loc)):
        a = loc[i]
        b = loc[(i+1)%len(loc)]
        j = 0
        while t < T[i]:
            xa = a["lat"] + j*(b["lat"]-a["lat"])/(max(1, int((T[i]-T[i-1])/dt)))
            xb = a["lng"] + j*(b["lng"]-a["lng"])/(max(1, int((T[i]-T[i-1])/dt)))
            fixedLoc.append({"lat": xa, "lng": xb})
            j += 1
            t += dt
    clock = time.time()
    for i in fixedLoc:
        utils.setLoc(i)
        while time.time()-clock < dt:
            pass
        clock = time.time()

def run(loc, v):
    import tools.utils as utils
    newLoc = randLoc(loc)
    while True:
        run1(newLoc, v)