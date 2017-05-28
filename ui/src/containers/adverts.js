import m from "mithril";
import {CategoriesModel} from "../models/categories.js";
import {AdvertsModel} from "../models/adverts.js";
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
       AdvertsModel.CurrentAdvert.Image = e.target.result;
       m.redraw();
     };

    reader.readAsDataURL(file);

}

var Adverts = {
  SubmitNew:function(){

    AdvertsModel.AddAdvert(AdvertsModel.CurrentAdvert).then(function(){
      AdvertsModel.CurrentAdvert = {}
      iziToast.success({
        position:"topRight",
        title:"Success",
        message:"Added advert successfully"
      })
    })
  },
  oncreate:function(){
    AdvertsModel.GetAdverts()
  },
  view:function(){

    var adverts = AdvertsModel.Adverts.Posts?AdvertsModel.Adverts.Posts.map(function(item, key){
        if(item.Type==="advert"){
          let ad = item.Advert;
          return (<div class=" dib w-100 w-50-ns  pa1 v-top" key={key}>
            <h4 class="fw4">{ad.Name}</h4>
            <img src={ad.Image} class="w-100" />
          </div>);
        }else{
          return
        }
      }):[]

    return (
      <section>
        <div class="pa3 bg-white shadow-m2 tc">
          <h3>Adverts</h3>
        </div>
        <div class="pa3 bg-white shadow-m2 mt3 cf">
            <input class="pa3 w-100 db border-box mb1" placeholder="Advert Title" oninput={m.withAttr("value",(update)=>{
                AdvertsModel.CurrentAdvert.Name = update;
              })}/>
            <div class="pv3">
                <label for="LogoImage" class="fw6">Logo Image</label>
                <input id="LogoImage" type="file" class="w-100 pv2 ph3 mt2" aria-invalid="false" onchange={handleLogoChange}/>
                <img class="w4" src={AdvertsModel.CurrentAdvert.Image} oninput={m.withAttr("value",(update)=>{AdvertsModel.CurrentAdvert.Image = update})}/>
            </div>
            <button class="bg-navy white-80 shadow-xs hover-shadow-m1 pv3 ph4 fr" onclick={Adverts.SubmitNew}>add</button>
        </div>
        <section class="pa3 bg-white shadow-m2 mt3 cf" >
          {adverts}
        </section>
      </section>
    )
  }
}

export default Adverts;
