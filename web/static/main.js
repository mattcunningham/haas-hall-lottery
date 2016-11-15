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
    jqXHR.forEach(function(entry) {
	    $("#fromServer").append(createElement(entry))
    });
    $("#getDataButton").css("display","none");
}

function createElement(entry) {
    index = 0
    output = "<div class=\"entry\">"
    for (var key in entry) {
	if (entry.hasOwnProperty(key)) {
	    if (entry[key] === Object(entry[key])) { 
		for (var objKey in entry[key]) {
		    if (index % 2 == 0) { // even
			output += "<span class=\"entryItem-even\">" + entry[key][objKey] + "</span>";
		    } else {
			output += "<span class=\"entryItem-odd\">" + entry[key][objKey] + "</span>";
		    }
		    index++
		}
	    } else {
		if (index % 2 == 0) { // even
		    output += "<span class=\"entryItem-even\">" + entry[key] + "</span>";
		} else {
		    output += "<span class=\"entryItem-odd\">" + entry[key] + "</span>";
		}
		index++
	    }
	}
    }
    output += "</div>\n"
	console.log(output);
    return output
}