<template>

  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">

    <Header />

    <div class="container mt-5">

      <div class="row">

        <div class="col-md-4">

          <div class="card">
            <div class="card-header">
              Election Info
            </div>
            <div class="card-body">
              <p class="card-text">{{ electionDate }}</p>
              <p class="card-text">{{ votingLocations }}</p>
              <button class="btn btn-primary" @click="learnMore">
                Learn More
              </button>
            </div>
          </div>

        </div>

        <div class="col-md-8">

          <div class="jumbotron">
            <h2>Welcome to Online Elections!</h2>
            <p>Cast your vote from home.</p>
            <div class="buttons">
              <a href="/#/voters/view_candidates">
                <button class="btn btn-secondary" @click="viewCandidates">
                  View Candidates
                </button>
              </a>

            <a href="/#/voters/vote_now">
              <button class="btn btn-primary" @click="voteNow">
                Vote Now
              </button>
            </a>

            </div>
            <p><br></p>Or
             <button class="btn btn-success" @click="signUp">
               Sign up as Candidate
             </button>

          </div>

          <h2>Announcements</h2>

          <div class="list-group">
            <div class="list-group-item" v-for="item in announcements" :key="item.id">
              {{ item.content }}
            </div>
          </div>

        </div>

      </div>
      <div class="past-elections">

        <h2>Past Elections</h2>

        <ul class="list-group">
          <li class="list-group-item" v-for="election in pastElections" :key="election.id">
            {{ election.name }} - {{ election.date }}
          </li>
        </ul>
      </div>
    </div>
    <Footer/>
  </div>

</template>
<script>

import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'

export default {
  name: 'VoterHome',
  components: {
    'Header': Header,
    'Footer': Footer
  },
  data () {
    return {
      electionDate: 'Nov 8, 2022',
      votingLocations: 'List of locations...',
      announcements: [
        {
          id: 1,
          content: 'Announcement 1'
        },
        {
          id: 2,
          content: 'Announcement 2'
        },
        {
          id: 3,
          content: 'Announcement 3'
        }
      ],
      pastElections: [
        {
          id: 1,
          name: 'Local Election',
          date: 'Nov 2019'
        },
        {
          id: 2,
          name: 'Midterm Election',
          date: 'Nov 2018'
        }
      ]
    }
  },
  methods: {
    viewProfile () {
      // view profile logic
    },
    logout () {
      // logout logic
    },
    learnMore () {
      // learn more logic
    },
    viewCandidates () {
      console.log()
      // view candidate logic
    },
    voteNow () {
      // vote now logic
    },
    signUp () {
      this.$router.push('/register_candidate')
      // Sign up as candidate logic
    }
  }
}
</script>

<style scoped>
/* CSS styling */
.buttons {
  display: flex;
}

.buttons button {
  margin-right: 10px;
}

</style>
