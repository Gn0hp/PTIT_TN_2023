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
        {{ data.item.role_elect }}
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

export default {
  name: 'VerifyUser',
  components: {Header, Footer},
  data () {
    return {
      fields: ['id', 'name', 'cccd', 'status', 'actions'],
      fieldsCandidate: ['id', 'name', 'cccd', 'roleElect', 'status', 'actions'],
      items: []
      // unverifiedVoters: [
      //   {
      //     id: 1,
      //     name: 'John Doe',
      //     cccd: '1234567890123456'
      //   },
      //   {
      //     id: 2,
      //     name: 'Jane Doe',
      //     cccd: '1234567890123456'
      //   },
      //   {
      //     id: 3,
      //     name: 'John Smith',
      //     cccd: '1234567890123456'
      //   }
      // ]
    }
  },
  created () {
    this.items = JSON.parse(this.$route.query.users)
    console.log(this.items)
    console.log('axios111 ', AxiosInstance.defaults.headers.common)
  },
  methods: {
    verifyVoter (id) {
      // call API to verify
      console.log(id)
    },
    rejectVoter (id) {
      // call API to reject
      console.log(id)
    },
    verifyCandidate (id) {
      // call API to verify
      console.log(id)
    },
    rejectCandidate (id) {
      // call API to reject
      console.log(id)
    }
  }
}
</script>
