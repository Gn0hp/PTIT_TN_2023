<template>
  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">
    <Header />
    <div class="container mt-5" style="height: 80vh;">

      <h1>Candidates</h1>

      <div v-for="group in candidates" :key="group.role">

        <h3>{{ group.role }}</h3>

        <b-table striped hover :items="group.candidates" :fields="fields"
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

    <!-- Additional tables for other roles -->
    <Footer/>
  </div>
</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'
import {RequestParams} from '../../config/request'
import {AxiosInstance} from '../../config/auth'

export default {
  name: 'ViewCandidate',
  components: {
    'Header': Header,
    'Footer': Footer
  },
  async created () {
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

        console.log(this.candidates)
      })
      .catch(error => {
        console.log(error)
      })
  },
  data () {
    return {
      fields: ['id', 'full_name', 'party'],
      candidates: [
        // {
        //   role: 'President',
        //   candidates: [
        //     {id: 1, name: 'John Doe', party: 'Democrat'},
        //     {id: 2, name: 'Jane Smith', party: 'Republican'}
        //   ]
        // },
        // {
        //   role: 'Vice President',
        //   candidates: [
        //     {id: 1, name: 'John Doe', party: 'Democrat'},
        //     {id: 2, name: 'Jane Smith', party: 'Republican'}
        //   ]
        // }
      ]
    }
  },
  methods: {
    onRowSelected (item) {
      this.$router.push(`/voters/view_candidates/${item[0].id}?name=${item[0].full_name}`)
    }
  }
}
</script>
