package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

// scoreBoard keeps track of the score
type scoreBoard map[string]*team

// team keeps track a team's outcomes
type team struct {
	name                            string
	played, win, loss, draw, points int
}

// Tally reads an input file to create a scoreBoard
func Tally(reader io.Reader, writer io.Writer) error {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields
	board := make(scoreBoard)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := board.addRecord(record); err != nil {
			return err
		}
	}

	teams := byScore(board.getTeams())
	sort.Sort(teams)
	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range teams {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.name, team.played, team.win, team.draw, team.loss, team.points)
	}
	return nil
}

// addGame validates and adds a game record to the scoreboard
func (b scoreBoard) addRecord(record []string) error {
	if len(record) != 3 {
		return fmt.Errorf("Invalid record: %v", record)
	}
	home, homeInBoard := b[record[0]]
	away, awayInBoard := b[record[1]]
	if !homeInBoard {
		home = &team{name: record[0]}
		b[record[0]] = home
	}
	if !awayInBoard {
		away = &team{name: record[1]}
		b[record[1]] = away
	}
	switch record[2] {
	case "win":
		addWin(home, away)
	case "loss":
		addWin(away, home)
	case "draw":
		addDraw(home, away)
	default:
		return fmt.Errorf("Invalid outcome: %q", record[2])
	}
	return nil
}

// getTeams converts the scoreboard into a list
func (b scoreBoard) getTeams() []team {
	var teams []team
	for _, team := range b {
		teams = append(teams, *team)
	}
	return teams
}

// byScore sorts the team by points, then by wins, then aphabetically
type byScore []team

func (t byScore) Len() int      { return len(t) }
func (t byScore) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t byScore) Less(i, j int) bool {
	switch {
	case t[i].points != t[j].points:
		return t[i].points > t[j].points
	case t[i].win != t[j].win:
		return t[i].win > t[j].win
	default:
		return t[i].name < t[j].name
	}
}

// addWin adds win to winner team & loss to the loser one
func addWin(winner, loser *team) {
	winner.played++
	loser.played++
	winner.win++
	loser.loss++
	winner.points += 3
}

// addDraw adds draw to both teams
func addDraw(home, away *team) {
	home.played++
	away.played++
	home.draw++
	away.draw++
	home.points++
	away.points++
}
