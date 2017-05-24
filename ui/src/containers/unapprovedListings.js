import m from "mithril";
import {ListingsModel} from '../models/listings.js';

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
          {ListingsModel.UnApprovedListings.map(function(listing, key){
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
                  <button class="pv2 ph3 shadow-4 bg-dark-red white bw0 white ma2 grow">Delete</button>
                  <button class="pv2 ph3 shadow-4 bg-dark-green white bw0 ma2 grow">Approve</button>
                </div>
              </div>
            </div>);
          })

        }
        </section>
      </section>
    )
  }
}

export default UnApprovedListings;
