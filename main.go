package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func main() {

	// init router
	router := mux.NewRouter()

	//handle end point / routing
	router.HandleFunc("/polindrom/{args}", checkPolindrom).Methods("GET")
	router.HandleFunc("/polindromtext/", checkPolindromText).Methods("GET")
	router.HandleFunc("/getLang", getLang).Methods("GET")
	router.HandleFunc("/addLang", addLang).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(notFound)
	log.Fatal(http.ListenAndServe(":5000", router))

}

func checkPolindrom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	pali := isPalindrome(params["args"])

	json.NewEncoder(w).Encode(pali)

}
func checkPolindromText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := r.URL.Query()
	text := v.Get("text")
	pali := isPalindrome(text)
	if pali == true {
		w.Header().Set("Response-Desc", "Success")
		w.Write([]byte(`{"message":"Palindrome" }`))
	} else {
		w.Header().Set("Response-Desc", "Failed")
		w.Write([]byte(`{"message":"Not Palindrome" }`))
	}
	//json.NewEncoder(w).Encode(pali)

}
func isPalindrome(arg string) bool {
	argPalin := strings.ToLower(arg)
	arrStr := strings.Split(argPalin, "")
	countFail := true
	for i, rangeArg := range arrStr {
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

	//if !Auth(w, r) {
	//	return
	//}
	if !AllowOnlyGET(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var structInput DetailLanguage
	var influenStruct Influen

	influenStruct.InfluencedBy = append(influenStruct.InfluencedBy, "B")
	influenStruct.InfluencedBy = append(influenStruct.InfluencedBy, "ALGOL 68")
	influenStruct.InfluencedBy = append(influenStruct.InfluencedBy, "Assembly")
	influenStruct.InfluencedBy = append(influenStruct.InfluencedBy, "FORTRAN")

	influenStruct.Influences = append(influenStruct.Influences, "C++")
	influenStruct.Influences = append(influenStruct.Influences, "Objective-C")
	influenStruct.Influences = append(influenStruct.Influences, "C#")
	influenStruct.Influences = append(influenStruct.Influences, "Java")
	influenStruct.Influences = append(influenStruct.Influences, "Javascript")
	influenStruct.Influences = append(influenStruct.Influences, "PHP")
	influenStruct.Influences = append(influenStruct.Influences, "Go")

	structInput.Language = "C"
	structInput.Appeared = 1972
	structInput.Created = append(structInput.Created, "Dennis Ritchie")
	structInput.Functional = true
	structInput.Objectorient = false
	structInput.Relation.InfluencedBy = influenStruct.InfluencedBy
	structInput.Relation.Influences = influenStruct.Influences

	msg = append(msg, structInput)
	json.NewEncoder(w).Encode(structInput)

}

var ListStoredData []*StoredData
var msg []DetailLanguage

func addLang(w http.ResponseWriter, r *http.Request) {

	//if !Auth(w, r) {
	//	return
	//}
	//if !AllowOnlyGET(w, r) {
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")
	idx := 1

	lengthStored := len(ListStoredData)
	json.NewDecoder(r.Body).Decode(&ListStoredData)

	idx += lengthStored
	detailStored := new(StoredData)
	for _, rangeMsg := range msg {
		detailStored.ID = idx
		detailStored.ListDetailLang.Language = rangeMsg.Language
		detailStored.ListDetailLang.Appeared = rangeMsg.Appeared
		detailStored.ListDetailLang.Created = rangeMsg.Created
		detailStored.ListDetailLang.Functional = rangeMsg.Functional
		detailStored.ListDetailLang.Objectorient = rangeMsg.Objectorient
		detailStored.ListDetailLang.Relation = rangeMsg.Relation
		ListStoredData = append(ListStoredData, detailStored)
		idx++
	}
	json.NewEncoder(w).Encode(detailStored)

}

func notFound(w http.ResponseWriter, r *http.Request) {

	if !AllowOnly(w, r) {
		return
	}

	//if !Auth(w, r) {
	//	return
	//}
	//if !AllowOnlyGET(w, r) {
	//	return
	//}

	//if id := r.URL.Query().Get("id"); id != "" {
	//	OutputJSON(w, addLang(id))
	//	return
	//}

	//OutputJSON(w, getLang)
	//OutputJSON(w, addLang)
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}
func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
func AllowOnly(w http.ResponseWriter, r *http.Request) bool {
	if r.RequestURI != "/getLang" {
		w.Write([]byte("Method not allowed"))
		return false
	}
	if r.RequestURI != "/addLang" {
		w.Write([]byte("Method not allowed"))
		return false
	}

	return true
}

// data for add
/*
{
    "language": "C",
    "appeared": 1972,
    "created": [
        "Dennis Ritchie"
    ],
    "functional": true,
    "object-oriented": false,
    "relation": {
        "influenced-by": [
            "B",
            "ALGOL 68",
            "Assembly",
            "FORTRAN"
        ],
        "influences": [
            "C++",
            "Objective-C",
            "C#",
            "Java",
            "Javascript",
            "PHP",
            "Go"
        ]
    }
}

*/
