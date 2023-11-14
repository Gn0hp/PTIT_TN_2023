<template>
  <div>
    <Header />
      <h1>Stat Role Elect</h1>
      <div class="container">
        <b-table striped hover :items="roleElects"
                 select-mode="single" ref="selectableTable"
                 selectable @row-selected="onRowSelected($event)"
                 :fields="fields">
          <template v-slot:cell(id)="data">
            {{ data.index + 1 }}
          </template>
                 <template v-slot:cell(name)="data">
            {{ data.item.voter_name }}
          </template>

          <template v-slot:cell(date)="data">
            {{ data.item.vote_time }}
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
  name: 'AdminStatRoleElect',
  components: { Header, Footer },
  async created () {
    const candidateId = this.$route.params.id
    await AxiosInstance.get(`${RequestParams.host}${RequestParams.path.stat_roleElect}/${candidateId}`, {
      headers: {
        Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
      }
    })
      .then(res => {
        console.log(res)
        this.roleElects = res.data.data.data
      })
  },
  data () {
    return {
      roleElects: [],
      fields: ['id', 'name', 'date', 'cccd_id', 'note']
    }
  },
  methods: {
    async onRowSelected (item) {
      console.log(item)
      await this.$router.push({
        name: 'AdminStatBallot',
        params: {
          id: item[0].ballot_id
        }
      })
    }
  }
}

</script>
