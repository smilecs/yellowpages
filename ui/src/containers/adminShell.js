import m from "mithril";
import localforage from 'localforage';
import {UserModel} from '../models/userAuth.js';

var AdminShell = {
  oncreate:function(){

  },
  view:function(vnode){
	return (
	  <section>
		<section class="pt4 pb5 ph5 black-80 bg-yellow">
		  <div class="pv3">
			<div class="fr ">
			  <span>20th May, 2016</span>
			</div>
			<div>
			  <img src="/assets/img/logo-white.png"  class="h2 h3-ns"/>
			</div>
		  </div>

		  <div class="cf ">
			<div class="dib pt4">
			  <h3 class="db mv1">
				      {UserModel.User.Name} (<small>{UserModel.User.Username}</small>)
			  </h3>

			  <small>
				      {UserModel.User.Email}
			  </small>
			</div>
			<div class="fr ">
			  <div class="pa4 shadow-m2 br1">
				<div class="tc dib ph3">
				  <span class="db">
					22
				  </span>
				  <span class="db">
					Credits
				  </span>
				</div>
				<div class="tc dib ph3">
				  <span class="db">
					22
				  </span>
				  <span class="db">
					Credits
				  </span>
				</div>
				<div class="tc dib ph3">
				  <span class="db">
					22
				  </span>
				  <span class="db">
					Credits
				  </span>
				</div>
			  </div>
			</div>
		  </div>

		</section>

		<section style="margin-top:-2.5rem;" class="ph5">
		  <section class="dib w-30-ns ph3 v-top">
			<div class="bg-white shadow-m2 pa3 br1">
			  <div class="">
          <a class="db pa2 bb b--light-gray hover-bg-light-gray link" href="/" oncreate={m.route.link}>Dashboard</a>
          <a class="db pa2 bb b--light-gray hover-bg-light-gray link" href="/categories" oncreate={m.route.link}>Categories</a>
          <a class="db pa2 bb b--light-gray hover-bg-light-gray link">Users</a>
          <a class="db pa2 bb b--light-gray hover-bg-light-gray link">Admins</a>
          <div class=" bb b--light-gray">
            <a class="db pa2  link">Listings</a>
            <a class="db pa2  hover-bg-light-gray link">&gt; Add Listing</a>
            <a class="db pa2 hover-bg-light-gray link"> &gt; Unapproved Listings</a>
            <a class="db pa2 hover-bg-light-gray link">&gt; All Listings</a>
          </div>
			  </div>
			</div>
		  </section><section class="dib w-70-ns ph3 v-top br1">
      {m.fragment(vnode.attrs,[vnode.children])}
		  </section>
		</section>
	  </section>
	)
  }
}
export default AdminShell;
