<template>
  <div>
    <Header/>
    <div class="container mt-5">

      <h1>Candidates</h1>

      <div v-for="group in candidates" :key="group.role">

        <h3>{{ group.role }}</h3>

        <b-table :id="group.role" striped hover :items="group.candidates"
                 select-mode="single" ref="selectableTable"
                 selectable @row-selected="onRowSelected($event, group.role)" :fields="fields"
        >
          <template v-slot:cell(id)="data">
           {{ data.index + 1 }}
          </template>
          <template v-slot:cell(vote)="data">
            <b-form-checkbox :checked="data.item.selected"
                             v-model="data.item.selected"></b-form-checkbox>
          </template>
        </b-table>

      </div>

      <div class="d-flex justify-content-end">
        <b-form @submit.prevent="onSubmitBallot" class="mr-3">
          <b-button type="reset" variant="danger" @click="onReset">Cancel</b-button>
          <b-button type="submit" variant="primary">Submit</b-button>
        </b-form>
      </div>
    </div>
    <b-modal v-if="selectedCandidate" @hide="selectedCandidate = null"
             v-model="candidateModalVisible"
             id="modal-candidate-info" centered hide-backdrop
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning"
             :title="selectedCandidate.full_name">
      <p>{{ selectedCandidate.party }}</p>
    </b-modal>
    <Footer/>
  </div>
</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'
import {AxiosInstance} from '../../config/auth'
import {RequestParams} from '../../config/request'

export default {
  name: 'VoteNow',
  components: {
    Header,
    Footer
  },
  async created () {
    this.ballot.voter_id = JSON.parse(sessionStorage.getItem('userGlobal')).id
    await AxiosInstance.get(RequestParams.host + RequestParams.path.view_candidate, {
      headers: {
        Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
      },
      params: {
        election_id: sessionStorage.getItem('electionId')
      }
    })
      .then(response => {
        const mapCandidates = response.data.data.data.map_candidate // or supported for other map
        for (const [key, val] of Object.entries(mapCandidates)) {
          const tmp = {
            role: key,
            candidates: val
          }
          this.candidates.push(tmp)
        }
        this.candidates.map((item) => {
          if (item.candidates != null) {
            item.candidates.map((candidate) => {
              candidate.selected = false
            })
          }
        })
        console.log(this.candidates)
      })
      .catch(error => {
        console.log(error)
      })
  },
  data () {
    return {
      fields: ['id', 'full_name', 'party', 'vote'],
      candidates: [],
      selectedCandidate: null,
      candidateModalVisible: false,
      ballot: {
        voter_id: null,
        role_elects: []
      }
    }
  },
  methods: {
    showModal (candidate) {
      this.selectedCandidate = candidate
    },
    async onRowSelected (item, tblId) {
      console.log(item, tblId)
      const tblSelected = this.candidates.find((item1) => item1.role === tblId)
      tblSelected.candidates.map((candidate) => {
        candidate.selected = false
      })
      item[0].selected = true
      this.showModal(item[0])
      this.candidateModalVisible = !this.candidateModalVisible
    },
    async onSubmitBallot () {
      this.ballot.role_elects = this.candidates.flatMap((g) => {
        return g.candidates.filter((c) => c.selected).map((c) => {
          console.log('electionId: ', sessionStorage.getItem('electionId'))
          return {
            role_elect_id: c.reId,
            candidate_id: c.id
          }
        })
      })
      this.ballot.election_id = parseInt(sessionStorage.getItem('electionId'))
      console.log('ballot: ', this.ballot)
      await AxiosInstance.post(RequestParams.host + RequestParams.path.vote, this.ballot, {
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
        }
      })
        .then(response => {
          console.log(response)
          this.$router.push('/voters/home')
        })
        .catch(error => {
          console.log(error)
        })
    },
    async onReset () {
      await this.$router.push('/voters/home')
    }
  }
}
</script>
