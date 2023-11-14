<template>
  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">

    <Header />

    <div class="container mt-5" style="height: 80vh;">

      <h2>{{ candidate.name }}'s Posts</h2>

      <div v-if="candidate.posts.length > 0 ">
        <b-card v-for="post in candidate.posts" :key="post.id" :title="post.title" @click="showModal(post)"
          v-b-modal.modal-post class="mb-2">

          <b-card-text>
            {{ post.content }}
          </b-card-text>

          <small class="text-muted">{{ post.publishedDate }}</small>

        </b-card>
      </div>
      <div v-else>
        <b-card>
          <b-card-text >
            No posts found
          </b-card-text>
        </b-card>
      </div>

    </div>
    <b-modal v-if="selectedPost" @hide="selectedPost = null" id="modal-post" centered hide-backdrop
      header-bg-variant="dark" header-text-variant="light" footer-bg-variant="dark" body-bg-variant="warning"
      :title="selectedPost.title">

      <p>{{ selectedPost.content }}</p>

    </b-modal>
    <Footer />

  </div>
</template>

<script>
import Header from '../layouts/Header'
import Footer from '../layouts/Footer'
import { AxiosInstance } from '../../config/auth'
import { RequestParams } from '../../config/request'

export default {
  name: 'CandidateDetails',
  components: { Footer, Header },
  async created () {
    const candidateId = this.$route.params.id
    this.candidate.name = this.$route.query.name
    console.log(candidateId)
    await AxiosInstance.get(`${RequestParams.host}${RequestParams.path.get_post_by_candidate_id}?candidate_id=${candidateId}`, {
      headers: {
        Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
      }
    })
      .then(res => {
        console.log(res.data.data.data)
        this.candidate.posts = res.data.data.data
      })
      .catch(err => {
        console.log(err)
      })
  },
  data () {
    return {
      selectedPost: null,
      postModalVisible: false,
      candidate: {
        name: '',
        posts: []
      }
    }
  },
  methods: {
    showModal (post) {
      this.selectedPost = post
      this.postModalVisible = true
    }
  }
}
</script>
