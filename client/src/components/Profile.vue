<template>
  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">
    <Header />
    <div style="height: 80vh;">
      <div class="container-fluid h-custom">
        <div class="row d-flex justify-content-center align-items-center h-100">
          <div class="col-md-9 col-lg-6 col-xl-5">
            <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp"
                 class="img-fluid" alt="Sample image">
          </div>
          <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
            <form>
              <div class="divider d-flex align-items-center my-4">
                <h1 class="text-center fw-bold mx-3 mb-0">Your Info</h1>
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
            </form>
          </div>
        </div>
      </div>
    </div>
    <Footer/>
  </div>
</template>

<script>
import Header from './layouts/Header'
import Footer from './layouts/Footer'
import PlaceSelection from './utils/PlaceSelect'
export default {
  name: 'Profile',
  components: {PlaceSelection, Footer, Header},
  data () {
    return {
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
  }
}
</script>

<style scoped>

</style>
