<template>
  <div>
    <Header />
    <h1> Election Result Statistic </h1>
    <button class="btn btn-primary" @click="statByCandidate">By Candidate</button>
    <button class="btn btn-primary" @click="statByTime">By Time</button>
    <div class="container" v-show="show1" style="padding-top: 2%;">
      <b-table striped hover :items="candidates"
               select-mode="single" ref="selectableTable"
               selectable @row-selected="onRowSelectedCandidate($event)"
               :fields="fields_candidate" >

        <template v-slot:cell(id)="data">
          {{ data.index + 1 }}
        </template>

        <template v-slot:cell(full_name)="data">
          {{ data.item.candidate.full_name }}
        </template>

        <template v-slot:cell(cccd_id)="data">
          {{ data.item.candidate.cccd_id }}
        </template>

        <template v-slot:cell(address)="data">
          {{ data.item.candidate.address }}
        </template>
        <template v-slot:cell(role)="data">
          {{ data.item.election_role.name }}
        </template>

        <template v-slot:cell(ballots)="data">
          {{ data.item.total_vote }}
        </template>
      </b-table>
    </div>

    <div class="container" v-show="show2">
      <b-table striped hover :items="candidates"
               select-mode="single" ref="selectableTable"
               selectable @row-selected="onRowSelectedTime($event)"
               :fields="fields_time">

               <template v-slot:cell(id)="data">
          {{ data.index + 1 }}
        </template>

        <template v-slot:cell(full_name)="data">
          {{ data.item.candidate.full_name }}
        </template>

        <template v-slot:cell(cccd_id)="data">
          {{ data.item.candidate.cccd_id }}
        </template>

        <template v-slot:cell(address)="data">
          {{ data.item.candidate.address }}
        </template>
        <template v-slot:cell(role)="data">
          {{ data.item.election_role.name }}
        </template>

        <template v-slot:cell(ballots)="data">
          {{ data.item.total_vote }}
        </template>

      </b-table>
    </div>
    <Footer />
  </div>
</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'
import {AxiosInstance} from '../../config/auth'
import {RequestParams} from '../../config/request'

export default {
  name: 'AdminStat',
  components: {Footer, Header},
  async created () {
    await AxiosInstance.get(RequestParams.host + RequestParams.path.stat + `/${sessionStorage.getItem('electionId')}`, {
      headers: {
        Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
      }
    })
      .then((res) => {
        console.log(res.data.data.data)
        this.candidates = res.data.data.data
      })
      .catch((err) => {
        console.log(err)
      })
  },
  data () {
    return {
      show1: false,
      show2: false,
      fields_candidate: ['id', 'full_name', 'cccd_id', 'address', 'role', 'ballots'],
      fields_time: ['id', 'full_name', 'cccd_id', 'address', 'role', 'time'],
      candidates: []
    }
  },
  methods: {
    statByCandidate () {
      this.show1 = true
      this.show2 = false
    },
    statByTime () {
      this.show1 = false
      this.show2 = true
    },
    onRowSelectedCandidate (item) {
      console.log(item[0])
      this.$router.push({
        name: 'AdminStatRoleElect',
        params: {
          id: item[0].candidate.id
        }
      })
    },
    onRowSelectedTime (item) {
      console.log(item[0])
    }
  }
}
</script>
