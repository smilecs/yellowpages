import m from "mithril";
import {ListingsModel} from '../models/listings.js';
import {ListingItem} from '../components/listingItem.js';

var UnApprovedListings = {
  oncreate:function(){
    ListingsModel.GetUnApprovedListings()
  },
  view:function(){
    return (
      <section>
        <div class="pa3 bg-white shadow-m2 tc">
          <div>
            <h3>Unapproved Listings</h3>
          </div>
        </div>
        <section class="pv3 ">
          {
          ListingsModel.UnApprovedListings.map(function(listing, key){
            return (<ListingItem listing={listing} key={key} />);
          })
        }
        </section>
      </section>
    )
  }
}

export default UnApprovedListings;
