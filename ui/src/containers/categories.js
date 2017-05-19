import m from "mithril";
import {CategoriesModel} from "../models/categories.js";

var Categories = {
  SubmitNew:function(){
    var category = document.getElementById("categoryInput").value;
    console.log(category)
    CategoriesModel.AddCategory(category).then(function(){
      document.getElementById("categoryInput").value = ""
    })
  },
  oncreate:function(){
    CategoriesModel.GetCategories()
  },
  view:function(){

    var categories = CategoriesModel.Categories.map(function(category, key){
        return (<div class="pa2" key={key}>
        {category.Category}
      </div>);
      })

    return (
      <section>
        <div class="pa3 bg-white shadow-m2 tc">
          <h3>Categories</h3>
        </div>
        <div class="pa3 bg-white shadow-m2 mt3 cf">
            <input class="pa3 w-100 db border-box mb1" id="categoryInput"/>
            <button class="bg-navy white-80 shadow-xs hover-shadow-m1 pv3 ph4 fr" onclick={Categories.SubmitNew}>add</button>
        </div>
        <section class="pa3 bg-white shadow-m2 mt3 cf" >
          {categories}
        </section>
      </section>
    )
  }
}

export default Categories;
