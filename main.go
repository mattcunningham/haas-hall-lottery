package lottery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//	"google.golang.org/appengine"
)

const (
	RED    = "\u001B[31m"
	YELLOW = "\u001B[33m"
	GREEN  = "\u001B[32m"
	END    = "\u001B[0m"
	HOME   = `<!doctype html5>
<html>
  <head>
    <title>Haas Hall Academy Lottery</title>
    <link href="https://www.haashall.org/lottery/style.css" rel="stylesheet" type="text/css" />
  </head>
  <body>
    <main>
      <form id="container" action="/import" method="post" enctype="multipart/form-data">
<div id="import">
  <label for="getData" id="getDataButton">Import CSV Data</label>
  <input value="Import CSV Data" accept=".csv" type="file" id="getData"/>
</div>
<a href="" id="export">export</a>
<div id="fromServer"></div>
      </form>
    </main>
    <script
  src="https://code.jquery.com/jquery-3.1.1.min.js"
  integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8="
  crossorigin="anonymous"></script>
    <script type="text/javascript">
$(':file').change(function(){
var file = this.files[0];
var name = file.name;
var size = file.size;
var type = file.type;
var result;
      if (file) {
    var reader = new FileReader();
    reader.readAsText(file);
    reader.onload = function(e) {
    result = e.target.result;
$.ajax({
url: 'post',  //Server script to process data
type: 'POST',
dataType: 'json',
        success: completeHandler,
        data: {'data': result, 'entries': prompt("How many people do you want to admit?")},
});
    }
}
});

function completeHandler(jqXHR, textStatus) {
    outputFile = "";
    var header = createHeader(jqXHR[0]);
    console.log(header);
    var headerOutput = header[0];
    var headerFile = header[1];
    $("#fromServer").append(headerOutput);
    outputFile += headerFile;
    jqXHR.forEach(function(entry) {
        var ret = createElement(entry);
        var forOutput = ret[0];
        var forFile = ret[1];
        $("#fromServer").append(forOutput);
        $("#export").css("display", "block");
        outputFile += forFile;
    });
    $("#getDataButton").css("display","none");
    var a = document.getElementById("export");
    var file = new Blob([outputFile], {type: "text/plain"});
    a.href = URL.createObjectURL(file);
    a.download = "results.csv";
}

function createHeader(entry) {
    fileData = ""
    output = "<div class=\"entry header\">"
    for (var key in entry) {
        if (key === "Grade" || key === "Priority") {
            fileData += key + ", ";
            continue;
        }
        if (entry[key] === Object(entry[key])) {
            for (var objKey in entry[key]) {
                fileData += objKey + ", ";
            }
        } else {
            fileData += key + ", ";
            output += "<span class=\"entry-item header\">" + key + "</span>";
        }
    }
    fileData = fileData.substring(0, fileData.length - 2) + "\n";
    return [output, fileData];
}

function createElement(entry) {
    index = 0
    output = "<div class=\"entry\">"
    fileData = ""
    for (var key in entry) {
        if (entry.hasOwnProperty(key)) {
            if (key === "Grade" || key === "Priority") {
                fileData += entry[key] + ", ";
                continue;
            }
            if (entry[key] === Object(entry[key])) { 
                for (var objKey in entry[key]) {
                    fileData += entry[key][objKey] + ", ";
                }
            } else { // this is for the real information
                if (index % 2 == 0) { // even
                    fileData += formatCSV(key, entry[key]) + ", ";
                    output += "<span class=\"entry-item even\">" + formatAdmittance(key, entry[key]) + "</span>";
                } else {
                    fileData += formatCSV(key, entry[key]) + ", ";
                    output += "<span class=\"entry-item odd\">" + formatAdmittance(key, entry[key]) + "</span>";
                }
                index++;
            }
        }
    }
    fileData = fileData.substring(0, fileData.length - 2) + "\n";
    output += "</div>\n"
    return [output, fileData];
}

function formatCSV(key, info) {
    if (key !== "Status") {
        return info
    }
    if (info == 2) {
        return "Admitted";
    } else if (info == 1) {
        return "Waitlisted";
    }
}

function formatAdmittance(key, info) {
    console.log(key);
    if (key !== "Status" && key !== "Priority") {
       return info;
    }
    if (key === "Status") {
if (info == 2) {
    return "<span id=\"admitted\">Admitted</span>";
} else if (info == 1) {
    return "<span id=\"waitlisted\">Waitlisted</span>";
}
    }
    if (key === "Priority") {
if (info == 2) {
    return "<span id=\"faculty\">Faculty</span>";
} else if (info == 1) {
    return "<span id=\"sibling\">Sibling</span>";
} else {
    return "";
}
    }
}
    </script>
    <!-- also I'm not a supporter of school choice -->
  </body>
</html>`
)

type fileError struct {
	error string
}

func (e fileError) Error() string {
	return fmt.Sprintf("%s", e.error)
}

// Returns a string formatted to be a specific color when printed in the console
// The only available colors are listed as constants. Add more as necessary.
// ex. formatPrint("my string", RED)
func formatPrint(s string, color string) string {
	return color + s + END
}

// Writes a 404 message to the page
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatPrint("REQUEST[404]: ", RED), r.URL)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("nothing to see here"))
	w.WriteHeader(http.StatusNotFound)
}

// Main function to serve the site
func init() {
	// serves for the webpage "/" — i.e. the homepage

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { // we want "/" with no appending text
			NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(HOME)) // prints to webpage
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/post" && r.Method != "POST" {
			NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		entries := CSVtoEntries(r.FormValue("data"))
		sorted := Sort(entries)
		prioritized := Prioritize(sorted)
		i, err := strconv.Atoi(r.FormValue("entries"))
		if err != nil {
			return
		}
		admitted := Admit(prioritized, i)
		data, err := json.Marshal(admitted) // the error needs to do something
		if err != nil {
			fmt.Println(err)
		}
		w.Write(data)
	})
}
