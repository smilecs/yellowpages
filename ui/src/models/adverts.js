import m from "mithril";
export var AdvertsModel = {
  CurrentAdvert:{},
  Adverts:{},
  GetAdverts:function(){
    return  m
      .request({
        method: 'GET',
        url: '/api/adverts/all',
      })
      .then(function(response) {
        console.log(response)
        AdvertsModel.Adverts = response;
      })
      .catch(function(error) {
        console.error(error);
      });
  },
  AddAdvert:function(advert){
    return m
      .request({
        method: 'POST',
        url: '/api/adverts/new',
        data:advert,
      })
      .then(function(response) {
        AdvertsModel.Adverts.Posts.unshift({Type:"advert",Advert:advert})
      })
      .catch(function(error) {
        console.error(error);
      });
  }
}
