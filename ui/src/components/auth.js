import m from 'mithril';
import localforage from 'localforage';
import {UserModel} from '../models/userAuth.js';

export var AdminAuth = {
 oncreate:function(){
  UserModel.GetTokenFromStorage()
  UserModel.GetUserfromStorage()

 },
 view:function(vnode){
   console.log(UserModel.Token)
   if (UserModel.Token==null){
       m.route.set("/login")
       return m("div")
   }
   return m("div",vnode.attrs,m.fragment(vnode.attrs,[vnode.children]));
 }
}
