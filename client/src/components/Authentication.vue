<template>
  <div>
    <div class="container-fluid h-custom tmp-body" v-if="showSignin">
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
              <label class="form-label" for="form3Example3">CCCD ID</label>
              <input v-model="loginForm.cccd_id" type="text" id="form3Example3" class="form-control form-control-lg"
                     placeholder="CCCD ID"/>
            </div>

            <!-- Password input -->
            <div class="form-outline mb-3">
              <label class="form-label" for="form3Example4">Password</label>
              <input v-model="loginForm.password" type="password" id="form3Example4"
                     class="form-control form-control-lg"
                     placeholder="Enter password"/>
            </div>
            <div v-if="error" class="alert alert-danger">
              {{ error }}
            </div>
            <div class="d-flex justify-content-between align-items-center">
              <!-- Checkbox -->
              <div class="form-check mb-0">
                <input class="form-check-input me-2" type="checkbox" value="" id="form2Example3"/>
                <label class="form-check-label" for="form2Example3">
                  Remember me
                </label>
              </div>
              <a href="#!" class="text-body">Forgot password?</a>
            </div>

            <div class="text-center text-lg-start mt-4 pt-2">
              <button type="submit" class="btn btn-primary btn-lg"
                      style="padding-left: 2.5rem; padding-right: 2.5rem;">Login
              </button>
              <p class="small fw-bold mt-2 pt-1 mb-0">Don't have an account?
                <button @click="handleSignupEvent"
                        class="btn btn-light">Register
                </button>
              </p>
            </div>
          </form>
        </div>
      </div>
    </div>
    <div class="container-fluid h-custom" v-if="showSignup">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col-md-9 col-lg-6 col-xl-5">
          <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp"
               class="img-fluid" alt="Sample image">
        </div>
        <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
          <form @submit.prevent="onSubmitRegisterForm">
            <div class="divider d-flex align-items-center my-4">
              <h1 class="text-center fw-bold mx-3 mb-0">Sign Up</h1>
            </div>
            <div class="d-flex">

              <!-- Full name input -->
              <div class="form-outline me-2 w-50">
                <label class="form-label" for="signupForm.fullName">Full Name</label>
                <input
                  type="text"
                  id="signupForm.fullName"
                  class="form-control form-control-lg"
                  v-model="signUpForm.full_name"
                  placeholder="Your Full Name"
                >
              </div>

              <!-- Date input -->
              <div class="form-outline w-50">
                <label class="form-label" for="signupForm.dateOfBirth">Date of Birth</label>
                <input
                  type="date"
                  id="signupForm.dateOfBirth"
                  class="form-control form-control-lg"
                  v-model="signUpForm.date_of_birth"
                >
              </div>
            </div>
            <!-- Select Gender input-->
            <label class="form-label" for="signupForm.gender">Gender</label>
            <div id="signupForm.gender">
              <div class="form-check form-check-inline">
                <input class="form-check-input" type="radio" id="genderMale"
                       v-model="signUpForm.gender" value="male">
                <label class="form-check-label" for="genderMale">Male</label>
              </div>

              <div class="form-check form-check-inline">
                <input class="form-check-input" type="radio" id="genderFemale"
                       v-model="signUpForm.gender" value="female">
                <label class="form-check-label" for="genderFemale">Female</label>
              </div>

              <div class="form-check form-check-inline">
                <input class="form-check-input" type="radio" id="genderOther"
                       v-model="signUpForm.gender" value="other">
                <label class="form-check-label" for="genderOther">Other</label>
              </div>
            </div>
            <!--             Select Place -->
            <label class="form-label" for="signupForm.address">Address</label>
            <PlaceSelection id="signupForm.address"
                            @place-selected="handlePlaceSelected"
            />

            <div class="d-flex">
              <div class="form-outline w-50">
                <label class="form-label" for="signupForm.cccd">CCCD</label>
                <input
                  type="text"
                  id="signupForm.cccd"
                  class="form-control form-control-lg"
                  v-model="signUpForm.cccd_id"
                  placeholder="Your CCCD Id"
                >
              </div>
              <div class="form-outline w-50">
                <label class="form-label" for="signupForm.cccd_date">Date</label>
                <input
                  type="date"
                  id="signupForm.cccd_date"
                  class="form-control form-control-lg"
                  v-model="signUpForm.cccd_date"
                >
              </div>
              <div class="form-outline w-50">
                <label class="form-label" for="signupForm.cccd_place">Issued By</label>
                <input
                  type="text"
                  id="signupForm.cccd_place"
                  class="form-control form-control-lg"
                  placeholder="Your CCCD Issuer"
                  v-model="signUpForm.cccd_place"
                >
              </div>
            </div>
            <div class="form-outline mb-4">
              <label class="form-label" for="signupForm.email">Email address</label>
              <input
                type="email"
                id="signupForm.email"
                class="form-control form-control-lg"
                v-model="signUpForm.email"
                placeholder="Enter a valid email address"
              >
              <div class="form-outline mb-4">
                <label class="form-label" for="signupForm.phone">Phone number</label>
                <input
                  type="text"
                  id="signupForm.phone"
                  class="form-control form-control-lg"
                  v-model="signUpForm.phone"
                  placeholder="Enter a valid phone number"
                >
              </div>
              <div class="form-outline mb-4">
                <label class="form-label" for="signupForm.password">Password</label>
                <input
                  type="password"
                  id="signupForm.password"
                  class="form-control form-control-lg"
                  v-model="signUpForm.password"
                  placeholder="Your password"
                >
              </div>
            </div>

            <div class="d-flex justify-content-between align-items-center">
              <!-- Checkbox -->
              <div class="form-check mb-0">
                <input class="form-check-input me-2" type="checkbox" value="" id="form31Example3"/>
                <label class="form-check-label" for="form31Example3">
                  Accept all terms and conditions
                </label>
              </div>
            </div>
            <div class="text-center text-lg-start mt-4 pt-2">
              <button type="submit" class="btn btn-primary btn-lg"
                      style="padding-left: 2.5rem; padding-right: 2.5rem;" @click="showModal">Register
              </button>
              <p class="small fw-bold mt-2 pt-1 mb-0">Already have account?
                <button @click="handleSigninEvent"
                        class="btn btn-light">Sign In
                </button>
              </p>
            </div>
          </form>
          <b-modal ref="myModalRef" hide-footer title="Identity Verification"
                   centered hide-backdrop
                   header-bg-variant="dark"
                   header-text-variant="light" @hide="null">
            <div class="d-block text-center">
              <h3>The user should wait for Identity verification, we will notify you until verified</h3>
            </div>
          </b-modal>
        </div>
      </div>
    </div>
    <Footer/>
  </div>
</template>

<script>
import {RequestParams} from '../config/request'
import PlaceSelection from './utils/PlaceSelect.vue'
import Footer from './layouts/Footer.vue'
import auth, { AxiosInstance } from '../config/auth'

export default {
  name: 'Authentication',
  components: {
    Footer,
    'PlaceSelection': PlaceSelection
  },
  data () {
    return {
      error: null,
      showSignin: true,
      showSignup: false,
      loginForm: {
        cccd_id: '',
        password: ''
      },
      signUpForm: {
        full_name: '',
        gender: '',
        date_of_birth: null,
        address: '',
        cccd_id: '',
        cccd_date: null,
        cccd_place: '',
        email: '',
        phone: '',
        password: ''
      }
    }
  },
  mixins: [auth],
  methods: {
    handleSigninEvent () {
      this.showSignin = true
      this.showSignup = false
    },
    handleSignupEvent () {
      this.showSignin = false
      this.showSignup = true
    },
    showModal () {
      this.$refs.myModalRef.show()
    },
    handlePlaceSelected (place) {
      this.signUpForm.address = `${place.ward.Name} - ${place.district.Name} - ${place.city.Name}`
    },
    async onSubmitLoginForm () {
      if (this.loginForm.cccd_id !== '') {
        await AxiosInstance.post(RequestParams.host + RequestParams.path.login, this.loginForm).then(data => {
          AxiosInstance.defaults.headers.common['Authorization'] = `Bearer ${data.data.data.access_token}`
          if (data.data.data.data.id) {
            if (data.data.data.type === 'voter') {
              this.$router.push({
                path: '/voters/home',
                params: data.data.data
              })
            } else if (data.data.data.type === 'candidate') {
              this.$router.push({
                path: '/candidates/home',
                params: data.data.data
              })
            }
            sessionStorage.setItem('userGlobal', JSON.stringify(data.data.data.data))
            sessionStorage.setItem('accessToken', data.data.data.access_token)
            console.log(data.data.data.access_token)
          }
        }).then(async () => {
          await AxiosInstance.get(RequestParams.host + RequestParams.path.check_election,
            {
              headers: {
                Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
              }
            }).then(data => {
            sessionStorage.setItem('electionId', data.data.data.data.election_id)
            sessionStorage.setItem('hasElection', data.data.data.data.has_election)
            console.log(sessionStorage.getItem('accessToken'))
          })
        })
          .catch(err => {
            console.log(err)
            this.handleError(err.response.data.error)
          })
      }
    },
    async onSubmitRegisterForm () {
      if (this.signUpForm.cccd_id !== '') {
        this.signUpForm.date_of_birth = new Date(this.signUpForm.date_of_birth).toISOString()
        this.signUpForm.cccd_date = new Date(this.signUpForm.cccd_date).toISOString()
        console.log(this.signUpForm)
        try {
          const res = await AxiosInstance.post(RequestParams.host + RequestParams.path.register, this.signUpForm)
          console.log(res)
        } catch (e) {
          console.error(e)
        }
      }
    },
    handleError (err) {
      this.error = err
    }
  }
}
</script>

<style scoped>
.tmp-body{
  height: 80vh !important;
}
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
