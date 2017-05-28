import m from 'mithril';
import localforage from 'localforage';

export var UserModel = {
  User:{},
  Token:"",
  GetUserfromStorage: function(){
    if (!UserModel.User || !UserModel.User.Username){
      return localforage.getItem('AuthUser').then(function(user){
        console.log(user)
        if (user!=null){
          UserModel.User = user
          m.redraw()
          return
        }
        UserModel.User = null
        m.redraw()
      })
    }
  },
  GetTokenFromStorage:function(){
    console.log("get token from storage")
    if (!UserModel.Token){
      return localforage.getItem('jwtToken').then(function(token){
        console.log(token)
        if (token!=null){
          UserModel.Token = token
          m.redraw()
          return
        }
        UserModel.Token = null
        m.redraw()
      })

    }
  },
  Logout:function(){
    return localforage.removeItem('jwtToken').then(function(){
      return localforage.removeItem('jwtToken').then(function(){
        m.route.set("/login")
      })
    })
  }

}

export var UserLogin = {
  User: {},
  Submit: function() {
    UserLogin.User.Username = document.getElementById('username').value;
    UserLogin.User.Password = document.getElementById('password').value;

    console.log(UserLogin.User)
    m
      .request({
        method: 'POST',
        url: '/api/admin/login',
        data: UserLogin.User,
      })
      .then(function(response) {
        console.log(response);
          localforage.setItem('jwtToken', response.Token).then(function(){
          return localforage.setItem('AuthUser', response.User)
        }).then(function(){
          UserModel.GetTokenFromStorage().then(function(){
            return UserModel.GetUserfromStorage()
          }).then(function(){
            m.route.set("/")
          })
        })
      })
      .catch(function(error) {
        console.error(error);
      });
  },
};
