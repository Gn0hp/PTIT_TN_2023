<template>

  <div>
    <Header/>

    <h1> User Management</h1>
    <button class="btn btn-primary" @click="showVoter">Voter</button>
    <button class="btn btn-primary" @click="showCandidate">Candidate</button>

    <div class="container" v-show="show1">
      <h2> Manage Voters</h2>
      <b-table striped hover :items="voters" :fields="fields">

        <template v-slot:cell(name)="data">
          {{ data.item.name }}
        </template>

        <template v-slot:cell(actions)="data">
          <button class="btn btn-primary" @click="editVoter(data.item.id, data.item)" v-b-modal.edit-modal>
            Edit
          </button>
          <button class="btn btn-dark" @click="deleteCandidate(data.item.id)">Block</button>
          <button class="btn btn-danger" @click="deleteVoter(data.item.id)">Delete</button>
        </template>

      </b-table>
    </div>
    <div class="container" v-show="show2">
      <h2> Manage Candidates</h2>
      <b-table striped hover :items="voters" :fields="fields">
        <template v-slot:cell(name)="data">
          {{ data.item.name }}
        </template>

        <template v-slot:cell(actions)="data">
          <button class="btn btn-primary" @click="editCandidate(data.item.id, data.item)" v-b-modal.edit-modal>Edit</button>
          <button class="btn btn-dark" @click="deleteCandidate(data.item.id)">Block</button>
          <button class="btn btn-danger" @click="deleteCandidate(data.item.id)">Delete</button>
        </template>

      </b-table>
    </div>
    <b-modal v-if="selectedUser" v-model="showEditModal" id="edit-modal"
            @hide="selectedUser = null"
            :title="selectedUser.full_name"
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning"
             centered hide-backdrop
        hide-footer>
        <b-form @submit="submitEdit">
          <label class="form-label" for="formEditFn1">Full Name</label>
          <b-form-input
            id="formEditFn1"
            v-model="selectedUser.full_name">
          </b-form-input>
          <label class="form-label" for="formEditCccd1">CCCD</label>
          <b-form-input
          id="formEditCccd1"
            v-model="selectedUser.cccd_id">
          </b-form-input>
          <label class="form-label" for="formEditAddr1">Address</label>
          <b-form-input
          id="formEditAddr1"
            v-model="selectedUser.address">
          </b-form-input>
          <!-- other form fields for selected user -->
          <div class="text-center text-lg-start mt-4 pt-2">
            <button type="submit" class="btn btn-primary btn-lg" style="padding-left: 2.5rem; padding-right: 2.5rem; margin-left: 35%">
              Update
            </button>
          </div>
        </b-form>
      </b-modal>
    <Footer/>
  </div>

</template>

<script>
// import table data

import Footer from '../layouts/Footer.vue'
import Header from '../layouts/Header.vue'
import {AxiosInstance} from '../../config/auth'
import {RequestParams} from '../../config/request'

export default {
  name: 'UserManagement',
  components: {Header, Footer},
  data () {
    return {
      fields: ['id', 'full_name', 'cccd_id', 'address', 'status', 'actions'],
      voters: [],
      candidates: [],
      show1: false,
      show2: false,
      showEditModal: false,
      selectedUser: null
    }
  },
  methods: {
    editVoter (id, user) {
      this.showEditModal = !this.showEditModal
      this.selectedUser = user
      // open edit modal
    },
    deleteVoter (id) {
      // call API to delete
    },
    editCandidate (id, user) {
      console.log(user)
      this.showEditModal = !this.showEditModal
      this.selectedUser = user
    },
    deleteCandidate (id) {
      // call API to delete
    },
    async showVoter () {
      // show voter table
      this.show1 = true
      this.show2 = false
      this.selectedUser = null

      await AxiosInstance.get(RequestParams.host + RequestParams.path.getVoterList)
        .then((res) => {
          console.log(res.data)
          this.voters = res.data.data.data
        })
        .catch((err) => {
          console.log(err)
        })
    },
    async showCandidate () {
      // show candidate table
      this.show1 = false
      this.show2 = true
      await AxiosInstance.get(RequestParams.host + RequestParams.path.getCandidateList)
        .then((res) => {
          console.log(res.data)
          this.voters = res.data.data.data
        })
        .catch((err) => {
          console.log(err)
        })
    },
    async submitEdit () {

    }
  }
}
</script>
