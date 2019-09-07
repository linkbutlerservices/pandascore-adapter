# Pandascore esports match result adapter #

**Author:** Link Butler  
**Author URI:** https://link-butler.com  
**License:** MIT License  
**Version:** 1.0     
_Built with [Bridges](https://github.com/linkpoolio/bridges)._

## Description ##
A simple chainlink external adapter, allowing smart contracts to fetch esports match results.  

Supported leagues

| Games             | Leagues                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| LOL               | All-Star, CBLOL, Challenger Korea, Demacia Cup, European Masters, Intel Extreme Masters, KeSPA Cup, LCK, LCS, League of Origin, LEC, LJL, LLA, LMS, LPL, LVP SLO, Mid-Season Invitational, NA Scouting Grounds, NEST, OPL, Rift Rivals: CB vs. CLS vs. LLN, Rift Rivals: GPL vs. LJL vs. OPL, Rift Rivals: LCK vs. LPL vs. LMS, Rift Rivals: LCK vs LPL vs LMS/VCS, Rift Rivals: LCL vs. TCL, Rift Rivals: LCL vs. TCL vs. VCS, Rift Rivals: NA vs. EU, TCL, VCS, World Championship, |
| DOTA 2            | Captains Draft,  China Dota2 Winter Cup, Dota Asia Championships, Dota Pit League, DreamLeague, EPICENTER, ESL One, GESC, Mars Dota League, Perfect World, PGL, Professional League, StarLadder, The International, The Summit                                                                                                                                                                                                                                                        |
| OVERWATCH         | Contenders Europe, Contenders Korea, Contenders North America, Overwatch League, World Cup                                                                                                                                                                                                                                                                                                                                                                                            |
| CS:GO             | BLAST Pro Series, Charleroi Esports, Copenhagen Games, cs_summit, DreamHack, ECS, ELEAGUE, EMS, EPICENTER, ESL, ESL ESEA, ESL One, ESWC, FACEIT League, GG.Bet Ice Challenge, iBUYPOWER, IEM, Intel Challenge, MLG Columbus, PGL Major, PLG Grand Slam, StarLadder StarSeries, StarSeries & i-League, TOYOTA Master, United Masters, WESG                                                                                                                                             |
| PUBG              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |

## Usage ##
##### Parameters
* **match_id_or_slug** - (string) (required) The id or slug of the match. Default: None
##### Result
* **match_result**  - (Bytes32) a Bytes32 representation of the match result where each attribute of the match is separated by :.   
    * **Format** : matchStatus:matchResult:matchNumberOfGames:matchWinnerId:matchWinnerAcronym
     *  **matchStatus** - (string) current status of the match, value can be one of:
        * fin  -  match finished
        * tba -  match not started
        * rng -  match in progress
        * cxl -  match cancelled
     *  **matchResult** - (string) result of the match, value can be one of :
        * drw  -  match concluded with a draw (matchWinnerId is empty, matchWinnerAcronym is empty)
        * ff -  match concluded with a forfeit from one of the teams (matchWinnerId is defined, matchWinnerAcronym is defined)
        * win -  match concluded with a winner (matchWinnerId is defined, matchWinnerAcronym is defined)
     *  **matchNumberOfGames** - (int) number of games played during the match
     *  **matchWinnerId** - (int) id of the winner 
     *  **matchWinnerAcronym** - (string) acronym of the winner (slug was too long to store in a Bytes32)
     
#### Examples
##### Request with smart contract
```
Chainlink.Request memory req = buildChainlinkRequest(jobId, this, this.fulfill.selector);
req.add("match_id_or_slug", "virtus-pro-vs-infamous-2019-08-17");
req.add("copyPath", "match_result");
```
Result :
```
<!-- fin:win:2:1651:VP Encoded in Bytes32 -->
0x66696e3a77696e3a323a313635313a5650000000000000000000000000000000

```
##### Request with curl
```
curl -X POST -H 'Content-Type: application/json' \
-d @- << EOF
{
	"jobRunId": "1234",
	"data": {
		"match_id_or_slug": "virtus-pro-vs-infamous-2019-08-17",
	}
}
EOF
```
Result :
```json
{
    "jobRunId": "1234",
    "status": "completed",
    "error": null,
    "pending": false,
    "data": {
        "match_id_or_slug": "virtus-pro-vs-infamous-2019-08-17",
        "match_result": "fin:win:2:1651:VP"
    }
}
```

## Setup ##
#### Local Install
**Requirements:** [Golang](https://golang.org/pkg/) is installed.

##### Build
```
make build
```

##### Run the adapter
```
API_KEY=apikey ./pandascore-adapter
```

#### Docker
##### Run the container:
```
docker run -it -e API_KEY=apikey -p 8080:8080 linkbutlerservices/pandascore-adapter
```
