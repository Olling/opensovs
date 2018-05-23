$(document).ready(function () {

	$(function(){
		$('#recipeListRefresh').click(function() {

			$.ajax({
					url: '/api/recipes',
					type: "GET",
					beforeSend: function (xhrObj) {
					},
					complete: function () {
					},
					success: function (data) {
						$("#recipeList").empty();
						$('#recipeList').append('<thead id="recipeListTableHead">\
									<tr>\
										<th>\
											ID\
										</th>\
										<th>\
											Title\
										</th>\
										<th>\
											Added\
										</th>\
										<th>\
											Blog\
										</th>\
										<th>\
											Instructions\
										</th>\
									</tr>\
								</thead>\
								<tbody id="recipeListTableBody">\
								</tbody>\
		            		');
						var isJSON = true;
						try {
							var response = JSON.parse(data);
						}
						catch (err) {
							isJSON = false
						}
						if (isJSON == true) {
							var recipes = JSON.parse(data);
							jQuery.each(recipes, function (index, item) {
								$('<tr>\
										<td>\
											'+ item.ID + '\
										</td>\
										<td>\
											'+ item.Title + '\
										</td>\
										<td>\
											'+ item.Added + '\
										</td>\
										<td>\
											'+ item.Blog + '\
										</td>\
										<td>\
											'+ item.Instructions + '\
										</td>\
									</tr>\
								').appendTo('#recipeListTableBody');
							});
						}
					},
					error: function (xhr, resp, text) {
						console.log(xhr, resp, text);
					}
				});
			});
		});
});