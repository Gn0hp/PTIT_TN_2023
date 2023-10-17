<template>

  <div>
    <Header />

    <h1> User Management</h1>
    <button class="btn btn-primary" @click="showVoter">Voter</button>
    <button class="btn btn-primary" @click="showCandidate">Candidate</button>

    <div v-show="show1">
      <h2> Manage Voters</h2>
      <b-table striped hover :items="voters" :fields="fields">

        <template v-slot:cell(name)="data">
          {{ data.item.name }}
        </template>

        <template v-slot:cell(actions)="data">
          <button class="btn btn-primary" @click="editVoter(data.item.id)">Edit</button>
          <button class="btn btn-danger" @click="deleteVoter(data.item.id)">Delete</button>
        </template>

      </b-table>
    </div>
    <div v-show="show2">
        <h2> Manage Candidates</h2>
      <b-table striped hover :items="voters" :fields="fields" >
        <template v-slot:cell(name)="data">
          {{ data.item.name }}
        </template>

        <template v-slot:cell(actions)="data">
          <button class="btn btn-primary" @click="editCandidate(data.item.id)">Edit</button>
          <button class="btn btn-danger" @click="deleteCandidate(data.item.id)">Delete</button>
        </template>

      </b-table>
    </div>
    <Footer />
  </div>

</template>

<script>
// import table data

import Footer from '../layouts/Footer.vue'
import Header from '../layouts/Header.vue'

export default {
  name: 'UserManagement',
  components: {Header, Footer},
  data () {
    return {
      fields: ['id', 'name', 'cccd', 'actions'],
      voters: [
        {
          id: 1,
          name: 'John Doe',
          cccd: '1234567890123456'
        },
        {
          id: 2,
          name: 'Jane Doe',
          cccd: '1234567890123456'
        },
        {
          id: 3,
          name: 'John Smith',
          cccd: '1234567890123456'
        }
      ],
      show1: false,
      show2: false
    }
  },
  methods: {
    editVoter (id) {
      // open edit modal
    },
    deleteVoter (id) {
      // call API to delete
    },
    editCandidate (id) {
      // open edit modal
    },
    deleteCandidate (id) {
      // call API to delete
    },
    showVoter () {
      // show voter table
      this.show1 = true
      this.show2 = false
    },
    showCandidate () {
      // show candidate table
      this.show1 = false
      this.show2 = true
    }
  }
}
</script>
