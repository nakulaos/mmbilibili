package metric

import "backend/app/common/constant"

type JETCacheMetric struct {
	Name        string
	Description string
	cHit        CacheCallback
	cMiss       CacheCallback
	cLocalHit   CacheCallback
	cLocalMiss  CacheCallback
	cRemoteHit  CacheCallback
	cRemoteMiss CacheCallback
	cQuery      CacheCallback
	cQueryFail  CacheCallback
}

func (J *JETCacheMetric) IncrHit() {
	J.cHit()
}

func (J *JETCacheMetric) IncrMiss() {
	J.cMiss()
}

func (J *JETCacheMetric) IncrLocalHit() {
	J.cLocalHit()
}

func (J *JETCacheMetric) IncrLocalMiss() {
	J.cLocalMiss()
}

func (J *JETCacheMetric) IncrRemoteHit() {
	J.cRemoteHit()
}

func (J *JETCacheMetric) IncrRemoteMiss() {
	J.cRemoteMiss()
}

func (J *JETCacheMetric) IncrQuery() {
	J.cQuery()
}

func (J *JETCacheMetric) IncrQueryFail(err error) {
	J.cQueryFail()
}

func NewJETCacheMetric(name, description string) *JETCacheMetric {

	return &JETCacheMetric{
		Name:        name,
		Description: description,
		cHit:        newGauge(CacheHit, name+"."+constant.ALL, description),
		cMiss:       newGauge(CacheMiss, name+"."+constant.ALL, description),
		cLocalHit:   newGauge(CacheHit, name+"."+constant.Local, description),
		cLocalMiss:  newGauge(CacheMiss, name+"."+constant.Local, description),
		cRemoteHit:  newGauge(CacheHit, name+"."+constant.Remote, description),
		cRemoteMiss: newGauge(CacheMiss, name+"."+constant.Remote, description),
		cQuery:      newGauge(CacheHit, name+"."+constant.Query, description),
		cQueryFail:  newGauge(CacheMiss, name+"."+constant.Query, description),
	}
}
