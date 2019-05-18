package gigasecond

import "time"

const nanosecond time.Duration = 1
const second time.Duration = nanosecond * 1000000000
const gigasecond time.Duration = second * 1000000000

func AddGigasecond(t time.Time) time.Time {
    return t.Add(gigasecond)
}
