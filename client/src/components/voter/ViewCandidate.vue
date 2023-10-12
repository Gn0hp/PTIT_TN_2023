<template>
  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">
    <Header :user-name="this.voterName"></Header>
    <div class="container mt-5" style="height: 80vh;">

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

    <!-- Additional tables for other roles -->
    <Footer/>
  </div>
</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'

export default {
  name: 'ViewCandidate',
  components: {
    'Header': Header,
    'Footer': Footer
  },
  data () {
    return {
      voterName: 'John Doe',
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
    onRowSelected (item) {
      this.$router.push(`/voters/view_candidates/${item[0].id}`)
    }
  }
}
</script>
