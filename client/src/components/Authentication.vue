<template>
  <div>
    <div class="container-fluid h-custom" v-if="showSignin">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col-md-9 col-lg-6 col-xl-5">
          <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp"
               class="img-fluid" alt="Sample image">
        </div>
        <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
          <form>
            <div class="divider d-flex align-items-center my-4">
              <p class="text-center fw-bold mx-3 mb-0">Sign In</p>
            </div>

            <!-- Email input -->
            <div class="form-outline mb-4">
              <label class="form-label" for="form3Example3">CCCD ID</label>
              <input type="text" id="form3Example3" class="form-control form-control-lg"
                     placeholder="CCCD ID"/>
            </div>

            <!-- Password input -->
            <div class="form-outline mb-3">
              <label class="form-label" for="form3Example4">Password</label>
              <input type="password" id="form3Example4" class="form-control form-control-lg"
                     placeholder="Enter password"/>
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
              <p class="small fw-bold mt-2 pt-1 mb-0">Don't have an account? <button @click="handleSignupEvent"
                                                                                class="btn btn-light">Register</button></p>
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
          <form>
            <div class="divider d-flex align-items-center my-4">
              <p class="text-center fw-bold mx-3 mb-0">Sign Up</p>
            </div>
            <!-- Full name input -->
            <div class="form-outline mb-4">
              <label class="form-label" for="signup-fullName-inp">Full Name</label>
              <input type="text" id="signup-fullName-inp" class="form-control form-control-lg"
                     placeholder="Your Full Name"/>
            </div>
            <!-- Cccd Id input -->
            <div class="form-outline mb-4">
              <label class="form-label" for="signup-cccd-inp">CCCD</label>
              <input type="text" id="signup-cccd-inp" class="form-control form-control-lg"
                     placeholder="Your CCCD Id"/>
            </div>
            <!-- Email input -->
            <div class="form-outline mb-4">
              <label class="form-label" for="form3Example3">Email address</label>
              <input type="email" id="form3Example3" class="form-control form-control-lg"
                     placeholder="Enter a valid email address"/>
            </div>
            <div class="form-outline mb-4">
              <label class="form-label" for="signup-phone">Phone Number</label>
              <input type="text" id="signup-phone" class="form-control form-control-lg"
                     placeholder="Phone Number"/>
            </div>

            <!-- Password input -->
            <div class="form-outline mb-3">
              <label class="form-label" for="form3Example4">Password</label>
              <input type="password" id="form3Example4" class="form-control form-control-lg"
                     placeholder="Enter password"/>
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
                      style="padding-left: 2.5rem; padding-right: 2.5rem;">Register
              </button>
              <p class="small fw-bold mt-2 pt-1 mb-0">Already have account? <button @click="handleSigninEvent"
                                                                                class="btn btn-light">Sign In</button></p>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Authentication',
  data () {
    return {
      showSignin: true,
      showSignup: false
    }
  },
  methods: {
    handleSigninEvent () {
      this.showSignin = true
      this.showSignup = false
    },
    handleSignupEvent () {
      this.showSignin = false
      this.showSignup = true
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
