$(document).ready(function () {
	$("#recipeListDelete").on('click', function () {
        var recipes = $("input[name='recipeListDeleteCheck']:checked").map(function () { 
        	return this.value; 
        }).get();

        recipes.forEach(function(item) {
			$.ajax({
				url: '/api/recipes/'+item+'',
				type: "DELETE",
				contentType: "application/json",
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
});
