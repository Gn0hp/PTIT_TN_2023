<template>

  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">

    <Header/>

    <div class="container mt-5" style="height: 80vh;">

      <h2>{{ candidate.name }}</h2>

      <b-card v-for="post in candidate.posts" :key="post.id" :title="post.title" @click="showModal(post)" v-b-modal.modal-post class="mb-2">

        <b-card-text>
          {{ post.content }}
        </b-card-text>

        <small class="text-muted">{{ post.publishedDate }}</small>

      </b-card>

    </div>
    <b-modal v-if="selectedPost" @hide="selectedPost = null"
             id="modal-post" centered hide-backdrop
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning"
      :title="selectedPost.title">

      <p>{{ selectedPost.content }}</p>

    </b-modal>
    <Footer />

  </div>

</template>

<script>
import Header from '../layouts/Header'
import Footer from '../layouts/Footer'

export default {
  name: 'CandidateDetails',
  components: {Footer, Header},
  data () {
    return {
      selectedPost: null,
      postModalVisible: false,
      candidate: {
        name: 'John Doe',
        posts: [
          {
            id: 1,
            title: 'My First Post',
            content: 'Lorem ipsum...',
            publishedDate: 'Jan 1, 2021'
          },
          {
            id: 2,
            title: 'My Second Post',
            content: 'Lorem ipsum...',
            publishedDate: 'Jan 1, 2021'
          },
          {
            id: 3,
            title: 'My Third Post',
            content: 'Lorem ipsum...',
            publishedDate: 'Jan 1, 2021'
          }
        ]
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
