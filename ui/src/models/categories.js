import m from "mithril";
export var CategoriesModel = {
  Categories:[],
  GetCategories:function(){
    return  m
      .request({
        method: 'GET',
        url: '/api/categories',
      })
      .then(function(response) {
        CategoriesModel.Categories = response;
      })
      .catch(function(error) {
        console.error(error);
      });
  },
  AddCategory:function(category){
    var data  = {Category:category}
    return m
      .request({
        method: 'POST',
        url: '/api/categories',
        data:data,
      })
      .then(function(response) {
        CategoriesModel.Categories.unshift(data)

      })
      .catch(function(error) {
        console.error(error);
      });
  }
}
