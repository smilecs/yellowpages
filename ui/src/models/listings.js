import m from "mithril";

export var ListingsModel = {
  UnApprovedListings:[],
  ShowFormSubmissionLoader:false,
  SearchListings:[],
  SearchListingsPagination:{},
  CurrentListing:{Image:"",Images:[]},

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
  EditListing:function(slug,listing){
    console.log(slug)
    console.log(listing)
    return m
      .request({
        method: 'POST',
        url: '/api/listings/edit/'+slug,
        data:listing,
      })
      .then(function(response) {
        console.log(response)

      })
      .catch(function(error) {
        console.error(error);
      });
  },
  GetListing:function(slug){
    return m
      .request({
        method: 'GET',
        url: '/api/listings/slug/'+slug,
      })
      .then(function(response) {
        console.log(response)
        ListingsModel.CurrentListing = response
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
  ApproveListing:function(slug){
    return m.request({
      method:'GET',
      url: '/api/listings/approve?q='+ slug,
    }).then(function(response) {
      console.log(response)
      var element = ListingsModel.UnApprovedListings.find((listing)=>{
        return listing.Slug == slug;
      })
      var index = ListingsModel.UnApprovedListings.indexOf(element)
      ListingsModel.UnApprovedListings.splice(index,1)
    })
  },
  DeleteListing:function(slug){
    return m.request({
      method:'GET',
      url: '/api/listings/delete?q='+ slug,
    }).then(function(response) {
      console.log(response)
      var element = ListingsModel.UnApprovedListings.find((listing)=>{
        return listing.Slug == slug;
      })
      var index = ListingsModel.UnApprovedListings.indexOf(element)
      ListingsModel.UnApprovedListings.splice(index,1)
    })
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
        ListingsModel.SearchListingsPagination = response.Page;
      })
  },
  SearchLoadmore:function(){
    console.log(ListingsModel.SearchListingsPagination)
    return m
      .request({
        method: 'GET',
        url: ListingsModel.SearchListingsPagination.NextURL,
      })
      .then(function(response) {
        console.log(response)
        console.log(ListingsModel.SearchListings.length)
        ListingsModel.SearchListings = ListingsModel.SearchListings.concat(response.Posts);
        console.log(ListingsModel.SearchListings.length)
        ListingsModel.SearchListingsPagination = response.Page;
        m.redraw()
      })
  }
}
