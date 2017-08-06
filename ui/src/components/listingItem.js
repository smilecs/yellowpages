import m from "mithril";
import {ListingsModel} from '../models/listings.js';
import iziToast from "iziToast";

export var ListingItem = {
  ApproveListing:function(slug){
    ListingsModel.ApproveListing(slug).then(function(){
      iziToast.success({
        position:"topRight",
        title:"Success",
        message:"Approved Listing successfully"
      })
    }).catch(function(err){
      console.log(err)
      iziToast.error({
        position:"topRight",
        title:"Error",
        message:"Unable to approve listing"
      })
    })
  },
  DeleteListing:function(slug){
    ListingsModel.ApproveListing(slug).then(function(){
      iziToast.success({
        position:"topRight",
        title:"Success",
        message:"Deleted Listing successfully"
      })
    }).catch(function(err){
      console.log(err)
      iziToast.error({
        position:"topRight",
        title:"Error",
        message:"Unable to delete listing"
      })
    })
  },
    view:function(vnode){
      var {listing, key} = vnode.attrs;
      return (<div class=" dib w-100 w-50-ns  pa1 " key={key}>
        <div class=" bg-white shadow-m2 pa2">
          <div class="">
            <h4>{listing.CompanyName}</h4>
            <small>{listing.Specialisation}</small>
            <p>{listing.About}</p>
            <p>{listing.Hotline}</p>
            <p>{listing.Address}</p>
          </div>
          <div class="tr">
            <button class="dib pv2 ph3 shadow-4  bw0 ma2 grow link black bg-light-silver pointer" onclick={()=>ListingItem.DeleteListing(listing.Slug)}>Delete</button>
            <button class="dib pv2 ph3 shadow-4  bw0 ma2 grow link black bg-light-silver pointer" onclick={()=>ListingItem.ApproveListing(listing.Slug)}>Approve</button>
            <a class="dib pv2 ph3 shadow-4  bw0 ma2 grow link black bg-light-silver" oncreate={m.route.link} href={"/listings/edit/"+listing.Slug} >Edit</a>
          </div>
        </div>
      </div>)
    }
}
