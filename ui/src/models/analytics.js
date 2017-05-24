import m from "mithril";

export var Analytics = {
  Data:{}
}

Analytics.GetAnalytics = function(){
  return  m
    .request({
      method: 'GET',
      url: '/api/analytics',
    })
    .then(function(response) {
      Analytics.Data = response;
    })
    .catch(function(error) {
      console.error(error);
    });
}
