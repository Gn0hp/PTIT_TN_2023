<template>
  <div>
    <Header />
    <h1>Ballot Details by {{ username }}</h1>
    <div class="container">
      <b-table striped hover :items="ballots"
               select-mode="single" ref="selectableTable"
               selectable @row-selected="onRowSelected($event)"
               :fields="fields">
        <template v-slot:cell(id)="data">
          {{ data.index + 1 }}
        </template>
        <template v-slot:cell(candidate)="data">
          {{ data.item.candidate_name }}
        </template>
        <template v-slot:cell(role)="data">
          {{ data.item.role_name }}
        </template>
        <template v-slot:cell(cccd)="data">
          {{ data.item.candidate_cccd }}
        </template>
      </b-table>
    </div>
    <Footer />
  </div>
</template>
<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'
import { AxiosInstance } from '../../config/auth'
import { RequestParams } from '../../config/request'
export default {
  name: 'AdminStatBallot',
  components: { Header, Footer },
  async created () {
    const ballotId = this.$route.params.id
    await AxiosInstance.get(`${RequestParams.host}${RequestParams.path.stat_ballot}/${ballotId}`)
      .then(res => {
        console.log(res.data.data.data)
        this.ballots = res.data.data.data
        this.username = this.ballots[0].voter_name
      })
  },
  data () {
    return {
      username: '',
      ballots: [],
      fields: ['id', 'candidate', 'role', 'cccd', 'note']
    }
  },
  methods: {
    async onRowSelected (item) {
      console.log(item)
    }
  }
}

</script>
