import m from "mithril";

export var ListingsModel = {
  UnApprovedListings:[],
  SelectedLogo:"",
  SelectedListingImages:[],
  ShowFormSubmissionLoader:false,

  SearchListings:[],

  NewListing:function(listing){
    return m
      .request({
        method: 'POST',
        url: '/api/listings/add',
        data:listing,
      })
      .then(function(response) {
        console.log(response)

      })
      .catch(function(error) {
        console.error(error);
      });
  },
  GetUnApprovedListings:function(){
    return m
      .request({
        method: 'GET',
        url: '/api/listings/unapproved',
      })
      .then(function(response) {
        ListingsModel.UnApprovedListings = response.Data;
        console.log(response)
      })
      .catch(function(error) {
        console.error(error);
      });
  },
  SearchForListings:function(query){
    console.log('/api/search?q='+query)
    return m
      .request({
        method: 'GET',
        url: '/api/search?q='+query,
      })
      .then(function(response) {
        console.log(response)
        ListingsModel.SearchListings = response.Posts;
        console.log(response)
      })
      .catch(function(error) {
        console.error(error);
      });
  }
}
