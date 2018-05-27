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
							var recipes = JSON.parse(data);
						}
						catch (err) {
							isJSON = false
						}
						if (isJSON == true) {
							jQuery.each(recipes, function (index, item) {
								$('<tr>\
										<td>\
											'+ item.id + '\
										</td>\
										<td>\
											'+ item.title + '\
										</td>\
										<td>\
											'+ item.added + '\
										</td>\
										<td>\
											'+ item.blog + '\
										</td>\
										<td>\
											'+ item.instructions + '\
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