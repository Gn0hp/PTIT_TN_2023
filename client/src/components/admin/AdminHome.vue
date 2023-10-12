<template>

  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">

    <Header />

    <div class="container" style="height: 80vh;">

      <h1>Admin Dashboard</h1>
      <div class="card">
        <div class="card-header">
          Actions
        </div>
        <div class="card-body d-flex">
          <button class="btn btn-primary">Verify New Candidate</button>

          <button class="btn btn-primary ml-2">User Management</button>
        </div>
      </div>
      <div class="card" style="margin-top: 16px">
        <div class="card-header">
          Elections
        </div>
        <div class="card-body">
          <button class="btn btn-primary">Create New Election</button>
          <table class="table">
            <thead>
            <tr>
              <th>Name</th>
              <th>Date</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="election in elections" :key="election.id">
              <td>{{ election.name }}</td>
              <td>{{ election.date }}</td>
              <td>{{ election.status }}</td>
              <td>
                <button class="btn btn-secondary btn-sm">Edit</button>
                <button class="btn btn-danger btn-sm">Delete</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
        <!-- Other admin widgets -->
    </div>
  <Footer/>
  </div>

</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'

export default {
  name: 'AdminHome',
  components: {
    Footer,
    'Header': Header
  },
  data () {
    return {
      candidates: [],
      users: [],
      elections: [
        {
          id: 1,
          name: '2020 Presidential Election',
          date: '11/3/2020'
        },
        {
          id: 2,
          name: '2020 Senate Election',
          date: '11/3/2020'
        },
        {
          id: 3,
          name: '2020 House Election',
          date: '11/3/2020'
        }
      ]
    }
  },
  methods: {
    verifyCandidate () {
      // Logic to verify candidate
    },

    manageUsers () {
      // Logic to manage users
    }
  }
}
</script>

<style>
/* Styles */
</style>
