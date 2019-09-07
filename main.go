package main

import (
	"github.com/linkpoolio/bridges/bridge"
	"net/http"
	"os"
	"strconv"
)

type PandaScore struct {}

func (av *PandaScore) Run(h *bridge.Helper) (interface{}, error) {
	matchId := h.GetParam("match_id_or_slug")
	result := make(map[string]interface{})
	matchStatus := ""
	matchResult := ""
	matchWinnerId := ""
	matchWinnerAcronym := ""
	matchNumberOfGames := ""

	response := make(map[string]interface{})
	err := h.HTTPCallWithOpts(
		http.MethodGet,
		"https://api.pandascore.co/matches/"+matchId,
		&response,
		bridge.CallOpts{
			Auth: bridge.NewAuth(bridge.AuthParam, "token", os.Getenv("API_KEY")),
		},
	)

	if err == nil{
		if status, ok := response["status"].(string); ok {
			switch status {
				case "canceled": matchStatus = "cxl";
				case "finished": matchStatus = "fin";
				case "not_started": matchStatus = "tba";
				case "running": matchStatus = "rng";
			}
		}

		if isDraw, ok := response["draw"].(bool); ok {
			if !!isDraw{
				matchResult = "drw"
			}
		}

		if isForfeit, ok := response["forfeit"].(bool); ok {
			if isForfeit{
				matchResult = "ff"
			}
		}

		if winner , ok := response["winner"]; ok {
			if(winner != nil){
				matchResult = "win"
			}
		}

		if winnerId, ok := response["winner_id"].(float64); ok {
			matchWinnerId = floatToString(winnerId)
		}

		if numberOfGames, ok := response["number_of_games"].(float64); ok {
			matchNumberOfGames = floatToString(numberOfGames)
		}

		if winner, ok := response["winner"].(map[string]interface{}); ok {
			if winner != nil {
				if acronym, ok := winner["acronym"].(string); ok {
					matchWinnerAcronym = acronym
				}
			}
		}
	}

	result["match_result"] = matchStatus+":"+matchResult+":"+matchNumberOfGames+":"+matchWinnerId+":"+matchWinnerAcronym

	return result,err
}

func floatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 0, 64)
}


func (av *PandaScore) Opts() *bridge.Opts {
	return &bridge.Opts{
		Name:   "PandaScore",
		Lambda: true,
	}
}

func main() {
	bridge.NewServer(&PandaScore{}).Start(8080)
}
