	package cache

	import (
		"sync"
		"time"
	)

	type CacheEntry struct {
		createdAt time.Time
		val 	  []byte  
	}

	type Cache struct {
		data map[string]CacheEntry
		mu	 sync.RWMutex
	}

	func NewCache(interval time.Duration) *Cache {

		cache := &Cache{
			data:	make(map[string]CacheEntry),
		}
		go cache.reapLoop(interval)
		return cache
	}

	func (c *Cache) Add(key string, val []byte) {

		//DONT DOA ADD VALUES TO A MAP LIKE THIS!
		// c.data[key].val = val
		// c.data[key].createdAt = time.Now()

		//REMEMBER MAPS ARE NOT SAFE FOR CONCURRENT USE
		//REMEMBER TO LOCK YOUR MAPS WITH A MUTEX
		c.mu.Lock()
		defer c.mu.Unlock()
		
		c.data[key] = CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
		return
	}

	func (c *Cache) Get(key string) ([]byte, bool) {	
		//Remember when using read intensive functions to always use RLock
		//This is because many goroutines can safely read from the map at the same time
		//However only 1 writer can access an RWMutex at a time
		c.mu.RLock()
		defer c.mu.RUnlock()
		
		entry, ok := c.data[key]  
		
		if ok {
			return entry.val, true
		} else {
			return nil, false
		}
	}

	//What is the function of reapLoop?
	// - It is to remove any entries made at any time older than the interval
	// - This is the how we implement the cache in the sense that we ant to store data that
	//has been received for a certain interval
	//How does the reapLoop function work?
	// 1) It starts a ticker that "ticks" every interval.
	// 2) In an infinite loop, it waits for each tick.
	// 3) When a tick occurs, it locks the mutex.
	// 4) It then iterates through all entries in the cache.
	// 5) If an entry is older than the interval, it's deleted.
	// 6) Finally, it unlocks
	func (c *Cache) reapLoop(interval time.Duration) {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			c.mu.Lock()
			for key, entry := range c.data {
				if time.Since(entry.createdAt) > interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}