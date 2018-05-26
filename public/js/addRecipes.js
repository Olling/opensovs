$(document).ready(function () {
	// On submit, parse and call the api with the form data.
    $("#addRecipesForm").on('click', function () {


    	var recipe = {
    		title:document.forms['addRecipesForm'].elements['title'].value,
    		added:document.forms['addRecipesForm'].elements['blog'].value,
    		blog:document.forms['addRecipesForm'].elements['blog'].value,
    		instructions:document.forms['addRecipesForm'].elements['instructions'].value
    	};

    	$.ajax({
            url: '/api/recipes',
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(recipe),
            beforeSend: function (xhrObj) {
            },
            success: function (data) {
            },
            error: function (xhr, resp, text) {
                console.log(xhr, resp, text);
            }
        });
    });

		
});