package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// id
// name
// descripsi
// bahan

// Roll  is model for sushi pai
type Roll struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingridient  string `json:"ingridient"`
}

// Init rolls var as slice
var rolls []Roll

// show all sushi
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

// Show single sushi
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, roll := range rolls {
		if roll.ID == params["id"] {
			json.NewEncoder(w).Encode(roll)
			return
		}
	}
}

// add single sushi
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)

	json.NewEncoder(w).Encode(newRoll)

}

func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID==params["id"]{
			rolls = append(rolls[:i], rolls[i+1:] ...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls  = append(rolls,newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return


		}
	}
}
func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID==params["id"]{
			rolls = append(rolls[:i], rolls[i+1:] ...)
			break

		}
	}


}

func checkPolindrom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	pali := isPalindrome(params["args"])

	json.NewEncoder(w).Encode(pali)

}
func isPalindrome(arg string)bool{
	argPalin := strings.ToLower(arg)
	arrStr := strings.Split(argPalin,"")
	countFail := true
	for i,rangeArg := range arrStr{
		if i >= len(arrStr)-1-i {
			break
		}
		pairStr := arrStr[len(arrStr)-1-i]
		if rangeArg != pairStr {
			countFail = false
			break
		}
	}
	if !countFail {
		return false
	}
	return true
}

func getLang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var structInput DetailLanguage
	var influenStru Influen

	influenStru.InfluencedBy = append(influenStru.InfluencedBy,"B")
	influenStru.InfluencedBy = append(influenStru.InfluencedBy,"ALGOL 68")
	influenStru.InfluencedBy = append(influenStru.InfluencedBy,"Assembly")
	influenStru.InfluencedBy = append(influenStru.InfluencedBy,"FORTRAN")

	influenStru.Influences = append(influenStru.Influences,"C++")
	influenStru.Influences = append(influenStru.Influences,"Objective-C")
	influenStru.Influences = append(influenStru.Influences,"C#")
	influenStru.Influences = append(influenStru.Influences,"Java")
	influenStru.Influences = append(influenStru.Influences,"Javascript")
	influenStru.Influences = append(influenStru.Influences,"PHP")
	influenStru.Influences = append(influenStru.Influences,"Go")

	structInput.Language = "C"
	structInput.Appeared = 1972
	structInput.Created = append(structInput.Created,"Dennis Ritchie")
	structInput.Functional = true
	structInput.Objectorient = false
	structInput.Relation.InfluencedBy = influenStru.InfluencedBy
	structInput.Relation.Influences = influenStru.Influences

	//json.NewDecoder(r.Body).Decode(&detailLanguage)
	msg =append(msg,structInput)
	json.NewEncoder(w).Encode(structInput)

}
var ListStoredData []*StoredData
var msg []DetailLanguage

func addLang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var newRoll Roll
	//json.NewDecoder(r.Body).Decode(&newRoll)
	//newRoll.ID = strconv.Itoa(len(rolls) + 1)
	//rolls = append(rolls, newRoll)
	//
	//json.NewEncoder(w).Encode(newRoll)

	idx := 1

	lengthStored := len(ListStoredData)
	json.NewDecoder(r.Body).Decode(&ListStoredData)

	idx += lengthStored
	detailStored := new(StoredData)
	for _,rangeMsg := range msg{
		detailStored.ID = idx
		detailStored.ListDetailLang.Language = rangeMsg.Language
		detailStored.ListDetailLang.Appeared = rangeMsg.Appeared
		detailStored.ListDetailLang.Created  = rangeMsg.Created
		detailStored.ListDetailLang.Functional  = rangeMsg.Functional
		detailStored.ListDetailLang.Objectorient = rangeMsg.Objectorient
		detailStored.ListDetailLang.Relation = rangeMsg.Relation
		ListStoredData = append(ListStoredData,detailStored)
		idx++
	}
	json.NewEncoder(w).Encode(detailStored)


}

func main() {

	//Generate Mock data

	rolls = append(rolls,
		Roll{
			ID:          "1",
			Name:        "salmon",
			Description: "crab salmon",
			Ingridient:  "salmon nori, rice",
		}, Roll{
			ID:          "2",
			Name:        "salmon 2",
			Description: "crab salmon 2",
			Ingridient:  "salmon nori, rice 2",
		},

	)

	// init router
	router := mux.NewRouter()

	//handle end point / routing
	router.HandleFunc("/polindrom/{args}", checkPolindrom).Methods("GET")
	router.HandleFunc("/getLang", getLang).Methods("GET")
	router.HandleFunc("/addLang", addLang).Methods("POST")



	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}
