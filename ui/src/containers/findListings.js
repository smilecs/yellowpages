import m from "mithril";
import {ListingsModel} from '../models/listings.js';

var FindListings = {
  Search:function(){
    var query = document.getElementById("SearchBox").value;
    ListingsModel.SearchForListings(query)
  },
  oncreate:function(){
    ListingsModel.SearchForListings()
  },
  view:function(){
    return (
      <section>
        <div class="pa3 bg-white shadow-m2">
          <div>
            <h3>Find Listings</h3>
          </div>
        </div>
        <div class="pa3 bg-white shadow-m2">
          <div>
            <input type="search" class="pa3 w-100" id="SearchBox"/>
            <button  class="bw0 shadow-4 pv2 ph3" onclick={FindListings.Search}>Search</button>
          </div>
        </div>
        <section class="pv3 ">
          {
            ListingsModel.SearchListings.map(function(listing, key){
            return (<div class=" dib w-100 w-50-ns  pa1 " key={key}>
              <div class=" bg-white shadow-m2 pa2">
                <div class="">
                  <h4>{listing.Listing.CompanyName}</h4>
                  <small>{listing.Listing.Specialisation}</small>
                  <p>{listing.Listing.About}</p>
                  <p>{listing.Listing.Hotline}</p>
                  <p>{listing.Listing.Address}</p>
                </div>
                <div class="tr">
                  <button class="pv2 ph3 shadow-4 bg-dark-red white bw0 white ma2 grow">Edit</button>
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

export default FindListings;
