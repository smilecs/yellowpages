import m from 'mithril';
import {UserLogin} from '../models/userAuth.js';

var LoginPage = {
  view: function() {
    return (
      <section class="vh-100 bg-near-white flex flex-row justify-center items-center">
        <div class="dib bg-white w-90 w-50-m w-40-l shadow-m5">
          <div class="pa4 gray">
            <h2>LOGIN</h2>
            <div class="pv2">
              <label class="dn" for="email">EMAIL</label>
              <input
                type="text"
                placeholder="EMAIL"
                class="pa3 bw0 bb w-100 input-reset shadow-0 b--light-gray"
                id="username"
                style="border-bottom-width:2px;"
              />
            </div>
            <div class="pv2">
              <label class="dn" for="email">PASSWORD</label>
              <input
                type="password"
                placeholder="PASSWORD"
                class="pa3 bw0 bb w-100 input-reset shadow-0 b--light-gray"
                id="password"
                style="border-bottom-width:2px;"
              />
            </div>
          </div>
          <div>
            <button
              class="pa3 w-50 dib btn-reset bn  bg-navy white-80 fr"
              onclick={UserLogin.Submit}
            >
              SIGN IN
            </button>
          </div>

        </div>
      </section>
    );
  },
};

export default LoginPage;
