package smartfblib
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
func SwitchOn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Switched on")
}
func SwitchOff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Switched on")
}
func SwitchChannel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var channel string
	channel = params["channel"]
	if channel != ""{
		fmt.Fprintf(w, "Switched to channel %d", channel)		
	} else {
		fmt.Fprintf(w, "Channel not found or invalid %d", channel)
	}
}
func SetVolume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var volume string
	volume = params["volume"]
	if volume != ""{
		fmt.Fprintf(w, "Volume set to %d", volume)		
	} else {
		fmt.Fprintf(w, "Volume Parameter not found or invalid %d", volume)
	}
}