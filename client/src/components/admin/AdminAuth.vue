<template>
  <div>
    <div class="container-fluid h-custom" style="height: 80vh;">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col-md-9 col-lg-6 col-xl-5">
          <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp"
               class="img-fluid" alt="Sample image">
        </div>
        <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
          <form @submit.prevent="onSubmitLoginForm">
            <div class="divider d-flex align-items-center my-4">
              <h1 class="text-center fw-bold mx-3 mb-0">Sign In</h1>
            </div>
            <!-- Email input -->
            <div class="form-outline mb-4">
              <label class="form-label" for="form3Example3">Username</label>
              <input v-model="loginForm.username" type="text" id="form3Example3" class="form-control form-control-lg"
                     placeholder="Your username"/>
            </div>

            <!-- Password input -->
            <div class="form-outline mb-3">
              <label class="form-label" for="form3Example4">Password</label>
              <input v-model="loginForm.password" type="password" id="form3Example4"
                     class="form-control form-control-lg"
                     placeholder="Enter password"/>
            </div>

            <div class="d-flex justify-content-between align-items-center">
              <!-- Checkbox -->
              <a href="#!" class="text-body">Forgot password?</a>
            </div>

            <div class="text-center text-lg-start mt-4 pt-2">
              <button type="submit" class="btn btn-primary btn-lg"
                      style="padding-left: 2.5rem; padding-right: 2.5rem;">Login
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
    <Footer/>
  </div>
</template>

<script>
import {RequestParams} from '../../config/request'
import Footer from '../layouts/Footer.vue'
import auth, { AxiosInstance } from '../../config/auth'
import Vue from 'vue'

export default {
  name: 'AdminAuthentication',
  components: {
    Footer
  },
  data () {
    return {
      loginForm: {
        username: '',
        password: ''
      }
    }
  },
  mixins: [auth],
  methods: {
    async onSubmitLoginForm () {
      await AxiosInstance.post(RequestParams.host + RequestParams.path.adminLogin, this.loginForm).then(data => {
        // axios.defaults.headers.common['Authorization'] = `Bearer ${data.data.data.access_token}`
        AxiosInstance.defaults.headers.common['Authorization'] = `Bearer ${data.data.data.access_token}`
        Vue.prototype.$globalData = {
          userGlobal: data.data.data.data,
          accessToken: data.data.data.access_token
        }
        Vue.prototype.$globalData.userGlobal.full_name = data.data.data.data.username
        console.log(Vue.prototype.$globalData)
        if (data.data.data.data.id) {
          this.$router.push('/admin/home')
        }
      }).catch(err => {
        console.log(err)
        throw (err)
      })
    }
  }
}
</script>

<style scoped>
.divider:after,
.divider:before {
  content: "";
  flex: 1;
  height: 1px;
  background: #eee;
}

.h-custom {
  height: calc(100% - 73px);
}

@media (max-width: 450px) {
  .h-custom {
    height: 100%;
  }
}
</style>
