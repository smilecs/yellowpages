import m from 'mithril';

import LoginPage from './containers/loginPage.js';
import AdminShell from './containers/adminShell.js';
import ListUsers from './containers/listUsers.js';
import Categories from './containers/categories.js';
import AddListing from './containers/addListing.js';

import {AdminAuth} from './components/auth.js';

var root = document.getElementById('appContainer');

m.route.prefix('/dashboard');
m.route(root, '/', {
  '/': {
        view: function(vnode) {
            return m(AdminAuth,vnode.attrs,
                m(AdminShell,vnode.attrs, m(ListUsers))
              );
        },
      },
  '/categories': {
        view: function(vnode) {
            return m(AdminAuth,vnode.attrs,
                m(AdminShell,vnode.attrs, m(Categories))
              );
        },
      },
  '/listings/new': {
        view: function(vnode) {
            return m(AdminAuth,vnode.attrs,
                m(AdminShell,vnode.attrs, m(AddListing))
              );
        },
      },
  '/login': LoginPage,
});
