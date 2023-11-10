<template>

  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">

    <Header/>

    <div class="container" style="height: 80vh">

      <h1>Candidate Dashboard</h1>

      <div class="jumbotron" v-if="isElecting">
        <h2>Status</h2>
        <p>Current votes: {{ votes }}</p>
      </div>

      <div class="card">
        <div class="card-header">
          Campaign
        </div>
        <div class="card-body">
          <button class="btn btn-primary" @click="postModalShow = !postModalShow">Create Post</button>
          <a href="/#/view_result">
            <button class="btn btn-primary">View Latest Election Result</button>
          </a>
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
              <td>{{ post.created_at }}</td>
              <td>
                <button class="btn btn-secondary btn-sm" @click.stop @click="updatePost(post)" v-b-modal.modal-update-post>Edit</button>
                <button class="btn btn-danger btn-sm" @click.stop @click="deletePost(post.id)">Delete</button>
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

    <b-modal v-if="selectedPost"
             id="modal-update-post"
             :title="selectedPost.title"
             @hide="selectedPost = null"
             centered
             hide-backdrop
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning">

      <form @submit.prevent="submitUpdatePostForm">
        <label class="form-label" for="form3Example4132">Content</label>
        <textarea id="form3Example4132" v-model="selectedPost.content" rows="15"
                  class="form-control"></textarea>
        <div class="text-center text-lg-start mt-4 pt-2">
          <button class="btn btn-primary" type="submit"  style="padding-left: 2.5rem; padding-right: 2.5rem; margin-left: 35%">
            Update
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
import {RequestParams} from '../../config/request'
import {AxiosInstance} from '../../config/auth'

export default {
  components: {
    Footer,
    'Header': Header
  },
  async created () {
    this.candidateInstance = JSON.parse(localStorage.getItem('userGlobal'))
    await AxiosInstance.get(RequestParams.host + RequestParams.path.get_post_by_candidate_id, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('accessToken')}`
      },
      params: {
        candidate_id: this.candidateInstance.id
      }
    })
      .then(response => {
        this.posts = response.data.data.data
      })
      .then(async () => {
        await AxiosInstance.get(`${RequestParams.host}${RequestParams.path.candidate_watch_result}/${this.candidateInstance.id}`,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem('accessToken')}`
            }
          })
          .then(res => {
            console.log(res.data.data.data)
            this.votes = res.data.data.data
          })
      })
      .catch(error => {
        console.log(error)
      })
  },
  data () {
    return {
      selectedPost: null,
      isElecting: true,
      postModalShow: false,
      votes: 0, // sample data,
      candidateInstance: null,
      newPost: {
        title: null,
        content: null
      },
      posts: [
        // {
        //   id: 1,
        //   title: 'My First Post',
        //   content: 'Content for my first post',
        //   publishedDate: '2021-01-01'
        // },
        // {
        //   id: 2,
        //   title: 'My Second Post',
        //   content: 'Content for my second post',
        //   publishedDate: '2021-01-02'
        // },
        // {
        //   id: 3,
        //   title: 'My Third Post',
        //   content: 'Content for my third post',
        //   publishedDate: '2021-01-03'
        // }
      ]
    }
  },
  methods: {
    showModal (post) {
      this.selectedPost = post
      console.log(this.selectedPost)
    },
    async submitCreatePostForm () {
      this.postModalShow = false
      this.newPost.candidate_id = this.candidateInstance.id
      await AxiosInstance.post(RequestParams.host + RequestParams.path.create_post,
        this.newPost, {
          headers:
          {
            Authorization: `Bearer ${localStorage.getItem('accessToken')}`
          }
        }
      )
        .then(response => {
          this.posts.push(response.data.data.data)
        })
        .catch(error => {
          console.log(error)
        })
    },
    async updatePost (post) {
      this.showModal(post)
    },
    async deletePost (id) {
      console.log('post id: ', id)
      await AxiosInstance.delete(RequestParams.host + RequestParams.path.delete_post, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        },
        params: {
          id: id
        }
      })
        .then(response => {
          console.log(response)
          this.posts = this.posts.filter(post => post.id !== id)
        })
        .catch(error => {
          console.log(error)
        })
    },
    async submitUpdatePostForm () {
      console.log(this.selectedPost)
      await AxiosInstance.patch(RequestParams.host + RequestParams.path.patch_post, this.selectedPost, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        },
        params: {
          id: this.selectedPost.id
        }
      }).then(response => {
        console.log(response)
        this.posts = this.posts.map(post => {
          if (post.id === this.selectedPost.id) {
            return this.response.data.data.data
          }
          return post
        })
      }).catch(error => {
        console.log(error)
      })
    }

  }

}
</script>

<style>
/* Styles */
</style>
