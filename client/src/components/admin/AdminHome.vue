<template>

  <div style="background-image: url('https://mdbootstrap.com/img/Photos/Others/images/76.jpg');">
    <Header/>

    <div class="container" style="height: 80vh;">
      <b-toast id="my-toast" solid static toaster="b-toaster-top-center">
      <template #toast-title>
        <div class="d-flex flex-grow-1 align-items-baseline">
          <b-img blank blank-color="#ff5555" class="mr-2" width="12" height="12"></b-img>
          <strong class="mr-auto">{{ toast_title }}</strong>
          <small class="text-muted mr-2">42 seconds ago</small>
        </div>
      </template>
      {{ toast_message }}
    </b-toast>
    <div style="padding-top: 4%">
      <h1>Admin Dashboard</h1>
      <div class="card">
        <div class="card-header">
          Actions
        </div>
        <div class="card-body d-flex">
          <button class="btn btn-primary" @click="verifyUser">Verify New User</button>

          <button class="btn btn-primary ml-2" @click="userManagementDirect">User Management</button>

          <button class="btn btn-primary ml-2" @click="statDirect">Election
            Result Stat
          </button>

          <button class="btn btn-primary ml-2" v-bind:class="{disabled: hasNoElection}" @click="closeElectionRegister">
            Close Election Register
          </button>

          <button class="btn btn-primary ml-2" v-bind:class="{disabled: hasNoElection}" @click="closeElection">Close
            Election
          </button>

          <a href="/#/view_result">
            <button class="btn btn-primary ml-2">View Latest Election Result</button>
          </a>
        </div>
      </div>
      <div class="card" style="margin-top: 16px">
        <div class="card-header">
          Elections
        </div>
        <div class="card-body">
          <button class="btn btn-primary" @click="createElection">Create New Election</button>
          <table class="table">
            <thead>
            <tr>
              <th>Name</th>
              <th>Date</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="election in elections" :key="election.id">
              <td>{{ election.name }}</td>
              <td>{{ election.date }}</td>
              <td>{{ election.status }}</td>
              <td>
                <button class="btn btn-secondary btn-sm">Edit</button>
                <button class="btn btn-danger btn-sm">Delete</button>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!-- Other admin widgets -->
    </div>
    </div>
    <Footer/>
  </div>

</template>

<script>
import Header from '../layouts/Header.vue'
import Footer from '../layouts/Footer.vue'
import {RequestParams} from '../../config/request'
import {AxiosInstance} from '../../config/auth'
import userManagement from './UserManagement.vue'

export default {
  name: 'AdminHome',
  computed: {
    userManagement () {
      return userManagement
    }
  },
  created () {
    this.hasNoElection = sessionStorage.getItem('hasElection') === 'false'
  },
  components: {
    Footer,
    'Header': Header
  },
  data () {
    return {
      hasNoElection: false,
      toast_message: '',
      toast_title: '',
      candidates: [],
      users: [],
      elections: [
        {
          id: 1,
          name: '2020 Presidential Election',
          date: '11/3/2020'
        },
        {
          id: 2,
          name: '2020 Senate Election',
          date: '11/3/2020'
        },
        {
          id: 3,
          name: '2020 House Election',
          date: '11/3/2020'
        }
      ]
    }
  },
  methods: {

    manageUsers () {
      // Logic to manage users
    },
    async verifyUser () {
      try {
        const res = await AxiosInstance.get(RequestParams.host + RequestParams.path.adminGetPendingUsers, {
          headers: {
            Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
          }
        })
        console.log(JSON.stringify(res.data.data.data))
        await this.$router.push({
          path: '/admin/verify',
          query: {
            users: JSON.stringify(res.data.data.data)
          }
        })
      } catch (error) {
        console.log(error)
      }
    },
    async createElection () {
      await this.$router.push('/admin/create_election')
    },
    async closeElectionRegister () {
      const postBody = {
        id: parseInt(sessionStorage.getItem('electionId')),
        date_start_electing: Math.floor(Date.now() / 1000),
        duration: RequestParams.default_elect_duration
      }
      console.log(postBody)
      await AxiosInstance.post(RequestParams.host + RequestParams.path.close_election + `/${sessionStorage.getItem('electionId')}`, postBody, {
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
        }
      })
        .then(res => {
          console.log(res.data.data)
          const success = res.data.data.success
          this.toast_title = success ? 'Successfully' : 'Failed'
          this.toast_message = 'Close election register' + (success ? 'successfully' : 'failed')
          this.$bvToast.show('my-toast', {
            toaster: 'b-toaster-top-right',
            autoHideDelay: 3000,
            title: this.toast_title,
            variant: success ? 'success' : 'danger'
            // 'no-auto-hide':true,
          })
        })
        .catch(err => {
          console.log(err)
        })
    },
    async userManagementDirect () {
      await this.$router.push('/admin/user-management')
    },
    async statDirect () {
      await this.$router.push('/admin/stat')
    },
    showToast (success, message) {
      console.log('go here with', success, message)
    },
    async closeElection () {
      const electionId = sessionStorage.getItem('electionId')
      await AxiosInstance.post(`${RequestParams.host}${RequestParams.path.finish_election}/${electionId}`, {
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('accessToken')}`
        }
      })
        .then(async res => {
          const success = res.data.data.success
          this.toast_title = success ? 'Successfully' : 'Failed'
          this.toast_message = 'Close election ' + (success ? 'successfully' : 'failed')
          this.$bvToast.show('my-toast', {
            toaster: 'b-toaster-top-right',
            autoHideDelay: 3000,
            title: this.toast_title,
            variant: success ? 'success' : 'danger'
            // 'no-auto-hide':true,
          })
        })
    }
  }
}
</script>

<style>
/* Styles */
.toast:not(.show) {
   display: block;
   position: fixed !important;
   right: 2% !important;
}
</style>
