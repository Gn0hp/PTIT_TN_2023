<template>
  <div>
    <Header/>
    <div class="container mt-5">

      <h1>Candidates</h1>

      <div v-for="group in candidates" :key="group.role">

        <h3>{{ group.role }}</h3>

        <b-table striped hover :items="group.candidates"
                 select-mode="single" ref="selectableTable"
                 selectable @row-selected="onRowSelected"
        >
          <template v-slot:cell(name)="data">
            {{ data.item.name }}
          </template>

          <template v-slot:cell(party)="data">
            {{ data.item.party }}
          </template>
        </b-table>

      </div>

    </div>
    <b-modal v-if="selectedCandidate" @hide="selectedCandidate = null"
             v-model="candidateModalVisible"
             id="modal-candidate-info" centered hide-backdrop
             header-bg-variant="dark"
             header-text-variant="light"
             footer-bg-variant="dark"
             body-bg-variant="warning"
             :title="selectedCandidate.name">
      <p>{{ selectedCandidate.party }}</p>
    </b-modal>
    <Footer/>
  </div>
</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'

export default {
  name: 'VoteNow',
  components: {
    Header,
    Footer
  },
  data () {
    return {
      selectedCandidate: null,
      candidateModalVisible: false,
      candidates: [
        {
          role: 'President',
          candidates: [
            {id: 1, name: 'John Doe', party: 'Democrat'},
            {id: 2, name: 'Jane Smith', party: 'Republican'}
          ]
        },
        {
          role: 'Vice President',
          candidates: [
            {id: 1, name: 'John Doe', party: 'Democrat'},
            {id: 2, name: 'Jane Smith', party: 'Republican'}
          ]
        }
      ]
    }
  },
  methods: {
    showModal (candidate) {
      this.selectedCandidate = candidate
    },
    onRowSelected (item) {
      console.log(item[0])
      this.showModal(item[0])
      this.candidateModalVisible = !this.candidateModalVisible
    }
  }
}
</script>
