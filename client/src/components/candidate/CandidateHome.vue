<template>

  <div>

    <Header />

    <div class="container">

      <h1>Candidate Dashboard</h1>

      <div class="jumbotron">
        <h2>Status</h2>
        <p>Current votes: {{votes}}</p>
      </div>

      <div class="card">
        <div class="card-header">
          Campaign
        </div>
        <div class="card-body">
          <button class="btn btn-primary">Create Post</button>
          <table class="table">
            <thead>
            <tr>
              <th>Title</th>
              <th>Published</th>
              <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="post in posts" :key="post.id">
              <td>{{ post.title }}</td>
              <td>{{ post.publishedDate }}</td>
              <td>
                <button class="btn btn-secondary btn-sm">Edit</button>
                <button class="btn btn-danger btn-sm">Delete</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>

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
      votes: 100, // sample data
      posts: [
        {
          id: 1,
          title: 'My First Post',
          publishedDate: '2021-01-01'
        },
        {
          id: 2,
          title: 'My Second Post',
          publishedDate: '2021-01-02'
        },
        {
          id: 3,
          title: 'My Third Post',
          publishedDate: '2021-01-03'
        }
      ]
    }
  }
}
</script>

<style>
/* Styles */
</style>
