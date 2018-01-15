package main
 
import(
"fmt"
"net/http"
"os"
"time"
"html/template"
)

type Instance struct {
        Pod        string
        TS         string
        Version    string
}

 
func indexHandler( w http.ResponseWriter, r *http.Request){
tpl := template.Must(template.New("out").Parse(html))
i := &Instance{}
i = newInstance()
tpl.Execute(w, i)
 
}

func healthHandler( w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "I'm fine Version : %s ", version)

}


func newInstance() *Instance {
	var i = new(Instance)
	i.Pod, _ = os.Hostname()
	i.TS = time.Now().Format(time.UnixDate)
	i.Version = version
	return i
}
const version string = "1.1.5"
 
func main(){
http.HandleFunc("/", indexHandler)
http.HandleFunc("/healthz", healthHandler)
http.ListenAndServe(":8000",nil)
}
