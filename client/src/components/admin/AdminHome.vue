<template>

  <div>

    <Header />

    <div class="container">

      <h1>Admin Dashboard</h1>

      <div class="card">
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

  </div>

</template>

<script>
import Header from '../layouts/Header.vue'

export default {
  components: {
    'Header': Header
  },
  data () {
    return {
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
  }
}
</script>

<style>
/* Styles */
</style>
