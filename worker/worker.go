package worker

import (
	"bufio"
	"strings"
)

type PGN struct {
	Event     string
	Site      string
	Date      string
	Round     string
	White     string
	Black     string
	Result    string
	ECO       string
	EventDate string
	PlyCount  string
	Source    string
	EventType string
	Board     string
}

type PGNs struct {
	Inner []PGN
}

func NewPGN(event, site, date, round, white, black, result, eco, eventDate, plyCount, source, eventType, board string) PGN {
	return PGN{
		Event:     event,
		Site:      site,
		Date:      date,
		Round:     round,
		White:     white,
		Black:     black,
		Result:    result,
		ECO:       eco,
		EventDate: eventDate,
		PlyCount:  plyCount,
		Source:    source,
		EventType: eventType,
		Board:     board,
	}
}

func parseStringToPGN(pgnData string) (PGN, error) {
	pgn := PGN{}
	scanner := bufio.NewScanner(strings.NewReader(pgnData))
	boardContent := []string{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check if this line is part of the metadata
		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {
			content := line[1 : len(line)-1]
			parts := strings.SplitN(content, " ", 2)
			if len(parts) < 2 {
				continue
			}

			key := parts[0]
			value := strings.Trim(parts[1], "\"") // Remove quotes around value

			// Map to struct fields
			switch key {
			case "Event":
				pgn.Event = value
			case "Site":
				pgn.Site = value
			case "Date":
				pgn.Date = value
			case "Round":
				pgn.Round = value
			case "White":
				pgn.White = value
			case "Black":
				pgn.Black = value
			case "Result":
				pgn.Result = value
			case "ECO":
				pgn.ECO = value
			case "EventDate":
				pgn.EventDate = value
			case "PlyCount":
				pgn.PlyCount = value
			case "Source":
				pgn.Source = value
			case "EventType":
				pgn.EventType = value
			}
		} else if line != "" {
			// This line is part of the board content
			boardContent = append(boardContent, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return PGN{}, err
	}

	// Join all board content lines with a space or newline if needed
	pgn.Board = strings.Join(boardContent, " ")

	return pgn, nil
}
