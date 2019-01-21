package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/rs/cors"
    "strconv"
    "math/rand"
)

type myData struct {
    Username string
    Password  string
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

type Stats struct {
    Statcards Statcards `json:"statcards"`
    Userchart Serieslist `json:"userchart"`
    Activitychart Serieslist `json:"activitychart"`
    PreferencechartSeries Singleseries `json:"preferencechartSeries"`
    PreferencechartLabel SingleseriesString `json:"preferencechartLabel"`
}

type Statcard struct {
    Type string `json:"type"`
    Icon string `json:"icon"`
    Title string `json:"title"`
    Value string `json:"value"`
    FooterText string `json:"footerText"`
    FooterIcon string `json:"footerIcon"`
}

type Statcards []Statcard

type Singleseries []int

type Serieslist []Singleseries

type SingleseriesString []string

func login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    var t myData
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
    if (t.Username == "admin@gmail.com" && t.Password == "adminpass") {
        fmt.Fprintf(w, "true")
    } else {
        http.Error(w, "Invalid Username or Password!!", 401)
    }
}

func stats(w http.ResponseWriter, r *http.Request) {
    stat_cards := Statcards{
        Statcard{Type: "warning", Icon: "ti-server", Title: "Capacity", Value: strconv.Itoa(randomInt(50,150)) + "GB" , FooterText: "Updated now", FooterIcon: "ti-reload"},
        Statcard{Type: "success", Icon: "ti-wallet", Title: "Revenue", Value: "$" + strconv.Itoa(randomInt(1000,1500)), FooterText: "Last day", FooterIcon: "ti-calendar"},
        Statcard{Type: "danger", Icon: "ti-pulse", Title: "Errors", Value: strconv.Itoa(randomInt(0,30)), FooterText: "In the last hour", FooterIcon: "ti-timer"},
        Statcard{Type: "info", Icon: "ti-twitter-alt", Title: "Followers", Value: "+" + strconv.Itoa(randomInt(20,50)), FooterText: "Updated now", FooterIcon: "ti-reload"},
    }

    user_charts := Serieslist{
        Singleseries {randomInt(200,1000), randomInt(200,1000), randomInt(200,1000), randomInt(200,1000), randomInt(200,1000), randomInt(200,1000), randomInt(200,1000), randomInt(200,1000), randomInt(200,1000)},
        Singleseries {randomInt(50,800), randomInt(50,800), randomInt(50,800), randomInt(50,800), randomInt(50,800), randomInt(50,800), randomInt(50,800), randomInt(50,800), randomInt(50,800)},
        Singleseries {randomInt(10,500), randomInt(10,500), randomInt(10,500), randomInt(10,500), randomInt(10,500), randomInt(10,500), randomInt(10,500), randomInt(10,500), randomInt(10,500)},
    }

    activity_chart := Serieslist{
        Singleseries {randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900), randomInt(500,900)},
        Singleseries {randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900), randomInt(200,900)},
    }

    first_series := randomInt(5,100)
    second_series := randomInt(5,100)
    third_series := randomInt(5,100)

    pref_series := Singleseries{first_series, second_series, third_series}
    pref_label := SingleseriesString{strconv.Itoa(first_series) + "%", strconv.Itoa(second_series) + "%", strconv.Itoa(third_series) + "%"}

    stats := Stats{Statcards: stat_cards, Userchart: user_charts, Activitychart: activity_chart, PreferencechartSeries: pref_series, PreferencechartLabel: pref_label}
    
    json.NewEncoder(w).Encode(stats)
}

func handleRequests() {
    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/login", login).Methods("POST")
    r.HandleFunc("/statscards", stats).Methods("GET")
    handler := cors.Default().Handler(r)
    log.Fatal(http.ListenAndServe(":8000", handler))
}

func main() {
    handleRequests()
}