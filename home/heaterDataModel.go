package home

import "encoding/json"

/*
LIMIT max amount of history data calues
*/
const LIMIT = 10000

/*
HData is set of home data values, can be parsed from json
*/
type HData struct {
	TempInside  float32 `json:"TempInside, float"`
	TempOutside float32 `json:"TempOutside, float"`
	TempHeater  float32 `json:"TempHeater, float"`
	TempReverse float32 `json:"TempReverse, float"`

	PumpState   bool `json:"PumpState, bool"`
	HeaterState bool `json:"HeaterState, bool"`
	Timestamp   int  `json:"Timestamp, int"`
	Index       int  `json:"index, int"`
}

/*
HistoryData is storage for recent states
*/
type HistoryData []*HData

/*
Len of HistoryData
*/
func (q HistoryData) Len() int { return len(q) }

/*
func (q historyData) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}


func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
*/

/*
Push HomeData to HistoryData
*/
func (q *HistoryData) Push(x interface{}) {
	n := len(*q)
	item := x.(*HData)
	item.Index = n
	*q = append(*q, item)
	if n > LIMIT {
		old := *q
		item := old[n-1]
		item.Index = -1 // for safety
		*q = old[0 : n-1]
		item = nil
	}
}

/*
Pop HomeData from HistoryData
*/
func (q *HistoryData) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

/*
ToJSON returns serialized hash
*/
func (q *HistoryData) ToJSON() (d []byte, err error) {
	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	return b, nil
}

/*
// update modifies the priority and value of an Item in the queue.
func (q *historyData) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
*/
