def geodistance(p1, p2):
    lat1, lng1 = p1["lat"], p1["lng"]
    lat2, lng2 = p2["lat"], p2["lng"]
    from math import radians, cos, sin, asin, sqrt
    lng1, lat1, lng2, lat2 = map(radians, [float(lng1), float(lat1), float(lng2), float(lat2)]) # 经纬度转换成弧度
    dlon=lng2-lng1
    dlat=lat2-lat1
    a=sin(dlat/2)**2 + cos(lat1) * cos(lat2) * sin(dlon/2)**2
    distance=2*asin(sqrt(a))*6371*1000 # 地球平均半径，6371km
    return distance

def randLoc(loc: list, d=0.000025, n=7):
    import random
    import time
    import math
    # deepcopy loc
    result = []
    for i in loc:
        result.append(i.copy())

    center = {"lat": 0, "lng": 0}
    for i in result:
        center["lat"] += i["lat"]
        center["lng"] += i["lng"]
    center["lat"] /= len(result)
    center["lng"] /= len(result)
    random.seed(time.time())
    for i in range(n):
        for j in range(int(i*len(result)/n), int((i+1)*len(result)/n)):
            offset = (2*random.random()-1) * d
            distance = math.sqrt(
                (result[j]["lat"]-center["lat"])**2 + (result[j]["lng"]-center["lng"])**2
            )
            result[j]["lat"] +=  (result[j]["lat"]-center["lat"])/distance*offset
            result[j]["lng"] +=  (result[j]["lng"]-center["lng"])/distance*offset
    for j in range(int(i*len(result)/n), len(result)):
        offset = (2*random.random()-1) * d
        distance = math.sqrt(
            (result[j]["lat"]-center["lat"])**2 + (result[j]["lng"]-center["lng"])**2
        )
        result[j]["lat"] +=  (result[j]["lat"]-center["lat"])/distance*offset
        result[j]["lng"] +=  (result[j]["lng"]-center["lng"])/distance*offset
    return result

def fixLockT(loc: list, v, dt):
    fixedLoc = []
    t = 0
    T = []
    T.append(geodistance(loc[1],loc[0])/v)
    a = loc[0]
    b = loc[1]
    j = 0
    while t < T[0]:
        xa = a["lat"] + j*(b["lat"]-a["lat"])/(max(1, int(T[0]/dt)))
        xb = a["lng"] + j*(b["lng"]-a["lng"])/(max(1, int(T[0]/dt)))
        fixedLoc.append({"lat": xa, "lng": xb})
        j += 1
        t += dt
    for i in range(1, len(loc)):
        T.append(geodistance(loc[(i+1)%len(loc)],loc[i])/v + T[-1])
        a = loc[i]
        b = loc[(i+1)%len(loc)]
        j = 0
        while t < T[i]:
            xa = a["lat"] + j*(b["lat"]-a["lat"])/(max(1, int((T[i]-T[i-1])/dt)))
            xb = a["lng"] + j*(b["lng"]-a["lng"])/(max(1, int((T[i]-T[i-1])/dt)))
            fixedLoc.append({"lat": xa, "lng": xb})
            j += 1
            t += dt
    return fixedLoc

def run1(loc: list, v, dt=0.2):
    import time
    import tools.utils as utils
    fixedLoc = fixLockT(loc, v, dt)
    clock = time.time()
    for i in fixedLoc:
        utils.setLoc(i)
        while time.time()-clock < dt:
            pass
        clock = time.time()

def run(loc: list, v, d=15):
    import tools.utils as utils
    import random
    import time
    random.seed(time.time())
    while True:
        nList = (5, 7, 9)
        n = nList[random.randint(0, len(nList)-1)]
        newLoc = randLoc(loc, n=n)  # a path will be divided into n parts for random route
        vRand = 1000/(1000/v-(2*random.random()-1)*d)
        run1(newLoc, vRand)
        print("跑完一圈了")