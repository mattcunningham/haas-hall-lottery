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
    console.log(jqXHR);
}