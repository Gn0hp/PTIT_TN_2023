<template>

  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">

    <Header/>

    <div class="container" style="height: 80vh">

      <h1>Candidate Dashboard</h1>

      <div class="jumbotron">
        <h2>Status</h2>
        <p>Current votes: {{ votes }}</p>
      </div>

      <div class="card">
        <div class="card-header">
          Campaign
        </div>
        <div class="card-body">
          <button class="btn btn-primary" @click="postModalShow = !postModalShow">Create Post</button>
          <table class="table">
            <thead>
            <tr>
              <th>Title</th>
              <th>Published</th>
              <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="post in posts" :key="post.id" @click="showModal(post)" v-b-modal.modal-post>
              <td>{{ post.title }}</td>
              <td>{{ post.publishedDate }}</td>
              <td>
                <button class="btn btn-secondary btn-sm" @click.stop>Edit</button>
                <button class="btn btn-danger btn-sm" @click.stop>Delete</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    <b-modal v-if="selectedPost"
             id="modal-post" class="custom-modal"
             :title="selectedPost.title" @hide="selectedPost = null" centered hide-backdrop
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning"
    >
      <p>{{ selectedPost.content }}</p>
    </b-modal>
    </div>
    <b-modal v-model="postModalShow"
             id="modal-create-post"
             title="Create New Post" centered hide-backdrop
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning"
    >
      <form @submit.prevent="submitCreatePostForm">
        <label class="form-label" for="form3Example412">Title</label>
        <input v-model="newPost.title" type="text" id="form3Example412"
               class="form-control form-control-lg"
               placeholder="Enter title"/>
        <label class="form-label" for="form3Example413">Content</label>
        <textarea  v-model="newPost.content" type="text" id="form3Example413" rows="10"
               class="form-control form-control-lg"
               placeholder="What's on your mind?"></textarea>
        <div class="text-center text-lg-start mt-4 pt-2">
          <button type="submit" class="btn btn-primary btn-lg"
                  style="padding-left: 2.5rem; padding-right: 2.5rem; margin-left: 35%">Create
          </button>
        </div>
      </form>
    </b-modal>
    <Footer/>
  </div>

</template>

<script>

import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'

export default {
  components: {
    Footer,
    'Header': Header
  },
  data () {
    return {
      selectedPost: null,
      postModalShow: false,
      votes: 100, // sample data
      newPost: {
        title: null,
        content: null
      },
      posts: [
        {
          id: 1,
          title: 'My First Post',
          content: 'Content for my first post',
          publishedDate: '2021-01-01'
        },
        {
          id: 2,
          title: 'My Second Post',
          content: 'Content for my second post',
          publishedDate: '2021-01-02'
        },
        {
          id: 3,
          title: 'My Third Post',
          content: 'Content for my third post',
          publishedDate: '2021-01-03'
        }
      ]
    }
  },
  methods: {
    showModal (post) {
      this.selectedPost = post
      console.log(this.selectedPost)
    },
    submitCreatePostForm () {
      console.log(this.newPost)
    }
  }

}
</script>

<style>
/* Styles */
</style>
