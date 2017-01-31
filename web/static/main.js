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
			data: {'data': result},
		});
	    }
	}
});

function completeHandler(jqXHR, textStatus) {
    outputFile = "";
    $("#fromServer").append(createHeader(jqXHR[0]));
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
    output = "<div class=\"entry header\">"
    for (var key in entry) {
        if (entry[key] === Object(entry[key])) {
	    for (var objKey in entry[key]) {
		output += "<span class=\"entry-item header\">" + objKey + "</span>";
	    }
	} else {
	    output += "<span class=\"entry-item header\">" + key + "</span>";
	}
    }
    return output;
}

function createElement(entry) {
    index = 0
    output = "<div class=\"entry\">"
    fileData = ""
    for (var key in entry) {
	if (entry.hasOwnProperty(key)) {
	    if (entry[key] === Object(entry[key])) { 
		for (var objKey in entry[key]) { // this is for extraneous data; such as first name, last name, address — information that doesn't matter
		    if (index % 2 == 0) { // even
			fileData += entry[key][objKey] + ", ";
			output += "<span class=\"entry-item even\">" + entry[key][objKey] + "</span>";
		    } else {
			fileData += entry[key][objKey] + ", ";
			output += "<span class=\"entry-item odd\">" + entry[key][objKey] + "</span>";
		    }
		    index++
		}
	    } else { // this is for the real information
		if (index % 2 == 0) { // even
		    fileData += entry[key] + ", ";
		    output += "<span class=\"entry-item even\">" + formatAdmittance(key, entry[key]) + "</span>";
		} else {
		    fileData += entry[key] + ", ";
		    output += "<span class=\"entry-item odd\">" + formatAdmittance(key, entry[key]) + "</span>";
		}
		index++
	    }
	}
    }
    fileData += "\n";
    output += "</div>\n"
    return [output, fileData];
}

function formatAdmittance(key, info) {
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