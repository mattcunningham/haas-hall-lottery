package lottery

// written by Matthew Cunningham

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const HOME = `<!doctype html5>
<html>
  <head>
    <title>Haas Hall Academy Lottery</title>
    <style>
@charset "utf-8";

@import "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css";

body {
  background: url('http://haashall.org/lottery/assets/content.jpg');
}

.container {
  max-width: 980px;
}

.import-container {
  margin: 10px auto;
  position: relative;
  padding: 100px 0 0;
  width: 500px;
}

#rose {
  position: absolute;
  left: 0;
  top: 0;
  z-index: 10;
  width: 500px;
  height: 200px;
  background: url('http://haashall.org/lottery/assets/images/waxseal-300.png') no-repeat center center;
}

#import {
  margin: 0 auto 40px;
  width: 500px;
  background: #fff;
  text-align: center;
  padding: 100px 0 0;
  -webkit-box-shadow: 10px 10px 80px 10px rgba(0,0,0,0.16);
	-moz-box-shadow: 10px 10px 80px 10px rgba(0,0,0,0.16);
	box-shadow: 10px 10px 80px 10px rgba(0,0,0,0.16);
}

#import h1 {
  font-weight: bolder;
  font-size: 40px;
}

#import h2 {
  margin-top: 0;
  text-transform: uppercase;
}

#import h3 {
  margin-top: 0;
  font-weight: bold;
  color: #337ab7;
}

.drag {
  display: block;
  width: 452px;
  height: 216px;
	padding-top: 240px;
	overflow: hidden;
	-webkit-overflow: hidden;
	-moz-overflow: hidden;
  margin: 0 auto 20px;
  background: url('http://haashall.org/lottery/assets/images/drag.png') no-repeat center center;
}

label#getDataButton {
	font-family: Arial;
	color: #ffffff;
	font-size: 24px;
	font-weight: bold;
	text-transform: uppercase;
	padding: 10px 20px 10px 20px;
	text-decoration: none;
	margin-bottom: 30px;
}

.export {
	display: none;
}

.export h1 {
  font-weight: bolder;
  font-size: 54px;
  margin-top: 50px;
  text-align: right;
}
.export h2 {
  margin-top: 0;
  text-transform: uppercase;
  text-align: right;
}

.export-footer h3 span {
  font-weight: bold;
}

a#export {
	font-family: Arial;
	color: #ffffff;
	font-size: 24px;
	font-weight: bold;
	text-transform: uppercase;
	padding: 10px 20px 10px 20px;
	text-decoration: none;
	float: right;
}
		</style>
  </head>
  <body>
    <main>
      <form id="container" action="/import" method="post" enctype="multipart/form-data">
				<div class="import-container">
					<div id="rose"></div>
					<div id="import">
						<h1>Haas Hall Academy</h1>
						<h2>Admissions Lottery</h2>
						<h3>2017 &mdash; 2018</h3>
	  				<input class="drag" class="btn btn-primary" value="Import CSV Data" accept=".csv" type="file" id="getData"/>
						<label for="getData" id="getDataButton" class="btn btn-primary">Import CSV Data</label>
					</div>
				</div>
				<div class="export">
					<div class="container">
						<div class="row">
	          	<div class="col-sm-2">
	              <img src="http://haashall.org/lottery/assets/images/waxseal.png">
	          	</div>
	          	<div class="col-sm-10">
	              <h1>Haas Hall Academy</h1>
	              <h2>Admissions Lottery 2017 - 2018</h2>
	          	</div>
	        	</div>
						<table class="table table-bordered">
							<tbody id="fromServer">

							</tbody>
						</table>
						<div class="export-footer">
							<div class="row">
								<div class="col-sm-6"></div>
								<div class="col-sm-6"><a href="" id="export" class="btn btn-primary">Export</a></div>
							</div>
						</div>
					</div>
				</div>
      </form>
    </main>
    <script
		  src="https://code.jquery.com/jquery-3.1.1.min.js"
		  integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8="
		  crossorigin="anonymous"></script>
    <script type="text/javascript">
$(":file").change(function() {
  var file = this.files[0], result;
  if (file) {
    var reader = new FileReader;
    reader.readAsText(file);
    reader.onload = function(e) {
      result = e.target.result;
      $.ajax({ url: "post", type: "POST", dataType: "json", success: completeHandler,
        data: {
          data: result,
          entries: prompt("How many people do you want to admit?")
        }
      });
    };
  }
});
function completeHandler(serverData, textStatus) {
  outputFile = "";
  var header = createHeader(serverData[0]);
  var headerOutput = header[0], headerFile = header[1];
  $("#fromServer").append(headerOutput);
  outputFile += headerFile;
  serverData.forEach(function(entry) {
    var ret = createElement(entry);
    var forOutput = ret[0];
    var forFile = ret[1];
    $("#fromServer").append(forOutput);
    $("#export").css("display", "block");
    outputFile += forFile;
  });
  $("#getDataButton").css("display", "none");
	$(".import-container").css("display", "none");
	$(".export").css("display", "block");
  var a = document.getElementById("export");
  var file = new Blob([outputFile], { type: "text/plain" });
  a.href = URL.createObjectURL(file);
  a.download = "results.csv";
}
function createHeader(entry) {
  fileData = ""
  output = '<tr class="entry header">'
  for (var key in entry) {
    if (key === "Grade" || key === "Priority") {
      fileData += key + ",";
    } else if (entry[key] === Object(entry[key])) {
      for (var objKey in entry[key]) {
        fileData += objKey + ",";
      }
    } else {
      fileData += key + ",";
			output += '<th class="entry-item header">' + key + "</th>";
    }
  }
	output += "</tr>"
  fileData = fileData.substring(0, fileData.length - 1) + "\n";
  return [output, fileData];
}
function createElement(entry) {
  index = 0;
  output = '<tr class="entry">';
  fileData = "";
  for (var key in entry) {
    if (key === "Grade" || key === "Priority") {
      fileData += formatCSV(key, entry[key]) + ",";
    } else if (entry[key] === Object(entry[key])) {
      for (var objKey in entry[key]) {
        fileData += entry[key][objKey] + ",";
      }
    } else { // this is the real information
      fileData += formatCSV(key, entry[key]) + ",";
      output += '<td class="entry-item ' + ((index % 2 == 0) ? "even" : "odd") + '">' + formatAdmittance(key, entry[key]) + "</span>";
      index++;
    }
  }
  fileData = fileData.substring(0, fileData.length - 1) + "\n";
  output += "</tr>\n";
  return [output, fileData];
}
var waitListCounter = 0;
function formatCSV(key, info) {
  if (key !== "Status" && key !== "Priority") {
    return info
  }
  if (key === "Status") {
    if (info == 2) {
      return "Admitted";
    } else if (info == 1) {
      waitListCounter++;
      return "Wait list " + waitListCounter;
    }
  }
  if (key === "Priority") {
    return info == 2 ? "Faculty" : (info == 1 ?  "Sibling" : "");
  }
}
var waitListAdmitCounter = 0;
function formatAdmittance(key, info) {
  if (key !== "Status" && key !== "Priority") {
    return info;
  }
  if (key === "Status") {
    if (info == 2) {
      return "<span id=\"admitted\">Admitted</span>";
    } else if (info == 1) {
      waitListAdmitCounter++;
      return '<span id="waitlisted">Wait list <span id="waitlist-count">' + waitListAdmitCounter + "</span></span>";
    }
  }
  if (key === "Priority") {
    return 2 == info ? '<span id="faculty">Faculty</span>' : (1 == info ? '<span id="sibling">Sibling</span>' : "");
  }
}
    </script>
  </body>
</html>`

type fileError struct {
	error string
}

func (e fileError) Error() string {
	return fmt.Sprintf("%s", e.error)
}

// Writes a 404 message to the page
func NotFound(w http.ResponseWriter, r *http.Request) {
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
