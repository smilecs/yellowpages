import m from "mithril";

var ListingsModel = {
  NewListing:function(listing){
    return m
      .request({
        method: 'POST',
        url: '/api/categories',
        data:listing,
      })
      .then(function(response) {
        console.log(response)

      })
      .catch(function(error) {
        console.error(error); 
      });
  }
}
