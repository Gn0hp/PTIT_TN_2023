<template>
  <div>
    <Header/>
    <b-container>

      <h2>Create New Election</h2>

      <b-form @submit="onSubmit">

        <b-form-group label="Organization:">
          <b-form-input v-model="election.organization"/>
        </b-form-group>

        <b-form-group label="Date End Register:">
          <b-form-datepicker v-model="election.dateEndRegistering"/>
        </b-form-group>

        <b-form-group label="Date Start Voting:">
          <b-form-datepicker v-model="election.dateStartVoting"/>
        </b-form-group>

        <b-form-group label="Date End Voting:">
          <b-form-datepicker v-model="election.dateEndVoting"/>
        </b-form-group>

        <b-form-group
          v-for="(role, index) in election.roles"
          :key="index"
          label="Role">

          <b-form-input
            v-model="role.name"
            placeholder="Enter role name">
          </b-form-input>

        </b-form-group>

        <b-button @click="addRole">
          Add Role
        </b-button>

        <b-button type="submit" variant="primary">Create</b-button>
      </b-form>

    </b-container>
    <Footer/>
  </div>
</template>

<script>
import {BForm, BFormGroup, BFormInput, BFormDatepicker, BButton} from 'bootstrap-vue'
import Footer from '../layouts/Footer.vue'
import Header from '../layouts/Header.vue'
import {AxiosInstance} from '../../config/auth'
import {RequestParams} from '../../config/request'

export default {
  name: 'CreateElection',
  components: {
    Header,
    Footer,
    BForm,
    BFormGroup,
    BFormInput,
    BFormDatepicker,
    BButton
  },

  data () {
    return {
      election: {
        organization: '',
        dateEndRegistering: '',
        dateStartVoting: '',
        dateEndVoting: '',
        date_end_register: 0,
        date_start_electing: 0,
        duration: 0,
        roles: [
          {
            name: ''
          }]
      }
    }
  },
  methods: {
    async onSubmit () {
      // submit form
      this.election.date_start_electing = this.convertDateToTs(this.election.dateStartVoting)
      this.election.dateEndElecting = this.convertDateToTs(this.election.dateEndVoting)
      this.election.date_end_register = this.convertDateToTs(this.election.dateEndRegistering)
      this.election.duration = this.election.dateEndElecting - this.election.date_start_electing
      console.log(this.election)
      // TODO: post to server
      await AxiosInstance.post(RequestParams.host + RequestParams.path.createElection, this.election).then(data => {
        console.log(data)
        this.$router.push('/admin/home')
      }).catch(err => {
        console.log(err)
        throw (err)
      })
    },
    addRole () {
      this.election.roles.push({
        name: ''
      })
    },
    convertDateToTs (dateString) {
      const formattedDateString = dateString.replace(',', '') // Remove the comma
      const timestamp = Date.parse(formattedDateString)

      return timestamp / 1000
    }
  }

}
</script>
