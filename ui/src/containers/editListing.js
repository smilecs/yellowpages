import m from "mithril";
import {CategoriesModel} from "../models/categories.js";
import {ListingsModel} from "../models/listings.js";
import iziToast from 'iziToast';

function handleLogoChange(e){
    console.log(e)

    var file = e.target.files[0]

    var imageType = /^image\//;
    if (!imageType.test(file.type)) {
      return;
    }

    var reader = new FileReader();
    reader.onload = function(e) {
       ListingsModel.CurrentListing.Image = e.target.result;
       m.redraw();
     };

    reader.readAsDataURL(file);
}

function handleFilesChange(e){
    console.log(e)

    var files = e.target.files

    for (var i = 0; i < files.length; i++) {
      var file = files[i];
      var imageType = /^image\//;

      if (!imageType.test(file.type)) {
        continue;
      }



      var reader = new FileReader();
      reader.onload = function(e) {
         ListingsModel.CurrentListing.Images.unshift(e.target.result)
         m.redraw()
       };
      reader.readAsDataURL(file);
  }
}


var EditListing = {
  Listing:{},
  Submit:function(){
    console.log(ListingsModel.CurrentListing)

    ListingsModel.EditListing(m.route.param("slug"),ListingsModel.CurrentListing).then((message)=>{
      console.log(message)

      iziToast.success({
        position:"topRight",
        title:"Success",
        message:"Added Listing for `"+ListingsModel.CurrentListing.CompanyName+"` successfully"
      })
      ListingsModel.CurrentListing = {}
      m.redraw()
    })
  },
  oncreate:function(){
    ListingsModel.GetListing(m.route.param("slug"))
    CategoriesModel.GetCategories()
  },
  view:function(){
    let {CurrentListing} = ListingsModel;
    return (
      <section>
        <div class="pa3 bg-white shadow-m2 tc">
          <h3>Edit Listing</h3>
        </div>

        <section class="pa3 pa4-ns bg-white shadow-m2 mt3 cf">
          <div class="cf">
            <div class="pv3 w-100 w-50-ns fl">
                <label for="Category" class="fw6">Category</label>
                <select class="w-100 mt2" id="Category" onchange={m.withAttr("value",(update)=>{CurrentListing.Category = update})}>
                  {
                    CategoriesModel.Categories.map((category,key)=>{
                      return (<option value={category.Slug} key={key} selected={category.Slug==CurrentListing.Category?"selected":""}>{category.Category}</option>)
                    })
                  }
                </select>
            </div>

            <div class="pv3 w-100 w-50-ns  fl">
                <label for="CompanyName" class="fw6">Company Name</label>
                <input id="CompanyName" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.CompanyName} oninput={m.withAttr("value",(update)=>{CurrentListing.CompanyName = update})}/>
            </div>
          </div>
          <div class="pv3">
              <label for="PhysicalAddress" class="fw6">Physical Address</label>
              <input id="PhysicalAddress" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.Address} oninput={m.withAttr("value",(update)=>{CurrentListing.Address = update})}/>
          </div>
          <div class=" cf">
            <div class="pv3 w-100 w-50-ns fl">
                <label for="Hotline" class="fw6">Hotline</label>
                <input id="Hotline" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.Hotline} oninput={m.withAttr("value",(update)=>{CurrentListing.Hotline = update})}/>
            </div>
            <div class="pv3 w-100 w-50-ns fl">
                <label for="Email" class="fw6">Email</label>
                <input id="Email" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.Email} oninput={m.withAttr("value",(update)=>{CurrentListing.Email = update})}/>
            </div>
          </div>
          <div class="cf">
            <div class="pv3 w-100 w-50-ns fl dib v-btm">
                <label for="WorkingDaysAndTimes" class="fw6">Working Days/time (eg. Mon-Fri 8am-8pm)</label>
                <input id="WorkingDaysAndTimes" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.DHr} oninput={m.withAttr("value",(update)=>{CurrentListing.DHr = update})}/>
            </div>
            <div class="pv3 w-100 w-50-ns fl dib v-btm">
                <label for="Specialization" class="fw6">Specialization</label>
                <input id="Specialization" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.Specialisation} oninput={m.withAttr("value",(update)=>{CurrentListing.Specialisation = update})}/>
            </div>
          </div>
          <div class="pv3">
              <label for="Type" class="fw6">Type:&nbsp;&nbsp;&nbsp;</label>
              <select class="dib mt2" id="Type" aria-invalid="false" onchange={m.withAttr("value",(update)=>{CurrentListing.Plus = update})}>
                <option value="false" selected={CurrentListing.Plus=="false"?"selected":""}>Listing</option>
                <option value="true" selected={CurrentListing.Plus=="true"?"selected":""}>PlusListing</option>
              </select>
          </div>
          <div class="pv3">
              <label for="About" class="fw6">About</label>
              <textarea id="About"  class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.About} oninput={m.withAttr("value",(update)=>{CurrentListing.About = update})}></textarea>
          </div>
          <div class="pv3">
              <label for="WebsiteLink" class="fw6">Website Link</label>
              <input id="WebsiteLink" type="text" class="w-100 pv2 ph3 mt2" aria-invalid="false" value={CurrentListing.Website} oninput={m.withAttr("value",(update)=>{CurrentListing.Website = update})}/>
          </div>
          <div class="pv3">
              <label for="LogoImage" class="fw6">Logo Image</label>
              <input id="LogoImage" type="file" class="w-100 pv2 ph3 mt2" aria-invalid="false" onchange={handleLogoChange}/>
              <img class="w4" src={CurrentListing.Image} oninput={m.withAttr("value",(update)=>{CurrentListing.Image = update})}/>
          </div>
          <div class="pv3">
              <label for="Images" class="fw6">Images</label>
              <input id="Images" type="file" class="w-100 pv2 ph3 mt2" aria-invalid="false" onchange={handleFilesChange} multiple/>
              {
                CurrentListing.Images?CurrentListing.Images.map(function(image){
                  return(<img class="w4" src={image}/>)
                }):""
              }
          </div>
          <button type="button" class="white-80 shadow-4 grow bg-black dim pa3 fr ba0" onclick={EditListing.Submit} >Submit</button>

            <div class="tc" aria-hidden="true" class={ListingsModel.ShowFormSubmissionLoader?"db":"dn"}>
              <img src="/assets/ripple.gif" class="dib"/>
            </div>
        </section>
      </section>
    )
  }
}

export default EditListing;
