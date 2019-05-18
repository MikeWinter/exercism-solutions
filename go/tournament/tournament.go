package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	WIN = 1 << iota
	DRAW
	LOSS
)

var points = map[int]int{
	WIN:  3,
	DRAW: 1,
	LOSS: 0,
}

type stats struct {
	played, won, drawn, lost int
}

type result struct {
	team   string
	result int
}

type performance struct {
	team string
	stats
}

func (p stats) points() int {
	return p.won*points[WIN] + p.drawn*points[DRAW] + p.lost*points[LOSS]
}

func Tally(reader io.Reader, writer io.Writer) (err error) {
	results := make(chan result)
	errs := make(chan error, 2)
	go parseResults(reader, results, errs)

	teams := make(chan []performance)
	go accumulateResults(results, teams)

	writeResults(writer, <-teams, errs)
	for i := 0; i < 2; i++ {
		if e := <-errs; e != nil {
			return e
		}
	}
	return nil
}

func writeResults(writer io.Writer, teams []performance, errors chan<- error) {
	w := bufio.NewWriter(writer)
	defer func() {
		if e := w.Flush(); e != nil {
			errors <- e
		}
	}()

	if _, e := w.WriteString("Team                           | MP |  W |  D |  L |  P\n"); e != nil {
		errors <- e
		return
	}
	for _, t := range teams {
		if _, e := w.WriteString(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", t.team, t.played, t.won, t.drawn, t.lost, t.points())); e != nil {
			errors <- e
			return
		}
	}
	errors <- nil
}

func parseResults(reader io.Reader, results chan<- result, errors chan<- error) {
	defer close(results)
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				errors <- r
			default:
				errors <- fmt.Errorf("unknown error: %v", r)
			}
		}
	}()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			panic(fmt.Sprintf("malformed input: %s", line))
		}
		home, away, res := parts[0], parts[1], parts[2]

		switch res {
		case "win":
			results <- result{home, WIN}
			results <- result{away, LOSS}
		case "loss":
			results <- result{home, LOSS}
			results <- result{away, WIN}
		case "draw":
			results <- result{home, DRAW}
			results <- result{away, DRAW}
		default:
			panic(fmt.Errorf("unrecognised result: %v", res))
		}
	}
	errors <- scanner.Err()
}

func accumulateResults(results <-chan result, ts chan<- []performance) {
	defer close(ts)

	rs := make(map[string]stats)
	for res := range results {
		perf := rs[res.team]

		switch res.result {
		case WIN:
			perf.won++
		case LOSS:
			perf.lost++
		case DRAW:
			perf.drawn++
		}
		perf.played++

		rs[res.team] = perf
	}

	ps := make([]performance, 0, len(rs))
	for k, v := range rs {
		ps = append(ps, performance{k, v})
	}
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].points() == ps[j].points() {
			return ps[i].team <= ps[j].team
		}
		return ps[i].points() > ps[j].points()
	})
	ts <- ps
}

func (p performance) String() string {
	return fmt.Sprintf("%#v (%d)\n", p, p.points())
}
