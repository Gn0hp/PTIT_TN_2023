<template>

  <div>
    <Header/>
    <h2>Verify New Voter</h2>

    <b-table striped hover :items="items.voters" :fields="fields">

      <template v-slot:cell(name)="data">
        {{ data.item.full_name }}
      </template>
      <template v-slot:cell(cccd)="data">
        {{ data.item.cccd_id }}
      </template>

      <template v-slot:cell(actions)="data">
        <button class="btn btn-primary" @click="verifyVoter(data.item.id)">Verify</button>
        <button class="btn btn-danger" @click="rejectVoter(data.item.id)">Reject</button>
      </template>

    </b-table>

    <h2>Verify New Candidate</h2>
    <b-table striped hover :items="items.candidates" :fields="fieldsCandidate">

      <template v-slot:cell(name)="data">
        {{ data.item.full_name }}
      </template>
      <template v-slot:cell(cccd)="data">
        {{ data.item.cccd_id }}
      </template>
      <template v-slot:cell(roleElect)="data">
        {{ data.item.candidate_status }}
      </template>
      <template v-slot:cell(actions)="data">
        <button class="btn btn-primary" @click="verifyCandidate(data.item.id)">Verify</button>
        <button class="btn btn-danger" @click="rejectCandidate(data.item.id)">Reject</button>
      </template>

    </b-table>
    <Footer/>
  </div>

</template>

<script>
// import table data

import { AxiosInstance } from '../../config/auth'
import Footer from '../layouts/Footer.vue'
import Header from '../layouts/Header.vue'
import {RequestParams} from '../../config/request'

export default {
  name: 'VerifyUser',
  components: {Header, Footer},
  data () {
    return {
      fields: ['id', 'name', 'cccd', 'status', 'actions'],
      fieldsCandidate: ['id', 'name', 'cccd', 'status', 'candidate_status', 'actions'],
      items: [
      // unverifiedVoters: [
      ]
    }
  },
  created () {
    this.items = JSON.parse(this.$route.query.users)
    console.log(this.items)
  },
  methods: {
    async verifyVoter (id) {
      await AxiosInstance.post(RequestParams.host + RequestParams.path.verify_voter, {id: id}, {
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
        }
      })
        .then((response) => {
          console.log(response)
          this.items.voters = this.items.voters.filter((item) => item.id !== id)
        })
        .catch((error) => {
          console.log(error)
        })
    },
    rejectVoter (id) {
      // call API to reject
      console.log(id)
    },
    async verifyCandidate (id) {
      await AxiosInstance.post(RequestParams.host + RequestParams.path.verify_candidate, {id: id}, {
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
        }
      })
        .then((response) => {
          console.log(response)
          this.items.candidates = this.items.candidates.filter((item) => item.id !== id)
        })
        .catch((error) => {
          console.log(error)
        })
    },
    rejectCandidate (id) {
      // call API to reject
      console.log(id)
    }
  }
}
</script>
