<template>
<div>
  <Header />
  <div class="container-fluid h-custom">
    <div class="row d-flex justify-content-center align-items-center h-100">
      <div class="col-md-9 col-lg-6 col-xl-5">
        <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp"
             class="img-fluid" alt="Sample image">
      </div>
      <div class="col-md-8 col-lg-6 col-xl-4 offset-xl-1">
        <form @submit.prevent="submitForm">
          <div class="divider d-flex align-items-center my-4">
            <h1 class="text-center fw-bold mx-3 mb-0">Sign Up</h1>
          </div>
          <div class="form-outline">
            <label class="form-label" for="party">Party</label>
            <input type="text" id="party" class="form-control" placeholder="Your Party" v-model="candidate.party">
          </div>

          <div class="form-outline">
            <label class="form-label" for="objective">Objective</label>
            <input type="text" id="objective" class="form-control" placeholder="A little about your objective" v-model="candidate.objective">
          </div>

          <div class="form-outline">
            <label class="form-label" for="describe">Describe</label>
            <textarea id="describe" class="form-control" v-model="candidate.describe" placeholder="Your Abridged Biography "></textarea>
          </div>

          <div class="form-outline">
            <label class="form-label" for="patrimony">Patrimony</label>
            <textarea id="patrimony" class="form-control" placeholder="Your Patrimony" v-model="candidate.patrimony"></textarea>
          </div>

          <div class="form-outline">
            <label class="form-label" for="role">Role</label>
            <select id="role" class="form-select" v-model="candidate.role">
              <option value="" disabled selected>Select your elected Role</option>
              <option v-for="role in roles" :key="role.id" :value="role.id">
                {{ role.name }}
              </option>
              <!-- other roles -->
            </select>
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
          </div>
        </form>
      </div>
    </div>
  </div>
  <Footer/>
</div>
</template>

<script>
import Header from '../layouts/Header'
import PlaceSelection from '../utils/PlaceSelect'
import Footer from '../layouts/Footer'
import {RequestParams} from '../../config/request'
import { AxiosInstance } from '../../config/auth'
export default {
  name: 'RegisterCandidate',
  components: {Footer, PlaceSelection, Header},
  async created () {
    console.log(sessionStorage)
    await AxiosInstance.get(RequestParams.host + RequestParams.path.election_roles_by_election_id, {
      params: {
        election_id: sessionStorage.getItem('electionId')
      },
      headers: {
        Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
      }
    }).then(res => {
      this.roles = res.data.data.data
    }).catch(err => {
      console.error(err)
    })
  },
  data () {
    return {
      roles: [],
      candidate: {
        party: '',
        patrimony: '',
        objective: '',
        describe: '',
        role: ''
      }
    }
  },
  methods: {
    async submitForm () {
      const user = JSON.parse(sessionStorage.getItem('userGlobal'))
      console.log(user)
      const postBody = {
        address: user.address,
        cccd_id: user.cccd_id,
        cccd_date: user.cccd_date,
        cccd_place: user.cccd_place,
        date_of_birth: user.date_of_birth,
        email: user.email,
        full_name: user.full_name,
        gender: user.gender,
        password: user.password,
        phone: user.phone,
        patrimony: this.candidate.patrimony,
        status: user.status,
        party: this.candidate.party,
        note: `Objective: ${this.candidate.objective}, Describe: ${this.candidate.describe}`,
        role: this.candidate.role
      }
      console.log(postBody)
      await AxiosInstance.post(RequestParams.host + RequestParams.path.registerCandidate, postBody, {
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
        }
      }).then(data => {
        console.log(data)
        this.$router.push('/voters/home')
      }).catch(err => {
        console.error(err)
      })
    }
  }
}
</script>
