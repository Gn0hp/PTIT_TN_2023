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
          <b-form-datepicker v-model="election.dateEndRegister"/>
        </b-form-group>

        <b-form-group label="Date Start Voting:">
          <b-form-datepicker v-model="election.dateStartVoting"/>
        </b-form-group>

        <b-form-group label="Date End Voting:">
          <b-form-datepicker v-model="election.dateEndVoting"/>
        </b-form-group>

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
        dateEndRegister: '',
        dateStartVoting: '',
        dateEndVoting: ''
      }
    }
  },

  methods: {
    onSubmit () {
      // submit form
    }
  }

}
</script>
