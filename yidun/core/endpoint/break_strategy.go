package endpoint

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	cache "github.com/patrickmn/go-cache"
)

type BreakStrategy interface {
	RequestSuccess(productCode string, regionCode string, domain string)

	RequestFail(productCode string, regionCode string, domain string)

	AttemptAccess(productCode string, regionCode string, domain string) bool
}

type BreakerStatusEnum int32

const (
	OPEN BreakerStatusEnum = iota
	HALF_OPEN
	CLOSED
)

type FixedWindowBreakStrategy struct {
	clock                                 Clock
	statWindowMillis                      int64
	circuitBreakerFailThresholdPercentage int32
	circuitBreakerRequestVolumeThreshold  int32
	circuitBreakerSleepWindowMillis       int64
	metrics                               FixedWindowMetrics
	breakerMap                            SafeMap
}

type Clock interface {
	millis() int64
}

type DefaultClock struct{}

func (c *DefaultClock) millis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type Breaker struct {
	b_status      int32
	lastBreakTime *int64
}

// 安全的map
type SafeMap struct {
	mu  sync.Mutex
	Map map[string]*Breaker
}

// 添加值
func (sm *SafeMap) Set(key string, value *Breaker) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.Map[key] = value
}

// 获取值
func (sm *SafeMap) Get(key string) (*Breaker, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.Map[key]
	return value, ok
}

func (b *Breaker) status() *int32 {
	return &b.b_status
}

func (b *Breaker) lastBreakTimeMillis() *int64 {

	return b.lastBreakTime
}

type StrategyConfig struct {
	statWindowMillis                      int64
	circuitBreakerFailThresholdPercentage int32
	circuitBreakerRequestVolumeThreshold  int32
	circuitBreakerSleepWindowMillis       int64
}

var (
	DefaultStatWindowMillis                      = 10 * 60 * 1000
	DefaultCircuitBreakerFailThresholdPercentage = 40
	DefaultCircuitBreakerRequestVolumeThreshold  = 50
	DefaultCircuitBreakerSleepWindowMillis       = 20 * 60 * 1000
)

func (s *StrategyConfig) SetStatWindowMillis(value int64) {
	s.statWindowMillis = value
}

func (s *StrategyConfig) SetCircuitBreakerFailThresholdPercentage(value int32) {
	s.circuitBreakerFailThresholdPercentage = value
}

func (s *StrategyConfig) SetCircuitBreakerRequestVolumeThreshold(value int32) {
	s.circuitBreakerRequestVolumeThreshold = value
}

func (s *StrategyConfig) SetCircuitBreakerSleepWindowMillis(value int64) {
	s.circuitBreakerSleepWindowMillis = value
}

var DefaultStrategyConfig = &StrategyConfig{
	statWindowMillis:                      int64(DefaultStatWindowMillis),
	circuitBreakerFailThresholdPercentage: int32(DefaultCircuitBreakerFailThresholdPercentage),
	circuitBreakerRequestVolumeThreshold:  int32(DefaultCircuitBreakerRequestVolumeThreshold),
	circuitBreakerSleepWindowMillis:       int64(DefaultCircuitBreakerSleepWindowMillis),
}

func CreateDefaultFixWindowBreakStrategy() *FixedWindowBreakStrategy {
	return NewFixedWindowBreakStrategy(DefaultStrategyConfig)
}

func NewFixedWindowBreakStrategy(config *StrategyConfig) *FixedWindowBreakStrategy {
	return NewFixedWindowBreakStrategyWithClock(&DefaultClock{}, config)
}

func NewFixedWindowBreakStrategyWithClock(clock Clock, config *StrategyConfig) *FixedWindowBreakStrategy {
	if config == nil {
		config = &StrategyConfig{}
	}

	if config.statWindowMillis <= 0 {
		config.statWindowMillis = int64(DefaultStatWindowMillis)
	}

	if config.circuitBreakerFailThresholdPercentage <= 0 || config.circuitBreakerFailThresholdPercentage > 100 {
		config.circuitBreakerFailThresholdPercentage = int32(DefaultCircuitBreakerFailThresholdPercentage)
	}

	if config.circuitBreakerRequestVolumeThreshold <= 0 {
		config.circuitBreakerRequestVolumeThreshold = int32(DefaultCircuitBreakerRequestVolumeThreshold)
	}

	if config.circuitBreakerSleepWindowMillis <= 0 {
		config.circuitBreakerSleepWindowMillis = int64(DefaultCircuitBreakerSleepWindowMillis)
	}

	return &FixedWindowBreakStrategy{
		clock:                                 clock,
		statWindowMillis:                      config.statWindowMillis,
		circuitBreakerFailThresholdPercentage: config.circuitBreakerFailThresholdPercentage,
		circuitBreakerRequestVolumeThreshold:  config.circuitBreakerRequestVolumeThreshold,
		circuitBreakerSleepWindowMillis:       config.circuitBreakerSleepWindowMillis,
		metrics:                               *NewFixedWindowMetrics(config.statWindowMillis),
		breakerMap:                            SafeMap{Map: make(map[string]*Breaker)},
	}
}

func (s *FixedWindowBreakStrategy) RequestSuccess(productCode string, regionCode string, domain string) {
	currentTimeMillis := s.clock.millis()

	breaker := s.getBreaker(productCode, regionCode, domain)
	if atomic.CompareAndSwapInt32(breaker.status(), int32(HALF_OPEN), int32(CLOSED)) {
		// 熔断器从半开状态到关闭状态：域名恢复为可用
		// 重置相关统计数据
		s.metrics.reset(productCode, regionCode, domain, currentTimeMillis)
	}

	s.metrics.requestSuccess(productCode, regionCode, domain, currentTimeMillis)
}

func (s *FixedWindowBreakStrategy) RequestFail(productCode string, regionCode string, domain string) {
	currentTimeMillis := s.clock.millis()

	data := s.metrics.requestFail(productCode, regionCode, domain, currentTimeMillis)

	breaker := s.getBreaker(productCode, regionCode, domain)
	if atomic.CompareAndSwapInt32(breaker.status(), int32(HALF_OPEN), int32(OPEN)) {
		// 熔断器从半开状态到全开状态：域名继续被熔断
		// 设置新的熔断时间
		atomic.StoreInt64(breaker.lastBreakTimeMillis(), currentTimeMillis)
		return
	}
	// 根据统计信息确定是否需要熔断
	if s.breakTriggered(data.failCount, data.successCount) &&
		atomic.CompareAndSwapInt32(breaker.status(), int32(CLOSED), int32(OPEN)) {
		// 针对第一个触发熔断的请求，设置新的熔断时间
		atomic.StoreInt64(breaker.lastBreakTimeMillis(), currentTimeMillis)
	}
}

func (s *FixedWindowBreakStrategy) AttemptAccess(productCode string, regionCode string, domain string) bool {
	breaker := s.getBreaker(productCode, regionCode, domain)
	lastBreakTimeMillis := atomic.LoadInt64(breaker.lastBreakTimeMillis())
	status := BreakerStatusEnum(atomic.LoadInt32(breaker.status()))

	// 熔断器处于关闭状态：允许访问
	if status == CLOSED {
		return true
	}

	// 还处于熔断导致的禁用时间范围内：不允许访问
	if s.isInSleepWindow(lastBreakTimeMillis) {
		return false
	}

	// 已超过禁用时间范围：允许第1个请求访问
	return atomic.CompareAndSwapInt32(breaker.status(), int32(OPEN), int32(HALF_OPEN))
}

func (s *FixedWindowBreakStrategy) isInSleepWindow(lastBreakTimeMillis int64) bool {
	return s.clock.millis() <= lastBreakTimeMillis+int64(s.circuitBreakerSleepWindowMillis)
}

func (s *FixedWindowBreakStrategy) breakTriggered(failCount int32, successCount int32) bool {
	total := failCount + successCount
	return total >= s.circuitBreakerRequestVolumeThreshold && failCount*100/total >= s.circuitBreakerFailThresholdPercentage
}

func (s *FixedWindowBreakStrategy) getBreaker(productCode string, regionCode string, domain string) *Breaker {
	key := genBreakerKey(productCode, regionCode, domain)
	breaker, ok := s.breakerMap.Get(key)
	if !ok {
		initNum := int64(-1)
		breaker = &Breaker{
			b_status:      int32(CLOSED),
			lastBreakTime: &initNum,
		}
		s.breakerMap.Set(key, breaker)
		return breaker
	} else {
		return breaker
	}
}

func genBreakerKey(productCode string, regionCode string, domain string) string {
	return fmt.Sprintf("%s:%s:%s", productCode, regionCode, domain)
}

type Metrics interface {
	reset(productCode string, regionCode string, domain string, currentTimeMillis int64)
	requestSuccess(productCode string, regionCode string, domain string, currentTimeMillis int64) MetricsData
	requestFail(productCode string, regionCode string, domain string, currentTimeMillis int64) MetricsData
}

type FixedWindowMetrics struct {
	statMap          cache.Cache
	statWindowMillis int64
}

func NewFixedWindowMetrics(statWindowMillis int64) *FixedWindowMetrics {
	return &FixedWindowMetrics{
		// 使用开源的缓存库，内置过期数据清理
		statMap:          *cache.New(5*time.Minute, 10*time.Minute),
		statWindowMillis: statWindowMillis,
	}
}

func (fwm *FixedWindowMetrics) reset(productCode string, regionCode string, domain string, currentTimeMillis int64) {
	data := fwm.getData(productCode, regionCode, domain, currentTimeMillis)
	data.successCount.Store(0)
	data.failCount.Store(0)
}

func (fwm *FixedWindowMetrics) requestSuccess(productCode string, regionCode string, domain string, currentTimeMillis int64) MetricsData {
	data := fwm.getData(productCode, regionCode, domain, currentTimeMillis)

	successCount := data.successCount.Inc()
	failCount := data.failCount.Load()

	return MetricsData{successCount, failCount}
}

func (fwm *FixedWindowMetrics) requestFail(productCode string, regionCode string, domain string, currentTimeMillis int64) MetricsData {
	data := fwm.getData(productCode, regionCode, domain, currentTimeMillis)

	successCount := data.successCount.Load()
	failCount := data.failCount.Inc()

	return MetricsData{successCount, failCount}
}

func (fwm *FixedWindowMetrics) getData(productCode string, regionCode string, domain string, timeMillis int64) *Data {
	key := fwm.genStatKey(productCode, regionCode, domain, timeMillis)
	dataInterface, ok := fwm.statMap.Get(key)
	if ok {
		return dataInterface.(*Data)
	} else {
		data := &Data{
			successCount: AtomicInt{value: 0},
			failCount:    AtomicInt{value: 0},
		}
		fwm.statMap.SetDefault(key, data)
		return data
	}
}

func (fwm *FixedWindowMetrics) genStatKey(productCode string, regionCode string, domain string, timeMillis int64) string {
	statWindow := timeMillis - timeMillis%fwm.statWindowMillis
	return fmt.Sprintf("%s:%s:%s:%d", productCode, regionCode, domain, statWindow)
}

type Data struct {
	successCount AtomicInt
	failCount    AtomicInt
}

type MetricsData struct {
	successCount int32
	failCount    int32
}

type AtomicInt struct {
	value int32
	lock  sync.Mutex
}

func (ai *AtomicInt) Inc() int32 {
	ai.lock.Lock()
	defer ai.lock.Unlock()
	ai.value++
	return ai.value
}

func (ai *AtomicInt) Load() int32 {
	ai.lock.Lock()
	defer ai.lock.Unlock()
	return ai.value
}

func (ai *AtomicInt) Store(val int32) {
	atomic.StoreInt32(&ai.value, val)
}
