package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name   string
	played uint
	wins   uint
	draws  uint
	losses uint
}

func (team *Team) win() {
	team.played++
	team.wins++
}

func (team *Team) lose() {
	team.played++
	team.losses++
}

func (team *Team) draw() {
	team.played++
	team.draws++
}

func (team *Team) points() uint {
	return team.wins*3 + team.draws
}

type Teams map[string]*Team

func (teams Teams) getOrCreateTeam(name string) *Team {
	_, exists := teams[name]

	if !exists {
		teams[name] = &Team{name: name}
	}

	return teams[name]
}

func (teams Teams) fromFile(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line[0] == '#' {
			continue
		}

		stats := strings.Split(line, ";")

		if len(stats) != 3 {
			return errors.New("Invalid file")
		}

		homeTeam, awayTeam, outcome := stats[0], stats[1], stats[2]

		switch outcome {
		case "win":
			teams.getOrCreateTeam(homeTeam).win()
			teams.getOrCreateTeam(awayTeam).lose()
		case "loss":
			teams.getOrCreateTeam(homeTeam).lose()
			teams.getOrCreateTeam(awayTeam).win()
		case "draw":
			teams.getOrCreateTeam(homeTeam).draw()
			teams.getOrCreateTeam(awayTeam).draw()
		default:
			return errors.New("Invalid outcome")
		}
	}

	return nil
}

func (teams Teams) sort() []*Team {
	scores := []*Team{}

	for _, team := range teams {
		scores = append(scores, team)
	}

	sort.Slice(scores, func(a, b int) bool {
		if scores[a].points() == scores[b].points() {
			return scores[a].name < scores[b].name
		}

		return scores[a].points() > scores[b].points()
	})

	return scores
}

func Tally(reader io.Reader, writer io.Writer) error {
	tallies := Teams{}

	err := tallies.fromFile(reader)

	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")

	if err != nil {
		return err
	}

	sortedScores := tallies.sort()

	for _, score := range sortedScores {
		_, err := fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", score.name, score.played, score.wins, score.draws, score.losses, score.points())

		if err != nil {
			return err
		}
	}

	return nil
}
