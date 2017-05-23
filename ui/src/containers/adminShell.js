import m from "mithril";
import localforage from 'localforage';
import {UserModel} from '../models/userAuth.js';

var MenuComponent = {
  view:function(){
    return (
      <div class="">
        <a class="db pa2 bb b--light-gray hover-bg-light-gray link navy" href="/" oncreate={m.route.link}>Dashboard</a>
        <a class="db pa2 bb b--light-gray hover-bg-light-gray link navy" href="/categories" oncreate={m.route.link}>Categories</a>
        <div class=" bb b--light-gray">
          <a class="db pa2  link navy">Listings</a>
          <a class="db pa2  hover-bg-light-gray link navy" href="/listings/new" oncreate={m.route.link}>&gt; Add Listing</a>
          <a class="db pa2 hover-bg-light-gray link navy" href="/listings/unapproved" oncreate={m.route.link}> &gt; Unapproved Listings</a>
          <a class="db pa2 hover-bg-light-gray link navy" href="/listings/find" oncreate={m.route.link}>&gt; Find Listings</a>
        </div>
      </div>
    );
  }
}

var AdminShell = {
  fixNav:false,
  oncreate:function(vnode){
    var navBar = document.getElementById("fixedNav")
    console.log(navBar.offsetTop)

    var navBarOffset = navBar.offsetTop;

    var last_known_scroll_position = 0;
    var ticking = false;

    function CheckPostionAndUpdateNavClass(scroll_pos) {
      // do something with the scroll position
      console.log(scroll_pos)
      console.log(navBarOffset)

      if (scroll_pos>navBarOffset){
        console.log("fixed")
        vnode.state.fixNav = true
        m.redraw()
      }else{
        vnode.state.fixNav = false
        m.redraw()
      }

    }

    window.addEventListener('scroll', function(e) {
      last_known_scroll_position = window.scrollY;
      if (!ticking) {
        window.requestAnimationFrame(function() {
          CheckPostionAndUpdateNavClass(last_known_scroll_position);
          ticking = false;
        });
      }
      ticking = true;
    });
  },
  view:function(vnode){
    console.log(vnode.state.fixNav)
	return (
	  <section>
  		<section class="  pt4-ns  pb5-ns  ph5-ns black-80 bg-yellow">
  		  <div class={"pa2 pv3-ns relative-ns w-100  z-5 "+(vnode.state.fixNav===true?"fixed bg-yellow shadow-4":"")} id="fixedNav">
    			<div class="dib relative">
            <a href="#" class="dib black link v-mid mr3  pa2 ba relative" onclick={()=>vnode.state.showNav=!vnode.state.showNav}>☰</a>
              <div class={" right-0 buttom-0 absolute bg-white shadow-m2 pa3 br1 "+(vnode.state.showNav?"db":"dn")}>
                <MenuComponent />
              </div>

    			  <span class="f3 dib v-mid">Calabar<strong>Pages</strong></span>
    			</div>
          <div class="dib v-mid pv2 fr ">
            <div class="dib">
      			  <span class="dib v-mid">{UserModel.User.Username}</span>
              <small class="dib v-mid ph2" style="font-size:8px;">▼</small>
            </div>
    			</div>
  		  </div>
  		  <div class={"cf pa3 "+(vnode.state.fixNav===true?"pt5":"")}>
    			<div class="tc fr-ns ">
    			  <div class="pa2 pa4-ns shadow-m2 br1 mv3">
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


  		<section  class="ph2 ph5-ns mt-m25-ns ">
  		  <section class="dib w-30-ns ph3-ns v-top">
  			<div class="bg-white shadow-m2 pa3 br1 dn dib-ns">
          <MenuComponent />
  			</div>
      </section><section class="dib w-70-ns ph3-ns v-top br1">
        {m.fragment(vnode.attrs,[vnode.children])}
  		  </section>
  		</section>
	  </section>
	)
  }
}
export default AdminShell;
